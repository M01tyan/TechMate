package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"io/ioutil"
	"database/sql"

	"github.com/heroku/TechMate/modules"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	_"github.com/lib/pq"
)


func main() {
	port := os.Getenv("PORT")
	var mode string
	Db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
	    log.Print(err)
	    Db.Close()
	}

	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/callback", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Print(err)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					mode = modules.GetLineID(Db, event.Source.UserID)
					switch mode {
					case "init_new":
						r := regexp.MustCompile(`([sdm])1([0-9]{6})`)
						if r.MatchString(message.Text) {
							modules.InsertStudentID(Db, message.Text, event.Source.UserID)
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("名前を入力してください。")).Do(); err != nil {
								log.Print(err)
							}
							mode = modules.UpdateMode(Db, 2, event.Source.UserID)
						} else {
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("学籍番号が正しくありません。\nもう一度入力してください。")).Do(); err != nil {
								log.Print(err)
							}
						}
					case "init_name":
						modules.InsertName(Db, message.Text, event.Source.UserID)
						genres_json, err := ioutil.ReadFile("./modules/genre_flex.json")
						if err != nil {
					        log.Fatal(err)
					    }
						genre_flex, errs := linebot.UnmarshalFlexMessageJSON(genres_json)
						if errs != nil {
							log.Print(errs)
						}
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("ジャンル", genre_flex)).Do(); err != nil {
							log.Print(err)
						}
						mode = modules.UpdateMode(Db, 3, event.Source.UserID)
					case "init_genre":
						if message.Text == "終了" {
							confirm := modules.Confirm(modules.GetStudentID(Db, event.Source.UserID), modules.GetName(Db, event.Source.UserID), modules.GetGenres(Db, event.Source.UserID))
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("ジャンル", confirm)).Do(); err != nil {
								log.Print(err)
							}
							mode = modules.UpdateMode(Db, 4, event.Source.UserID)
						} else {
							modules.InsertGenre(Db, message.Text, event.Source.UserID)
						}
					case "init_continue":
						if message.Text == "はい" {
							//modules.InsertData(name, line_id, student_id, my_genre)
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("登録しました。")).Do(); err != nil {
								log.Print(err)
							}
							mode = modules.UpdateMode(Db, 5, event.Source.UserID)
						} else {
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("登録をキャンセルしました。\nもう一度学籍番号から入力してください。")).Do(); err != nil {
								log.Print(err)
							}
							modules.DeleteData(Db, event.Source.UserID)
							mode = modules.UpdateMode(Db, 1, event.Source.UserID)
						}
					case "default":
						if message.Text == "検索" {
							genres_json, err := ioutil.ReadFile("./modules/genre_flex.json")
							if err != nil {
						        log.Fatal(err)
						    }
							genre_flex, errs := linebot.UnmarshalFlexMessageJSON(genres_json)
							if errs != nil {
								log.Print(errs)
							}
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("ジャンル", genre_flex)).Do(); err != nil {
								log.Print(err)
							}
							mode = modules.UpdateMode(Db, 6, event.Source.UserID)
						}
					case "search":
						result := modules.GetPost(Db, message.Text)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(result)).Do(); err != nil {
							log.Print(err)
						}
						mode = modules.UpdateMode(Db, 5, event.Source.UserID)
					}
				}
			} else if event.Type == linebot.EventTypeFollow {
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ご登録ありがとうございます。\nあなたの学籍番号を入力してください。")).Do(); err != nil {
					log.Print(err)
				}
				mode = modules.GetLineID(Db, event.Source.UserID)
			}
		}
	})

	router.Run(":" + port)
}

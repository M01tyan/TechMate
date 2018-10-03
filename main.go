package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"io/ioutil"

	"github.com/heroku/TechMate/modules"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	_"github.com/lib/pq"
)


func main() {
	port := os.Getenv("PORT")
	mode := "init_new"
	student_id := ""
	line_id := ""
	name := ""
	my_genre := []string{}

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
					switch mode {
					case "init_new":
						line_id = event.Source.UserID
						r := regexp.MustCompile(`([sdm])1([0-9]{6})`)
						if r.MatchString(message.Text) {
							student_id = message.Text 
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("名前を入力してください。")).Do(); err != nil {
								log.Print(err)
							}
							mode = "init_name"
						} else {
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("学籍番号が正しくありません。\nもう一度入力してください。")).Do(); err != nil {
								log.Print(err)
							}
						}
					case "init_name":
						name = message.Text
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
						mode = "init_genre"
					case "init_genre":
						if message.Text == "終了" {
							confirm := modules.Confirm(student_id, name, my_genre)
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("ジャンル", confirm)).Do(); err != nil {
								log.Print(err)
							}
							mode = "init_continue"
						} else {
							my_genre = append(my_genre, message.Text)
						}
					case "init_continue":
						if message.Text == "はい" {
							modules.InsertData(name, line_id, student_id, my_genre)
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("登録しました。")).Do(); err != nil {
								log.Print(err)
							}
							my_genre = nil
							mode = "default"
						} else {
							student_id = ""
							name = ""
							my_genre = nil
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("登録をキャンセルしました。\nもう一度学籍番号から入力してください。")).Do(); err != nil {
								log.Print(err)
							}
							mode = "init_new"
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
							mode = "search"
						}
					case "search":
						var print_result string
						search := message.Text
						search_result := modules.GetPost(search)
						for _, r := range search_result {
							print_result += r.NAME + "\t" + r.STUDENT_ID + "\n"
						}
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(print_result)).Do(); err != nil {
							log.Print(err)
						}
						mode = "default"
					}
				}
			} else if event.Type == linebot.EventTypeFollow {
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ご登録ありがとうございます。\nあなたの学籍番号を入力してください。")).Do(); err != nil {
					log.Print(err)
				}
				mode = "init_new"
			}
		}
	})

	router.Run(":" + port)
}

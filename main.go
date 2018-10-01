package main

import (
	"log"
	"net/http"
	"os"
	"database/sql"

//	"github.com/heroku/TechMate/modules"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	_ "github.com/lib/pq"
)


func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
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
					log.Print("success")
					var Db *sql.DB
				    Db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
				    if err != nil {
				    	panic(err)
				        log.Print(err)
				        Db.Close()
				    }
				    var genre_name string
				    errs := Db.QueryRow("SELECT name FROM genres WHERE id=$1", 1).Scan(&genre_name)
				    if err != nil {
				        log.Print(errs)
				    }
				    log.Print(genre_name)
					text := message.Text + genre_name
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})

	router.Run(":" + port)
}

package main

import (
	"log"
	"net/http"
	_"os"
	"regexp"

	"github.com/heroku/TechMate/modules"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	_"github.com/lib/pq"
)


func main() {
	//port := os.Getenv("PORT")
	port := "8080"
	mode := "init_new"
	student_id := ""
	line_id := ""
	name := ""
	my_genre := []string{}
/*
	flex_message := []byte(`{
	  "type": "carousel",
	  "contents": [
	    {
	      "type": "bubble",
	      "body": {
	        "type": "box",
	        "layout": "vertical",
	        "spacing": "sm",
	        "contents": [
	          {
	            "type": "text",
	            "flex": 1,
	            "text": "ジャンルを選択してください"
	          },
	          {
	            "type": "text",
	            "flex": 1,
	            "text": "(※最大５個まで)"
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "C",
	              "text": "C"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "Java",
	              "text": "Java"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "C++",
	              "text": "C++"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "Ruby on Rails",
	              "text": "Ruby on Rails"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "Python",
	              "text": "Python"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "Swift",
	              "text": "Swift"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "color": "#aaaaaa",
	            "action": {
	              "type": "message",
	              "label": "終了",
	              "text": "終了"
	            }
	          }
	        ]
	      }
	    },
	    {
	      "type": "bubble",
	      "body": {
	        "type": "box",
	        "layout": "vertical",
	        "spacing": "sm",
	        "contents": [
	          {
	            "type": "text",
	            "flex": 1,
	            "text": "ジャンルを選択してください"
	          },
	          {
	            "type": "text",
	            "flex": 1,
	            "text": "(※最大５個まで)"
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "HTML",
	              "text": "HTML"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "CSS",
	              "text": "CSS"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "JavaScript",
	              "text": "JavaScript"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "Kotlin",
	              "text": "Kotlin"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "React",
	              "text": "React"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "Vue",
	              "text": "Vue"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "color": "#aaaaaa",
	            "action": {
	              "type": "message",
	              "label": "終了",
	              "text": "終了"
	            }
	          }
	        ]
	      }
	    },
	    {
	      "type": "bubble",
	      "body": {
	        "type": "box",
	        "layout": "vertical",
	        "spacing": "sm",
	        "contents": [
	          {
	            "type": "text",
	            "flex": 1,
	            "text": "ジャンルを選択してください"
	          },
	          {
	            "type": "text",
	            "flex": 1,
	            "text": "(※最大５個まで)"
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "PHP",
	              "text": "PHP"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "Go",
	              "text": "Go"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "SQL",
	              "text": "SQL"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "Ruby",
	              "text": "Ruby"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "Unity",
	              "text": "Unity"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "Perl",
	              "text": "Perl"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "color": "#aaaaaa",
	            "action": {
	              "type": "message",
	              "label": "終了",
	              "text": "終了"
	            }
	          }
	        ]
	      }
	    },
	    {
	      "type": "bubble",
	      "body": {
	        "type": "box",
	        "layout": "vertical",
	        "spacing": "sm",
	        "contents": [
	          {
	            "type": "text",
	            "flex": 1,
	            "text": "ジャンルを選択してください"
	          },
	          {
	            "type": "text",
	            "flex": 1,
	            "text": "(※最大５個まで)"
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "Git",
	              "text": "Git"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "Raspberry Pi",
	              "text": "Raspberry Pi"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "AWS",
	              "text": "AWS"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "AI",
	              "text": "AI"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "Deep Learning",
	              "text": "Deep Learning"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "action": {
	              "type": "message",
	              "label": "画像解析",
	              "text": "画像解析"
	            }
	          },
	          {
	            "type": "button",
	            "flex": 1,
	            "style": "primary",
	            "color": "#aaaaaa",
	            "action": {
	              "type": "message",
	              "label": "終了",
	              "text": "終了"
	            }
	          }
	        ]
	      }
	    }
	  ]
	}`)
*/
	bot, err := linebot.New(
		//os.Getenv("LINE_CHANNEL_SECRET"),
		//os.Getenv("LINE_CHANNEL_TOKEN"),
		"016d6abfada205ab9c9de38a7b6518d8",
		"pmpxhqajHvsM+vLfTuStk75zWOPP84bD/trCrT7uWSGWHOc3CQi/39s1snUpPLKBB0oSW3mI3Iw9dReOWPcKlkZOa/2TzhU7JAKiYd6+oUFD2pcjhf+cCFVzykrSk3qnJm5z8SqJR/dUWcg+nUXqEgdB04t89/1O/w1cDnyilFU=",
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
						/*
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ジャンルを入力してください。")).Do(); err != nil {
							log.Print(err)
						}
						*/
						container := &linebot.BubbleContainer{
							Type: linebot.FlexContainerTypeBubble,
        					Body: &linebot.BoxComponent{
        						Type:   linebot.FlexComponentTypeBox,
					            Layout: linebot.FlexBoxLayoutTypeHorizontal,
					            Contents: []linebot.FlexComponent{
					                &linebot.TextComponent{
					                    Type: linebot.FlexComponentTypeText,
					                    Text: "ジャンルを選択してください",
					                },
					                &linebot.ButtonComponent{
					                    Type: linebot.FlexComponentTypeButton,
					                    Flex: 1,
					                    Style: linebot.FlexButtonStyleTypePrimary,
					                    Color: "#aaaaaa",
					                    Action: linebot.ActionTypeMessage {
					                    	Lable: "C",
					                    	Text: "C",
					                    },
					                },
					            },
					        },
					    }
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewFlexMessage("ジャンル", container)).Do(); err != nil {
							log.Print(err)
						}
						mode = "init_genre"
					case "init_genre":
						if message.Text == "終了" {
							var print_genre string
							for _, g := range my_genre {
								print_genre += g + "\n"
							}
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("学籍番号：　"+student_id+"\n名前：　"+name+"\nジャンル\n"+print_genre+"\nこちらで登録してよろしいですか。")).Do(); err != nil {
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
						}
					case "default":
						if message.Text == "検索" {
						    genre := modules.GetPost(my_genre)
						    var print_genre string
						    for _, g := range genre {
						    	print_genre += g.NAME + "\n"
							}
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(print_genre)).Do(); err != nil {
								log.Print(err)
						    }
							mode = "default"
						}
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

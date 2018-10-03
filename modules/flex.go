package modules

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func Confirm(student_id string, name string, genres []string) (container *linebot.BubbleContainer) {
	for i:=len(genres); i<5; i++ {
		genres = append(genres, " ")
	}
	container = &linebot.BubbleContainer{
        Type: linebot.FlexContainerTypeBubble,
        Body: &linebot.BoxComponent{
            Type:   linebot.FlexComponentTypeBox,
            Layout: linebot.FlexBoxLayoutTypeVertical,
            Contents: []linebot.FlexComponent{
                &linebot.TextComponent{
                    Type: linebot.FlexComponentTypeText,
                    Text: "登録内容",
                    Weight: linebot.FlexTextWeightTypeBold,
                    Color: "#1DB446",
                    Size: linebot.FlexTextSizeTypeXxl,
                    Margin: linebot.FlexComponentMarginTypeMd,
                },
                &linebot.SeparatorComponent{
                	Type: linebot.FlexComponentTypeSeparator,
                	Margin: linebot.FlexComponentMarginTypeXxl,
                },
                &linebot.BoxComponent{
                    Type:   linebot.FlexComponentTypeBox,
            		Layout: linebot.FlexBoxLayoutTypeVertical,
            		Margin: linebot.FlexComponentMarginTypeXxl,
            		Spacing: linebot.FlexComponentSpacingTypeSm,
            		Contents: []linebot.FlexComponent{
            			&linebot.BoxComponent{
		                    Type:   linebot.FlexComponentTypeBox,
		            		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		            		Contents: []linebot.FlexComponent{
		            			&linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: "学籍番号",
				                    Color: "#555555",
				                    Size: linebot.FlexTextSizeTypeMd,
				                    Flex: 0,
				                },
				                &linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: student_id,
				                    Color: "#111111",
				                    Size: linebot.FlexTextSizeTypeMd,
				                    Align: linebot.FlexComponentAlignTypeEnd,
				                },
				            },
		            	},
		            	&linebot.BoxComponent{
		                    Type:   linebot.FlexComponentTypeBox,
		            		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		            		Contents: []linebot.FlexComponent{
		            			&linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: "名前",
				                    Color: "#555555",
				                    Size: linebot.FlexTextSizeTypeMd,
				                    Flex: 0,
				                },
				                &linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: name,
				                    Color: "#111111",
				                    Size: linebot.FlexTextSizeTypeMd,
				                    Align: linebot.FlexComponentAlignTypeEnd,
				                },
				            },
		            	},
		            	&linebot.BoxComponent{
		                    Type:   linebot.FlexComponentTypeBox,
		            		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		            		Margin: linebot.FlexComponentMarginTypeXxl,
		            		Contents: []linebot.FlexComponent{
		            			&linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: "ジャンル",
				                    Color: "#555555",
				                    Size: linebot.FlexTextSizeTypeMd,
				                },
				                &linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: genres[0],
				                    Color: "#111111",
				                    Size: linebot.FlexTextSizeTypeMd,
				                    Align: linebot.FlexComponentAlignTypeEnd,
				                },
				            },
		            	},
		            	&linebot.BoxComponent{
		                    Type:   linebot.FlexComponentTypeBox,
		            		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		            		Contents: []linebot.FlexComponent{
		            			&linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: " ",
				                    Color: "#555555",
				                    Size: linebot.FlexTextSizeTypeMd,
				                },
				                &linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: genres[1],
				                    Color: "#111111",
				                    Size: linebot.FlexTextSizeTypeMd,
				                    Align: linebot.FlexComponentAlignTypeEnd,
				                },
				            },
		            	},
		            	&linebot.BoxComponent{
		                    Type:   linebot.FlexComponentTypeBox,
		            		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		            		Contents: []linebot.FlexComponent{
		            			&linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: " ",
				                    Color: "#555555",
				                    Size: linebot.FlexTextSizeTypeMd,
				                },
				                &linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: genres[2],
				                    Color: "#111111",
				                    Size: linebot.FlexTextSizeTypeMd,
				                    Align: linebot.FlexComponentAlignTypeEnd,
				                },
				            },
		            	},
		            	&linebot.BoxComponent{
		                    Type:   linebot.FlexComponentTypeBox,
		            		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		            		Contents: []linebot.FlexComponent{
		            			&linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: " ",
				                    Color: "#555555",
				                    Size: linebot.FlexTextSizeTypeMd,
				                },
				                &linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: genres[3],
				                    Color: "#111111",
				                    Size: linebot.FlexTextSizeTypeMd,
				                    Align: linebot.FlexComponentAlignTypeEnd,
				                },
				            },
		            	},
		            	&linebot.BoxComponent{
		                    Type:   linebot.FlexComponentTypeBox,
		            		Layout: linebot.FlexBoxLayoutTypeHorizontal,
		            		Contents: []linebot.FlexComponent{
		            			&linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: " ",
				                    Color: "#555555",
				                    Size: linebot.FlexTextSizeTypeMd,
				                },
				                &linebot.TextComponent{
				                    Type: linebot.FlexComponentTypeText,
				                    Text: genres[4],
				                    Color: "#111111",
				                    Size: linebot.FlexTextSizeTypeMd,
				                    Align: linebot.FlexComponentAlignTypeEnd,
				                },
				            },
		            	},
            		},
                },
                &linebot.SeparatorComponent{
                	Type: linebot.FlexComponentTypeSeparator,
                	Margin: linebot.FlexComponentMarginTypeXxl,
                },
                &linebot.TextComponent{
                    Type: linebot.FlexComponentTypeText,
                    Text: "こちらの内容でよろしいですか？",
                    Color: "#aaaaaa",
                    Size: linebot.FlexTextSizeTypeSm,
                    Margin: linebot.FlexComponentMarginTypeMd,
                    Align: linebot.FlexComponentAlignTypeCenter,
                },
                &linebot.BoxComponent{
                    Type:   linebot.FlexComponentTypeBox,
            		Layout: linebot.FlexBoxLayoutTypeHorizontal,
            		Margin: linebot.FlexComponentMarginTypeMd,
            		Contents: []linebot.FlexComponent{
            			&linebot.ButtonComponent{
            				Type: linebot.FlexComponentTypeButton,
            				Flex: 1,
            				Style: linebot.FlexButtonStyleTypePrimary,
            				Action: &linebot.MessageAction {
            					Label: "はい",
            					Text: "はい",
            				},
            			},
            			&linebot.ButtonComponent{
            				Type: linebot.FlexComponentTypeButton,
            				Flex: 1,
            				Style: linebot.FlexButtonStyleTypePrimary,
            				Color: "#aaaaaa",
            				Action: &linebot.MessageAction {
            					Label: "いいえ",
            					Text: "いいえ",
            				},
            			},
            		},
            	},
            },
        },
    }
    return
}
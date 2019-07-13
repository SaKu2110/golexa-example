package main

import(
	"github.com/gin-gonic/gin"
	"github.com/SaKu2110/golexa"
)

func main(){
	gin := gin.Default()
	// golexa 初期化
	golexa := golexa.Default()
	golexa.SetIntent(intent)

	gin.POST("/", golexa.Handler)
	gin.Run()
}

func intent (c *golexa.Context) {
	switch c.Intent{
	case "LaunchRequest":
		c.Ask("テストスキルが起動しました。")
	case "PingIntent":
		c.Ask("ポング！")
	case "EchoColorIntent":
		c.Ask(c.Slots("color")+"ですね。")
	case "SetScheduleIntent":
		c.Ask(c.Slots("firstname")+"さんと"+c.Slots("city")+"に行くのですね。行ってらっしゃい。")
	default:
		c.Tell("すみません。よくわかりません。")
	}
}

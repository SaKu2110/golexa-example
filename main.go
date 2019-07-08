package main

import(
	"github.com/gin-gonic/gin"
	"github.com/SaKu2110/golexa"
)

func main(){
	gin := gin.Default()
	golexa := golexa.Default(intent)
	gin.POST("/", golexa.CallHandler)
	gin.Run()
}

func intent (request string) string {
	switch request {
	case "LaunchRequest":
		// リクエストに対する処理
		return golexa.Ask("起動しました。")
	default:
		return golexa.Tell(request+"が呼び出されました。")
	}
}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/basicGin/logic"
	"log"
)

type TelegramController struct{}

// 结构体必须大写 否则找不到
type TelegramRequestBody struct {
	Urls  []string `json:"urls"`
	Proxy string   `json:"proxy"`
}
type TelegramResponseBody struct {
	Output string `json:"output"`
	Msg    string `json:"msg"`
}

/*
curl -X POST http://127.0.0.1:8192/api/v1/telegram/download/telegram \
-H "Content-Type: application/json" \

	-d '{
	    "uri": "http://example.com/resource",
	    "proxy": "http://proxy.example.com:8080"
	}'
*/
func (t TelegramController) DownloadAll(ctx *gin.Context) {
	var requestBody TelegramRequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("url = %s \n proxy = %s\n", requestBody.Urls, requestBody.Proxy)
	var rep TelegramResponseBody
	log.Println("开始下载")
	go logic.Downloads(requestBody.Urls, requestBody.Proxy)
	rep.Msg = "已经开始下载"
	ctx.JSON(200, rep)
}

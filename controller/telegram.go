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
	Uri   string `json:"url"`
	Proxy string `json:"proxy"`
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
func (t TelegramController) PostTelegram(ctx *gin.Context) {

	var requestBody TelegramRequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("url = %s \n proxy = %s\n", requestBody.Uri, requestBody.Proxy)
	var rep TelegramResponseBody
	log.Println("开始下载")
	output, err := logic.Download(requestBody.Uri, requestBody.Proxy)
	if err != nil {
		rep.Msg = err.Error()
	}
	rep.Output = output
	ctx.JSON(200, rep)
}

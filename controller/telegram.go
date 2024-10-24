package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/basicGin/logic"
	"log"
)

type TelegramController struct{}

// 结构体必须大写 否则找不到
type Telegram struct {
	URLs  []string `json:"urls" binding:"required"`
	Proxy string   `json:"proxy" binding:"required"`
}
type TelegramResponseBody struct {
	URLs []string `json:"urls"`
	Msg  string   `json:"msg"`
}

/*
curl --location --request POST 'http://127.0.0.1:8193/api/v1/telegram/download' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \

	--data-raw '{
	    "urls": [
	        "string"
	    ],
	    "proxy": "string"
	}'
*/
func (t TelegramController) DownloadAll(ctx *gin.Context) {
	req := new(Telegram)
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//fmt.Printf("url = %s \nproxy = %s\n", req.URLs, req.Proxy)
	var rep TelegramResponseBody
	log.Println("开始下载")
	rep.URLs = req.URLs
	go logic.Downloads(req.URLs, req.Proxy)
	rep.Msg = "已经开始下载"
	ctx.JSON(200, rep)
}

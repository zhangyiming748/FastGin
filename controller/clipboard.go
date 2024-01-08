package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangyiming748/basicGin/logic"
)

type ClipboardController struct{}

/*
curl --location --request POST 'http://127.0.0.1:8192/api/v1/clipboard/upload' \
--header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from": "string",
    "msg": "string"
}'
*/

func (c ClipboardController) Upload(ctx *gin.Context) {
	var cp logic.Clipboard
	if err := ctx.BindJSON(&cp); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 处理用户数据
	logic.ClipBoard(cp)
	ctx.JSON(200, gin.H{
		"message": "成功接收JSON数据",
		"data":    cp,
	})
}

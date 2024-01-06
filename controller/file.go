package controller

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type FileController struct{}

func (f FileController) Upload(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "文件上传失败： %v", err)
		return
	}

	filename := header.Filename
	// 保存到当前目录下
	dst := filename
	//dst := "/path/to/save/" + filename

	out, err := os.Create(dst)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "创建文件失败： %v", err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "保存文件失败： %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "文件上传成功",
	})
}

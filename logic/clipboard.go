package logic

import (
	"log/slog"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type Clipboard struct {
	From string `json:"from"`
	Msg  string `json:"msg"`
}

func ClipBoard(cp Clipboard) {
	filename := strings.Join([]string{GetRoot(), string(os.PathSeparator), time.Now().Format("2006-01-02 15:04:05.000000"), ".txt"}, "")
	slog.Info("剪贴板任务", slog.String("保存的文件位置", filename), slog.Any("保存的文件内容", cp))
	storage, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		slog.Warn("保存剪贴板失败", slog.String("错误信息", err.Error()))
		return
	}
	defer storage.Close()
	storage.WriteString(cp.Msg)
}
func GetRoot() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}

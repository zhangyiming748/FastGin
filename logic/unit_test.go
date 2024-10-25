package logic

import (
	"fmt"
	"github.com/zhangyiming748/basicGin/util"
	"os"
	"testing"
)

func init() {
	util.SetLog("telegram.log")
}
func TestPATH(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("无法获取用户的个人文件夹目录:", err)

	}
	t.Log(home)
}

func TestDownloadsTelegram(t *testing.T) {
	urls := util.ReadByLine("post.link")
	proxy := "http://127.0.0.1:8889"
	Downloads(urls, proxy)
}

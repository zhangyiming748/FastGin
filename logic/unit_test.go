package logic

import (
	"fmt"
	"os"
	"testing"
)

func TestPATH(t *testing.T) {
	// 获取环境变量 PROFILE
	profile := os.Getenv("userprofile")

	// 检查环境变量是否存在
	if profile == "" {
		fmt.Println("环境变量 PROFILE 不存在")
	} else {
		fmt.Printf("PROFILE 环境变量的值: %s\n", profile)
	}
}

func TestDownloadTelegram(t *testing.T) {
	uri := "https://t.me/woshadiao/165436"
	proxy := "http://127.0.0.1:8889"
	download, err := Download(uri, proxy)
	if err != nil {
		return
	} else {
		t.Logf("%v\n", download)
	}
}

func TestDownloadsTelegram(t *testing.T) {
	urls := []string{"https://t.me/woshadiao/165436", "https://t.me/cosplayandanime/9190", "https://t.me/cosplayandanime/9191", "https://t.me/cosplayandanime/9192", "https://t.me/cosplayandanime/9191", "https://t.me/cosplayandanime/9193", "https://t.me/cosplayandanime/9194", "https://t.me/cosplayandanime/9195", "https://t.me/cosplayandanime/9196", "https://t.me/cosplayandanime/9197", "https://t.me/cosplayandanime/9198"}
	proxy := "http://127.0.0.1:8889"
	Downloads(urls, proxy)
}

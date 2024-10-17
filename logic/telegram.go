package logic

import (
	"fmt"
	"github.com/zhangyiming748/basicGin/util"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Downloads(urls []string, proxy string) {
	var status string
	var count int
	defer func() {
		fmt.Sprintf("全部下载结束,失败 %d / %d 个\n", count, len(urls))
	}()
	f, err := os.OpenFile("failed.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, url := range urls {
		output, fail := Download(url, proxy)
		if fail != nil {
			count++
			out := fmt.Sprintf("%s download fail ,err:%s:%v", url, output, fail)
			f.WriteString(out)
		}
	}
	f.Sync()
}
func Download(uri, proxy string) (string, error) {
	var status string
	defer func() {
		log.Println(status)
	}()
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("无法获取用户的个人文件夹目录:", err)
		return "", err
	}
	dir := filepath.Join(home, "Downloads")
	fmt.Printf("用户的个人文件夹目录: %s\n", home)
	fmt.Printf("用户的下载文件夹目录: %s\n", dir)
	target := filepath.Join(dir, "telegram")
	os.MkdirAll(target, 0755)
	tdl := util.WindowsTelegramLocation
	cmd := exec.Command(tdl, "download", "--proxy", proxy, "--url", uri, "--dir", target)
	fmt.Println(cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("下载命令执行出错", string(output))
		status = strings.Join([]string{status, "下载失败"}, "")
		return string(output), err
	} else {
		status = strings.Join([]string{status, "下载成功"}, "")
		return string(output), nil
	}
}

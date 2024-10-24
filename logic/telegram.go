package logic

import (
	"fmt"
	"github.com/zhangyiming748/basicGin/util"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func Downloads(urls []string, proxy string) {
	var status string
	var count int
	defer func() {
		status = fmt.Sprintf("全部下载结束,失败 %d / %d 个\n", count, len(urls))
		log.Println(status)
	}()
	f, err := os.OpenFile("failed.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, url := range urls {
		if strings.Contains(url, " ") {
			base := strings.Split(url, " ")[0]                  //https://t.me/acgr18/34406
			step, _ := strconv.Atoi(strings.Split(url, " ")[1]) //9
			//https://t.me/acgr18/34406 9
			prefix, suffix, _ := Split(base)
			for i := 0; i < step; i++ {
				uri := strings.Join([]string{prefix, strconv.Itoa(suffix + i)}, "/")
				log.Printf("Downloading %s to %s\n", uri, proxy)
				fail := Download(uri, proxy)
				if fail != nil {
					count++
					out := fmt.Sprintf("download fail :%s\n", url)
					f.WriteString(out)
				}
			}

		} else {
			fail := Download(url, proxy)
			if fail != nil {
				count++
				out := fmt.Sprintf("download fail :%s\n", url)
				f.WriteString(out)
			}
		}

	}
	f.Sync()
}

// https://github.com/iyear/tdl.git
func Download(uri, proxy string) error {
	var status string
	defer func() {
		log.Println(status)
	}()
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("无法获取用户的个人文件夹目录:", err)
		return err
	}
	dir := filepath.Join(home, "Downloads")
	fmt.Printf("用户的个人文件夹目录: %s\n", home)
	fmt.Printf("用户的下载文件夹目录: %s\n", dir)
	target := filepath.Join(dir, "telegram")
	os.MkdirAll(target, 0755)
	tdl := util.WindowsTelegramLocation
	if runtime.GOOS == "linux" {
		tdl = util.LinuxTelegramLocation
	}
	cmd := exec.Command(tdl, "download", "--proxy", proxy, "--url", uri, "--dir", target)
	fmt.Println(cmd.String())
	err = util.ExecCommand(cmd)
	if err != nil {
		log.Println("下载命令执行出错", uri)
		status = strings.Join([]string{status, "下载失败"}, "")
		return err
	} else {
		status = strings.Join([]string{status, "下载成功"}, "")
		return nil
	}
}

func Split(s string) (prefix string, suffix int, err error) {
	lastSlashIndex := strings.LastIndex(s, "/")

	if lastSlashIndex != -1 {
		// 分割字符串
		beforeLastSlash := s[:lastSlashIndex]
		afterLastSlash, _ := strconv.Atoi(s[lastSlashIndex+1:])

		return beforeLastSlash, afterLastSlash, nil
	} else {
		return "", -1, err
	}
}

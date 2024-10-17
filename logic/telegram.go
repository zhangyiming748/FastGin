package logic

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Download(uri, proxy string) (string, error) {
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
	tdl := "C:\\Users\\zen\\Downloads\\telegram\\tdl_Windows_64bit\\tdl.exe"
	cmd := exec.Command(tdl, "download", "--proxy", proxy, "--url", uri, "--dir", target)
	fmt.Println(cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("下载命令执行出错", string(output))
		return string(output), err
	} else {
		return string(output), nil
	}
}

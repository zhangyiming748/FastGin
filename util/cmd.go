package util

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

/*
执行命令过程中可以循环打印消息
*/
func ExecCommand(c *exec.Cmd) (key string, e error) {

	log.Printf("开始执行命令:%v\n", c.String())
	stdout, err := c.StdoutPipe()
	c.Stderr = c.Stdout
	if err != nil {
		log.Printf("连接Stdout产生错误:%v\n", err)
		return "", err
	}
	if err = c.Start(); err != nil {
		log.Printf("启动cmd命令产生错误:%v\n", err)
		return "", err
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		t = strings.Replace(t, "\u0000", "", -1)
		if key == "" {
			if n := GetKey(t); n != "" {
				key = n
				log.Printf("捕获到关键词:%v\n", key)
			}
		}
		fmt.Printf("\r%v", t)
		if err != nil {
			break
		}
	}
	if err = c.Wait(); err != nil {
		log.Printf("命令执行中产生错误:%v\n", err)
		return "", err
	}
	return key, nil
}

func GetPercentageSign(s string) int {
	if strings.HasSuffix(s, "]") {
		i, _ := getNumberBeforePercent(s)
		return i
	}
	return -1
}
func getNumberBeforePercent(s string) (int, error) {
	// 使用正则表达式匹配百分号前的数字
	re := regexp.MustCompile(`([0-9]+(?:\.[0-9]+)?)%`)
	matches := re.FindStringSubmatch(s)
	if len(matches) > 1 {
		// 将匹配到的数字转换为 float64
		number, err := strconv.ParseFloat(matches[1], 64)
		if err != nil {
			return 0, err
		}
		number = number * 10
		// 转换为 int
		return int(number), nil
	}

	return 0, fmt.Errorf("no percentage found in the string")
}
func GetKey(s string) string {
	//str := "(1249419900):6597 -> /h~ ... done! [184.88 MB in 43.347s; 4.26 MB/s]"

	// 正则表达式，匹配冒号后和箭头前的任意长度数字
	re := regexp.MustCompile(`:(\d+)\s+->`)

	// 查找匹配
	matches := re.FindStringSubmatch(s)
	if len(matches) > 1 {
		fmt.Println("匹配到的数字:", matches[1])
		return matches[1]
	} else {
		fmt.Println("没有匹配到数字")
		return ""
	}
}

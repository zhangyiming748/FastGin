package util

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

/*
执行命令过程中可以循环打印消息
*/
func ExecCommand(c *exec.Cmd) (e error) {
	log.Printf("开始执行命令:%v\n", c.String())
	stdout, err := c.StdoutPipe()
	c.Stderr = c.Stdout
	if err != nil {
		log.Printf("连接Stdout产生错误:%v\n", err)
		return err
	}
	if err = c.Start(); err != nil {
		log.Printf("启动cmd命令产生错误:%v\n", err)
		return err
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		t = strings.Replace(t, "\u0000", "", -1)
		fmt.Printf("\r%v", t)
		if err != nil {
			break
		}
	}
	if err = c.Wait(); err != nil {
		log.Printf("命令执行中产生错误:%v\n", err)
		return err
	}
	return nil
}

/*
执行命令过程中可以循环打印消息
*/
func ExecCommandWithBar(c *exec.Cmd) (e error) {
	log.Printf("开始执行命令:%v\n", c.String())
	bar := progressbar.New(1000)
	defer bar.Finish()
	stdout, err := c.StdoutPipe()
	c.Stderr = c.Stdout
	if err != nil {
		log.Printf("连接Stdout产生错误:%v\n", err)
		return err
	}
	if err = c.Start(); err != nil {
		log.Printf("启动cmd命令产生错误:%V\n", err)
		return err
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		bar.Set(GetPercentageSign(t))
		if err != nil {
			break
		}
	}
	if err = c.Wait(); err != nil {
		log.Printf("命令执行中产生错误:%v\n", err)
		return err
	}
	log.Printf("命令结束:%v\n", c.String())
	return nil
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

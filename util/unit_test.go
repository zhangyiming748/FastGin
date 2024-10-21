package util

import "testing"

func TestOffset(t *testing.T) {
	prefix := "https://t.me/Siwa2024/14935"
	GenerateUrl(1, 200, prefix)
}
func TestGetPercentageSign(t *testing.T) {
	s := "🔮 奇闻异录 与 沙雕时刻 meme collection~ ...  21.3% [........] [0 B in 297ms; 0 B/s]"
	ret := GetPercentageSign(s)
	t.Log(ret)
}

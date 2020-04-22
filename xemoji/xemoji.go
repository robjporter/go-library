package xemoji

import (
	"fmt"
)

func GetEmoji(code int) string {
	if code >-1 && code < len(emojiCodes) {
		return fmt.Sprintf("%s",emojiCodes[code])
	}
	return ""
}

func GetEmojis(codes []int) []string {
	ret := []string{}
	for a := 0; a < len(codes); a++ {
		if codes[a] >-1 && codes[a] < len(emojiCodes) {
			ret = append(ret,fmt.Sprintf("%s",emojiCodes[codes[a]]))
		}
	}
	return ret
}

func OutputAll() {
	for i := 0; i < len(emojiCodes); i++ {
		fmt.Println(emojiCodes[i])
	}
}
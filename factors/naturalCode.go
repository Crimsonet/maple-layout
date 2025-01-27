/**
 *	@Encode : UTF-8
 *	@Author : Crimsonet
 *	@Repo   : maple-layout
 *	@File   : naturalCode
 *	@Time   : 1/27/25
 *	@IDE    : GoLand
 */

package factors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

var naturalCodeMap = map[string]string{
	// 聲母映射
	"b": "b", "p": "p", "m": "m", "f": "f",
	"d": "d", "t": "t", "n": "n", "l": "l",
	"g": "g", "k": "k", "h": "h",
	"j": "j", "q": "q", "x": "x",
	"zh": "v", "ch": "i", "sh": "u",
	"r": "r", "z": "z", "c": "c", "s": "s",
	"y": "y", "w": "w",
	// 韻母映射
	"a": "a", "o": "o", "e": "e",
	"ai": "l", "ei": "z", "ao": "k", "ou": "b",
	"an": "j", "en": "f", "ang": "h", "eng": "g",
	"ong": "s", "er": "r", "i": "i",
	"ian": "m", "iao": "c", "in": "n", "ing": "y",
	"ua": "w", "uo": "w", "uai": "y", "ui": "v",
	"uan": "r", "un": "p", "uang": "d",
	"ue": "t", "u": "v",
}

func NaturalCode(pinyin string) string {
	if len(pinyin) == 1 {
		if code, ok := naturalCodeMap[pinyin]; ok {
			return code
		}
		return ""
	}
	initialList := []string{"zh", "ch", "sh", "b", "p", "m", "f", "d", "t", "n", "l", "g", "k", "h", "j", "q", "x", "r", "z", "c", "s", "y", "w"}
	var initial string
	for _, i := range initialList {
		if len(pinyin) >= len(i) && pinyin[:len(i)] == i {
			initial = i
			break
		}
	}
	if initial == "" {
		initial = "y"
		if pinyin[0] == 'w' {
			initial = "w"
		}
	}
	final := pinyin[len(initial):]
	if final == "" {
		final = string(pinyin[len(pinyin)-1])
	}
	initialCode, ok1 := naturalCodeMap[initial]
	finalCode, ok2 := naturalCodeMap[final]
	if ok1 && ok2 {
		return initialCode + finalCode
	}
	return ""
}

type ChineseResult struct {
	Number        int     `json:"number"`
	Frequency     float64 `json:"frequency"`
	Pronunciation string  `json:"pronunciation"`
}

func ProcessJSONFile(fileName string) map[rune]int {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("讀取文件 %s 時出錯: %v\n", fileName, err)
		return nil
	}
	var result struct {
		ChineseResult []ChineseResult `json:"chinese"`
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Printf("解析 JSON 數據時出錯: %v\n", err)
		return nil
	}
	letterCount := make(map[rune]int)
	for _, item := range result.ChineseResult {
		code := NaturalCode(item.Pronunciation)
		for _, char := range strings.ToLower(code) {
			letterCount[char] += int(math.Round(item.Frequency))
		}
	}
	return letterCount
}

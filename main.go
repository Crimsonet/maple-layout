package main

import (
	"fmt"
	"github.com/Crimsonet/maple-layout/editor"
	"github.com/Crimsonet/maple-layout/factors"
)

func main() {

	//如果想看自然碼原始輸出結果
	//或者想看加權平均數輸出結果
	//的話，這裏改一下就可以了。
	whetherPrintRawNaturalCode := false
	whetherPrintProNaturalCode := false

	if whetherPrintRawNaturalCode == true {
		fileName := "factors/chinese.json"
		letterCountRaw := factors.ProcessJSONFile(fileName)
		if letterCountRaw != nil {
			fmt.Println("自然碼數據統計結果:")
			for letter, count := range letterCountRaw {
				fmt.Printf("%c: %d\n", letter, count)
			}
		}
	}
	if whetherPrintProNaturalCode == true {
		chineseFilePath := "./editor/stats/chinese.json"
		englishFilePath := "./editor/stats/english.json"
		editor.Processor(chineseFilePath, englishFilePath)
	}

}

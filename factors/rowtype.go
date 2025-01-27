/**
 *	@Encode : UTF-8
 *	@Author : Crimsonet
 *	@Repo   : maple-layout
 *	@File   : rowtype
 *	@Time   : 1/27/25
 *	@IDE    : GoLand
 */

package factors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func RowType(filename string) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("讀取文件時出錯: %v\n", err)
		return
	}
	var result struct {
		ChineseResult []ChineseResult `json:"chinese_result"`
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Printf("解析 JSON 數據時出錯: %v\n", err)
		return
	}
	codeCount := make(map[string]int)
	for _, item := range result.ChineseResult {
		code := NaturalCode(item.Pronunciation)
		if code != "" {
			codeCount[code]++
		}
	}
	fmt.Println("自然碼鍵位組統計結果：")
	for code, count := range codeCount {
		fmt.Printf("%s: %d\n", code, count)
	}
	//如果某個自然碼組合的count計數
	//大於等於七，那麼結匯被計入json
	//位於./factors/row/*.json下
	//我默認這裏的組合是常用自然碼
	//七是因爲 70% 線正好在七的位置
}

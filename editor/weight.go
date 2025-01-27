/**
 *	@Encode : UTF-8
 *	@Author : Crimsonet
 *	@Repo   : maple-layout
 *	@File   : weight
 *	@Time   : 1/27/25
 *	@IDE    : GoLand
 */

package editor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readJSONFile(filePath string) (map[string]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var frequencyMap map[string]float64
	err = json.Unmarshal(data, &frequencyMap)
	if err != nil {
		return nil, err
	}

	return frequencyMap, nil
}

func calculateWeightedAverage(chineseFreq, englishFreq map[string]float64, chineseWeight, englishWeight float64) map[string]float64 {
	result := make(map[string]float64)
	allLetters := make(map[string]bool)
	for letter := range chineseFreq {
		allLetters[letter] = true
	}
	for letter := range englishFreq {
		allLetters[letter] = true
	}
	for letter := range allLetters {
		chineseVal, chineseExists := chineseFreq[letter]
		englishVal, englishExists := englishFreq[letter]

		if !chineseExists {
			chineseVal = 0
		}
		if !englishExists {
			englishVal = 0
		}

		result[letter] = (chineseVal*chineseWeight + englishVal*englishWeight) / (chineseWeight + englishWeight)
	}

	return result
}

func Processor(chinese string, english string) {
	chineseFilePath := chinese
	englishFilePath := english
	chineseFreq, err := readJSONFile(chineseFilePath)
	if err != nil {
		fmt.Printf("讀取漢語 JSON 時出錯: %v\n", err)
		return
	}
	englishFreq, err := readJSONFile(englishFilePath)
	if err != nil {
		fmt.Printf("讀取英語 JSON 時出錯: %v\n", err)
		return
	}

	//這裏的默認權值都是1
	//此權值會影響最終的鍵位計算結果
	chineseWeight := 1.0
	englishWeight := 1.0

	weightedAverage := calculateWeightedAverage(chineseFreq, englishFreq, chineseWeight, englishWeight)

	rows := [][]string{
		{"z", "x", "c", "v", "b", "n", "m"},
		{"a", "s", "d", "f", "g", "h", "j", "k", "l"},
		{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"},
		//如果在此上下亂改順序會導致R1-R3順序反掉
	}
	fmt.Println("加權平均結果:")
	for i, row := range rows {
		fmt.Printf("R%d: ", i+1)
		for _, letter := range row {
			if freq, exists := weightedAverage[letter]; exists {
				fmt.Println("")
				fmt.Printf("%s: %.2f ", letter, freq)
			}
		}
		fmt.Println()
	}
}

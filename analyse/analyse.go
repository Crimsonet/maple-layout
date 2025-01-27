/**
 *	@Encode : UTF-8
 *	@Author : Crimsonet
 *	@Repo   : maple-layout
 *	@File   : analyse
 *	@Time   : 1/27/25
 *	@IDE    : GoLand
 */

package analyse

import "github.com/Crimsonet/maple-layout/analyse/level"

// 1，鍵位組 JSON 數據在 ./factors/row/  下，英漢合併;
// 2，單鍵位 JSON 數據在 ./editor/stats/ 下，英漢隔離;
// 3，單鍵位 JSON 數據在 ./factors/      下，英漢合併;
// 默認的標號順序是每個鍵排從左往右，
// 暫時使用第 1，3 號數據進行計算。

func Analysis() {
	level.ROne()
	level.RTwo()
	level.RThree()
}

/*

這種東西還是手算才快(樂)。

"R1": {      5/7
	"v": 7.56, => H
	"n": 5.12, => H
	"m": 3.60, => H
	"b": 3.21, => H
	"c": 2.51, => H
	"z": 1.99, => L
	"x": 1.53  => L
},

"R2": {      7/9
	"h": 5.92, => H
	"a": 5.25, => H
	"s": 5.23, => H
	"l": 4.50, => H
	"j": 4.18, => H
	"d": 4.04, => H
	"g": 3.67, => H
	"f": 3.26, => L
	"k": 2.39  => L
},

"R3": {      6/10
	"i": 9.83, => H
	"y": 6.15, => H
	"t": 6.04, => H
	"e": 7.12, => H
	"r": 4.62, => H
	"u": 3.76, => H
	"o": 4.13, => L
	"w": 3.07, => L
	"p": 1.99, => L
	"q": 1.05  => L
}

*/

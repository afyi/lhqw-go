package core

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var Keys [14]string = [14]string{"连翘", "银花", "麻黄", "苦杏仁", "石膏", "板蓝根", "贯众", "鱼腥草", "藿香", "大黄", "红景天", "薄荷脑", "甘草", "淀粉"}

// 编码
func Encode(str string) string {
	var ret []string
	var re = regexp.MustCompile(`[A-Za-z0-9\-\_\.\!\~\*\'\(\)]`)
	for _, v := range str {
		if ok := re.MatchString(string(v)); ok {
			// 如果在ascii范围内，则直接取Ascii值
			ret = append(ret, fmt.Sprintf("%x", v))
		} else {
			// 如果不在，则取urlencode后的值,
			ret = append(ret, url.QueryEscape(string(v)))
		}
	}
	return hex2duo(strings.ToUpper(strings.Replace(strings.Join(ret, ""), "%", "", -1)))
}

// 转成核心价值观
func hex2duo(str string) string {
	var duo []string
	for _, v := range str {
		n, _ := strconv.ParseInt(string(v), 16, 64)
		if n < 13 {
			// 此处因为是16进制，所以最大值其实是15，只要是0-13，都是1位，可以直接从表中解析，把最后一位的值拿来当高位
			duo = append(duo, Keys[int(n)])
		} else {
			// 此处如果是14或者15，则做一个进位算法，先固定一种算法，让上一位为最后一个值，下一位为 n-2， 这样，2位数高位是第14位，也就是大黄，低位为甘草和淀粉，密码会比较短
			duo = append(duo, Keys[13], Keys[n-13])
		}
	}
	return strings.Join(duo, "")
}

func Decode(str string) string {
	//var splited []string
	// 把字符串转成rune
	var hex []int
	for i, v := range Keys {
		// 这里从字符串中返回加密串的数据，先找一次，如果没了，就说明整个串里都没这个数字
		ind := strings.Index(str, v)
		for ind > -1 {
			end := ind + len(v)
			if ind+len(v) > len(str) {
				end = len(str)
			}
			str = fmt.Sprintf("%s,%d,%s", str[:ind], i, str[end:])
			ind = strings.Index(str, v)
		}
	}

	numList := strings.Split(strings.Trim(strings.ReplaceAll(str, ",,", ","), ","), ",")
	for _, v := range numList {
		res, _ := strconv.Atoi(v)
		hex = append(hex, res)
	}
	return duo2hex(hex)
}

func duo2hex(hex []int) string {
	var ret []string
	var num int // 用于判定当前循环次数
	for i := 0; i < len(hex); i++ {
		// 这是个临时值
		tmp := 0
		if hex[i] == 13 {
			// 如果当前值=13，那么说明这个是个2位数，要么14，要么15，所以再取第后一位上的数字为当前值
			i++
			tmp = hex[i] + 13
		} else {
			tmp = hex[i]
		}
		if num&1 == 0 {
			ret = append(ret, fmt.Sprintf("%%%X", tmp))
		} else {
			ret = append(ret, fmt.Sprintf("%X", tmp))
		}
		num++
	}

	res, _ := url.QueryUnescape(strings.Join(ret, ""))
	return res
}

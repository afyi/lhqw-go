package core

import (
	"testing"
)

func TestEncode(t *testing.T) {
	got := Encode("你好,lhqw")
	want := "淀粉银花石膏大黄淀粉连翘藿香连翘淀粉银花板蓝根藿香板蓝根大黄淀粉连翘麻黄甘草贯众甘草贯众红景天鱼腥草银花鱼腥草鱼腥草"
	if got != want {
		t.Errorf("实际结果:%v, 期望值:%v", got, want) // 测试失败输出错误提示
	}
}

func TestDecode(t *testing.T) {
	got := Decode("淀粉银花石膏大黄淀粉连翘藿香连翘淀粉银花板蓝根藿香板蓝根大黄淀粉连翘麻黄甘草贯众甘草贯众红景天鱼腥草银花鱼腥草鱼腥草")
	want := "你好,lhqw"
	if got != want {
		t.Errorf("实际结果:%v, 期望值:%v", got, want) // 测试失败输出错误提示
	}
}

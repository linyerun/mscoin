package tool

import (
	"fmt"
	"testing"
)

func TestRand4Num(t *testing.T) {
	for i := 0; i < 10000; i++ {
		num := RandCode(6)
		if len(num) != 6 {
			fmt.Println("存在长度不为4的随机数")
			return
		}
	}
	fmt.Println("测试成功")
}

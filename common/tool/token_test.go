package tool

import (
	"fmt"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	sk := "111"
	token, err := GenerateToken(sk, time.Now().Unix(), 10000, 100)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("生成Token成功, token:", token)

	userId, err := ParseToken(token, sk)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("解析Token成功, userId:", userId)
}

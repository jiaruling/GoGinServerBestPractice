package utils

import "math/rand"

/*
   功能说明: 随机生成指定长度字符串
   参考:
   创建人: 贾汝凌
   创建时间: 2022/1.sql/19 15:42
*/

var letters = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

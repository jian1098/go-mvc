package common

import (
	"crypto/md5"
	"encoding/hex"
)

/*
*
md5 加密
*/
func MD5(input string) string {
	hasher := md5.New()                        // 创建一个新的MD5哈希器
	hasher.Write([]byte(input))                // 将输入字符串转换为字节切片并写入哈希器
	return hex.EncodeToString(hasher.Sum(nil)) // 返回十六进制格式的哈希值
}

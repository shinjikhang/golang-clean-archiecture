package helper

import (
	"encoding/hex"
	"math/rand"
)

// Hàm để tạo chuỗi ngẫu nhiên với chiều dài cho trước
func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

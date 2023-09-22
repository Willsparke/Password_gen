package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    passwordLength := 12 // 密碼長度
    password := generateRandomPassword(passwordLength)
    fmt.Println("隨機密碼：", password)
}

func generateRandomPassword(length int) string {
    var charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
    password := make([]byte, length)

    for i := range password {
        password[i] = charset[rand.Intn(len(charset))]
    }

    return string(password)
}

package main

import (
	"crypto/md5"
	"fmt"
	"os"

	"anime-skip.com/backend/internal/utils/auth"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("Password must be passed in as an argument")
		os.Exit(1)
	}
	password := os.Args[1]
	fmt.Println("Password: " + password)

	md5 := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	fmt.Println("md5:      " + md5)

	bcrypt, err := auth.GenerateEncryptedPassword(md5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("bcrypt:   " + bcrypt)
}

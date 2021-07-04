package main

import (
	"fmt"
	"math/rand"
	"time"

	"anime-skip.com/backend/internal/utils"
)

func main() {
	source := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println("Client ID:", utils.ReallyRandomString(*source, 32))
}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"anime-skip.com/backend/internal/utils"
)

func main() {
	// Generate Client ID
	source := rand.New(rand.NewSource(time.Now().Unix()))
	id := utils.ReallyRandomString(*source, 32)
	fmt.Println("Client ID:", id)

	// Get additional info
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter an app name: ")
	appName, _ := reader.ReadString('\n')
	appName = strings.TrimSpace(appName)
	fmt.Print("\nEnter a description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	fmt.Print("\nEnter the requesting user's username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	fmt.Printf("\nSELECT id FROM users WHERE username = '%v';", username)
	fmt.Print("\nEnter the requesting user's ID: ")
	userId, _ := reader.ReadString('\n')
	userId = strings.TrimSpace(userId)

	fmt.Printf(`
Run and COMMIT; the following query:

  BEGIN TRANSACTION; INSERT INTO api_clients (
    id, created_at, created_by_user_id, updated_at, updated_by_user_id, user_id, app_name, description
  ) VALUES (
    '%s' NOW, '389a3749-c8f4-4e39-bf5f-96c1c2452074', NOW, '389a3749-c8f4-4e39-bf5f-96c1c2452074', '%s', '%s', '%s'
  );
`, id, userId, appName, description)
	fmt.Println()
}

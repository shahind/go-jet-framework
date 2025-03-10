package cli

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/shahind/go-jet-framework/register"
	"github.com/shahind/go-jet-framework/tool"
)

// GenerateKey will generate Go-Jet application key in main config.yml file
type GenerateKey struct {
	register.Command
}

// Register this command
func (c *GenerateKey) Register() {
	c.Signature = "generate:key"               // Change command signature
	c.Description = "Generate application key" // Change command description
}

// Run this command
func (c *GenerateKey) Run() {
	fmt.Println("Generating new application KEY...")
	path := tool.GetDynamicPath("config/server.go")
	read, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	appKey, err := generateNewToken()
	if err != nil {
		log.Fatal(err)
	}

	newContent := strings.Replace(string(read), "REPLACE_WITH_CUSTOM_APP_KEY", appKey, -1)

	if err = os.WriteFile(path, []byte(newContent), 0); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Complete!")
}

// generateNewToken will return a random sha256 hash
func generateNewToken() (string, error) {
	data := make([]byte, 10)
	if _, err := rand.Read(data); err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	hashStr := fmt.Sprintf("%x", hash[:])

	return hashStr, nil
}

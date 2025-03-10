package cli

import (
	"fmt"
	"github.com/shahind/go-jet-framework/register"
	"log"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
)

// GenerateKey will generate Go-Jet application key in main config.yml file
type ServiceCreate struct {
	register.Command
}

// Register this command
func (c *ServiceCreate) Register() {
	c.Signature = "service:create [service-name]" // Change command signature
	c.Description = "Create new Go-Jet service"   // Change command description
}

// Run this command
func (c *ServiceCreate) Run() {
	if len(c.Args) == 0 {
		c.Help()
		return
	}

	fmt.Printf("Creating service %s...\n", c.Args)
	if err := c.clone(c.Args); err != nil {
		log.Fatalf("Error: %s", err)
	}

	if err := c.reset_git(); err != nil {
		log.Fatalf("Error: %s", err)
	}

	if err := c.update(); err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Println("Service created successfully!")
}

// Help will show help for this command
func (c *ServiceCreate) Help() {
	log.Println("Usage: create-service [service-name]")
}

// Clones Go-Jet repository in destination folder
func (c *ServiceCreate) clone(destination string) error {
	_, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:      "https://github.com/shahind/go-jet.git",
		Progress: nil,
	})

	return err
}

// Reset git repository
func (c *ServiceCreate) reset_git() error {
	path := fmt.Sprintf("%s/.git", c.Args)
	if err := os.RemoveAll(path); err != nil {
		return err
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = c.Args
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// Updates Go-Jet Framework to the latest minor version
func (c *ServiceCreate) update() error {
	cmd := exec.Command("go", "get", "-u", "github.com/shahind/go-jet-framework")
	cmd.Dir = c.Args

	return cmd.Run()
}

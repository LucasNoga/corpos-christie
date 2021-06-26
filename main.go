package main

import (
	"corpos-christie/colors"
	"corpos-christie/config"
	"corpos-christie/core"
	"corpos-christie/user"
	"corpos-christie/utils"
	"fmt"
	"log"
	"os"
)

// Configuration of the application
var cfg *config.Config

// Start tax calculator from input user
func start(cfg *config.Config, user *user.User) bool {
	// Ask income's user
	_, err := user.AskIncome()
	if err != nil {
		log.Printf("Error: asking income for user, details: %v", err)
		return false
	}

	// Ask if user is in couple
	_, err = user.AskIsInCouple()
	if err != nil {
		log.Printf("Error: asking is in couple for user, details: %v", err)
		return false
	}

	// Ask if user hasChildren
	_, err = user.AskHasChildren()
	if err != nil {
		log.Printf("Error: asking has children, details: %v", err)
		return false
	}

	// calculate tax
	result := core.Process(user, cfg)

	user.Tax = result.Tax
	user.Remainder = result.Remainder

	// show user
	user.Show()

	return true
}

// Ask user if he wants to restart program
func askRestart() bool {
	for {
		fmt.Print("Would you want to enter a new income (Y/n): ")
		var input string = utils.ReadValue()
		if input == "Y" || input == "y" || input == "Yes" || input == "yes" {
			log.Printf("Restarting program...")
			return true
		} else {
			return false
		}
	}
}

// Init configuration file
func init() {
	cfg = new(config.Config)
	_, err := cfg.LoadConfiguration("./config.json")
	if err != nil {
		log.Printf(colors.Red("Unable to load config.json file, details: %v"), colors.Red(err))
		cfg.LoadDefaultConfiguration()
	}

	// get line and file log
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Starting program
func main() {
	log.Printf("Project: %v", colors.Yellow(cfg.Name))
	log.Printf("Version: %v", colors.Yellow(cfg.Version))

	// Init user
	var user *user.User = new(user.User)

	var keep bool
	for ok := true; ok; ok = keep {
		status := start(cfg, user)
		if status {
			log.Println("Core process successful")
		} else {
			log.Println("Core process failed")
		}
		fmt.Println("--------------------------------------------------------------")
		keep = askRestart()
	}

	log.Printf("Program exited...")
	os.Exit(0)
}

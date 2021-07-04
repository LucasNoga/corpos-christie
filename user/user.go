package user

import (
	"errors"
	"fmt"
	"log"

	"github.com/LucasNoga/corpos-christie/lib/colors"
	"github.com/LucasNoga/corpos-christie/lib/utils"
)

// Define a user
type User struct {
	Income     int     // Income (Revenu imposable) of the user
	Tax        float64 // Tax to pay for the user
	Remainder  float64 // Money remind after tax paid
	Parts      float64 // Parts of the user (calculate from isInCouple, childre)
	IsInCouple bool    // User is he in couple or not
	Children   int     // number of children of the user
}

// Ask income to the user if ok return true, else return false
func (user *User) AskIncome() (bool, error) {
	var input string = utils.ReadValue()
	income, err := utils.ConvertStringToInt(input)
	if err != nil {
		log.Printf("Error: Tax income is not convertible in int, details: %v", err)
		return false, err
	}
	user.Income = income
	return true, nil
}

// Ask remainder to the user if ok return true, else return false
func (user *User) AskRemainder() (bool, error) {
	var input string = utils.ReadValue()
	remainder, err := utils.ConvertStringToInt(input)
	if err != nil {
		log.Printf("Error: Tax remainder is not convertible in int, details: %v", err)
		return false, err
	}
	user.Remainder = float64(remainder)
	return true, nil
}

// Ask if the user is in couple, ok return true, else return false
func (user *User) AskIsInCouple() (bool, error) {
	response, err := askYesNo()
	if err != nil {
		return false, err
	}
	if response {
		user.IsInCouple = true
	} else {
		user.IsInCouple = false
	}
	return response, nil
}

// Ask if the user does have children and how many
func (user *User) AskHasChildren() (bool, error) {
	var input string = utils.ReadValue()

	// user can skip the question
	if input == "" {
		return true, nil
	}

	childrens, err := utils.ConvertStringToInt(input)
	if err != nil {
		log.Printf("Error: Childrens value is not convertible in int, details: %v", err)
		return false, err
	}

	user.Children = childrens
	return true, nil
}

// Ask if the user does have children and how many
func (user *User) AskTaxDetails() (bool, error) {
	fmt.Print("Do you want to see tax details (Y/n) ? ")
	response, err := askYesNo()
	if err != nil {
		return false, err
	}
	return response, nil
}

// Ask user if he wants to restart program
func (user *User) AskRestart() bool {
	response, _ := askYesNo()
	return response
}

// Get the parts of the user
func (user *User) GetParts() float64 {
	return user.Parts
}

// Show user fields
func (user *User) Show() {
	var isInCouple = "No"
	if user.IsInCouple {
		isInCouple = "Yes"
	}
	fmt.Printf("Income:\t\t%s €\n", colors.Red(user.Income))
	fmt.Printf("In couple:\t%s\n", colors.Red(isInCouple))
	fmt.Printf("Children:\t%s\n", colors.Red(user.Children))
	fmt.Printf("Parts:\t\t%s\n", colors.Red(user.Parts))
	fmt.Printf("Tax:\t\t%s €\n", colors.Green(user.Tax))
	fmt.Printf("Remainder:\t%s €\n", colors.Green(user.Remainder))
}

// askYesNo is a function to handle response user that should be yes or no
func askYesNo() (bool, error) {
	var input string = utils.ReadValue()
	if input == "Y" || input == "y" || input == "Yes" || input == "yes" {
		return true, nil
	} else if input == "" || input == "N" || input == "n" || input == "No" || input == "no" {
		return false, nil
	} else {
		return false, errors.New("invalid response you have to answer by (yes/Yes/Y/y or no/No/N/n)")
	}
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	err := CreateUser()
	fmt.Println(err)
	err = CreateOrg()
	fmt.Println(err)

}

func Connect() error {
	return errors.New("connection failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("error connecting: %w", err)
	}

	//...
	return nil
}

func CreateOrg() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("error creating org: %w", err)
	}

	//...
	return nil
}

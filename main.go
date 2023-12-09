package main

import (
	validator "aws-resource-inv-tool/validator"
	"fmt"
)

func main() {

	// User validor checker
	sts_output, err := validator.CheckCredetianls()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected as: ")
	fmt.Println(sts_output)
}

package main

import (
	sts "aws-resource-inv-tool/sts"
	"fmt"
)

func main() {

	// User validor checker
	sts_output, err := sts.CheckCredetianls()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected as: ")
	fmt.Println(sts_output)
}

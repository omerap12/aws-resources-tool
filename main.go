package main

import (
	ec2 "aws-resource-inv-tool/ec2"
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

	ec2.GetAllIntances()

}

package main

import (
	"aws-resource-inv-tool/ec2"
	"aws-resource-inv-tool/sts"
	"fmt"
)

func main() {

	// User validor checker
	sts_output, err := sts.CheckCredetianls()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected as: ", *sts_output.Arn)
	instances, err := ec2.GetInstancesByType("t2.micro")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(instances)
}

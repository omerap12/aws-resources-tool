package main

import (
	"aws-resource-inv-tool/ec2"
	"fmt"
)

func main() {

	// User validor checker
	// sts_output, err := sts.CheckCredetianls()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Connected as: ", *sts_output.Arn)
	query := make(ec2.RequestArgs)
	// query["instance-type"] = []string{"t2.micro"}

	instances, err := ec2.GetInstances(&query)
	if err != nil {
		fmt.Println(err.Error())
	}
 	ec2.PrintInstances(instances)
}

package ec2

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func CheckInstances() {
	return
}

func GetAllIntances() (*ec2.DescribeInstancesOutput, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	svc := ec2.New(sess)
	input := &ec2.DescribeInstancesInput{}
	result, err := svc.DescribeInstances(input)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return result, nil
}

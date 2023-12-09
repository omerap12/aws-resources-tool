package ec2

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type ec2Info struct {
    id string
    state string
}

func newEc2Info(id string, state string) *ec2Info {
	e := ec2Info{id:id,state: state}
	return &e
}

func CheckInstances() {
	return
}

func GetAllIntances() (*ec2.DescribeInstancesOutput, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	svc := ec2.New(sess)
	input := &ec2.DescribeInstancesInput{}
	result, err := svc.DescribeInstances(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetInstancesByType(instanceType string) ([]map[string]string, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	svc := ec2.New(sess)
	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-type"),
				Values: []*string{
					aws.String("t2.micro"),
				},
			},
		},
	}
	result, err := svc.DescribeInstances(input)
	if err != nil {
		return nil, err
	}

	instances_arr := []map[string]string{}
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			ele := map[string]string{
				"InstanceId": *instance.InstanceId,
				"State":      *instance.State.Name,
			}
			instances_arr = append(instances_arr, ele)
		}
	}
	return instances_arr, nil
}

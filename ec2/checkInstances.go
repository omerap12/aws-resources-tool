package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type ec2Info struct {
	id          string
	name        string
	state       string
	machineType string
}

type requestArgs map[string]interface{}

func newEc2Info(id string, name string, state string, machineType string) *ec2Info {
	e := ec2Info{id: id, name: name, state: state, machineType: machineType}
	return &e
}

func GetInstances(args requestArgs) ([]*ec2Info, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	// instanceType := "t2.micro"
	ec2Client := ec2.NewFromConfig(cfg)

	filters := ec2.Filter
	filters := []types.Filter{
		{
			Name: types.String("tag:Name"),
			Values: []string{
				"dev-*",
			},
		},
	}
	input := ec2.DescribeInstancesInput{Filters: filters}
	result, err := ec2Client.DescribeInstances(context.TODO(), &input)
	if err != nil {
		return nil, err
	}

	instancesList := make([]*ec2Info, 0)

	for _, resr := range result.Reservations {
		for _, instance := range resr.Instances {
			var name string

			for _, tag := range instance.Tags {
				if aws.ToString(tag.Key) == "Name" {
					name = aws.ToString(tag.Value)
					break
				}
			}

			state := string(instance.State.Name)
			machineType := string(instance.InstanceType)

			instancesList = append(instancesList, newEc2Info(
				aws.ToString(instance.InstanceId),
				name,
				state,
				machineType,
			))

			// fmt.Printf("%s, %s, %s, %s\n",
			// 	aws.ToString(instance.InstanceId),
			// 	name,
			// 	state,
			// 	machineType,
			// )
		}
	}
	for _, ins := range instancesList {
		fmt.Println(*ins)
	}

	return instancesList, nil
}

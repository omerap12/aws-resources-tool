package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type Ec2Info struct {
	id          string
	name        string
	state       string
	machineType string
}

type RequestArgs map[string]interface{}

func newEc2Info(id string, name string, state string, machineType string) *Ec2Info {
	e := Ec2Info{id: id, name: name, state: state, machineType: machineType}
	return &e
}

func GetInstances(args *RequestArgs) ([]*Ec2Info, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	// instanceType := "t2.micro"
	ec2Client := ec2.NewFromConfig(cfg)

	var instanceType []string
	if queryInstanceType, ok := (*args)["instance-type"]; ok && queryInstanceType != nil {
		instanceType = queryInstanceType.([]string)
	}

	var filters []types.Filter
	if len(instanceType) > 0 {
		filters = append(filters, types.Filter{
			Name:   aws.String("instance-type"),
			Values: instanceType,
		})
	}

	result, err := ec2Client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{
		Filters: filters,
	})
	if err != nil {
		return nil, err
	}

	instancesList := make([]*Ec2Info, 0)

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
		}
	}
	for _, ins := range instancesList {
		fmt.Println(*ins)
	}

	return instancesList, nil
}

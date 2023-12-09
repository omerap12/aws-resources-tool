package sts

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func CheckCredetianls() (*sts.GetCallerIdentityOutput, error) {
	aws_access_key_id := os.Getenv("AWS_ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")
	aws_region := os.Getenv("AWS_DEFAULT_REGION")
	if aws_access_key_id == "" {
		return nil, fmt.Errorf("Missing AWS_ACCESS_KEY_ID")
	}
	if aws_secret_access_key == "" {
		return nil, fmt.Errorf("Missing AWS_SECRET_ACCESS_KEY")
	}
	if aws_region == "" {
		return nil, fmt.Errorf("Missing AWS_DEFAULT_REGION")
	}

	// Create sesstion based on env variables
	sess, err := session.NewSession()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	svc := sts.New(sess)

	// calling sts get-caller-identity
	input := &sts.GetCallerIdentityInput{}
	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return result, nil
}

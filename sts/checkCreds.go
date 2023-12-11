package sts

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func CheckCredetianls() (*sts.GetCallerIdentityOutput, error) {
	aws_access_key_id := os.Getenv("AWS_ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("AWS_SECRET_ACCESS_KEY")
	aws_region := os.Getenv("AWS_REGION")
	if aws_access_key_id == "" {
		return nil, fmt.Errorf("Missing AWS_ACCESS_KEY_ID")
	}
	if aws_secret_access_key == "" {
		return nil, fmt.Errorf("Missing AWS_SECRET_ACCESS_KEY")
	}
	if aws_region == "" {
		return nil, fmt.Errorf("Missing AWS_DEFAULT_REGION")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	clientSts := sts.NewFromConfig(cfg)
	inputSts := &sts.GetCallerIdentityInput{}
	result, err := clientSts.GetCallerIdentity(context.TODO(), inputSts)

	if err != nil {
		return nil, err
	}
	return result, nil
}

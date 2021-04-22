package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	lambda.Start(cfn.LambdaWrap(handler))
}

type Event struct {
	Bucket         string
	FunctionArn    string
	NotificationId string
}

func handler(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	fmt.Println(event)

	physicalResourceID = "wojtkeTestingStuff"
	if event.RequestType != cfn.RequestCreate {
		return
	}

	bucket := event.ResourceProperties["Bucket"].(string)
	functionArn := event.ResourceProperties["FunctionArn"].(string)
	notificationId := event.ResourceProperties["NotificationId"].(string)

	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	_, err = svc.PutBucketNotificationConfiguration(&s3.PutBucketNotificationConfigurationInput{
		Bucket: aws.String(bucket),
		NotificationConfiguration: &s3.NotificationConfiguration{
			LambdaFunctionConfigurations: []*s3.LambdaFunctionConfiguration{
				{
					Events: []*string{aws.String("s3:CreateObject:*")},
					Filter: &s3.NotificationConfigurationFilter{
						Key: &s3.KeyFilter{
							FilterRules: []*s3.FilterRule{
								{
									Name:  aws.String("suffix"),
									Value: aws.String(".json"),
								},
							},
						},
					},
					Id:                aws.String(notificationId),
					LambdaFunctionArn: aws.String(functionArn),
				},
			},
		},
	})

	fmt.Println(physicalResourceID, data, err)

	return
}

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type instances struct{}

func (i instances) new(profile string, region string) ([]map[string]string, error) {
	// open session
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
		Config: aws.Config{
			Region: aws.String(region),
		},
		SharedConfigState: session.SharedConfigEnable,
	})

	// create service
	svc := ec2.New(sess)

	// get instances
	// Only grab instances that are running or just started
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
					aws.String("pending"),
				},
			},
		},
	}

	res, _ := svc.DescribeInstances(params)

	return sort(res), err
}

func nilGuard(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

// return instances as a map
func sort(i *ec2.DescribeInstancesOutput) []map[string]string {
	var instances []map[string]string

	for _, reservation := range i.Reservations {
		for _, instance := range reservation.Instances {
			record := make(map[string]string)
			record["instance-id"] = nilGuard(instance.InstanceId)
			record["public-hostname"] = nilGuard(instance.PublicDnsName)
			record["public-ip"] = nilGuard(instance.PublicIpAddress)
			record["private-hostname"] = nilGuard(instance.PrivateDnsName)
			record["private-ip"] = nilGuard(instance.PrivateIpAddress)

			for _, tag := range instance.Tags {
				record[fmt.Sprintf("tag:%s", *tag.Key)] = *tag.Value
			}

			instances = append(instances, record)
		}
	}

	return instances
}

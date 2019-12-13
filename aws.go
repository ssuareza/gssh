package gssh

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// NewService creates a service connection with ec2
func NewService(profile string, region string) (ec2iface.EC2API, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
		Config: aws.Config{
			Region: aws.String(region),
		},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}

	return ec2.New(sess), nil
}

// Get gets the instances from aws
func Get(svc ec2iface.EC2API) (*ec2.DescribeInstancesOutput, error) {
	// Only grab instances that are running or just started
	input := &ec2.DescribeInstancesInput{
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

	res, err := svc.DescribeInstances(input)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Filter gets only a few metadata fields from ec2.DescribeInstancesOutput
func Filter(i *ec2.DescribeInstancesOutput) []map[string]string {
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

func nilGuard(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

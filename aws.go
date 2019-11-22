package gssh

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// Get gets the instances from aws
func Get(profile string, region string) ([]map[string]string, error) {
	// open session
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
		Config: aws.Config{
			Region: aws.String(region),
		},
		SharedConfigState: session.SharedConfigEnable,
	})

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

	// create service
	svc := ec2.New(sess)

	// this is for testing
	StubEC2(svc, input)

	// get instances
	res, _ := svc.DescribeInstances(input)

	return Sort(res), err
}

// stub (this is for testing)
func StubEC2(svc ec2iface.EC2API, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	res, err := svc.DescribeInstances(input)
	if err != nil {
		return nil, err
	}
	return res, err
}

func nilGuard(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

// Sort gets only a few metadata fields from instances list
func Sort(i *ec2.DescribeInstancesOutput) []map[string]string {
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

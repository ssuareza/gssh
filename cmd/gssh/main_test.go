package main

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/ssuareza/gssh"
)

type MockEC2API struct {
	ec2iface.EC2API
}

func (m *MockEC2API) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	output := &ec2.DescribeInstancesOutput{}
	return output, nil

}

func TestMyFunc(t *testing.T) {
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

	svc := &MockEC2API{}
	res, err := gssh.StubEC2(svc, input)
	fmt.Println(res)
	if err != nil {
		t.Error("Not able to get instances")
	}

	//resp := Get("dev", "us-east-1")
	/*instancesMap, err := instances.DescribeInstances()
	if err != nil {
		t.Error("Not able to get instances")
	}

	sorted := Sort(instancesMap)
	//fmt.Printf("%s", sorted)
	if sorted == nil {
		t.Error("Not able to get sorted instances map")
	}*/
}

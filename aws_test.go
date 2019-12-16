package gssh

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// Define a mock struct to be used in your unit tests of myFunc.
type mockEC2Client struct {
	ec2iface.EC2API
}

// DescribeInstances mock
func (m *mockEC2Client) DescribeInstances(*ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return &ec2.DescribeInstancesOutput{
		Reservations: []*ec2.Reservation{
			&ec2.Reservation{
				OwnerId: aws.String("1234567890"),
				Instances: []*ec2.Instance{
					&ec2.Instance{
						InstanceId:       aws.String("i-0"),
						PrivateDnsName:   aws.String("test1.private.ec2"),
						PrivateIpAddress: aws.String("0.0.0.1"),
						PublicDnsName:    aws.String("test1.public.ec2"),
						PublicIpAddress:  aws.String("0.0.1.1"),
						VpcId:            aws.String("vpc-1"),
						State: &ec2.InstanceState{
							Name: aws.String("running"),
						},
						Tags: []*ec2.Tag{
							&ec2.Tag{
								Key:   aws.String("tag1"),
								Value: aws.String("tag1"),
							},
						},
					},
					&ec2.Instance{
						InstanceId:       aws.String("i-1"),
						PrivateDnsName:   aws.String("test2.private.ec2"),
						PrivateIpAddress: aws.String("0.0.0.2"),
						PublicDnsName:    aws.String("test2.public.ec2"),
						PublicIpAddress:  aws.String("0.0.1.2"),
						VpcId:            aws.String("vpc-2"),
						State: &ec2.InstanceState{
							Name: aws.String("running"),
						},
						Tags: []*ec2.Tag{
							&ec2.Tag{
								Key:   aws.String("tag2"),
								Value: aws.String("tag2"),
							},
						},
					},
				},
			},
		},
	}, nil
}

func TestFilter(t *testing.T) {
	mockSvc := &mockEC2Client{}
	var servers []*ec2.DescribeInstancesOutput
	instances, err := Get(mockSvc)
	if err != nil {
		t.Fail()
	}
	servers = append(servers, instances)

	if len(Filter(servers)) != 2 {
		t.Error("Number of instances expected 2")
	}
}

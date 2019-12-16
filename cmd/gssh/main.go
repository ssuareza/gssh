package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/fatih/color"
	"github.com/gosuri/uitable"
	"github.com/ssuareza/gssh"
)

func printTable(i []gssh.Server) {
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow(color.BlueString("InstanceID"), color.BlueString("Name"), color.BlueString("PrivateIP"), color.BlueString("PublicIP"))
	for _, instance := range i {
		table.AddRow(color.GreenString(instance.Values["instance-id"]), instance.Name, instance.Values["private-ip"], instance.Values["public-ip"])
	}
	fmt.Printf("%s\n\n", table)
}

func getIP(id string, i []gssh.Server, iptype string) (string, error) {
	for _, instance := range i {
		if instance.Values["instance-id"] == id {
			if iptype == "public" {
				return instance.Values["public-ip"], nil
			}
			return instance.Values["private-ip"], nil
		}
	}
	return "", errors.New("IP not found")
}

func main() {
	// args
	var filter string
	if len(os.Args) == 2 {
		filter = os.Args[1]
	}

	// get config
	c, err := gssh.GetConfig()
	if err != nil {
		log.Panic(err)
	}

	// get instances
	profiles := strings.Split(c.AWS.Profile, ",")
	var instances []*ec2.DescribeInstancesOutput
	for k := range profiles {
		svc, err := gssh.NewService(profiles[k], c.AWS.Region)
		if err != nil {
			log.Fatal(err)
		}

		list, err := gssh.Get(svc, filter)
		if err != nil {
			log.Fatal(err)
		}
		instances = append(instances, list)
	}

	i := gssh.Filter(instances)
	if err != nil {
		log.Panic(err)
	}

	// filter
	printTable(i)

	// select instance
	fmt.Print("Select InstanceID: ")
	var instanceID string
	_, err = fmt.Scanln(&instanceID)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// get ip
	iptype := "private"
	if len(c.SSH.Bastion) == 0 {
		iptype = "public"
	}

	ip, err := (getIP(instanceID, i, iptype))
	if err != nil {
		log.Fatal(err)
	}

	// and connect
	gssh.Shell(ip, c)
}

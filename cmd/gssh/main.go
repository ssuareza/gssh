package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gosuri/uitable"
	"github.com/ssuareza/gssh"
)

func print(i []map[string]string) {
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("InstanceID", "Name", "PrivateIP", "PublicIP")
	for _, instance := range i {
		table.AddRow(instance["instance-id"], instance["tag:Name"], instance["private-ip"], instance["public-ip"])
	}
	fmt.Println(table)
}

func getIP(id string, i []map[string]string, iptype string) (string, error) {
	for _, instance := range i {
		if instance["instance-id"] == id {
			if iptype == "public" {
				return instance["public-ip"], nil
			}
			return instance["private-ip"], nil
		}
	}
	return "", errors.New("IP not found")
}

func main() {
	// get config
	c, err := gssh.GetConfig()
	if err != nil {
		log.Panic(err)
	}

	// get instances
	svc, err := gssh.NewService(c.AWS.Profile, c.AWS.Region)
	if err != nil {
		log.Fatal(err)
	}

	instances, err := gssh.Get(svc)
	if err != nil {
		log.Fatal(err)
	}

	i := gssh.Filter(instances)
	if err != nil {
		log.Panic(err)
	}
	print(i)

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

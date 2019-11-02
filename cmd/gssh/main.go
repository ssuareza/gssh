package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gosuri/uitable"
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

func getIP(id string, i []map[string]string) (string, error) {
	for _, instance := range i {
		if instance["instance-id"] == id {
			return instance["private-ip"], nil
		}
	}
	return "", errors.New("IP not found")
}

func main() {
	// get config
	c, err := config{}.new()
	if err != nil {
		log.Panic(err)
	}

	// get instances
	i, err := instances{}.new(c.aws, c.region)
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
	ip, err := (getIP(instanceID, i))
	if err != nil {
		log.Fatal(err)
	}

	// and connect
	shell(ip, c)
}

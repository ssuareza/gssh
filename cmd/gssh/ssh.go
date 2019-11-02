package main

import (
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

//  return publickey content
func publicKeyFile(file string) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(key), nil
}

// connect and open terminal
func connect(conn *ssh.Client) {
	sess, err := conn.NewSession()
	if err != nil {
		panic(err)
	}

	sess.Stdout = os.Stdout
	sess.Stdin = os.Stdin
	sess.Stderr = os.Stderr

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,      // please print what I type
		ssh.ECHOCTL:       0,      // please don't print control chars
		ssh.TTY_OP_ISPEED: 115200, // baud in
		ssh.TTY_OP_OSPEED: 115200, // baud out
	}

	// open shell
	termFD := int(os.Stdin.Fd())
	w, h, _ := terminal.GetSize(termFD)
	termState, _ := terminal.MakeRaw(termFD)
	defer terminal.Restore(termFD, termState)
	sess.RequestPty("xterm-256color", h, w, modes)
	sess.Shell()
	sess.Wait()
}

// connect trought bastion
func proxy(bastion *ssh.Client, host string, clientCfg *ssh.ClientConfig) (*ssh.Client, error) {
	netConn, _ := bastion.Dial("tcp", host)

	conn, chans, reqs, err := ssh.NewClientConn(netConn, host, clientCfg)

	return ssh.NewClient(conn, chans, reqs), err
}

func shell(host string, c config) error {
	user := c.user
	port := c.port
	bastion := c.bastion

	// configure ssh connection
	publicKey, err := publicKeyFile(os.Getenv("HOME") + "/.ssh/id_rsa")
	if err != nil {
		log.Println(err)
	}
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			publicKey,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// connect
	// trough bastion
	if len(bastion) != 0 {
		conn, _ := ssh.Dial("tcp", bastion+":"+port, config)
		defer conn.Close()
		newConn, err := proxy(conn, host+":"+port, config)
		if err != nil {
			return err
		}

		// open terminal
		connect(newConn)
	} else { // or not
		conn, _ := ssh.Dial("tcp", host+":"+port, config)
		defer conn.Close()

		// open terminal
		connect(conn)
	}
	return err
}

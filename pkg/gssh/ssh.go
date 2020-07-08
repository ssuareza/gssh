package gssh

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// GetPublicKey returns publickey content
func GetPublicKey(file string) (ssh.AuthMethod, error) {
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

// Connect connects trough ssh
func Connect(conn *ssh.Client) {
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

// Proxy connects trought a bastion server
func Proxy(bastion *ssh.Client, host string, clientCfg *ssh.ClientConfig) (*ssh.Client, error) {
	netConn, _ := bastion.Dial("tcp", host)

	conn, chans, reqs, err := ssh.NewClientConn(netConn, host, clientCfg)

	return ssh.NewClient(conn, chans, reqs), err
}

// Shell opens a terminal in destination
func Shell(host string, c *Config) error {
	user := c.SSH.User
	port := fmt.Sprint(c.SSH.Port)
	bastion := c.SSH.Bastion

	// configure ssh connection
	publicKey, err := GetPublicKey(os.Getenv("HOME") + "/.ssh/id_rsa")
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
	var conn *ssh.Client
	if len(bastion) != 0 {
		conn, _ = ssh.Dial("tcp", bastion+":"+port, config)
		conn, err = Proxy(conn, host+":"+port, config)
		if err != nil {
			return err
		}
	} else { // or not
		conn, _ = ssh.Dial("tcp", host+":"+port, config)
	}
	defer conn.Close()

	// open terminal
	Connect(conn)
	return err
}

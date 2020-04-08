package main

import (
	"log"

	sshc "github.com/robjporter/go-library/xsshc2"
)

func main() {
	ip := ""
	port := "22"
	username := "root"
	pemFile := "/home/william/.ssh/test-unencrypted.pem"

	// Prepare ssh client
	client, err := xsshc2.NewClient(&xsshc2.Config{
		User:           username,
		PrivateKeyFile: pemFile,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Connect to remote host
	conn, err := client.Connect(ip, port)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	log.Println("Connected, running command(s)...")

	// Run command(s) on remote host
	cmds := []string{
		"echo 'Hello World!' > /root/hello.txt",
		"sleep 10",
		"echo 'Good bye!' >> /root/hello.txt",
	}
	if err = client.Run(conn, cmds); err != nil {
		log.Println(err)
	}
}
package utils

import (
	"fd/connectErr"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"time"
)

func ConnectSftp(user, passsword, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		addr         string
	)

	addr = fmt.Sprintf("%s:%d", host, port)
	auth = append(auth, ssh.Password(passsword))
	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	sshClient, err := ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		return nil, connectErr.GetConnectSSHErr(err)
	}
	sftpClient, err = sftp.NewClient(sshClient)
	if err != nil {
		return nil, connectErr.GetConnectSSHErr(err)
	}
	return sftpClient, nil
}

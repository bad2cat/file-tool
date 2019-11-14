package Impl

import (
	"fd/connectErr"
	"fd/utils"
	"fmt"
	"os"
	"path"
)

func UpFileFromServer(user, password, host string, port int, loc, remote string) error {
	sftpClient, err := utils.ConnectSftp(user, password, host, port)
	if err != nil {
		return err
	}
	defer sftpClient.Close()
	srcFile, err := os.Open(loc)
	if err != nil {
		return connectErr.GetConnectSSHErr(err)
	}
	defer srcFile.Close()
	buf := make([]byte, 2048)
	remoteFileName := path.Base(loc)
	dstFile, err := sftpClient.Create(path.Join(remote, remoteFileName))
	if err != nil {
		return connectErr.GetConnectSSHErr(err)
	}
	defer dstFile.Close()
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}
	fmt.Println("write success")
	return nil
}

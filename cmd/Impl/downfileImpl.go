package Impl

import (
	"fd/utils"
	"fmt"
	"os"
	"path"
)

func DownFileFromServer(user,password,host string,port int,loc,remote string) error {
	sftpClient,err := utils.ConnectSftp(user,password,host,port)
	if err != nil{
		return err
	}
	defer sftpClient.Close()

	srcFile,err := sftpClient.Open(remote)
	if err != nil{
		return err
	}
	defer srcFile.Close()

	var locFileName = path.Base(remote)
	dtFile,err := os.Create(path.Join(loc,locFileName))
	if err != nil{
		return err
	}
	defer dtFile.Close()
	if _,err = srcFile.WriteTo(dtFile);err !=nil{
		return err
	}
	fmt.Println("download success")
	return nil
}
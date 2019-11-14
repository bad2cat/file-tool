package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestConnectSftp(t *testing.T) {
	sftClient,err := ConnectSftp("sd","123456","172.21.81.173",22)
	if err != nil{
		fmt.Fprint(os.Stdout,err)
	}
	defer sftClient.Close()
}

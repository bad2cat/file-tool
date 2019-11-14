package cmd

import (
	"fmt"
	"path"
	"testing"
)

func TestBasePath(t *testing.T) {
	pathName := "D:/tmpDir/test.py"
	fileName := path.Base(pathName)
	fmt.Println(fileName)
	locPath := "/home/jx/"
	pathFull := path.Join(locPath, fileName)
	fmt.Println(pathFull)
}

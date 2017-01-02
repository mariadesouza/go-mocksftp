package main

import (
	"fmt"
	"testing"
)

type FakeSftpConnector struct {
	itemName string
}

// make sure it satisfies the interface
var _ SftpConnector = (*FakeSftpConnector)(nil)

// Implement Atleast One functin
func (f *FakeSftpConnector) RemoveFile(source string) error {
	fmt.Println("RemoveSource")
	return nil
}

func TestProcessInnerFiles_bad(t *testing.T) {
	tt := &FakeSftpConnector{}

	err := processInnerFiles(tt, "/home/ftptest/", "abc.json")
	if err == nil {
		t.Errorf("Error was expected in processInnerFiles: %s", err)
	}

}
//This is th original function that takes in an sftp connections
//func processInnerFiles(sftpConnection *sftphelper.SFTPConnection, itemPath, itemName string) error{
	


package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/moby/moby/api/types"
	"github.com/moby/moby/client"
	"io"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(cli.ClientVersion())
	fmt.Println("=====pull image=====")
	image := "tomcat:8-jdk8"
	var pullReader io.ReaderCloser
	pullReader, err = cli.ImagePull(context.Background(), image, type.ImagePullOptions{
		All:	false,
		RegistryAuth: "",
		PrivilegeFunc: nil,	
	})
	if err != nil {
		panic(err.Error())
	}
	defer pullReader.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(pullReader)
	s := buf.String()
	fmt.Println("info:", s)
	fmt.Println("image pull success")
}
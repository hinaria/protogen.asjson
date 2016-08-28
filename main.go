package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	d "unyaa.com/protogen.asjson/descriptor"
)

func main() {
	buffer, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		exitWithError(err)
	}

	descriptors := &d.FileDescriptorSet{}
	err = proto.Unmarshal(buffer, descriptors)
	if err != nil {
		exitWithError(err)
	}

	marshaller := jsonpb.Marshaler{
		EmitDefaults: true,
		EnumsAsInts:  false,
	}

	json, err := marshaller.MarshalToString(descriptors)
	if err != nil {
		exitWithError(err)
	}

	fmt.Println(json)
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(-1)
}

package main

import (
	"fmt"
	"log"
	"testing"
)

func TestEnumWindows(t *testing.T) {
	wRectList := getAllWindowRect()
	if len(wRectList) > 0 {
		fmt.Printf("%+v\n", wRectList)
	}
}

func TestJsonMarshallUnmarshall(t *testing.T) {
	dispList := &[]Display{}
	unmarshallJsonFile(PROFILE_NAME, dispList)
	log.Printf("%+v\n", dispList)
	marshallToJsonFile(PROFILE_NAME+"_test", dispList)
	unmarshallJsonFile(PROFILE_NAME+"_test", dispList)
	log.Printf("%+v\n", dispList)
}

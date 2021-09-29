package main

import (
	"log"
	"os"
	"testing"
)

func init() {
	log.SetOutput(os.Stdout)
}

func TestInit(t *testing.T) {
	log.Println("hello init")
}
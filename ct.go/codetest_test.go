package main

import (
	"os"
	"testing"
)

func driver(t *testing.T) {
	return os.Open(string(drive) + ":\\")

}

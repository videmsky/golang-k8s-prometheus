package server

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Hello Kubiya!!!" {
		t.Fatal("Test fail")
	}
}
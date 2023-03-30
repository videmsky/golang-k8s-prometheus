package server

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Hello GKP!!!" {
		t.Fatal("Test fail")
	}
}
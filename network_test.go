package main

import (
	"net"
	"testing"
)

func TestItReturnsUnavailablePorts(t *testing.T) {
	ln, _ := net.Listen("tcp", ":8072")
	defer ln.Close()

	unavailable := PortChecker([]string{"8071", "8072", "8073"})
	length := len(unavailable)

	if length != 1 {
		t.Errorf("Expected length of 1, got '%d'", length)
	}

	if length == 1 && unavailable[0] != "8072" {
		t.Errorf("Expected port 8072, got '%s'", unavailable[0])
	}
}

func TestItReturnsEmptySlice_whenGivenEmptySlice(t *testing.T) {
	unavailable := PortChecker([]string{})

	if len(unavailable) != 0 {
		t.Errorf("Expected empty slice, got '%v'", unavailable)
	}
}

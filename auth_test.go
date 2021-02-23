package main

import (
	"testing"
)

var (
	testUser   = "user"
	testPass   = "pass"
	testServer = "smtp.test.com"
	testPort   = 465
)

func TestAuth(t *testing.T) {
	d := NewDialer(testServer, testPort, testUser, testPass)
	err := d.DialAndAuth()
	if err != nil {
		t.Fatal(err)
	}
}

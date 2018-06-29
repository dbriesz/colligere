package cmd

import (
	"testing"

	"github.com/rendon/testcli"
)

func TestGreetings(t *testing.T) {
	// Using package functions
	testcli.Run("greetings")
	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s", testcli.Error())
	}

	if !testcli.StdoutContains("Hello?") {
		t.Fatalf("Expected %q to contain %q", testcli.Stdout(), "Hello?")
	}
}

func TestGreetingsWithName(t *testing.T) {
	// Using the struct version, if you want to test multiple commands
	c := testcli.Command("greetings", "--name", "John")
	c.Run()
	if !c.Success() {
		t.Fatalf("Expected to succeed, but failed with error: %s", c.Error())
	}

	if !c.StdoutContains("Hello John!") {
		t.Fatalf("Expected %q to contain %q", c.Stdout(), "Hello John!")
	}
}

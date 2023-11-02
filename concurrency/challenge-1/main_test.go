package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_printMessage(t *testing.T) {
	wg.Add(1)

	updateMessage("epsilon", &wg)

	wg.Wait()

	if !strings.Contains(msg, "epsilon") {
		t.Errorf("Error in Test_printMessage")
	}

}

func Test_updateMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "Hello, world!"

	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Error in Test_updateMessage")
	}
}

func Test_main(t *testing.T) {

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Errorf("Error in Test_main#1")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Errorf("Error in Test_main#2")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Errorf("Error in Test_main#2")
	}

}

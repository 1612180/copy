package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	testName1 := "test.txt"
	testName2 := "test_another.txt"

	// Quick write
	ioutil.WriteFile(testName1, []byte("Hello\nUniverse\n:)"), 0644)

	copy(testName1, testName2)

	// Test if 2 file is equal
	if !isEqual("test.txt", "test_another.txt") {
		t.Error("New copied file is not the same")
	}

	// Remove test file
	os.Remove(testName1)
	os.Remove(testName2)
}

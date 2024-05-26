package main

import "testing"

func TestValid(t *testing.T) {
	result := isSubsequence("abc", "ahbgdc")
	if !result {
		t.Errorf("Expected true, got false")
	}
}

func TestInvalid(t *testing.T) {
	result := isSubsequence("axc", "ahbgdc")
	if result {
		t.Errorf("Expected false, got true")
	}
}

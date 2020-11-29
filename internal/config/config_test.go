package config

import (
	"testing"
)

func TestValidConfig(t *testing.T) {
	_, err := LoadConfig("tests/valid.yaml")
	if err != nil {
		t.Error()
	}
}

func TestInvalidConfig(t *testing.T) {
	_, err := LoadConfig("tests/invalid.yaml")
	if err == nil {
		t.Error()
	}
}

func TestUnavailableConfig(t *testing.T) {
	_, err := LoadConfig("tests/notexist.yaml")
	if err == nil {
		t.Error()
	}
}

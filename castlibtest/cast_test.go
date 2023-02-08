package castlibtest

import (
	"github.com/Eclalang/cast"
	"testing"
)

func TestAtoi(t *testing.T) {
	expected := 12
	actual := cast.Atoi("12")
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFloatToInt(t *testing.T) {
	expected := 12
	actual := cast.FloatToInt(12.34)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestIntToFloat(t *testing.T) {
	expected := 12.0
	actual := cast.IntToFloat(12)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestParseBool(t *testing.T) {
	expected := true
	actual := cast.ParseBool("true")
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestParseFloat(t *testing.T) {
	expected := 12.34
	actual := cast.ParseFloat("12.34", 64)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

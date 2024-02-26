package cast

import (
	"testing"
)

func TestAtoi(t *testing.T) {
	expected := 12
	actual := Atoi("12")
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestFloatToInt(t *testing.T) {
	expected := 12
	actual := FloatToInt(12.34)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestIntToFloat(t *testing.T) {
	expected := 12.0
	actual := IntToFloat(12)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestParseBool(t *testing.T) {
	expected := true
	actual := ParseBool("true")
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestParseFloat(t *testing.T) {
	expected := 12.34
	actual := ParseFloat("12.34", 64)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

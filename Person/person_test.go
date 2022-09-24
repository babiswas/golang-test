package Person

import "testing"

func TestNameSetting(t *testing.T) {
	p := Person{}
	p.Set_name("Bapan")
	if p.name != "Bapan" {
		t.Error("Expected Bapan got something else:")
	}
}

func TestIdSetting(t *testing.T) {
	p := Person{}
	p.Set_id(12)
	if p.id != 12 {
		t.Error("Expected 12 got something else")
	}
}

func TestIdNameSetting(t *testing.T) {
	p := Person{}
	p.Set_id(12)
	p.Set_name("Bapan")
	if p.id != 12 || p.name != "Bapan" {
		t.Error("Expected name Bapan and Id 12 ,got something else")
	}
}

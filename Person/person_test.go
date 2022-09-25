package Person

import "testing"

type PersonTest struct {
	name string
	id   int
	want string
}

var persontest = []PersonTest{
	{"Bapan", 12, "Bapan 12"},
	{"Madan", 21, "Madan 21"},
	{"Kiran", 12, "Kiran 12"},
	{"Ajay", 2, "Ajay 2"},
}

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

func TestDisplay(t *testing.T) {
	p := Person{}
	p.Set_id(21)
	p.Set_name("Bapan")
	if p.Display() != "Bapan 21" {
		t.Error("Expected Bapan 21 got something else")
	}
}

func TestInterfaceWork(t *testing.T) {
	p := Person{}
	if p.Work() != "working" {
		t.Error("Inteface method Work failed")
	}
}

func TestInterfaceSalary(t *testing.T) {
	p := Person{}
	if p.Salary() != 15000 {
		t.Error("Expected salary 15000 got something else.")
	}
}

func TestPersonString(t *testing.T) {
	p := Person{}
	for _, val := range persontest {

		p.Set_name(val.name)
		p.Set_id(val.id)
		got := p.Display()
		want := val.want
		if got != want {
			t.Error("Test Failed")
		}
	}
}

func BenchmarkSetName(b *testing.B) {
	p := Person{}
	for n := 0; n < b.N; n++ {
		p.Set_name("bapan")
	}
}

func BenchmarkSetid(b *testing.B) {
	p := Person{}
	for n := 0; n < b.N; n++ {
		p.Set_id(12)
	}
}

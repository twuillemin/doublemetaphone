package doublemetaphone

import "testing"

func TestEmptyString(t *testing.T) {

	primary, secondary := DoubleMetaphone("")
	if len(primary) != 0 {
		t.Fatal("Primary DoubleMetaphone of empty String must be empty")
	}
	if len(secondary) != 0 {
		t.Fatal("Secondary DoubleMetaphone of empty String must be empty")
	}
}

func TestShortString(t *testing.T) {

	primary, secondary := DoubleMetaphone("g")
	if primary != "K" {
		t.Fatal("Primary DoubleMetaphone of \"g\" must be \"K\" but was " + primary)
	}
	if secondary != "K" {
		t.Fatal("Secondary DoubleMetaphone of \"g\" must be \"K\" but was " + secondary)
	}
}

func TestMyName(t *testing.T) {

	primary, secondary := DoubleMetaphone("Wuillemin")
	if primary != "AMN" {
		t.Fatal("Primary DoubleMetaphone of \"Wuillemin\" must be \"AMN\" but was " + primary)
	}
	if secondary != "FMN" {
		t.Fatal("Secondary DoubleMetaphone of \"Wuillemin\" must be \"FMN\" but was " + secondary)
	}
}

func TestBasicValues(t *testing.T) {

	primary, secondary := DoubleMetaphone("dog")
	if primary != "TK" {
		t.Fatal("Primary DoubleMetaphone of \"dog\" must be \"TK\" but was " + primary)
	}
	if secondary != "TK" {
		t.Fatal("Secondary DoubleMetaphone of \"dog\" must be \"TK\" but was " + secondary)
	}

	primary, secondary = DoubleMetaphone("cat")
	if primary != "KT" {
		t.Fatal("Primary DoubleMetaphone of \"cat\" must be \"KT\" but was " + primary)
	}
	if secondary != "KT" {
		t.Fatal("Secondary DoubleMetaphone of \"cat\" must be \"KT\" but was " + secondary)
	}
}
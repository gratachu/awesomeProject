package main

import (
	"reflect"
	"testing"
)

func TestUniq(t *testing.T) {
	strings := []string{"red", "green", "blue", "blue", "green", "blue"}

	got := uniq(strings)
	want := []string{"red", "green", "blue"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got error %q want %q", got , want)
	}
}
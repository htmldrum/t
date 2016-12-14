package main

import (
	"testing"
)

func TestUsage(t *testing.T){
	var test = struct {
		input int
		want int
	}{len(usage()), 2}
	
	if test.want != test.input {
		t.Errorf(`len_usage. Expected: %d. Got: 2, %d.`, test.want, test.input)
	}
}

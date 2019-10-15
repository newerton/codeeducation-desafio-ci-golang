package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(10, 20)
	if got != 30 {
		t.Errorf("Sum return = %d; want 30", got)
	}
}
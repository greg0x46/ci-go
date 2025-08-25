package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(2, 3)
	
	if total != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", total)
	
	}
}

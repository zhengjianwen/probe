package apm

import "testing"

func TestDeleteRule(t *testing.T) {
	if err := DeleteRule(1, 512); err != nil {
		t.Fatal(err)
	}
}
package main

import "testing"

func TestGetProtectedMachineId(t *testing.T) {
	id, err := GetProtectedMachineId()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

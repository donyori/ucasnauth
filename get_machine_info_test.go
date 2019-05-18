package main

import "testing"

func TestGetMacAddrs(t *testing.T) {
	addrs, err := GetMacAddrs()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(addrs)
}

func TestGetProtectedMachineId(t *testing.T) {
	id, err := GetProtectedMachineId()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}

package main

import "github.com/denisbrodbeck/machineid"

func GetProtectedMachineId() (string, error) {
	return machineid.ProtectedID(AppId)
}

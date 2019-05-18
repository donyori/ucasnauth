package main

import (
	"net"

	"github.com/denisbrodbeck/machineid"
)

func GetMacAddrs() ([]string, error) {
	nis, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	addrs := make([]string, 0, len(nis))
	for i := range nis {
		mac := nis[i].HardwareAddr.String()
		if mac != "" {
			addrs = append(addrs, mac)
		}
	}
	if len(addrs) == 0 {
		addrs = nil
	}
	return addrs, nil
}

func GetProtectedMachineId() (string, error) {
	return machineid.ProtectedID(AppId)
}

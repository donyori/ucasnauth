package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/crypto/scrypt"
)

func Save(data []byte) error {
	exeDir, err := getExeDir()
	if err != nil {
		return err
	}
	key, err := generateKey(true)
	if err != nil {
		return err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	nonce := make([]byte, 12)
	_, err = rand.Read(nonce)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(exeDir, NonceFilename), nonce, 0666)
	if err != nil {
		return err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	encrypted := aesgcm.Seal(nil, nonce, data, []byte(AppId))
	return ioutil.WriteFile(filepath.Join(exeDir, DataFilename),
		encrypted, 0666)
}

func Load() ([]byte, error) {
	exeDir, err := getExeDir()
	if err != nil {
		return nil, err
	}
	nonce, err := ioutil.ReadFile(filepath.Join(exeDir, NonceFilename))
	if err != nil {
		return nil, err
	}
	encrypted, err := ioutil.ReadFile(filepath.Join(exeDir, DataFilename))
	if err != nil {
		return nil, err
	}
	key, err := generateKey(false)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return aesgcm.Open(nil, nonce, encrypted, []byte(AppId))
}

func generateKey(isForSave bool) ([]byte, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	mId, err := GetProtectedMachineId()
	if err != nil {
		mId = ""
	}
	macAddrs, err := GetMacAddrs()
	if err != nil {
		macAddrs = nil
	}
	if mId == "" && macAddrs == nil {
		return nil, ErrCannotGetMachineInfo
	}
	seed := mId + strings.Join(macAddrs, "")
	var saltFilename string
	if runtime.GOOS == "windows" {
		saltFilename = "_" + SaltFilename
	} else {
		saltFilename = "." + SaltFilename
	}
	saltFilename = filepath.Join(userHome, saltFilename)
	var salt []byte
	if isForSave {
		salt = make([]byte, 32)
		_, err = rand.Read(salt)
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile(saltFilename, salt, 0600)
	} else {
		salt, err = ioutil.ReadFile(saltFilename)
	}
	if err != nil {
		return nil, err
	}
	return scrypt.Key([]byte(seed), salt, 32768, 8, 1, 32)
}

func getExeDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		exePath, err = filepath.Abs(os.Args[0])
		if err != nil {
			return "", err
		}
	}
	exePath, err = filepath.EvalSymlinks(exePath)
	if err != nil {
		return "", err
	}
	return filepath.Dir(exePath), nil
}

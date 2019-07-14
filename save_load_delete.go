package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/crypto/scrypt"
)

func Save(data []byte) error {
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
	err = ioutil.WriteFile(filepath.Join(DataDir, NonceFilename), nonce, 0600)
	if err != nil {
		return err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	encrypted := aesgcm.Seal(nil, nonce, data, []byte(AppId))
	return ioutil.WriteFile(filepath.Join(DataDir, DataFilename),
		encrypted, 0600)
}

func Load() ([]byte, error) {
	nonce, err := ioutil.ReadFile(filepath.Join(DataDir, NonceFilename))
	if err != nil {
		return nil, err
	}
	encrypted, err := ioutil.ReadFile(filepath.Join(DataDir, DataFilename))
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

func Delete() error {
	return os.RemoveAll(DataDir)
}

func generateKey(isForSave bool) ([]byte, error) {
	seed, err := GetProtectedMachineId()
	if err != nil {
		return nil, err
	}
	if _, oserr := os.Stat(DataDir); os.IsNotExist(oserr) {
		err = os.MkdirAll(DataDir, 0600)
		if err != nil {
			return nil, err
		}
	}
	saltFilename := filepath.Join(DataDir, SaltFilename)
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

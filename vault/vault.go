package vault

import (
	"encoding/json"
	"errors"
	"os"
)

func SaveVaultToFile(vault *Vault, password, filepath string) error {	
	data, err := json.Marshal(vault)
	if err != nil {
		return err
	}

	salt, err := GenerateSalt()
	if err != nil {
		return err
	}
	key := DeriveKey(password, salt)

	encrypted, err := Encrypt(data, key)
	if err != nil {
		return err
	}

	finalData := append(salt, encrypted...)
	return os.WriteFile(filepath, finalData, 0600)
}

func LoadVaultFromFile(password, filepath string) (*Vault, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if len(data) < saltSize {
		return nil, errors.New("fail is broken")
	}

	salt := data[:saltSize]
	encrypted := data[saltSize:]

	key := DeriveKey(password, salt)

	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		return nil, err
	}

	var vault Vault
	err = json.Unmarshal(decrypted, &vault)
	if err != nil {
		return nil, err
	}

	return &vault, nil
}


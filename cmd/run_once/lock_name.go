package main

import (
	"crypto/sha256"
	"encoding/base32"
	"encoding/json"
	"os"
	"path/filepath"
)

func getLockPath(cmd string, args []string) (string, error) {
	LockJson, err := json.Marshal(append([]string{cmd}, args...))
	if err != nil {
		return "", err
	}
	hash := sha256.Sum224(LockJson)
	base64EncodedHash := base32.StdEncoding.EncodeToString(hash[:])

	lockPath := filepath.Join(os.TempDir(), base64EncodedHash)

	return lockPath, nil
}

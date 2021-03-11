package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"

	//"math/rand"
	//"strconv"
	"github.com/gorilla/mux"
)

// Struct for version

type Version struct {
	Vers string `json:"GoLang-version"`
}

// Struct the JSON Single Value

type Value struct {
	Val string `json:"value"`
}

//Initialize empty array

var ver []Version
var val []Value

// Declare required variables
var passphrase string = "password123"     // Going to go with default passphrase
var fileName string = "encrypt-file.json" /* File gets created as part of the process and
where the encrypted data get stored. To decrypt the ecrypted content, this file will be called
and the content in there will be used. */

// Get Version of Go lang
func version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	info := runtime.Version()
	ver = append(ver, Version{Vers: info})
	json.NewEncoder(w).Encode(ver)
}

// Hashing Passwords to Compatible Cipher Keys; in this program we always go with default password: password123
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Encrypting Data with an AES Cipher
func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

// Decrypting Data that uses an AES Cipher
func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

// Encrypt JSON string and store it in a JSON file
func encryptFile(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	data, _ := ioutil.ReadAll(r.Body)
	//passphrase := "password123" // Going to go with default passphrase
	//fileName := "sample.json"
	f, _ := os.Create(fileName)
	defer f.Close()
	//Lets encrypt the string and store it in the sample.json file
	f.Write(encrypt(data, passphrase))
	//fmt.Fprintf(w, "%+v", string(data))
	json.NewEncoder(w).Encode(data)

}

// Decrypt the JSON file which was encryptd and return the value
func decryptFile(w http.ResponseWriter, r *http.Request) {

	data1, _ := ioutil.ReadFile(fileName)
	w.Header().Set("Content-Type", "application/json")
	//passphrase := "password123" // Going to go with default passphrase
	//fileName := "sample.json"
	pText := decrypt(data1, passphrase)

	//pText := strconv.Itoa(pText)
	//val = append(val, Value{Val: pText})
	json.NewEncoder(w).Encode(string(pText))
	//fmt.Println("%T", string(pText))

}

// Main function
func main() {
	// Initialize router
	r := mux.NewRouter()

	// Route handlers & endpoints
	r.HandleFunc("/api/version", version).Methods("GET")
	r.HandleFunc("/api/encrypt", encryptFile).Methods("POST")
	r.HandleFunc("/api/decrypt", decryptFile).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}

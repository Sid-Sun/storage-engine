package handlers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"github.com/sid-sun/storage-engine/pkg/api/contract/db"
	"golang.org/x/crypto/sha3"
	"io"
)

// Encrypt encrypts a note with the pass and returns an encrypted AAD and encrypted note
func Encrypt(note string, password string) ([]byte, [32]byte, []byte, error) {
	// Generate additional auth data of 256 Bits
	AAD := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, AAD); err != nil {
		return nil, [32]byte{}, nil, err
	}

	// Store in a var to avoid converting twice
	noteBytes := []byte(note)

	// Create dst with length of aes blocksize + note length
	// And initialize first BlockSize bytes randonly for IV
	dst := make([]byte, aes.BlockSize+len(noteBytes))
	if _, err := io.ReadFull(rand.Reader, dst[:aes.BlockSize]); err != nil {
		return nil, [32]byte{}, nil, err
	}

	// Create cipher and CFB with AAD then encrypt the note into dst
	blockCipher, err := aes.NewCipher(AAD)
	if err != nil {
		return nil, [32]byte{}, nil, err
	}
	cfb := cipher.NewCFBEncrypter(blockCipher, dst[:aes.BlockSize])
	cfb.XORKeyStream(dst[aes.BlockSize:], noteBytes)

	// Hash AAD - used for proper pass verification in get
	aadHash := sha3.Sum256(AAD)

	// Create blockCipher with hash of supplied password and encrypt AAD
	key := sha3.Sum256([]byte(password))
	blockCipher, err = aes.NewCipher(key[:])
	if err != nil {
		return nil, [32]byte{}, nil, err
	}
	blockCipher.Encrypt(AAD[:aes.BlockSize], AAD[:aes.BlockSize])
	blockCipher.Encrypt(AAD[aes.BlockSize:], AAD[aes.BlockSize:])

	return AAD, aadHash, dst, nil
}

// DecryptAAD checks if the data corresponds with the provided pass
// returns decrypted AAD and errors, if any
func DecryptAAD(data db.Data, pass string) ([]byte, error) {
	// Hash pass, make cipher from hash and make var for decryptedAAD output
	key := sha3.Sum256([]byte(pass))
	blockCipher, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	decryptedAAD := make([]byte, 32)
	// Decrypt AAD & Hash it
	blockCipher.Decrypt(decryptedAAD[:aes.BlockSize], data.AAD[:aes.BlockSize])
	blockCipher.Decrypt(decryptedAAD[aes.BlockSize:], data.AAD[aes.BlockSize:])
	decryptedAADHash := sha3.Sum256(decryptedAAD)

	// If hashed hash of decrypted AAD matches with stored AAD, we have a Match!
	// And yes - that is a Tinder reference thanks to @black-dragon74
	if bytes.Equal(data.Hash[:], decryptedAADHash[:]) {
		return decryptedAAD, nil
	}

	return []byte{}, errors.New(IncorrectPassError)
}

// Decrypt decrypts the provided note with the aad, returning decrypted note and errors, if any
func Decrypt(note []byte, decryptedAAD []byte) (string, error) {
	// Create new cipher with AAD
	blockCipher, err := aes.NewCipher(decryptedAAD)
	if err != nil {
		return "", err
	}

	// Create CFB Decrypter with cipher, instantiating with IV
	cfb := cipher.NewCFBDecrypter(blockCipher, note[:aes.BlockSize])
	// Create variable for storing decrypted note of shorter length taking into account IV
	decrypted := make([]byte, len(note)-aes.BlockSize)
	// Decrypt note
	cfb.XORKeyStream(decrypted, note[aes.BlockSize:])

	return string(decrypted), nil
}

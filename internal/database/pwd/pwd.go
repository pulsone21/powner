package pwd

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Argon2idHasher struct {
	// time represents the number of
	// passed over the specified memory.
	time uint32
	// cpu memory to be used.
	memory uint32
	// threads for parallelism aspect
	// of the algorithm.
	threads uint8
	// keyLen of the generate hash key.
	keyLen uint32
	// saltLen the length of the salt used.
	saltLen uint32
}

func NewArgon2idHasher() *Argon2idHasher {
	return &Argon2idHasher{
		memory:  64 * 1024,
		time:    1,
		threads: 1,
		saltLen: 16,
		keyLen:  32,
	}
}

// GenerateHash using the password and provided salt.
// If not salt value provided fallback to random value
// generated of a given length.
func (a *Argon2idHasher) GenerateHash(password, salt []byte) (*HashSalt, error) {
	var err error
	// If salt is not provided generate a salt of
	// the configured salt length.
	if len(salt) == 0 {
		salt, err = randomSecret(a.saltLen)
	}
	if err != nil {
		return nil, err
	}

	// Generate hash
	hash := argon2.IDKey(password, salt, a.time, a.memory, a.threads, a.keyLen)
	// Return the generated hash and salt used for storage.
	return &HashSalt{hash: hash, salt: salt}, nil
}

// Compare generated hash with store hash.
func (a *Argon2idHasher) Compare(hashSalt *HashSalt, password []byte) error {
	// Generate hash for comparison.
	hashSalt, err := a.GenerateHash(password, hashSalt.salt)
	if err != nil {
		return err
	}
	// Compare the generated hash with the stored hash.
	// If they don't match return error.
	if !bytes.Equal(hashSalt.hash, hashSalt.hash) {
		return errors.New("hash doesn't match")
	}
	return nil
}

type HashSalt struct {
	hash []byte
	salt []byte
}

func (h *HashSalt) String() string {
	return fmt.Sprintf("%v:%v", h.salt, h.hash)
}

func StringToHashSalt(s string) (*HashSalt, error) {
	if i := strings.Count(s, ":"); i == 1 {
		subs := strings.Split(s, ":")
		return &HashSalt{
			hash: []byte(subs[1]),
			salt: []byte(subs[0]),
		}, nil

	}
	return nil, fmt.Errorf("string is not HashSalt complaint", s)
}

func randomSecret(length uint32) ([]byte, error) {
	secret := make([]byte, length)

	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

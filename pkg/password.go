package pkg

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

/*
https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html
https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go

parameters from the article seem reasonable, using them
*/

// check argon2.IDKey for more information
type hashParametersType struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var generationHashParameters *hashParametersType = &hashParametersType{
	memory:      64 * 1024,
	iterations:  3,
	parallelism: 2,
	saltLength:  16,
	keyLength:   32,
}

func EncodedHashGenerate(password string) (string, error) {
	// generate a cryptographically secure random salt
	salt, err := generateRandomBytes(generationHashParameters.saltLength)
	if err != nil {
		return "", err
	}
	// this will generate a hash of the password using the Argon2id variant
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		generationHashParameters.iterations,
		generationHashParameters.memory,
		generationHashParameters.parallelism,
		generationHashParameters.keyLength,
	)

	// base64 encode the salt and hashed password
	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// return a string using the standard encoded hash representation
	// for more information check the article
	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		generationHashParameters.memory,
		generationHashParameters.iterations,
		generationHashParameters.parallelism,
		base64Salt,
		base64Hash,
	)

	return encodedHash, nil
}

func generateRandomBytes(length uint32) ([]byte, error) {
	bytes := make([]byte, length)
	// fill the array with random bytes
	_, err := rand.Read(bytes)
	return bytes, err
}

func EncodedHashCompare(password, encodedHash string) (
	match bool, err error,
) {
	// extract the parameters, salt and derived key from the encoded hash
	parameters, salt, hash, err := decodeEncodedHash(encodedHash)
	if err != nil {
		return false, err
	}

	// hash provided string using the same parameters
	otherHash := argon2.IDKey(
		[]byte(password),
		salt,
		parameters.iterations,
		parameters.memory,
		parameters.parallelism,
		parameters.keyLength,
	)

	// check whether the contents of the hashed passwords are identical
	// we are using subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks
	return subtle.ConstantTimeCompare(hash, otherHash) == 1, nil
}

func formatError(err error) error {
	return fmt.Errorf("the encoded hash is not in the correct format: %s", err)
}

func decodeEncodedHash(encodedHash string) (
	parameters *hashParametersType, salt, hash []byte, err error,
) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, formatError(errors.New("invalid number of '$'"))
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, formatError(err)
	}
	if version != argon2.Version {
		return nil, nil, nil, fmt.Errorf(
			"incompatible argon2id version: current - %d, desired - %d ",
			version,
			argon2.Version,
		)
	}

	parameters = &hashParametersType{}
	_, err = fmt.Sscanf(
		vals[3],
		"m=%d,t=%d,p=%d",
		&parameters.memory,
		&parameters.iterations,
		&parameters.parallelism,
	)
	if err != nil {
		return nil, nil, nil, formatError(err)
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, formatError(err)
	}
	parameters.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, formatError(err)
	}
	parameters.keyLength = uint32(len(hash))

	return parameters, salt, hash, nil
}

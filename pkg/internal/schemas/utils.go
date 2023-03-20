package schemas

import "golang.org/x/crypto/bcrypt"

func hashValue(value []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(value, 10)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func compareHashValue(value, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, value) == nil
}

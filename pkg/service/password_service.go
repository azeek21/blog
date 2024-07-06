package service

import "golang.org/x/crypto/bcrypt"

type PasswordServ struct {
}

func NewPasswordSerice() PasswordService {
	return PasswordServ{}
}

func (s PasswordServ) CreateHash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	return string(hash), err
}

func (s PasswordServ) VerifyPasswordAgainstHash(pwd string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		return false, err
	}
	return true, err
}

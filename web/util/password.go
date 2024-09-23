package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword băm mật khẩu bằng bcrypt
func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil

}

// CheckPassword kiểm tra xem mật khẩu có khớp với mật khẩu băm không
func CheckPassword(pwd string, hashedPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}

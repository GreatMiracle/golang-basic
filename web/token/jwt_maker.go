package token

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const minSecretKeySize = 32

// JWTMaker là struct chứa secret key để tạo và xác thực JWT tokens
type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (*JWTMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid secret key size: must be at least 32 characters")
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"id":         payload.ID,
		"username":   payload.Username,
		"issued_at":  payload.IssuedAt,
		"expired_at": payload.ExpiredAt,
	}

	// Tạo JWT claims dựa trên payload
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Đăng Ký token bằng secret key
	accessToken, err := jwtToken.SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", err
	}
	return accessToken, nil

}

// VerifyToken kiểm tra tính hợp lệ của JWT token
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	claims := jwt.MapClaims{}

	//jwtToken, err := jwt.ParseWithClaims(token, &claims, keyFunc)
	//myPayload, ok := jwtToken.Claims.(*Payload)
	//if ok {
	//	return nil, ErrInvalidToken
	//}

	_, err := jwt.ParseWithClaims(token, &claims, keyFunc)
	if err != nil {
		return nil, err
	}

	marshal, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	myPayload := &Payload{}
	err = json.Unmarshal(marshal, &myPayload)
	if err != nil {
		return nil, err
	}

	// Kiểm tra thời gian hết hạn
	if claims["expired_at"] != nil {
		expirationTime, ok := claims["expired_at"].(string)
		if !ok {
			return nil, ErrInvalidToken
		}

		// Chuyển đổi string sang time.Time
		expiredAtTime, err := time.Parse(time.RFC3339, expirationTime)
		if err != nil {
			return nil, ErrInvalidToken
		}

		// So sánh thời gian hiện tại với thời gian hết hạn
		if time.Now().After(expiredAtTime) {
			return nil, ErrExpiredToken
		}
	}

	return myPayload, err
}

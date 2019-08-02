package jwtauth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

var JwtConfig = Config{
	SigningMethod: "HS256",
	SigningKey: "ak19@lml",
	Expired: 7*86400,
}
type Config struct {
	SigningMethod string        `json:"SigningMethod"`
	SigningKey    string        `json:"SigningKey"`
	Expired       time.Duration `json:"Expired"`
}

func GenerateToken(userID int64) (string, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(JwtConfig.Expired) * time.Second).Unix()
	method := getSigningMethod(JwtConfig.SigningMethod)
	token := jwt.NewWithClaims(method, &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt,
		NotBefore: now.Unix(),
		Subject:   strconv.FormatInt(userID, 10),
	})
	tokenString, err := token.SignedString([]byte(JwtConfig.SigningKey))
	if err != nil {
		return "", errors.Wrap(err, "GenerateToken SignedString")
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtConfig.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, errors.New("token.Claims.(*jwt.StandardClaims) not ok")
	}
	return claims, nil
}

func getSigningMethod(method string) (m *jwt.SigningMethodHMAC) {
	switch method {
	case "HS256":
		m = jwt.SigningMethodHS256
	case "HS384":
		m = jwt.SigningMethodHS384
	case "HS512":
		m = jwt.SigningMethodHS512
	default:
		m = jwt.SigningMethodHS256
	}
	return m
}

func ChangeConfig(c Config)  {
	JwtConfig = c
}

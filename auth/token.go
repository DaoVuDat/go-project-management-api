package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	db "project-management/db/sqlc"
	"project-management/domain"
	"time"
)

type JwtCustomPayload struct {
	jwt.RegisteredClaims
	Role   string `json:"role"`
	UserId int    `json:"userId"`
}

func CreateToken(account *db.UserAccount, duration time.Duration, privateKey string) (string, *JwtCustomPayload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return "", nil, err
	}

	payload := &JwtCustomPayload{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    account.Username,
			ID:        tokenID.String(),
		},
		Role:   string(account.Type),
		UserId: int(account.UserID),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(privateKey))
	return token, payload, err
}

func VerifyToken(token string, privateKey string) (*JwtCustomPayload, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, domain.ErrInvalidToken
		}
		return []byte(privateKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &JwtCustomPayload{}, keyFunc)
	if err != nil {
		log.Println(err.Error())
		//verr, ok := err.(*jwt.ValidationError)
		//if ok && errors.Is(verr.Inner, ErrExpiredToken) {
		//	return nil, ErrExpiredToken
		//}
		return nil, domain.ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*JwtCustomPayload)
	if !ok {
		return nil, domain.ErrInvalidToken
	}
	return payload, nil
}

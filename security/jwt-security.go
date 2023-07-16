package security

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"

	"time"
)

type JWTService struct {
	Key []byte
}

func NewJWTService(key []byte) *JWTService {
	return &JWTService{
		Key: key,
	}
}

var tokenAuth *jwtauth.JWTAuth

func Init() {
	tokenAuth = jwtauth.New("HS256", []byte("try-me"), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": "23ope"})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func (js JWTService) EncryptPassword(password string) (string, error) {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(byte), nil
}

func (js JWTService) ComparePasswords(hashPass, pass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
	if err != nil {
		return false, err
	}
	return true, nil
}

type JWTClaim struct {
	Email string `json:"email"`
	Id    string `json:"email"`
	Date  jwt.StandardClaims
}

func (J JWTClaim) Valid() error {
	if J.Email == "" && J.Date.ExpiresAt == 0 {
		return errors.New("mess up with data")
	}
	return nil
}

var jwtkey = []byte("supesecretKey")

func (js JWTService) GenerateJWT(email, id string) (string, error) {
	expiryTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email: email, Id: id, Date: jwt.StandardClaims{ExpiresAt: expiryTime.Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//tokenString, err := token.SignedString(jwtkey)
	tokenString, err := token.SignedString(js.Key)
	return tokenString, err
}

func (js JWTService) ValidateToken(signedToken string) (JWTClaim, error) {
	result := JWTClaim{}
	token, err := jwt.ParseWithClaims(
		signedToken,
		&result,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(js.Key), nil
		})
	if err != nil {
		return result, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return result, err
	}
	if claims.Date.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
	}
	return result, nil
}

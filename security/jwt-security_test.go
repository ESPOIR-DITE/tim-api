package security

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	Init()
}

//func TestEncryptPassword(t *testing.T) {
//	string, err := EncryptPassword("espoir")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string)
//}

func TestComparePasswords(t *testing.T) {
	instance := NewJWTService(nil)
	result, err := instance.ComparePasswords("$2a$14$423DJLIkerROAaVns1ZvNOr7pqYk21JNoar3O98YZfZrfH.Ai9.DG", "1234")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

//func TestGenerateJWT(t *testing.T) {
//	result, err := GenerateJWT("espoir@mn", "espoir")
//	if err != nil {
//		fmt.Println("err")
//		fmt.Println(err)
//	}
//	fmt.Println(result)
//
//}
//func TestValidateToken(t *testing.T) {
//	result, err := ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImVzcG9pciIsImVtYWlsIjoiZXNwb2lyQG1uIiwiRGF0ZSI6eyJleHAiOjE2NjM0MTU1NjZ9fQ.GnIOvyQBWoscg5YQqmJYLIlOCBx6bnErFDH_pB5a-ek")
//	if err != nil {
//		fmt.Println("err")
//		fmt.Println(err)
//	}
//	fmt.Println(result)
//}

package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtClaim struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
	Exp int64 `json:"exp"`
}

func CreateToken(user_id int64, signingKey string) (string, error){
 claims := jwtClaim{
	UserID: user_id,
	Exp: time.Now().Add(time.Minute * 30).Unix(),
 }

 token := jwt.NewWithClaims(jwt.SigningMethodES256 , claims)


 tokenString,  err := token.SignedString([]byte(signingKey))

if err != nil {
	return "" , err
}
return string(tokenString) ,  nil

}

func VerifyToken(tokenString , signingKey string) (int64 , error){
	token , err := jwt.ParseWithClaims(tokenString , &jwtClaim{} , func(t *jwt.Token) (interface{} , error){
    if  _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
		return nil , fmt.Errorf("invalid authorization token")
	}
	return []byte(signingKey), nil
	
	})

	if err != nil{
		return 0 , fmt.Errorf("Invalid authentication token")
	}

	claims , ok :=  token.Claims.(*jwtClaim)

	if !ok{
		return 0 , fmt.Errorf("Invalid authentication token")
	}
	if claims.Exp < time.Now().Unix(){
		return 0 , fmt.Errorf("token has expired")
	}

	return claims.UserID , nil

}
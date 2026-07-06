package auth

import (
	"fmt"
	"gotickets/internal/domain/user/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
const (
	jwtSecretKey         = "your_secret"
	defaultDuration = 24 * time.Hour // 7 days
)

type JWTClaims struct{
	UserId uint
	Name string
	Email string
	Role types.Role
	jwt.RegisteredClaims
}


type jwtService struct{
	secretKey string
	tokenDuration time.Duration
}

type JWTService interface{
    GenerateToken(userId uint,name string,email string)(string,error)
	ValidateToken(token string)(*JWTClaims,error)
}

func NewJwtService(secret string) JWTService{
	return &jwtService{
		secretKey: secret ,
		tokenDuration: defaultDuration,
	}
} 

func (js *jwtService) GenerateToken(userId uint,name string,email string)(string,error){
	claims:=JWTClaims{
		UserId:userId,
		Name:name,
		Email: email,
		Role:types.Driver,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(js.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gotickets",
		},
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	tokenStr,err:=token.SignedString([]byte(js.secretKey))
	if err!=nil{
		return "",err
	}

	return tokenStr,nil

}

func (js *jwtService) ValidateToken(token string)(*JWTClaims,error){
	tokenStr,err:=jwt.ParseWithClaims(token,&JWTClaims{},func(token *jwt.Token)(any,error){
		if _,ok:=token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil,fmt.Errorf("unexpected signing method %v",token.Header["alg"])
		}
		return []byte(js.secretKey),nil
	})

	if err!=nil{
		return nil,fmt.Errorf("invalid token: %w",err)
	}

	if claims,ok:=tokenStr.Claims.(*JWTClaims);ok && tokenStr.Valid{
		return claims,nil
	}
	return nil,fmt.Errorf("invalid token")
}
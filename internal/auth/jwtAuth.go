package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTAuthenticator ...
type JWTAuthz interface {
	GenerateToken(userName string, isUser bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issuer    string
}

//JWTAuthService ...
func JWTAuthService() JWTAuthz {
	return &jwtServices{
		secretKey: getSecretKey(),
		issuer:    "sachin-verma",
	}
}

func getSecretKey() string {

	//To be replaced with Kubernetes secrets
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {

		//Temporary secret. Need a strong secret here
		secret = "sachinvermaandsons-secretkey"
	}
	return secret
}

//GenerateToken ...
func (service *jwtServices) GenerateToken(userName string, isUser bool) (string, error) {
	claims := &authCustomClaims{
		userName,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)

	//returning encoded string after signing
	t, err := token.SignedString([]byte(service.secretKey))
	// fmt.Println(t)
	if err != nil {
		panic(err)
	}
	return t, err
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %v", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}

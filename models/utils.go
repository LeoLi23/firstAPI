package models

import(
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWT : HEADER PAYLOAD SIGNATURE
const (
	SecretKEY              string = "JWT-Secret-Key"
	DEFAULT_EXPIRE_SECONDS int    = 600 // default expired 10 minutes
	PasswordHashBytes             = 16
)

// This struct is the payload
type MyCustomClaims struct {
	UserID int `json:"UserID"`
	jwt.StandardClaims
}

// This struct is the parsing of token payload
type JwtPayload struct {
	Username  string `json:"Username"`
	UserID    int    `json:"UserID"`
	IssuedAt  int64  `json:"Iat"`
	ExpiresAt int64  `json:"Exp"`
}

//generate token
func GenerateToken(loginInfo *LoginRequest, userID int, expiredSeconds int) (tokenString string, err error) {
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}

	// Create the Claims
	mySigningKey := []byte(SecretKEY)
	expireAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	logs.Info("Token will be expired at ", time.Unix(expireAt, 0))

	user := *loginInfo
	claims := MyCustomClaims{
		userID,
		jwt.StandardClaims{
			Issuer:    user.Username,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireAt,
		},
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Signs the token with a secret
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", errors.New("error: failed to generate token")
	}

	return tokenStr, nil
}

//validate token
func ValidateToken(tokenString string) (*JwtPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKEY), nil
		})

	claims, ok := token.Claims.(*MyCustomClaims)
	if ok && token.Valid {
		logs.Info("%v %v", claims.UserID, claims.StandardClaims.ExpiresAt)
		logs.Info("Token will be expired at ", time.Unix(claims.StandardClaims.ExpiresAt, 0))

		return &JwtPayload{
			Username:  claims.StandardClaims.Issuer,
			UserID:    claims.UserID,
			IssuedAt:  claims.StandardClaims.IssuedAt,
			ExpiresAt: claims.StandardClaims.ExpiresAt,
		}, nil
	} else {
		fmt.Println(err)
		return nil, errors.New("error: failed to validate token")
	}
}
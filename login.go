package main

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)
//Token Claim
type Claim struct {
	UserName	string
	Password	string
	jwt.StandardClaims
}
//Login api request body
type loginReqBody struct {
	UserName string `json:"username"`
	Password	string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request){

	//get request body into loginreqbody struct
	var reqBody loginReqBody
	err := BodyToJson(r, &reqBody)
	if err != nil {
		response(nil, 500, w)
		return
	}

	//Username & password check
	if (reqBody.UserName != username) || (reqBody.Password != password) {
		response(nil, 401, w)
		return
	}

	//set token claim struct
	expirationTime := time.Now().Add(5 * time.Hour)
	jwtData := Claim{
		UserName:       "ahmetgul",
		Password:       "deneme123",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtData)
	stringToken, err := token.SignedString([]byte(secret))
	if err != nil {
		response(nil, 500, w)
		return
	}

	//response data include token for login api's response
	respData := struct {
		UserName string
		Token string
	}{
		UserName: "ahmetgul",
		Token: stringToken,
	}

	//end of api response
	err = response(respData, 200, w)
	if err != nil {
		response(nil, 500, w)
		return
	}
}

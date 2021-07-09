package main

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"strings"
)

//general response function
func response(responseData interface{}, statusCode int, w http.ResponseWriter) error {
	//response data
	var respData struct{
		Success		bool			//is api finished succesfully?
		Data 		interface{}		//spesific response data
	}

	//status code check for respData.Success
	if statusCode == 200 {
		respData.Success = true
	} else {
		respData.Success = false
	}
	respData.Data = responseData

	//marshall the spesific response data
	jsonResponse, err := json.Marshal(respData)
	if err != nil {
		return err
	}

	//set response header and body
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(jsonResponse)
	return err
}

//for read request body and set to spesific struct
func BodyToJson(r *http.Request, data interface{}) error {

	//read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	//unmarshall the bıdy into our spesific reqBody
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	return nil
}

//sunfunction that returns secret key as []byte array
func JwtKeyGetter (token *jwt.Token) (interface{}, error){
	return []byte(secret), nil
}

//token validation check
func TokenIsValid(r *http.Request) bool {
	//read token from request header
	reqToken := r.Header.Get("Authorization")

	//read token include "Bearer" so lets exclude that
	splitToken := strings.Split(reqToken, "Bearer ")

	//parse token with secret_key
	token, err := jwt.Parse(splitToken[1], JwtKeyGetter)
	if err == nil && token.Valid {
		return true
	} else {
		return false
	}
}
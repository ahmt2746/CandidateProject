package main

import (
	"net/http"
)

//this api returns user ınformation
func UsrInfo (w http.ResponseWriter, r *http.Request) {

	//request's token check
	if !TokenIsValid(r){
		response(nil, 401, w)
		return
	}

	//response data
	respData := struct {
		UserName	string
		FirstName	string
		LastName	string
	}{
		username,
		firstname,
		lastname,
	}

	response(respData, 200, w)

}

//this api changes usr password
func UpdateUsrPassword (w http.ResponseWriter, r *http.Request) {

	//request's token check
	if !TokenIsValid(r){
		response(nil, 401, w)
		return
	}

	//get request body into passwordreqbody struct
	var reqBody struct{
		Password	string `json:"password"`
	}
	BodyToJson(r, &reqBody)

	//update usr password
	password = reqBody.Password

	response(w, 200, w)
}

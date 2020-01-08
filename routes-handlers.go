package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func getUsers(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Error Occured",
	}
	jsonResponse := getUsersFromDB()

	if jsonResponse == nil {
		returnErrorResponse(response, request, httpError)
	} else {
		response.Header().Set("Content-Type", "application/json")
		response.Write(jsonResponse)
	}
}


func registerUser(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Error Occured",
	}
	var userDetails Userdetails
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userDetails)
	defer request.Body.Close()
	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if userDetails.Username == "" {
			httpError.Message = "First Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if userDetails.FirstName == "" {
			httpError.Message = "Last Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if userDetails.LastName == "" {
			httpError.Message = "Country can't be empty"
			returnErrorResponse(response, request, httpError)
		}else if userDetails.Password == "" {
			httpError.Message = "Last Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else {
			isInserted := registerUserInDB(userDetails)
			if isInserted {
				getUsers(response, request)
			} else {
				returnErrorResponse(response, request, httpError)
			}
		}
	}
}
func deleteUser(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Error Occured",
	}
	userID := mux.Vars(request)["username"]
	if userID == "" {
		httpError.Message = "Username can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {
		isdeleted := deleteUserFromDB(userID)
		if isdeleted {
			getUsers(response, request)
		} else {
			returnErrorResponse(response, request, httpError)
		}
	}
}


func loginUser(response http.ResponseWriter, request *http.Request){
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Error Occured",
	}
	var logDetails LoginUser
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&logDetails)
	defer request.Body.Close()
	if err != nil {
		returnErrorResponse(response, request, httpError)
	}else {
		httpError.Code = http.StatusBadRequest
		if logDetails.Username == "" {
			httpError.Message = "UserName can't be empty"
			returnErrorResponse(response, request, httpError)
		}else if logDetails.Password == "" {
			httpError.Message = "password can't be empty"
			returnErrorResponse(response, request, httpError)
		} else {
			isLogin := LoginUserFromDB(logDetails)
			if isLogin == nil {
				returnErrorResponse(response, request, httpError)
			} else {
				response.Header().Set("Content-Type", "application/json")
				response.Write(isLogin)
			}
		}
	}
}
func validateToken(response http.ResponseWriter, request *http.Request){
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Error Occured",
	}
	userID := mux.Vars(request)["token"]
	if userID == "" {
		httpError.Message = "Token can't be empty"
		returnErrorResponse(response, request, httpError)
	} else {
		istoken := checkTokenFromDB(userID)
		if istoken == nil{
			returnErrorResponse(response, request, httpError)
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.Write(istoken)
		}
	}
}

func updateUser(response http.ResponseWriter, request *http.Request) {
	var httpError = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Error Occured",
	}
	var userDetails Userdetails
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userDetails)
	defer request.Body.Close()
	if err != nil {
		returnErrorResponse(response, request, httpError)
	} else {
		httpError.Code = http.StatusBadRequest
		if userDetails.FirstName == "" {
			httpError.Message = "First Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if userDetails.Username == "" {
			httpError.Message = "username can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if userDetails.LastName == "" {
			httpError.Message = "Last Name can't be empty"
			returnErrorResponse(response, request, httpError)
		} else if userDetails.Password == "" {
			httpError.Message = "Password can't be empty"
			returnErrorResponse(response, request, httpError)
		} else {
			isUpdated := updateUserInDB(userDetails)
			if isUpdated {
				getUsers(response, request)
			} else {
				returnErrorResponse(response, request, httpError)
			}
		}
	}
}

func returnErrorResponse(response http.ResponseWriter, request *http.Request, errorMesage ErrorResponse) {
	httpResponse := &ErrorResponse{Code: errorMesage.Code, Message: errorMesage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMesage.Code)
	response.Write(jsonResponse)
}
package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
func String(length int) string {
	return StringWithCharset(7, charset)
}
func getUsersFromDB() []byte {
	var (
		user  Userdetails
		users []Userdetails
	)
	rows, err := db.Query("SELECT username,first_name,last_name,token,password FROM users")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&user.Username,&user.FirstName,&user.LastName,&user.Token,&user.Password)
		users = append(users, user)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(users)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}
func LoginUserFromDB(log LoginUser) []byte {
	var lg LoginUser
	var users[] string
	rows, err := db.Query("SELECT username from users where username=? and password=?",log.Username,log.Password)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&lg.Username)
		users = append(users,lg.Username)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(users)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	if jsonResponse !=nil {
		stmt, err := db.Prepare("DELETE FROM login_tokens WHERE username=?")
		if err != nil {
			fmt.Println(err)
			return jsonResponse
		}
		_, queryError := stmt.Exec(lg.Username)
		if queryError != nil {
			fmt.Println(queryError)
			return jsonResponse
		}
		stmt, err = db.Prepare("INSERT into login_tokens SET username=?,token=?")
		if err != nil {
			fmt.Println(err)
			return jsonResponse
		}
		token_string := generateRandomToken()
		_, queryError = stmt.Exec(lg.Username, token_string)
		if queryError != nil {
			fmt.Println(queryError)
			return jsonResponse
		}
	}
	return jsonResponse
}

func registerUserInDB(userDetails Userdetails) bool {
	stmt, err := db.Prepare("INSERT into users SET first_name=?,last_name=?,username=?,password=?,token=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	token_string:=generateRandomToken()
	_, queryError := stmt.Exec(userDetails.FirstName, userDetails.LastName, userDetails.Username, userDetails.Password,token_string)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}
func generateRandomToken()string{
	return String(4)
}

func deleteUserFromDB(userID string) bool {
	stmt, err := db.Prepare("DELETE FROM users WHERE username=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userID)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}
func checkTokenFromDB(token string) []byte {
	var (
		logindetails Login_token
		users []Login_token
	)

	rows, err := db.Query("select username,token FROM login_tokens WHERE token=?",token)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&logindetails.Username,&logindetails.Token)
		users = append(users, logindetails)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(users)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}


func updateUserInDB(userDetails Userdetails) bool {
	stmt, err := db.Prepare("UPDATE users SET first_name=?,last_name=?,password=? WHERE username=?")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, queryError := stmt.Exec(userDetails.FirstName, userDetails.LastName, userDetails.Password, userDetails.Username)
	if queryError != nil {
		fmt.Println(queryError)
		return false
	}
	return true
}

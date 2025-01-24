package api

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// api endpoints handler funcs

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	db.NamedExec("INSERT INTO users (username, password) VALUES (:username, :password)", &user)
	json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	var dbUser User
	db.Get(&dbUser, "SELECT * FROM users WHERE username=$1", user.Username)
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token, _ := GenerateJWT(dbUser.Username, dbUser.ID)
	json.NewEncoder(w).Encode(JWTToken{Token: token})
}

func CreateProperty(w http.ResponseWriter, r *http.Request) {
	var property Property
	_ = json.NewDecoder(r.Body).Decode(&property)
	claims, err := ValidateToken(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	property.UserID = claims.UserID
	db.NamedExec("INSERT INTO properties (title, description, price, user_id) VALUES (:title, :description, :price, :user_id)", &property)
	json.NewEncoder(w).Encode(property)
}

func GetProperties(w http.ResponseWriter, r *http.Request) {
	var properties []Property
	db.Select(&properties, "SELECT * FROM properties")
	json.NewEncoder(w).Encode(properties)
}

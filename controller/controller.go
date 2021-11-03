package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	config "github.com/nicolascancino/web-service-go/jwt"
	"github.com/nicolascancino/web-service-go/models"
	"github.com/nicolascancino/web-service-go/repository"
)

func HolaMundo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	fmt.Fprint(w, "hello world")
}

func Registro(w http.ResponseWriter, r *http.Request) {

	var user models.Usuario

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error decoding body %v"+err.Error(), http.StatusBadRequest)
		//fmt.Fprintf(w, "Error decoding Body %v", err.Error())
		return
	}

	if err := user.Validation(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.Header().Add("Content-type", "application/json")
		return
	}

	/*_, founded, _ := bd.CheckIfExistUser(user.Email)
	if founded {
		http.Error(w, "This email alredy has beed used", http.StatusBadRequest)
		return
	}
	_, status, err := bd.CheckIfExistUser(user.Email)
	if err != nil {
		http.Error(w, "Has been happened an error inserting the user "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "", http.StatusBadRequest)
	}*/
	w.WriteHeader(http.StatusCreated)

	//w.Header().Add("Content-type", "application/json")
	//fmt.Fprint(w, "hello world")
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	var user *models.Usuario

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "User Or password invalid"+err.Error(), http.StatusBadRequest)
	}

	document, exist := repository.Login(user.Email, user.Password)

	if !exist {
		http.Error(w, "User Or password invalid"+err.Error(), http.StatusBadRequest)
	}

	jwtKey, err := config.GenerateJWT(&document)
	if err != nil {
		http.Error(w, "JWT has not been generated"+err.Error(), http.StatusBadRequest)
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}

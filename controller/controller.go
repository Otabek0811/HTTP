package controller

import (
	"app/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type Controller struct {
	
}

var users []models.User = []models.User{
	{
		Id:       "5fde231f-29eb-469c-93a1-ca7f4c16fcfe",
		FullName: "Peter Parker",
		Login:    "peter",
		Password: "admin",
	},
	{
		Id:       "7fde231f-29eb-469c-93a1-ca7f4c16fcfe",
		FullName: "John Wick",
		Login:    "wick",
		Password: "1234567",
	},
}

func (c *Controller) Get_User_by_id(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	arr := strings.Split(url, "/")
	fmt.Println(url)
	fmt.Println(arr)
	fmt.Println(len(arr))

	if arr[len(arr)-1] != "" {
		id := arr[len(arr)-1]

		for _, user := range users {
			if user.Id == id {
				data, err := json.Marshal(user)
				if err != nil {
					w.Write([]byte(err.Error()))
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				w.Write(data)
				w.WriteHeader(http.StatusOK)
				return
			}
		}

		w.Write([]byte("Not found user with this id"))
		w.WriteHeader(http.StatusNotFound)
	}
}

func Get_Users(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(users)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func Update_User(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	arr := strings.Split(url, "/")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error while ioutil:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("error while user update unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if arr[len(arr)-1] != "" {
		id := arr[len(arr)-1]
		for ind, val := range users {
			if val.Id == id {
				users[ind].FullName = user.FullName
				users[ind].Login = user.Login
				users[ind].Password = user.Password
			}
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(body)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func Create_User(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error while ioutil:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("error while user unmarshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Id = uuid.New().String()
	users = append(users, user)

	body, err = json.Marshal(user)
	if err != nil {
		log.Println("error while user create marshal:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(body)
	return
}

func Delete_User(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	arr := strings.Split(url, "/")
	if arr[len(arr)-1] != "" {
		id := arr[len(arr)-1]
		for ind, user := range users {
			if user.Id == id {
				users = append(users[:ind], users[ind+1:]...)
				return
			}
		}
		w.Write([]byte("Not found user with this id"))
		w.WriteHeader(http.StatusNotFound)
		return
	}

}

package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"strconv"

	"net/http"

	models "github.com/MikeCodeSun/go-mux-api/model"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type App struct {
	Db *sql.DB
	Router *mux.Router
}

var validate = validator.New()

func RwE(err error, w http.ResponseWriter) {
  jsonData, _ := json.Marshal(err)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonData)

}

func(a *App) HomePage(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
  var user models.User
	// get input from req.body and pass value to user
	err := json.NewDecoder(r.Body).Decode(&user)
	// fmt.Println(user)
  defer r.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		RwE(err,w)
		return
	}

	// validata input user
	err = validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
		RwE(err,w)
		return
	}


	// insert user data to database & if err handle err
	err = user.CreateUser(a.Db)
	if err != nil {
		fmt.Println(err.Error())
		RwE(err,w)
		return
	}

  // response with new user json data & if err handle err
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println(err.Error())
		RwE(err,w)
		return
	}

}

func(a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// get user id from route path params user mux
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	user.Id = id
	if err != nil {
		fmt.Println(err.Error())
		RwE(err,w)
		return
	}
//  get user data use id from db
	err = user.GetUser(a.Db)
	if err != nil {
		fmt.Println(err.Error())
		RwE(err,w)
		return
	}
	// response with user data user data
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println(err.Error())
		RwE(err,w)
		return
	}
}


func(a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
// get user id from route path params
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err.Error())
		RwE(err,w)
		return
	}
	user.Id = id
	// delete user from db
	err = user.DeleteUser(a.Db)
	if err != nil {
		fmt.Println(err.Error())
		RwE(err,w)
		return
	}
//res with msg
fmt.Fprintf(w, "User delete OK!")
}

func(a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
// get id
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
  if err != nil {
		fmt.Println(err.Error())
		RwE(err,w)
		return
	}
err = json.NewDecoder(r.Body).Decode(&user)
if err != nil {
	fmt.Println(err.Error())
	RwE(err,w)
	return
}
user.Id = id
fmt.Println(user)
defer r.Body.Close()
// update to db
err = user.UpdateUser(a.Db)
if err != nil {
	fmt.Println(err.Error())
	RwE(err,w)
	return
}
// res with new data
fmt.Fprintf(w, "Update Ok!")
}

func(a *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	// get all users data from database
	users, err := models.GetUsers(a.Db)
	if err != nil {
		fmt.Println(err.Error())
		RwE(err,w)
		return
	}
	// res with users
  jsonData,err := json.Marshal(users)
	if err != nil {
		fmt.Println(err.Error())
		RwE(err,w)
		return
	}
	fmt.Fprintf(w, string(jsonData))
}
// connect db 
func(a *App) Initilize(host, port, user, password, dbName string) {
	var err error
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	a.Db, err = sql.Open("postgres", dbInfo)
	// defer a.Db.Close()
	if err != nil {
		log.Fatal(err.Error())
	} 
	
	a.Router = mux.NewRouter()

	a.AppRouters()
}

func(a *App)  AppRouters() {
	
	a.Router.HandleFunc("/", a.HomePage).Methods("GET")
	a.Router.HandleFunc("/user/create", a.CreateUser).Methods("POST")
	a.Router.HandleFunc("/user/all", a.GetUsers).Methods("GET")
	a.Router.HandleFunc("/user/{id}", a.GetUser).Methods("GET")
	a.Router.HandleFunc("/user/{id}", a.DeleteUser).Methods("DELETE")
	a.Router.HandleFunc("/user/{id}", a.UpdateUser).Methods("PATCH")
	
}
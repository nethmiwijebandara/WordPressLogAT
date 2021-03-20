package controller

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"../config/db"
	"../model"
	"../unzip"

	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

//Function for unzip the files and pass data to the DB
func FileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var res model.ResponseResult

	files, err := unzip.Unzip(r)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		f, err := os.Open(file)

		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)

		doc := make(map[string]interface{})
		var i int = 1
		for scanner.Scan() {
			s := "line" + " " + strconv.Itoa(i)
			i += 1
			fmt.Println(s, scanner.Text())
			doc[s] = scanner.Text()
		}
		collectionUnzip, err := db.GetDBCollectionUnzip("docs")
		if err != nil {
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}
		_, err = collectionUnzip.InsertOne(context.TODO(), doc)

		i = 1
		//Remove the file
		os.Remove(file)

		if err != nil {
			res.Error = "Error While Creating  Doc, Try Again"
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Result = "Insert Successful"
		json.NewEncoder(w).Encode(res)
		return
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	//	db.GetDBCollectionLogin("loginDetails")
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("%#v\n", string(body))
	err := json.Unmarshal(body, &user)
	var res model.ResponseResult
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	collectionLogin, err := db.GetDBCollectionLogin("loginList")

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	// var user model.User
	err = collectionLogin.FindOne(context.TODO(), bson.D{{"email", user.Email}}).Decode(&user)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

			if err != nil {
				res.Error = "Error While Hashing Password, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}
			user.Password = string(hash)

			_, err = collectionLogin.InsertOne(context.TODO(), user)
			if err != nil {
				res.Error = "Error While Creating User, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}
			res.Result = "Registration Successful"
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "Username already Exists!!"
	json.NewEncoder(w).Encode(res)
	return
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// //	w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "POST")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var login model.Login
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &login)
	if err != nil {
		log.Fatal(err)
	}
	// log.Fatal(body, "body")
	collection, err := db.GetDBCollectionLogin("loginList")
	// log.Fatal(collection, "body")
	if err != nil {
		log.Fatal(err)
	}
	var user model.User
	var res model.ResponseResult
	fmt.Println(user.Email)

	err = collection.FindOne(context.TODO(), bson.D{{"email", login.Email}}).Decode(&user)
	// log.Fatal("err")
	if err != nil {
		res.Error = "Invalid email"
		json.NewEncoder(w).Encode(res)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))

	if err != nil {
		res.Error = "Invalid password"
		json.NewEncoder(w).Encode(res)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//	"username":  result.Username,
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"role":      user.Role,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		res.Error = "Error while generating token,Try again"
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "Login Successful"
	login.Token = tokenString
	json.NewEncoder(w).Encode(login)

}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	var result model.User
	var res model.ResponseResult
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result.FirstName = claims["firstname"].(string)
		result.LastName = claims["lastname"].(string)
		result.Role = claims["role"].(string)

		json.NewEncoder(w).Encode(result)
		return
	} else {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

}

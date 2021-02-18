package controller

import (
	"../config/db"
	"../model"
	"../unzip"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"log"
	"net/http"
)

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
			s:= "line" + " " + strconv.Itoa(i)
			i+=1
			fmt.Println(s,scanner.Text())
			doc[s] =  scanner.Text()
		}
		collection, err := db.GetDBCollection("docs")
		if err != nil {
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}
		_, err = collection.InsertOne(context.TODO(), doc)

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



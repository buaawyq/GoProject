package main

import (
	"GoProject/Domain"
	"GoProject/MySql"
	"GoProject/Redis"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./frontend/index.html")
	//dir,_:= os.Getwd()
	//print(dir)
	t.Execute(w, nil)
}

func Article(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./frontend/vote.html")
	t.Execute(w, nil)
}

func Vote(writer http.ResponseWriter, request *http.Request) {
	article_id := request.PostFormValue("articleid")
	//redis
	client_redis := Redis.GetRedisClient()
	client_redis.SAdd(article_id, "user: 17033")

}
func Login(writer http.ResponseWriter, request *http.Request) {
	//vars:=request.URL.Query()
	//id,name:=request.URL.Query()["id"],request.URL.Query()["name"]
	var number, username sql.NullString
	fmt.Println("handler begin....")
	bT := time.Now()
	user := request.PostFormValue("user")
	//redis
	client_redis := Redis.GetRedisClient()
	res, err := client_redis.Get(user).Result()
	if err != redis.Nil {
		eT := time.Since(bT)
		//JsonToStruct
		var card Domain.Card
		json.Unmarshal([]byte(res), &card)
		number, username = card.Number, card.User
		fmt.Println(eT, card)
	} else {
		//mysql
		DB := MySql.InitDataBase()
		card := MySql.QueryOne(DB, user)
		//StructToJson
		cardJson, _ := json.Marshal(card)
		client_redis.Set(user, cardJson, 1000*time.Second)
		eT := time.Since(bT)
		number, username = card.Number, card.User
		fmt.Println(eT, card)
	}

	// 设置Content-Type
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	data := map[string]sql.NullString{
		"number":   number,
		"userName": username,
	}
	json.NewEncoder(writer).Encode(data)
}

func DownLoad(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("download begin...")
	path := "E:\\golang\\message.txt"
	//open file
	msgTxt, _ := os.Open(path)
	defer func() {
		msgTxt.Close()
	}()

	msg, err := ioutil.ReadAll(msgTxt)
	if err != nil {
		return
	}
	fmt.Println(string(msg))
	//request.FormFile()
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("content-disposition", "attchment;filename=message.txt")
	writer.Write(msg)
}

func main() {
	mux := http.NewServeMux()
	//mux.HandleFunc("/",handler)
	mux.HandleFunc("/index", Index)
	mux.Handle("/bootstrap-3.3.7-dist/", http.FileServer(http.Dir("frontend")))
	mux.HandleFunc("/api/login", Login)
	mux.HandleFunc("/api/download", DownLoad)
	mux.HandleFunc("/article", Article)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

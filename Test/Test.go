package main

func main() {
	/*//IO
		PATH:="C:\\Users\\wyq\\go\\src\\GoWeb\\resources\\img_wechat.jpg"
	    f,err:=ioutil.ReadFile(PATH)
	    if err!=nil{
	    	log.Println(err)
		}

		//DB
		DB:=MySql.InitDataBase()
		str:="insert into test (name,img) values (?,?)"
		if _,err:=DB.Exec(str,"wyq",f);err!=nil{
			log.Println(err)
		}*/
	/*op:=&[]byte{}
	DB:=MySql.InitDataBase()
	str:="select img from test where id='7'"
	res:=DB.QueryRow(str)
	res.Scan(&op)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//writer.Header().Set("Content-Type","application/x-download;image/jpeg")
		writer.Header().Set("Content-Disposition", "attachment;filename=wang.jpg")
		writer.Write(*op)
	})
	http.ListenAndServe(":8080",nil)*/
}

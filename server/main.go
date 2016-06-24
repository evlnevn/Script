package main

import (
	"absensi-element/server/handler"
	"absensi-element/server/serversql"
	"absensi-element/session"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	serversql.OpenDB()
	http.HandleFunc("/test", index)
	http.Handle("/", http.FileServer(http.Dir("../app")))
	http.HandleFunc("/login", loginpost)
	http.HandleFunc("/get-user/", handler.GetUser)
	fmt.Println("server run localhost:3000")
	http.ListenAndServe(":3000", nil)

}
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("a"))
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Print("login")
	if r.Method == http.MethodPost {
		loginpost(w, r)
		return

	}
	loginget(w, r)
}

func loginpost(w http.ResponseWriter, r *http.Request) {
	if http.MethodPost == r.Method {
		r.ParseForm()
		name := r.FormValue("name")
		password := r.FormValue("password")
		user := new(serversql.User)
		err := user.GetBy(name)
		if err != nil {
			ResponseJson(w, http.StatusForbidden, []byte("user tidak terdaftar"))
			return
		}
		fmt.Println(password, name)
		if user.ID == 0 {
			ResponseJson(w, http.StatusForbidden, []byte("ID tidak ada"))
			return
		}
		if user.Password != password {
			ResponseJson(w, http.StatusForbidden, []byte("password salah"))
			return
		}
		s := session.Session{
			User:  user.ID,
			Waktu: time.Now(),
		}
		err = s.Insert()
		if err != nil {
			fmt.Print("pp", err)
			ResponseJson(w, http.StatusForbidden, []byte("gagal memasukkan session"))
			return
		}
		w.WriteHeader(http.StatusOK)

	}

}

func loginget(w http.ResponseWriter, r *http.Request) {
}

func ResponseJsonData(w http.ResponseWriter, data interface{}) {
	w.Header().Set("content-Type", "application/json")
	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
	}
	w.WriteHeader(http.StatusOK)

	w.Write(b)
}

func ResponseJson(w http.ResponseWriter, status int, err []byte) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(err)
}

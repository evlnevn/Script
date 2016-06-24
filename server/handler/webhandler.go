package handler

import (
	"absensi-element/server/serversql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	host := r.Host
	id := strings.Replace(url, host+"/get-user/", "", -1)
	u := serversql.User{}
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response(nil, http.StatusNoContent, w, err)
		return
	}
	err = u.Get(id64)
	if err != nil {
		response(nil, http.StatusNoContent, w, err)
		return
	}

	data, err := json.Marshal(u)
	if err != nil {
		response(nil, http.StatusNoContent, w, err)
		return
	}
	response(data, http.StatusOK, w, err)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if http.MethodPost != r.Method {
		err := fmt.Errorf("method not available")
		response(nil, http.StatusForbidden, w, err)
		return
	}
	r.ParseForm()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response(nil, http.StatusNoContent, w, err)
		return
	}
	user := serversql.User{}
	fmt.Print(string(b))
	err = json.Unmarshal(b, &user)
	if err != nil {
		response(nil, http.StatusNoContent, w, err)
		return
	}
	err = user.Insert()
	if err != nil {
		response(nil, http.StatusNoContent, w, err)
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		response(nil, http.StatusNoContent, w, err)
		return
	}
	response(data, http.StatusOK, w, err)
}

func response(data []byte, status int, w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil {
		log.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(data)
}

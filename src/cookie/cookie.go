package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func DrawMenu(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, "<a href='/'>HOME <ba><br/>"+"\n")
	io.WriteString(w, "<a href='/read'>Read Cookie <ba><br/>"+"\n")
	io.WriteString(w, "<a href='/write'>Write Cookie <ba><br/>"+"\n")
	io.WriteString(w, "<a href='/delete'>Delete Cookie <ba><br/>"+"\n")
}

func IndexServer(w http.ResponseWriter, req *http.Request) {
	DrawMenu(w)
}

func WriteCookieServer(w http.ResponseWriter, r *http.Request) {
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "rancongjie", Value: "rancongjiecookie", Expires: expire, MaxAge: 86400}

	http.SetCookie(w, &cookie)

	DrawMenu(w)
}
func ReadCookieServer(w http.ResponseWriter, r *http.Request) {
	DrawMenu(w)
	var cookie, err = r.Cookie("rancongjie")
	if err == nil {
		var cookievalue = cookie.Value
		io.WriteString(w, "<b>get cookie value is "+cookievalue+"</b>\n")
	}
}

func DeleteCookieServer(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "rancongjie", Value: "rancongjiecookie", Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)

	DrawMenu(w)
}

func main() {
	http.HandleFunc("/", IndexServer)
	http.HandleFunc("/write", WriteCookieServer)
	http.HandleFunc("/read", ReadCookieServer)
	http.HandleFunc("/delete", DeleteCookieServer)
	fmt.Println("listen on 9090")
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}

}

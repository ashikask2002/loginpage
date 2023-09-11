package main

import (
	"html/template"
	"net/http"
	
)


var tmpl *template.Template

type UserSt struct{
	Username string
	Password string
}

var users=make(map[string]UserSt)
var session=make(map[string]string)

type errors struct {
	UsernameError string
	PasswordError string
} 


func init() {
	tmpl = template.Must(template.ParseGlob("templates/*"))

	users["ashu@gmail.com"] = UserSt{"ashu","1234"}
	users["ashi@gmail.com"] = UserSt{"ashi","1234"}
	users["ikru@gmail.com"]= UserSt{"ikru","1234"}

}


func main() {

	http.HandleFunc("/", loginPage)
	http.HandleFunc("/home", home)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8000", nil)
}

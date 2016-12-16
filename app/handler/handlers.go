package handler

import (
	"net/http"
	"encoding/json"
	"github.com/jaymoulin/http/app/helper"
	"github.com/gorilla/sessions"
	"html/template"
)

var templateHtml *template.Template
var store *sessions.CookieStore

func init() {
	templateHtml = template.Must(template.ParseGlob("./app/view/*.tmpl"))
	store = sessions.NewCookieStore([]byte("something-very-secret"))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var b struct{
		Err string
	}
	b.Err = r.FormValue("err")
	templateHtml.ExecuteTemplate(w, "home", b)
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	var g struct{UserList []helper.User}
	g.UserList = helper.GetUserList()
	templateHtml.ExecuteTemplate(w, "products", g)
}
func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Articles !\n"))
}
func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Validate !\n"))
}
func ResultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Result !\n"))
}

func ChoicesHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	var u []helper.User
	if _, ok := session.Values["logged"]; ok {
		u = helper.GetUserList()
	} else {
		u = make([]helper.User, 1, 1)
		u[0] = helper.User{}
	}
	res, _ := json.Marshal(u)
	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	b := false
	for _,u := range helper.GetUserList() {
		if u.Password == r.FormValue("passwd") && u.Email == r.FormValue("login") {
			session, _ := store.Get(r, "session-name")
			session.Values["logged"] = u
			session.Save(r, w)
			http.Redirect(w, r, "/choices", 303)
			b = true
		}
	}
	if (!b) {
		http.Redirect(w, r, "/?err=1", 303)
	}
}
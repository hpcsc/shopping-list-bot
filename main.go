package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"io"
	"fmt"
	"strings"
	"net/url"
	"io/ioutil"
)

func main() {
	port := "8080"
	appId := ""
	appPassword := ""

	fmt.Print("App id: " + appId)
	fmt.Print("App password: " + appPassword)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)

	response := GetAccessToken(appId, appPassword)
	print(response)

	fmt.Print("Listenning at " + port)
	http.ListenAndServe(":" + port, r)
}

func GetAccessToken(appId string, appPassword string) string {
	address := "https://login.microsoftonline.com/botframework.com/oauth2/v2.0/token"
	hc := http.Client{}

	form := url.Values{}
	form.Add("grant_type", "client_credentials")
	form.Add("client_id", appId)
	form.Add("client_secret", appPassword)
	form.Add("scope", "https://api.botframework.com/.default")
	req, err := http.NewRequest("POST", address, strings.NewReader(form.Encode()))
	if err != nil {
		print("Failed to create request", err)
		return ""
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := hc.Do(req)
	if err != nil {
		print("Failed to post: ", err)
		return ""
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	return string(bodyBytes)
}

func ProductsHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Product")
}

func ArticlesHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Articles")
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Home")
}
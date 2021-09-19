package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/dghubble/oauth1"
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	client *http.Client
	cfg    *config
}

type secretKeys struct {
	consumerKey    string
	consumerSecret string
	accessToken    string
	tokenSecret    string
	bearerToken    string
}

type config struct {
	twitterUrlV2    string
	twitterUrlV1    string
	getAllTweetPath string
	createTweetPath string
	deleteTweetPath string
}

func Init() *Server {
	r := mux.NewRouter()
	secretKeys := &secretKeys{
		consumerKey:    os.Getenv("CONSUMER_KEY"),
		consumerSecret: os.Getenv("CONSUMER_SECRET"),
		accessToken:    os.Getenv("ACCESS_TOKEN"),
		tokenSecret:    os.Getenv("TOKEN_SCERET"),
		bearerToken:    os.Getenv("BEARER_TOKEN"),
	}

	oauthConfig := oauth1.NewConfig(secretKeys.consumerKey, secretKeys.consumerSecret)
	token := oauth1.NewToken(secretKeys.accessToken, secretKeys.tokenSecret)

	httpClient := oauthConfig.Client(oauth1.NoContext, token)

	s := &Server{
		router: r,
		client: httpClient,
		cfg: &config{
			twitterUrlV2:    "api.twitter.com/2/users/",
			twitterUrlV1:    "api.twitter.com",
			getAllTweetPath: "1.1/statuses/home_timeline.json",
			createTweetPath: "1.1/statuses/update.json",
			deleteTweetPath: "1.1/statuses/destroy/%s.json",
		},
	}

	r.HandleFunc("/tweet/{id}", s.getTweet).Methods("GET")
	r.HandleFunc("/tweet", s.getAllTweets).Methods("GET")
	r.HandleFunc("/tweet", s.createTweet).Methods("POST")
	r.HandleFunc("/tweet/{id}", s.deleteTweet).Methods("DELETE")
	return s

}

func (s *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), s.router))
}

func (s *Server) getTweet(w http.ResponseWriter, r *http.Request) {
	jsonResponse, err := json.Marshal("success")
	status := http.StatusOK

	if err != nil {
		status = http.StatusInternalServerError
		jsonResponse = []byte("fail")
	}

	w.WriteHeader(status)
	w.Write(jsonResponse)
}

func (s *Server) getAllTweets(w http.ResponseWriter, r *http.Request) {

	url := url.URL{
		Scheme: "https",
		Host:   s.cfg.twitterUrlV1,
		Path:   s.cfg.getAllTweetPath,
	}

	resp, err := s.client.Get(url.String())
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (s *Server) createTweet(w http.ResponseWriter, r *http.Request) {

	var data map[string]string

	_ = json.NewDecoder(r.Body).Decode(&data)
	url := url.URL{
		Scheme:   "https",
		Host:     s.cfg.twitterUrlV1,
		Path:     s.cfg.createTweetPath,
		RawQuery: fmt.Sprintf("status=%v", url.QueryEscape(data["status"])),
	}

	res, err := s.client.Post(url.String(), "application/json", strings.NewReader(``))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))

}

func (s *Server) deleteTweet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	url := url.URL{
		Scheme: "https",
		Host:   s.cfg.twitterUrlV1,
		Path:   fmt.Sprintf(s.cfg.deleteTweetPath, params["id"]),
	}

	res, err := s.client.Post(url.String(), "application/json", strings.NewReader(``))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}

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
	"github.com/gorilla/handlers"
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

type Tweet struct {
	CreatedAt string `json:"created_at"`
	ID        int64  `json:"id"`
	IDStr     string `json:"id_str"`
	Text      string `json:"text"`
	Truncated bool   `json:"truncated"`
	Entities  struct {
		Hashtags     []interface{} `json:"hashtags"`
		Symbols      []interface{} `json:"symbols"`
		UserMentions []interface{} `json:"user_mentions"`
		Urls         []interface{} `json:"urls"`
	} `json:"entities"`
	Source               string      `json:"source"`
	InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
	InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
	User                 struct {
		ID          int64       `json:"id"`
		IDStr       string      `json:"id_str"`
		Name        string      `json:"name"`
		ScreenName  string      `json:"screen_name"`
		Location    string      `json:"location"`
		Description string      `json:"description"`
		URL         interface{} `json:"url"`
		Entities    struct {
			Description struct {
				Urls []interface{} `json:"urls"`
			} `json:"description"`
		} `json:"entities"`
		Protected                      bool          `json:"protected"`
		FollowersCount                 int           `json:"followers_count"`
		FriendsCount                   int           `json:"friends_count"`
		ListedCount                    int           `json:"listed_count"`
		CreatedAt                      string        `json:"created_at"`
		FavouritesCount                int           `json:"favourites_count"`
		UtcOffset                      interface{}   `json:"utc_offset"`
		TimeZone                       interface{}   `json:"time_zone"`
		GeoEnabled                     bool          `json:"geo_enabled"`
		Verified                       bool          `json:"verified"`
		StatusesCount                  int           `json:"statuses_count"`
		Lang                           interface{}   `json:"lang"`
		ContributorsEnabled            bool          `json:"contributors_enabled"`
		IsTranslator                   bool          `json:"is_translator"`
		IsTranslationEnabled           bool          `json:"is_translation_enabled"`
		ProfileBackgroundColor         string        `json:"profile_background_color"`
		ProfileBackgroundImageURL      interface{}   `json:"profile_background_image_url"`
		ProfileBackgroundImageURLHTTPS interface{}   `json:"profile_background_image_url_https"`
		ProfileBackgroundTile          bool          `json:"profile_background_tile"`
		ProfileImageURL                string        `json:"profile_image_url"`
		ProfileImageURLHTTPS           string        `json:"profile_image_url_https"`
		ProfileLinkColor               string        `json:"profile_link_color"`
		ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string        `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool          `json:"profile_use_background_image"`
		HasExtendedProfile             bool          `json:"has_extended_profile"`
		DefaultProfile                 bool          `json:"default_profile"`
		DefaultProfileImage            bool          `json:"default_profile_image"`
		Following                      bool          `json:"following"`
		FollowRequestSent              bool          `json:"follow_request_sent"`
		Notifications                  bool          `json:"notifications"`
		TranslatorType                 string        `json:"translator_type"`
		WithheldInCountries            []interface{} `json:"withheld_in_countries"`
	} `json:"user"`
	Geo           interface{} `json:"geo"`
	Coordinates   interface{} `json:"coordinates"`
	Place         interface{} `json:"place"`
	Contributors  interface{} `json:"contributors"`
	IsQuoteStatus bool        `json:"is_quote_status"`
	RetweetCount  int         `json:"retweet_count"`
	FavoriteCount int         `json:"favorite_count"`
	Favorited     bool        `json:"favorited"`
	Retweeted     bool        `json:"retweeted"`
	Lang          string      `json:"lang"`
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
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	headersOK := handlers.AllowedHeaders([]string{"Content-Type"})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(headersOK, methods, origins)(s.router)))
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
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp != nil && resp.StatusCode != http.StatusOK {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
	}

	w.Header().Add("content-type", "application/json")
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

	resp, err := s.client.Post(url.String(), "application/json", strings.NewReader(``))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp != nil && resp.StatusCode != http.StatusOK {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
	}

	w.Header().Add("content-type", "application/json")
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

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/json"
	"github.com/disgoorg/disgo/oauth2"
	"github.com/disgoorg/disgo/rest"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake"
)

var (
	letters      = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	clientID     = snowflake.GetSnowflakeEnv("client_id")
	clientSecret = os.Getenv("client_secret")
	baseURL      = os.Getenv("base_url")
	logger       = log.Default()
	httpClient   = http.DefaultClient
	client       oauth2.Client
)

func main() {
	rand.Seed(time.Now().UnixNano())

	logger.SetLevel(log.LevelDebug)
	logger.Info("starting example...")
	logger.Infof("disgo %s", disgo.Version)

	client = oauth2.New(clientID, clientSecret, oauth2.WithLogger(logger), oauth2.WithRestClientConfigOpts(rest.WithHTTPClient(httpClient)))

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/trylogin", handleTryLogin)
	_ = http.ListenAndServe(":6969", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	var body string
	cookie, err := r.Cookie("token")
	if err == nil {
		session := client.SessionController().GetSession(cookie.Value)
		if session != nil {
			var user *discord.OAuth2User
			user, err = client.GetUser(session)
			if err != nil {
				writeError(w, "error while getting user data", err)
				return
			}
			var userJSON []byte
			userJSON, err = json.MarshalIndent(user, "<br />", "&ensp;")
			if err != nil {
				writeError(w, "error while formatting user data", err)
				return
			}

			var connections []discord.Connection
			connections, err = client.GetConnections(session)
			if err != nil {
				writeError(w, "error while getting connections data", err)
				return
			}
			var connectionsJSON []byte
			connectionsJSON, err = json.MarshalIndent(connections, "<br />", "&ensp;")
			if err != nil {
				writeError(w, "error while formatting connections data", err)
				return
			}
			body = fmt.Sprintf("user:<br />%s<br />connections: <br />%s", userJSON, connectionsJSON)
		}
	}
	if body == "" {
		body = `<button><a href="/login">login</a></button>`
	}
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte(body))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, client.GenerateAuthorizationURL(baseURL+"/trylogin", discord.PermissionsNone, "", false, discord.ApplicationScopeIdentify, discord.ApplicationScopeGuilds, discord.ApplicationScopeEmail, discord.ApplicationScopeConnections, discord.ApplicationScopeWebhookIncoming), http.StatusMovedPermanently)
}

func handleTryLogin(w http.ResponseWriter, r *http.Request) {
	var (
		query = r.URL.Query()
		code  = query.Get("code")
		state = query.Get("state")
	)
	if code != "" && state != "" {
		identifier := randStr(32)
		_, err := client.StartSession(code, state, identifier)
		if err != nil {
			writeError(w, "error while starting session", err)
			return
		}
		http.SetCookie(w, &http.Cookie{Name: "token", Value: identifier})
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func writeError(w http.ResponseWriter, text string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(text + ": " + err.Error()))
}

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

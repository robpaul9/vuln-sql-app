package server

import (
	"fmt"
	"net/http"

	"github.com/robpaul9/vulnsqlapp/adapters/records"

	log "github.com/robpaul9/golog"
)

type Config struct {
	DDAgentHost    string
	DDAgentPort    string
	ServiceName    string
	ServicePort    string
	Logger         log.Logger
	RecordsService records.Service
}

type Server struct {
	*Config
}

func New(config *Config) *Server {
	return &Server{
		Config: config,
	}
}

func (s *Server) Start() {
	http.HandleFunc("/", s.MessageHandler)

	address := fmt.Sprintf("0.0.0.0:%s", s.ServicePort)
	s.Logger.Infof("application is serving on %s", address)

	s.Logger.Fatal(http.ListenAndServe(address, nil))

}

func (s *Server) MessageHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		http.ServeFile(w, r, "/static/form.html")
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			s.Logger.Error(err)
			return
		}

		formEmail := r.FormValue("email")

		r, err := s.RecordsService.GetUser(formEmail)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			s.Logger.Error(err)
			return
		}

		if len(r) == 0 {
			http.Error(w, "user not found", http.StatusUnauthorized)
			s.Logger.Warn("user %s not found", formEmail)
			return
		}

		fmt.Fprintf(w, "welcome. Here is your user information...\n\r")

		for _, item := range r {
			fmt.Fprintf(w, "Email: %s\n\r", item.Email)
			fmt.Fprintf(w, "ID: %v\n\r", item.ID)
			fmt.Fprintf(w, "Role: %s\n\r", item.Role)
			fmt.Fprintf(w, "Account created on: %v\n\r\n", item.CreatedAt)
		}

	} else {
		http.Error(w, "only GET and POST requests allowed", http.StatusMethodNotAllowed)

	}
}

package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"regexp"
	"net/url"
	"fmt"
)

func Start(listenPort int, tourPort int, gamblerPort int, scorePort int, collectorPort int) error {
	server, err := NewServer()
	if err != nil {
		log.Fatal(err)
	}

	server.useRules( []Rule {
		{
			IfMatches:"/api/tour",
			Forward:fmt.Sprintf("localhost:%d",tourPort),
		},
		{
			IfMatches:"/api/gambler",
			Forward:fmt.Sprintf("localhost:%d",gamblerPort),
		},
		{
			IfMatches:"/api/score",
			Forward:fmt.Sprintf("localhost:%d",scorePort),
		},
		{
			IfMatches:"/admin/events",
			Forward:fmt.Sprintf("localhost:%d",collectorPort),
		},
		{
			IfMatches:"/ui",
			Serve:"../ui/",
		},
	})

	return http.ListenAndServe(fmt.Sprintf(":%d",listenPort), nil)
}

func NewServer() (*Server, error) {
	s := new(Server)
	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := s.handler(r); handler != nil {
		log.Printf( "Serve %s", r.RequestURI)
		handler.ServeHTTP(w, r)
		return
	}
	log.Printf( "Not found %s", r.RequestURI)

	http.Error(w, "Not found.", http.StatusNotFound)
}

func (s *Server) handler(req *http.Request) http.Handler {
	log.Printf( "Got %s", req.RequestURI)
	for _, rule := range s.rules {
	    matched,_ := regexp.MatchString(rule.IfMatches, req.RequestURI)
		if matched {
			log.Printf( "Matched %s -> forward: %s, serve:%s", rule.IfMatches, rule.Forward, rule.Serve)
			return rule.handler
		}
	}
	return nil
}


func (s *Server) useRules( rules []Rule) {
	s.rules = rules
	for _,rule := range rules {
		log.Printf( "Rule for %s -> forward: %s, serve:%s", rule.IfMatches, rule.Forward, rule.Serve)
		rule.handler = makeHandler(rule)
	}
}

func makeHandler(r Rule) http.Handler {
	if forw := r.Forward; forw != "" {
		log.Printf( "Proxy for %s -> forward: %s", r.IfMatches, r.Forward)
		return &httputil.ReverseProxy{
				Director: func(req *http.Request) {
				u,_ := url.Parse(forw)
				req.URL.Scheme = "http"
				req.URL.Host = u.Host
				req.URL.Path = u.Path
			},
		}
	}
	if dir := r.Serve; dir != "" {
		log.Printf( "Proxy for %s -> serve: %s", r.IfMatches, r.Serve)
		return http.FileServer(http.Dir(dir))
	}
	return nil
}

type Server struct {
	rules []Rule
}

type Rule struct {
	IfMatches string // to match against request Host header
	Forward string // non-empty if reverse proxy
	Serve   string // non-empty if file server

	handler http.Handler
}


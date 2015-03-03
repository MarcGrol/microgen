package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

func Start(baseDir string, listenPort int, tourPort int, gamblerPort int, scorePort int, collectorPort int) error {
	var err error
	server := newServer()

	targetHost := "localhost"

	server.addForwardRule("/api/tour", fmt.Sprintf("%s:%d", targetHost, tourPort))
	server.addForwardRule("/api/gambler", fmt.Sprintf("%s:%d", targetHost, gamblerPort))
	server.addForwardRule("/api/score", fmt.Sprintf("%s:%d", targetHost, scorePort))
	server.addForwardRule("/admin/events", fmt.Sprintf("%s:%d", targetHost, collectorPort))
	server.addServeRule("/static", fmt.Sprintf("%s/tourApp/ui/", baseDir))
	if server.err != nil {
		log.Printf("Error registrering handler %s", server.err)
		return err
	}

	return server.listenAndServe(listenPort)
}

type server struct {
	err   error
	rules []*rule
}

func newServer() *server {
	s := new(server)
	s.err = nil
	s.rules = make([]*rule, 0, 10)
	return s
}

func (s *server) addForwardRule(urlPattern string, targetHost string) {
	if s.err == nil {
		_, s.err = url.Parse(targetHost)
		if s.err != nil {
			log.Printf("Error parsing url %s:%s", targetHost, s.err)
		}

		s.rules = append(s.rules, newForwardRule(urlPattern, targetHost))
	}
}

func (s *server) addServeRule(urlPattern string, dir string) {
	if s.err == nil {
		s.rules = append(s.rules, newServeRule(urlPattern, dir))
	}
}

type rule struct {
	urlPattern string // to match against requestUri
	forward    string // non-empty if reverse proxy
	serve      string // non-empty if file server

	handler http.Handler
}

func newForwardRule(urlPattern string, targetHost string) *rule {
	rule := new(rule)
	rule.urlPattern = urlPattern
	rule.forward = targetHost
	log.Printf("Create proxy for %s -> forward: %s", rule.urlPattern, rule.forward)
	rule.handler = &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = rule.forward
			log.Printf("Forward %s to: %s", req.RequestURI, req.URL.String())
		},
	}

	return rule
}

func newServeRule(urlPattern string, targetDir string) *rule {
	rule := new(rule)
	rule.urlPattern = urlPattern
	rule.serve = targetDir
	log.Printf("Create file server for %s -> serve: %s", rule.urlPattern, rule.serve)
	rule.handler = http.StripPrefix(urlPattern, http.FileServer(http.Dir(rule.serve)))

	return rule
}

func (s *server) handlerForRequest(req *http.Request) http.Handler {
	for _, rule := range s.rules {
		matched, _ := regexp.MatchString(rule.urlPattern, req.RequestURI)
		if matched {
			//log.Printf("Matched %s -> forward: %s, serve:%s", rule.urlPattern, rule.forward, rule.serve)
			return rule.handler
		}
	}
	log.Printf("%s not matched", req.RequestURI)
	return nil
}

func (s *server) listenAndServe(listenPort int) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler := s.handlerForRequest(r)
		if handler == nil {
			http.Error(w, "Not found.", http.StatusNotFound)
		}
		log.Printf("Found %s", r.RequestURI)
		handler.ServeHTTP(w, r)
	})
	return http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil)
}
package apiserver

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"stats/internal/client/asocks"
	"stats/internal/client/clouvider"
	"stats/internal/client/dataimpulse"
	"stats/internal/client/sms_activate"
	"stats/internal/model"
	"sync"
)

type server struct {
	logger *zap.Logger
	router *http.ServeMux
	cv     *clouvider.Cloud
	smsAct *sms_activate.SmsActivate
	asocks *asocks.Asocks
	di     *dataimpulse.DataImpulse
	//store  *sqlstore.Store
}

func newServer(c *Config) *server {
	s := &server{
		router: http.NewServeMux(),
		cv:     clouvider.New(c.Clouvider.ApiKey),
		smsAct: sms_activate.New(c.SmsActivate.ApiKey, c.ProxyUrl),
		asocks: asocks.New(c.Asocks.ApiKey, c.ProxyUrl),
		di:     dataimpulse.New(c.DataImpulse.ApiKey),
		//logger:       zap.New(),
	}

	s.configureRouter()

	return s
}

func (s *server) configureRouter() {
	//s.router.Use(s.setRequestID)
	//s.router.Use(s.logRequest)
	//s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.HandleFunc("GET /stats", s.allStats())
	//s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")
	//
	//private := s.router.PathPrefix("/private").Subrouter()
	//private.Use(s.authenticateUser)
	//private.HandleFunc("/whoami", s.handleWhoami())
}

func (s *server) allStats() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		stat := model.Stat{}
		lock := sync.Mutex{}
		wg := sync.WaitGroup{}
		wg.Add(4)
		go func() {
			ans := s.cv.Stat()
			lock.Lock()
			stat.Clouvider = ans
			defer wg.Done()
			defer lock.Unlock()
		}()
		go func() {
			ans := s.smsAct.Stat()
			lock.Lock()
			stat.Smsactivate = ans
			defer wg.Done()
			defer lock.Unlock()
		}()
		go func() {
			ans := s.asocks.Stat()
			lock.Lock()
			stat.Asocks = ans
			defer wg.Done()
			defer lock.Unlock()
		}()
		go func() {
			ans := s.di.Stat()
			lock.Lock()
			stat.Dataimpulse = ans
			defer wg.Done()
			defer lock.Unlock()
		}()
		wg.Wait()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stat)
	}
}

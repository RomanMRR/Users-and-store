package apiserver

import (
	"encoding/json"
	"http-rest-api/internal/app/model"
	"http-rest-api/internal/app/store"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		var result = make(map[string]interface{})
		result["result"] = data
		json.NewEncoder(w).Encode(result)
	}
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
		})
		logger.Infof("started %s", r.Method)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (s *server) configureRouter() {
	s.router.Use(s.logRequest)
	s.router.HandleFunc("/create_user", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/create_shop", s.handleStoresCreate()).Methods("POST")
	s.router.HandleFunc("/update_user", s.handleUsersUpdate()).Methods("PUT")
	s.router.HandleFunc("/update_shop", s.handleShopsUpdate()).Methods("PUT")
	s.router.HandleFunc("/delete_user", s.handleUsersDelete()).Methods("DELETE")
	s.router.HandleFunc("/delete_shop", s.handleShopsDelete()).Methods("DELETE")
	s.router.HandleFunc("/find_user", s.handleUserFind()).Methods("GET")
	s.router.HandleFunc("/find_shop", s.handleShopFind()).Methods("GET")
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic"`
		Age        int16  `json:"age"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Name:              req.Name,
			Surname:           req.Surname,
			Patronymic:        req.Patronymic,
			Age:               req.Age,
			Registration_date: time.Now(),
		}

		if err := s.store.GetRepository().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleStoresCreate() http.HandlerFunc {
	type request struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Working bool   `json:"working"`
		Owner   string `json:"owner"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		shop := &model.Shop{
			Name:    req.Name,
			Address: req.Address,
			Working: req.Working,
			Owner:   req.Owner,
		}

		if err := s.store.GetRepository().Create(shop); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, shop)
	}
}

func (s *server) handleUsersUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input = &model.UpdateUserInput{}
		if err := json.NewDecoder(r.Body).Decode(input); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.GetRepository().Update(input); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, "ok")
	}
}

func (s *server) handleShopsUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input = &model.UpdateShopInput{}
		if err := json.NewDecoder(r.Body).Decode(input); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.GetRepository().Update(input); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, "ok")
	}
}

func (s *server) handleUsersDelete() http.HandlerFunc {
	type request struct {
		ID int
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var id int = req.ID

		if err := s.store.GetRepository().Delete(id, model.UserTable); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, "deleted user "+strconv.Itoa(id))
	}
}

func (s *server) handleShopsDelete() http.HandlerFunc {
	type request struct {
		ID int
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		var id int = req.ID

		if err := s.store.GetRepository().Delete(id, model.ShopTable); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, "deleted shop "+strconv.Itoa(id))
	}
}

func (s *server) handleUserFind() http.HandlerFunc {
	type request struct {
		Surname string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		surname := req.Surname

		u, err := s.store.GetRepository().Find(surname, "users")
		if err != nil {
			s.error(w, r, http.StatusServiceUnavailable, err)
			return
		}
		s.respond(w, r, http.StatusOK, u)

	}
}

func (s *server) handleShopFind() http.HandlerFunc {
	type request struct {
		Name string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		name := req.Name

		shop, err := s.store.GetRepository().Find(name, "shops")
		if err != nil {
			s.error(w, r, http.StatusServiceUnavailable, err)
			return
		}
		s.respond(w, r, http.StatusOK, shop)

	}
}

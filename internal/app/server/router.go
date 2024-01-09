package server

import (
	"encoding/json"
	"fmt"
	"l0_ms/internal/dao"
	"log"
	"net/http"
)

type Config struct {
	Port int `yaml:"port"`
}

type HttpRouter struct {
	port   int
	server *http.ServeMux
	client *dao.Client
}

func NewHttpRouter(config Config, client *dao.Client) *HttpRouter {
	serveMux := http.NewServeMux()
	fs := http.FileServer(http.Dir("web"))
	serveMux.Handle("/", fs)
	httpRouter := HttpRouter{
		port:   config.Port,
		server: serveMux,
		client: client,
	}
	serveMux.HandleFunc("/order", GetOrder(&httpRouter))
	return &httpRouter
}

func (r *HttpRouter) Start() error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", r.port), r.server)
	if err != nil {
		return err
	}
	return nil
}

func GetOrder(httpRouter *HttpRouter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		orderUid := r.URL.Query().Get("order_uid")
		id, err := httpRouter.client.GetOrder(orderUid)
		if err != nil {
			resp := make(map[string]string)
			resp["error"] = "order not found"
			err = json.NewEncoder(w).Encode(resp)
			if err != nil {
				http.NotFound(w, r)
				log.Printf("error json encode: %v", err)
				return
			}
			return
		}

		err = json.NewEncoder(w).Encode(id)
		if err != nil {
			log.Printf("error json encode: %v", err)
			http.NotFound(w, r)
		}
	}
}


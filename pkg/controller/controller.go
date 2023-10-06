package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/notEpsilon/shorty/pkg/service"
)

type Map map[string]any

type Controller struct {
	service.ShortyService

	Host string
	Port int
}

func NewController(host string, port int, service service.ShortyService) *Controller {
	return &Controller{
		ShortyService: service,
		Host:          host,
		Port:          port,
	}
}

func registerHandlers(c *Controller) {
	// Register the `/shorten` handler.
	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		url := r.Form.Get("url")
		if url == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "err: you need to provide a `url`")
			return
		}
		shortUrl, err := c.ShortyService.Shorten(url)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "err: couldn't shortent the provided `url`: %s\n", err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Map{
			"slug": shortUrl.Slug,
		})
	})

	// Register the `/r` (redirect) handler.
	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		slug := r.Form.Get("s")
		if slug == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "err: you need to provide a `slug` in the `s` query param")
			return
		}
		if err := c.ShortyService.Redirect(slug, w, r); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "err: couldn't redirect to that slug (invalid slug): %s\n", err.Error())
			return
		}
	})
}

func (c *Controller) Start() error {
	registerHandlers(c)
	log.Printf("started listening on %s:%d...\n", c.Host, c.Port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", c.Host, c.Port), nil)
}

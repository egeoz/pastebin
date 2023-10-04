package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pastebin/src/dto"
	"pastebin/src/logger"
	"pastebin/src/repos"
	"pastebin/src/tools"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

type Controller struct {
	repository repos.Repository
	host       string
	port       int
}

func NewController(r repos.Repository, h string, p int) *Controller {
	return &Controller{repository: r, host: h, port: p}
}

func (c *Controller) GetEntry(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	entry, err := c.repository.GetEntry(r.Context(), uuid)
	if err != nil {
		http.Error(w, "entry not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entry)
}

func (c *Controller) CreateEntry(w http.ResponseWriter, r *http.Request) {
	var req dto.Entry
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Title == "" {
		http.Error(w, "parameter \"title\" is mandatory", http.StatusBadRequest)
		return
	}
	if req.Content == "" {
		http.Error(w, "parameter \"content\" is mandatory", http.StatusBadRequest)
		return
	}
	if req.ContentType == "" {
		http.Error(w, "parameter \"content_type\" is mandatory", http.StatusBadRequest)
		return
	}
	if !tools.CheckIfValidContentType(req.ContentType) {
		http.Error(w, "parameter \"content_type\" has invalid value", http.StatusBadRequest)
		return
	}
	if req.Encrypted == nil {
		http.Error(w, "parameter \"is_encrypted\" is mandatory", http.StatusBadRequest)
		return
	}

	req.UUID = tools.GenerateUUID()
	uuid, err := c.repository.CreateEntry(r.Context(), req)
	if err != nil {
		http.Error(w, "failed to create entry", http.StatusInternalServerError)
		return
	}

	logger.Log.Debug(fmt.Sprintf("Created entry with the UUID: %s", uuid))

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"uuid": uuid,
	})
}

func (c *Controller) Serve() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(httprate.LimitByIP(50, 10*time.Second))
	router.Use(middleware.Timeout(30 * time.Second))

	router.Get("/api/entries/{uuid}", c.GetEntry)
	router.Post("/api/entries/new", c.CreateEntry)

	address := fmt.Sprintf("%s:%d", c.host, c.port)
	logger.Log.Info(fmt.Sprintf("Serving at %s", address))
	http.ListenAndServe(address, router)
}

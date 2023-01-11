package event

import (
	"encoding/json"
	"github.com/s1ovac/dev11/internal/handlers"
	"github.com/s1ovac/dev11/internal/storage"
	"log"
	"net/http"
)

var _ handlers.Handler = &handler{}

type handler struct {
	cache *storage.Storage
}

func NewHandler() handlers.Handler {
	return &handler{
		cache: storage.NewCache(),
	}
}

func (h *handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/create_event", h.CreateEvent)
	mux.HandleFunc("/update_event", h.UpdateEvent)
	mux.HandleFunc("/delete_event", h.DeleteEvent)
	mux.HandleFunc("/events_for_day", h.EventForDay)
	mux.HandleFunc("/events_for_week", h.EventForWeek)
	mux.HandleFunc("/events_for_month", h.EventForMonth)
}
func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
	}
	data, err := json.Marshal(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	w.Write(data)
}

func (h *handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
	}
	data, err := json.Marshal(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	w.Write(data)
}

func (h *handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
	}
	data, err := json.Marshal(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	w.Write(data)
}

func (h *handler) EventForDay(w http.ResponseWriter, r *http.Request) {
	//event, err := h.cache.Get()

}

func (h *handler) EventForWeek(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	w.Write(data)
}

func (h *handler) EventForMonth(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	w.Write(data)
}

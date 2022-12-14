package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type Event struct {
	User  int       `json:"user"`
	Event int       `json:"event"`
	Title string    `json:"title"`
	Info  string    `json:"Info"`
	Date  time.Time `json:"date"`
}

func (e *Event) Decode(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&e)
	if err != nil {
		return err
	}

	return nil
}

type Store struct {
	mu       *sync.Mutex
	stevents map[int][]Event
}

func (s *Store) Create(e *Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if stevents, ok := s.stevents[e.User]; ok {
		for _, stevent := range stevents {
			if stevent.Event == e.Event {
				return fmt.Errorf("event already created")
			}
		}
	}

	s.stevents[e.User] = append(s.stevents[e.User], *e)

	return nil
}

func (s *Store) Update(e *Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := -1

	stevents := make([]Event, 0)
	ok := false

	if stevents, ok = s.stevents[e.User]; !ok {
		return fmt.Errorf("no such user")
	}

	for idx, stevent := range stevents {
		if stevent.Event == e.Event {
			index = idx
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("no event for user")
	}

	s.stevents[e.User][index] = *e

	return nil
}

func (s *Store) Delete(e *Event) (*Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := -1

	stevents := make([]Event, 0)
	ok := false

	if stevents, ok = s.stevents[e.User]; !ok {
		return nil, fmt.Errorf("no such user")
	}

	for idx, stevent := range stevents {
		if stevent.Event == e.Event {
			index = idx
			break
		}
	}

	if index == -1 {
		return nil, fmt.Errorf("no event for user")

	}

	eventsLength := len(s.stevents[e.User])
	deletedEvent := s.stevents[e.User][index]
	s.stevents[e.User][index] = s.stevents[e.User][eventsLength-1]
	s.stevents[e.User] = s.stevents[e.User][:eventsLength-1]
	return &deletedEvent, nil
}

func (s *Store) GetEventsForDay(userID int, date time.Time) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Event

	stevents := make([]Event, 0)
	ok := false

	if stevents, ok = s.stevents[userID]; !ok {
		return nil, fmt.Errorf("user does not exists")
	}

	for _, stevent := range stevents {
		if stevent.Date.Year() == date.Year() && stevent.Date.Month() == date.Month() && stevent.Date.Day() == date.Day() {
			result = append(result, stevent)
		}
	}

	return result, nil
}

func (s *Store) GetEventsForWeek(userID int, date time.Time) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Event

	stevents := make([]Event, 0)
	ok := false

	if stevents, ok = s.stevents[userID]; !ok {
		return nil, fmt.Errorf("user does not exists")
	}

	for _, stevent := range stevents {
		y1, w1 := stevent.Date.ISOWeek()
		y2, w2 := date.ISOWeek()
		if y1 == y2 && w1 == w2 {
			result = append(result, stevent)
		}
	}

	return result, nil
}

func (s *Store) GetEventsForMonth(userID int, date time.Time) ([]Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Event

	stevents := make([]Event, 0)
	ok := false

	if stevents, ok = s.stevents[userID]; !ok {
		return nil, fmt.Errorf("user does not exists")
	}

	for _, stevent := range stevents {
		if stevent.Date.Year() == date.Year() && stevent.Date.Month() == date.Month() {
			result = append(result, stevent)
		}
	}

	return result, nil
}

// middleware to log requests.
type Logger struct {
	handler http.Handler
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handler: handlerToWrap}
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

const dateLayout = "2020-12-01"

var storage Store = Store{stevents: make(map[int][]Event), mu: &sync.Mutex{}}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/events_for_day", EventsForDayHandler) //GET
	mux.HandleFunc("/events_for_week", EventsForWeekHandler)
	mux.HandleFunc("/events_for_month", EventsForMonthHandler)
	mux.HandleFunc("/create_event", CreateEventHandler) //POST
	mux.HandleFunc("/update_event", UpdateEventHandler)
	mux.HandleFunc("/delete_event", DeleteEventHandler)

	wrappedMux := NewLogger(mux)

	port := ":8080"
	func() {
		temp := os.Getenv("PORT")
		if temp != "" {
			port = temp
		}
	}()

	log.Printf("Server is listening for incoming requests at: %v", port)
	log.Fatalln(http.ListenAndServe(port, wrappedMux))
}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var e Event

	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := storage.Create(&e); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Event has been created!", []Event{e}, http.StatusCreated)

	fmt.Println(storage.stevents)
}

func UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	var e Event

	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := storage.Update(&e); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Event has been updated!", []Event{e}, http.StatusOK)

	fmt.Println(storage.stevents)
}

func DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var e Event

	if err := e.Decode(r.Body); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var deletedEvent *Event
	var err error
	if deletedEvent, err = storage.Delete(&e); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Event has been deleted!", []Event{*deletedEvent}, http.StatusOK)
}

func EventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var stevents []Event
	if stevents, err = storage.GetEventsForDay(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Request has been executed!", stevents, http.StatusOK)
}

func EventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var stevents []Event
	if stevents, err = storage.GetEventsForWeek(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Request has been executed!", stevents, http.StatusOK)
}

func EventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	date, err := time.Parse(dateLayout, r.URL.Query().Get("date"))
	if err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	var stevents []Event
	if stevents, err = storage.GetEventsForMonth(userID, date); err != nil {
		errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultResponse(w, "Request has been executed!", stevents, http.StatusOK)
}

func errorResponse(w http.ResponseWriter, e string, status int) {
	errorResponse := struct {
		Error string `json:"error"`
	}{Error: e}

	js, err := json.Marshal(errorResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func resultResponse(w http.ResponseWriter, r string, e []Event, status int) {
	resultResponse := struct {
		Result string  `json:"result"`
		Events []Event `json:"stevents"`
	}{Result: r, Events: e}

	js, err := json.Marshal(resultResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

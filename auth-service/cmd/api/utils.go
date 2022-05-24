package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) readJson(w http.ResponseWriter, r *http.Request, data any) error {
	maxSize := int64(1024 * 1024 * 10) // 10MB
	r.Body = http.MaxBytesReader(w, r.Body, maxSize)
	doc := json.NewDecoder(r.Body)
	err := doc.Decode(data)
	if err != nil {
		return err
	}
	err = doc.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("Body must have only a single JSON object")
	}
	return nil
}

func (app *Config) writeJson(w http.ResponseWriter,status int, data any, headers ...http.Header) error{
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	if(len(headers)>0){	
		for k, v := range headers[0] {
			w.Header().Set(k, v[0])
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
	return nil
}

func (app *Config) errorJson(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if(len(status) > 0){
		statusCode = status[0]
	}
	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()
	return app.writeJson(w, statusCode, payload)
}

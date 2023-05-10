package main

import (
	"log-service/data"
	"net/http"

	"github.com/tsawler/toolbox"
)

var tools toolbox.Tools

type jsonPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {

	var requestPayload jsonPayload

	_ = tools.ReadJSON(w, r, &requestPayload)

	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		tools.ErrorJSON(w, err)
		return
	}

	resp := toolbox.JSONResponse{
		Error:   false,
		Message: "logged",
	}

	tools.WriteJSON(w, http.StatusAccepted, resp)
}

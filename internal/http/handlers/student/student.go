package studentpkg

import (
	"log/slog"
	"net/http"
)

func New() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Healthy like hanumaji!!"))
		slog.Info("Creating Student")
		
	}
}

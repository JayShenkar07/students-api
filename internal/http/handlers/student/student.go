package studentpkg

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/JayShenkar07/students-api/internal/types"
	"github.com/JayShenkar07/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//serializing the request - put it in struct
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			slog.Error("Failed to create student")
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		slog.Info("Creating Student")
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
		slog.Info("Student Created")
	}
}

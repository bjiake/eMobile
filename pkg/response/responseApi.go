package responseApi

import (
	"eMobile/internal/domain/models/response"
	"fmt"
	"github.com/go-chi/render"
	"net/http"
)

func WriteJson(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	render.Status(r, status)
	render.JSON(w, r, response.Success{
		Data: data,
	})
}

func WriteError(w http.ResponseWriter, r *http.Request, status int, err interface{}) {
	render.Status(r, status)

	var response map[string]interface{}

	// for map validator errors
	if errMap, ok := err.(map[string]string); ok {
		response = map[string]interface{}{
			"error": errMap,
		}
	} else {
		// for simple errors
		response = map[string]interface{}{
			"error": fmt.Sprintf("%v", err),
		}
	}

	render.JSON(w, r, response)
}

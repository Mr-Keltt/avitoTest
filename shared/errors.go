package shared

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrResponse - structure for describing errors.
type ErrResponse struct {
	Err            error `json:"-"` // internal error field
	HTTPStatusCode int   `json:"-"` // HTTP status code

	StatusText string `json:"status"`          // short description of the error
	ErrorText  string `json:"error,omitempty"` // detailed error description
}

// Render allows you to use ErrResponse with chi/render.
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// Err Invalid Request creates an Err Response for invalid requests.
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Invalid request",
		ErrorText:      err.Error(),
	}
}

// ErrInternal creates an ErrResponse for internal server errors.
func ErrInternal(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Internal server error",
		ErrorText:      err.Error(),
	}
}

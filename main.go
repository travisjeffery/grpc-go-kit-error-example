package example

import (
	"encoding/json"
	"log"
	"net/http"

	"context"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	example "github.com/travisjeffery/grpc-go-kit-error-example"
)

type UserService interface {
	GetUser(*User) (*User, error)
}

type userService struct{}

var (
	ErrNotFound   = &Error{Message: "not found", Code: http.StatusNotFound}
	ErrBadRequest = &Error{Message: "bad request", Code: http.StatusBadRequest}
)

func (userService) GetUser(u *User) (*User, error) {
	if u.Name == "" {
		return u, ErrBadRequest
	}
	return u, nil
}

func main() {
	svc := userService{}

	getUserHandler := httptransport.NewServer(
		makeGetUserEndpoint(svc),
		decodeGetUserRequest,
		encodeResponse,
	)

	http.Handle("/user", getUserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func makeGetUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		u, err := svc.GetUser(req.User)
		if err != nil {
			return getUserResponse{Error: err.(*Error)}, nil
		}
		return getUserResponse{User: u}, nil
	}
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type errorer interface {
	GetError() *example.Error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok {
		if err := e.GetError(); err != nil {
			w.WriteHeader(int(err.Code))
		}
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type getUserRequest struct {
	User *User `json:"user"`
}

type getUserResponse struct {
	User  *User  `json:"user,omitempty"`
	Error *Error `json:"error,omitempty"`
}

package domain

import (
	"context"
	"net/http"
)

type (
	User struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}

	UserEntity struct {
		ID       string
		Username string
		Password string
	}

	UserRepository interface {
		FetchByUsername(ctx context.Context, username string) (*UserEntity, error)
	}

	UserService interface {
		FetchByUsername(ctx context.Context, username string) (*User, error)
	}

	UserHandler interface {
		FetchByUsername() http.HandlerFunc
	}
)

package user

import (
	"context"
	"database/sql"
	"example/domain"
)

type repository struct {
	db *sql.DB
}

func (r *repository) FetchByUsername(ctx context.Context, username string) (*domain.UserEntity, error) {
	panic("implement me")
}

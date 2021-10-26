package user

import (
	"context"

	"app_microservice/internal/pkg/repository/root"
)

type Repository struct {
	root *root.Repository
}

func NewRepository(root *root.Repository) *Repository {
	return &Repository{root: root}
}

func (r *Repository) CreateOrUpdate(ctx context.Context, sql string, args ...interface{}) (uint, error) {
	return r.root.CreateOrUpdate(ctx, sql, args...)
}

func (r *Repository) Get(ctx context.Context, sql string, args ...interface{}) ([]map[string]interface{}, error) {
	return r.root.Get(ctx, sql, args...)
}

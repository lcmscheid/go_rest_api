package service

import (
	"context"
	"errors"

	"github.com/lcmscheid/go_rest_api/internal/todo/model"
	"github.com/lcmscheid/go_rest_api/pkg/db"
	"github.com/lcmscheid/go_rest_api/pkg/erru"
)

func (s Service) Get(ctx context.Context, id int) (model.ToDo, error) {
	todo, err := s.repo.Find(ctx, id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return model.ToDo{}, erru.ErrArgument{Wrapped: errors.New("todo object not found")}
	default:
		return model.ToDo{}, err
	}
	return todo, nil
}

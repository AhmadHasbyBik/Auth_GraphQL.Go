package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"graphql/graph/model"
	"time"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUserInput) (*model.User, error) {

	now := int(time.Now().Unix())

	u := &model.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: now,
		UpdatedAt: 0,
	}

	result, err := DB.Exec("INSERT INTO `users` (username, email, password, created_at, updated_at) VALUES(?,?,?,?,?)", u.Username, u.Email, u.Password, u.CreatedAt, u.UpdatedAt)

	if err != nil {
		return nil, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = int(lastId)

	return u, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var result []*model.User
	rows, err := DB.Query("SELECT id, username, email, created_at, updated_at FROM `users`")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u model.User
		err = rows.Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		result = append(result, &u)
	}

	return result, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

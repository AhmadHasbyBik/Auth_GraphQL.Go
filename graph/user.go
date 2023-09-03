package graph

import "graphql/graph/model"

func FindUserByEmail(email string) (*model.User, error) {
	rows := DB.QueryRow("SELECT id, username, email, password, created_at, updated_at FROM `users` WHERE email =?", email)

	var u model.User

	err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

package repository

import (
	"clean-arch/app/model"
	"database/sql"
	"errors"
)

type UserRepository interface {
	GetAllUsers() ([]*model.User, error)
	GetUserByID(id string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
	UpdateUserRole(userID string, roleID string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]*model.User, error) {
	users := []*model.User{}
	query := `
		SELECT u.id, u.username, u.email, u.password_hash, u.full_name, 
		       u.is_active, u.role_id, u.created_at, u.updated_at,
		       r.id, r.name, r.description
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
		ORDER BY u.created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &model.User{
			Role: &model.Role{},
		}
		err := rows.Scan(
			&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FullName,
			&user.IsActive, &user.RoleID, &user.CreatedAt, &user.UpdatedAt,
			&user.Role.ID, &user.Role.Name, &user.Role.Description,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

func (r *userRepository) GetUserByID(id string) (*model.User, error) {
	user := &model.User{
		Role: &model.Role{},
	}
	query := `
		SELECT u.id, u.username, u.email, u.password_hash, u.full_name, 
		       u.is_active, u.role_id, u.created_at, u.updated_at,
		       r.id, r.name, r.description
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
		WHERE u.id = $1
	`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FullName,
		&user.IsActive, &user.RoleID, &user.CreatedAt, &user.UpdatedAt,
		&user.Role.ID, &user.Role.Name, &user.Role.Description,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{
		Role: &model.Role{},
	}
	query := `
		SELECT u.id, u.username, u.email, u.password_hash, u.full_name, 
		       u.is_active, u.role_id, u.created_at, u.updated_at,
		       r.id, r.name, r.description
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
		WHERE u.username = $1
	`

	err := r.db.QueryRow(query, username).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FullName,
		&user.IsActive, &user.RoleID, &user.CreatedAt, &user.UpdatedAt,
		&user.Role.ID, &user.Role.Name, &user.Role.Description,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user *model.User) (*model.User, error) {
	query := `
		INSERT INTO users (username, email, password_hash, full_name, role_id, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRow(
		query,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.FullName,
		user.RoleID,
		true,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(user *model.User) error {
	query := `
		UPDATE users
		SET username = $1, email = $2, full_name = $3, is_active = $4, updated_at = NOW()
		WHERE id = $5
	`

	result, err := r.db.Exec(
		query,
		user.Username,
		user.Email,
		user.FullName,
		user.IsActive,
		user.ID,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *userRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *userRepository) UpdateUserRole(userID string, roleID string) error {
	// Validate role exists
	roleQuery := `SELECT id FROM roles WHERE id = $1`
	err := r.db.QueryRow(roleQuery, roleID).Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("role not found")
		}
		return err
	}

	query := `
		UPDATE users
		SET role_id = $1, updated_at = NOW()
		WHERE id = $2
	`

	result, err := r.db.Exec(query, roleID, userID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

package repository

import (
	"clean-arch/app/model"
	"database/sql"
	"errors"
)

type AuthRepository interface {
	GetUserByUsername(username string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) error
	GetUserPermissions(userID string) ([]string, error)
	GetRoleByID(roleID string) (*model.Role, error)
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{
		Role: &model.Role{},  // Initialize Role to prevent nil pointer
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

func (r *authRepository) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{
		Role: &model.Role{},  // Initialize Role to prevent nil pointer
	}
	query := `
		SELECT u.id, u.username, u.email, u.password_hash, u.full_name, 
		       u.is_active, u.role_id, u.created_at, u.updated_at,
		       r.id, r.name, r.description
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
		WHERE u.email = $1
	`

	err := r.db.QueryRow(query, email).Scan(
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

func (r *authRepository) GetUserByID(id string) (*model.User, error) {
	user := &model.User{
		Role: &model.Role{},  // Initialize Role to prevent nil pointer
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

func (r *authRepository) CreateUser(user *model.User) (*model.User, error) {
	roleID := user.RoleID.String()
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
		roleID,
		true,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *authRepository) UpdateUser(user *model.User) error {
	query := `
		UPDATE users
		SET username = $1, email = $2, password_hash = $3, 
		    full_name = $4, is_active = $5, updated_at = NOW()
		WHERE id = $6
	`

	result, err := r.db.Exec(
		query,
		user.Username,
		user.Email,
		user.PasswordHash,
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

func (r *authRepository) GetUserPermissions(userID string) ([]string, error) {
	permissions := []string{}
	query := `
		SELECT p.name
		FROM permissions p
		JOIN role_permissions rp ON p.id = rp.permission_id
		JOIN users u ON u.role_id = rp.role_id
		WHERE u.id = $1
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var permission string
		if err := rows.Scan(&permission); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}

	return permissions, rows.Err()
}

func (r *authRepository) GetRoleByID(roleID string) (*model.Role, error) {
	role := &model.Role{}
	query := `SELECT id, name, description FROM roles WHERE id = $1`

	err := r.db.QueryRow(query, roleID).Scan(&role.ID, &role.Name, &role.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("role not found")
		}
		return nil, err
	}

	return role, nil
}

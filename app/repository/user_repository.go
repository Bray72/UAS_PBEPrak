package repository

import (
	"clean-arch/app/model"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

// GetAllUsersWithPagination fetches users with pagination and search
func GetAllUsersWithPagination(db *sql.DB, search, sortBy, order string, limit, offset int) ([]*model.User, error) {
	validSortColumns := map[string]bool{
		"id": true, "username": true, "email": true, "full_name": true, "created_at": true,
	}
	if !validSortColumns[sortBy] {
		sortBy = "created_at"
	}

	if strings.ToLower(order) != "desc" {
		order = "asc"
	}

	var query string
	var args []interface{}

	if search != "" {
		query = fmt.Sprintf(`
			SELECT u.id, u.username, u.email, u.password_hash, u.full_name, 
			       u.is_active, u.role_id, u.created_at, u.updated_at,
			       r.id, r.name, r.description
			FROM users u
			LEFT JOIN roles r ON u.role_id = r.id
			WHERE (u.username ILIKE $1 OR u.email ILIKE $1 OR u.full_name ILIKE $1)
			ORDER BY %s %s
			LIMIT $2 OFFSET $3
		`, sortBy, order)
		searchParam := "%" + search + "%"
		args = []interface{}{searchParam, limit, offset}
	} else {
		query = fmt.Sprintf(`
			SELECT u.id, u.username, u.email, u.password_hash, u.full_name, 
			       u.is_active, u.role_id, u.created_at, u.updated_at,
			       r.id, r.name, r.description
			FROM users u
			LEFT JOIN roles r ON u.role_id = r.id
			ORDER BY %s %s
			LIMIT $1 OFFSET $2
		`, sortBy, order)
		args = []interface{}{limit, offset}
	}

	log.Printf("[DEBUG] Query: %s", query)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
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

	return users, nil
}

// CountUsers counts total users matching search criteria
func CountUsers(db *sql.DB, search string) (int, error) {
	var total int
	var query string

	if search != "" {
		query = `SELECT COUNT(*) FROM users WHERE (username ILIKE $1 OR email ILIKE $1 OR full_name ILIKE $1)`
		err := db.QueryRow(query, "%"+search+"%").Scan(&total)
		if err != nil {
			return 0, err
		}
	} else {
		query = `SELECT COUNT(*) FROM users`
		err := db.QueryRow(query).Scan(&total)
		if err != nil {
			return 0, err
		}
	}

	return total, nil
}

// GetUserByID fetches user by ID with role information
func GetUserByID(db *sql.DB, id string) (*model.User, error) {
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

	err := db.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FullName,
		&user.IsActive, &user.RoleID, &user.CreatedAt, &user.UpdatedAt,
		&user.Role.ID, &user.Role.Name, &user.Role.Description,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByUsername fetches user by username
func GetUserByUsername(db *sql.DB, username string) (*model.User, error) {
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

	err := db.QueryRow(query, username).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.FullName,
		&user.IsActive, &user.RoleID, &user.CreatedAt, &user.UpdatedAt,
		&user.Role.ID, &user.Role.Name, &user.Role.Description,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser inserts a new user into the database
func CreateUser(db *sql.DB, req *model.CreateUserRequest, hashedPassword string) (*model.User, error) {
	now := time.Now()
	var id string

	query := `
		INSERT INTO users (username, email, password_hash, full_name, role_id, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`

	err := db.QueryRow(query, req.Username, req.Email, hashedPassword, req.FullName, req.RoleID, true, now, now).Scan(&id)
	if err != nil {
		return nil, err
	}

	return GetUserByID(db, id)
}

// UpdateUser updates user information
func UpdateUser(db *sql.DB, id string, req *model.UpdateUserRequest) (*model.User, error) {
	now := time.Now()
	query := `
		UPDATE users
		SET username = $1, email = $2, full_name = $3, is_active = $4, updated_at = $5
		WHERE id = $6
	`

	result, err := db.Exec(query, req.Username, req.Email, req.FullName, req.IsActive, now, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return GetUserByID(db, id)
}

// DeleteUser deletes a user from the database
func DeleteUser(db *sql.DB, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// AssignRole assigns a role to user
func AssignRole(db *sql.DB, userID, roleID string) (*model.User, error) {
	roleQuery := `SELECT id FROM roles WHERE id = $1`
	var roleIDVar string
	err := db.QueryRow(roleQuery, roleID).Scan(&roleIDVar)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("role not found")
		}
		return nil, err
	}

	now := time.Now()
	query := `
		UPDATE users
		SET role_id = $1, updated_at = $2
		WHERE id = $3
	`

	result, err := db.Exec(query, roleID, now, userID)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return GetUserByID(db, userID)
}

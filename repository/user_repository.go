package repository

import (
	"database/sql"
	"sayembara/entity/model"
	"sayembara/utils"
)

type UserRepository interface {
	Create(user model.UserPassword) (string, error)
	IsEmailAvailable(email string) bool
	GetUserByEmail(email string) (model.EmailPassword, error)
	GetUserById(id string) (model.User, error)
	GetUsers() ([]model.User, error)
}

type userRepository struct {
	idGenerator utils.IdGenerator
	db          *sql.DB
}

func NewUserRepository(idGenerator utils.IdGenerator, db *sql.DB) *userRepository {
	return &userRepository{idGenerator, db}
}

func (r *userRepository) Create(user model.UserPassword) (string, error) {
	id := r.idGenerator()
	query := "INSERT INTO users(id, name, email, password, profile, banner, category) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, id, user.Name, user.Email, user.Password, user.Profile, user.Banner, user.Category)

	if user.Category == "UMKM" {
		roleId := r.idGenerator()
		query = "INSERT INTO umkm(id, id_user, verified) VALUES (?, ?, ?)"
		_, err = r.db.Exec(query, roleId, id, false)
	}

	if user.Category == "MAHASISWA" {
		roleId := r.idGenerator()
		query = "INSERT INTO mahasiswa(id, id_user, expert) VALUES (?, ?, ?)"
		_, err = r.db.Exec(query, roleId, id, false)
	}

	return id, err
}

func (r *userRepository) IsEmailAvailable(email string) bool {
	query := "SELECT * FROM users WHERE email = ?"
	rows, _ := r.db.Query(query, email)

	if rows.Next() {
		return false
	}

	return true
}

func (r *userRepository) GetUserByEmail(email string) (model.EmailPassword, error) {
	query := "SELECT id, email, password, category FROM users WHERE email = ?"
	rows, err := r.db.Query(query, email)

	var user model.EmailPassword
	if rows.Next() {
		err = rows.Scan(
			&user.Id, &user.Email, &user.Password, &user.Category,
		)
	}

	return user, err
}

func (r *userRepository) GetUserById(id string) (model.User, error) {
	query := "SELECT id, name, email, profile, banner, category FROM users WHERE id = ?"
	rows, err := r.db.Query(query, id)

	var user model.User
	if rows.Next() {
		err = rows.Scan(
			&user.Id, &user.Name, &user.Email, &user.Profile, &user.Banner, &user.Category,
		)
	}

	return user, err
}

func (r *userRepository) GetUsers() ([]model.User, error) {
	query := "SELECT id, name, profile, category FROM users"
	rows, err := r.db.Query(query)

	var users []model.User

	for rows.Next() {
		var user model.User
		rows.Scan(
			&user.Id, &user.Name, &user.Profile, &user.Category,
		)
		users = append(users, user)
	}

	return users, err
}

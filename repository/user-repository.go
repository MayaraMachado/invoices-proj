package repository

import (
	"github.com/mayaramachado/invoice-api/entity"
	"database/sql"
	"time"
	"fmt"
)

type UserRepository interface {
	Save(user entity.User) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	CloseDbConnection()
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(dbConnection *sql.DB) UserRepository {
	return &userRepository{
			db: dbConnection,
		}
}

func (repository *userRepository) CloseDbConnection() {
	err := repository.db.Close()
	if err != nil{
			panic("Failed to close database!")
		}
}

func (repository *userRepository) Save(user entity.User) (entity.User, error) {
	newUser := entity.User{}
	active := 1
	query_string := "INSERT INTO users (email, password, is_active, created_at) VALUES ($1, $2, $3, $4) RETURNING *;"
	result := repository.db.QueryRow(query_string, user.Email, user.Password, active, time.Now())
	
	err := result.Scan(&newUser.Id, &newUser.Email, &newUser.Password, &newUser.IsActive, &newUser.CreatedAt, &newUser.DeactivatedAt )
	if err != nil {
		fmt.Println(err)
		return newUser, err
	}
	return newUser, nil
}

func (repository *userRepository) GetByEmail(email string) (entity.User, error){
	user := entity.User{}
	query_string := `SELECT * FROM user WHERE email=$1 and is_active=1;`
	result := repository.db.QueryRow(query_string, email)
	err := result.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		return user, err
	}
	return user, nil
}

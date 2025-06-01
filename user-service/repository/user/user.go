package repository

import (
	"context"
	"database/sql"

	"github.com/shoelfikar/finpay/models"

)

type userRepository struct {
	DB *sql.DB
}

type UserRepository interface {
   CreateUser(user *models.User) *models.User
   GetUserByID(userId string) *model.User
   GetUserByEmail(email string) *model.User
}

func NewUserRepository(db *sql.DB) UserRepository {
   return &userRepository{
      DB: db,
   }
}

func (u *userRepository) CreateUser(user *model.User) *model.User {
   tx, err := u.DB.Begin()
   ctx := context.Background()
   if err != nil {
      panic("Error database transaction"+ err.Error())
   }

   query := `
      INSERT INTO users (id, email, password, phone_number, balance, point, status, created_by)
      VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
      RETURNING id;
   `
   _, err = tx.ExecContext(ctx, query, &user.Id, &user.Email, &user.Password, &user.PhoneNumber, 0, 0, &user.Status, &user.CreatedBy)

   if err != nil {
      tx.Rollback()
      panic("Error insert user "+ err.Error())
   }

   tx.Commit()

   user.Password = nil

   return user
}

func (u *userRepository) GetUserByID(userId string) *model.User {
   ctx := context.Background()
   query := `
      SELECT id, email, role, status, created_at, updated_at, created_by 
      FROM users WHERE id = $1
   `

   var user model.User
   result := u.DB.QueryRowContext(ctx, query, &userId).Scan(&user.Id, &user.Email, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy)

   if result == sql.ErrNoRows {
      return nil
   }

   return &user
}

func (u *userRepository) GetUserByEmail(email string) *model.User {
   ctx := context.Background()
   query := `
      SELECT id, email, password, status, role, created_at, updated_at, created_by
      FROM users WHERE email = $1
   `

   var user model.User
   result := u.DB.QueryRowContext(ctx, query, &email).Scan(&user.Id, &user.Email, &user.Password, &user.Status, &user.Role, &user.CreatedAt, &user.UpdatedAt, &user.CreatedBy)

   if result == sql.ErrNoRows {
      return nil
   }

   return &user
}
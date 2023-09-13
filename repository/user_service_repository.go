package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SawitProRecruitment/UserService/bootstrap"
	"github.com/SawitProRecruitment/UserService/errors"
	"github.com/SawitProRecruitment/UserService/models"
)

type userServiceRepository struct {
	db *bootstrap.PostgresClient
}

//go:generate mockery --name IUserServicePointRepository
type IUserServicePointRepository interface {
	GetUserByPhone(ctx context.Context, phoneNumber string) (user *models.User, err error)
	UpdateUser(ctx context.Context, userID int, fields map[string]interface{}) error
	InsetUser(ctx context.Context, data models.User) (int64, error)
	GetUserByUserID(ctx context.Context, userID int) (user *models.User, err error)
}

func NewUserPointRepo(db *bootstrap.PostgresClient) IUserServicePointRepository {
	return &userServiceRepository{
		db: db,
	}
}

func (u *userServiceRepository) GetUserByPhone(ctx context.Context, phoneNumber string) (*models.User, error) {
	var user models.User
	err := u.db.Db.QueryRowContext(ctx,
		fmt.Sprintf(`SELECT user_id, user_phone_number, user_full_name, user_password, user_logged FROM %s WHERE user_phone_number = $1`,
			user.TableName()), phoneNumber).
		Scan(&user.UserID, &user.UserPhoneNumber, &user.UserFullName, &user.UserPassword, &user.UserLogged)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (u *userServiceRepository) UpdateUser(ctx context.Context, userID int, fields map[string]interface{}) error {
	setClause := ""
	values := []interface{}{}
	tx, err := u.db.Db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	i := 1
	for fieldName, fieldValue := range fields {
		setClause += fmt.Sprintf("%s = $%d, ", fieldName, i)
		values = append(values, fieldValue)
		i++
	}

	setClause = setClause[:len(setClause)-2]

	updateQuery := fmt.Sprintf("UPDATE users SET %s WHERE user_id = $%d", setClause, i)
	values = append(values, userID)

	_, err = u.db.Db.Exec(updateQuery, values...)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return errors.ErrorInternalServer
	}
	return err
}

func (u *userServiceRepository) InsetUser(ctx context.Context, data models.User) (int64, error) {
	var lastInsertID int64
	tx, err := u.db.Db.Begin()
	if err != nil {
		return 0, err
	}
	stmt, err := u.db.Db.PrepareContext(ctx, "INSERT INTO users (user_phone_number, user_full_name, user_password, user_logged) VALUES ($1, $2, $3, $4)  RETURNING user_id")
	if err != nil {
		return 0, errors.ErrorInternalServer
	}

	err = stmt.QueryRowContext(ctx, data.UserPhoneNumber, data.UserFullName, data.UserPassword, data.UserLogged).Scan(&lastInsertID)
	if err != nil {
		// Rollback the transaction if there's an error
		tx.Rollback()
		return 0, errors.ErrorUserAlreadyExist
	}
	err = tx.Commit()
	if err != nil {
		return 0, errors.ErrorInternalServer
	}
	return lastInsertID, nil
}

func (u *userServiceRepository) GetUserByUserID(ctx context.Context, userID int) (*models.User, error) {
	var user models.User
	err := u.db.Db.QueryRowContext(ctx,
		fmt.Sprintf(`SELECT user_id, user_phone_number, user_full_name, user_password, user_logged FROM %s WHERE user_id = $1`,
			user.TableName()), userID).
		Scan(&user.UserID, &user.UserPhoneNumber, &user.UserFullName, &user.UserPassword, &user.UserLogged)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrorUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

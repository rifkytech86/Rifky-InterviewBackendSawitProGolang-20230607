package repository

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SawitProRecruitment/UserService/bootstrap"
	"github.com/SawitProRecruitment/UserService/commons"
	"github.com/SawitProRecruitment/UserService/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_userServiceRepository_GetUserByPhone(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("error mocking sql")
	}
	defer db.Close()

	repo := &userServiceRepository{
		db: &bootstrap.PostgresClient{
			Db: db,
		},
	}
	tests := []struct {
		phoneNumber    string
		expectedResult *models.User
		expectedError  error
		mockBehavior   func()
	}{
		{
			phoneNumber: "+6285722811111",
			expectedResult: &models.User{
				UserID:          1,
				UserPhoneNumber: "+6285722811111",
				UserFullName:    "sawit pro",
				UserPassword:    "asdqwe1A@",
				UserLogged:      1,
			},
			expectedError: nil,
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"user_id", "user_phone_number", "user_full_name", "user_password", "user_logged"}).
					AddRow(1, "+6285722811111", "sawit pro", "asdqwe1A@", 1)
				mock.ExpectQuery("SELECT user_id, user_phone_number, user_full_name, user_password, user_logged").
					WithArgs("+6285722811111").
					WillReturnRows(rows)
			},
		},
		{
			phoneNumber:    "123456789",
			expectedResult: nil,
			expectedError:  commons.ErrorUserNotFound,
			mockBehavior: func() {
				mock.ExpectQuery("SELECT user_id, user_phone_number, user_full_name, user_password, user_logged").
					WithArgs("123456789").
					WillReturnError(sql.ErrNoRows)
			},
		},
		{
			phoneNumber:    "123456789",
			expectedResult: nil,
			expectedError:  sql.ErrConnDone,
			mockBehavior: func() {
				mock.ExpectQuery("SELECT user_id, user_phone_number, user_full_name, user_password, user_logged").
					WithArgs("123456789").
					WillReturnError(sql.ErrConnDone)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.phoneNumber, func(t *testing.T) {
			test.mockBehavior()
			user, err := repo.GetUserByPhone(context.Background(), test.phoneNumber)
			assert.Equal(t, test.expectedResult, user)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func Test_userServiceRepository_InsetUser(t *testing.T) {
	// Initialize a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("error mocking sql")
	}
	defer db.Close()

	repo := &userServiceRepository{
		db: &bootstrap.PostgresClient{
			Db: db,
		},
	}
	tests := []struct {
		userData      models.User
		expectedID    int64
		expectedError error
		mockBehavior  func()
	}{
		{
			userData: models.User{
				UserPhoneNumber: "123456789",
				UserFullName:    "sawit pro",
				UserPassword:    "asdqwe1A@",
				UserLogged:      1,
			},
			expectedID:    1,
			expectedError: nil,
			mockBehavior: func() {
				mock.ExpectBegin()

				mock.ExpectPrepare(`INSERT INTO users \(user_phone_number, user_full_name, user_password, user_logged\) VALUES \(\$1, \$2, \$3, \$4\)  RETURNING user_id`)

				mock.ExpectQuery(`INSERT INTO users \(user_phone_number, user_full_name, user_password, user_logged\) VALUES \(\$1, \$2, \$3, \$4\)  RETURNING user_id`).
					WithArgs("123456789", "sawit pro", "asdqwe1A@", 1).
					WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(1))

				mock.ExpectCommit()
			},
		},
		{
			userData: models.User{
				UserPhoneNumber: "123456789",
				UserFullName:    "sawit pro",
				UserPassword:    "asdqwe1A@",
				UserLogged:      1,
			},
			expectedID:    0,
			expectedError: commons.ErrorUserAlreadyExist,
			mockBehavior: func() {
				mock.ExpectBegin()

				mock.ExpectPrepare(`INSERT INTO users \(user_phone_number, user_full_name, user_password, user_logged\) VALUES \(\$1, \$2, \$3, \$4\)  RETURNING user_id`)

				mock.ExpectQuery(`INSERT INTO users \(user_phone_number, user_full_name, user_password, user_logged\) VALUES \(\$1, \$2, \$3, \$4\)  RETURNING user_id`).
					WithArgs("123456789", "sawit pro", "asdqwe1A@", 1).
					WillReturnError(sql.ErrConnDone)

				mock.ExpectCommit()
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			test.mockBehavior()
			userID, err := repo.InsetUser(context.Background(), test.userData)
			assert.Equal(t, test.expectedID, userID)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func Test_userServiceRepository_GetUserByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("error mocking sql")
	}
	defer db.Close()

	repo := &userServiceRepository{
		db: &bootstrap.PostgresClient{
			Db: db,
		},
	}

	tests := []struct {
		userID         int
		expectedResult *models.User
		expectedError  error
		mockBehavior   func()
	}{
		{
			userID: 1,
			expectedResult: &models.User{
				UserID:          1,
				UserPhoneNumber: "123456789",
				UserFullName:    "sawit pro",
				UserPassword:    "asdqwe1A@",
				UserLogged:      1,
			},
			expectedError: nil,
			mockBehavior: func() {
				rows := sqlmock.NewRows([]string{"user_id", "user_phone_number", "user_full_name", "user_password", "user_logged"}).
					AddRow(1, "123456789", "sawit pro", "asdqwe1A@", 1)
				mock.ExpectQuery("SELECT user_id, user_phone_number, user_full_name, user_password, user_logged").
					WithArgs(1).
					WillReturnRows(rows)
			},
		},
		{
			userID:         2,
			expectedResult: nil,
			expectedError:  commons.ErrorUserNotFound,
			mockBehavior: func() {
				mock.ExpectQuery("SELECT user_id, user_phone_number, user_full_name, user_password, user_logged").
					WithArgs(2).
					WillReturnError(sql.ErrNoRows)
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			test.mockBehavior()
			user, err := repo.GetUserByUserID(context.Background(), test.userID)
			assert.Equal(t, test.expectedResult, user)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestNewUserPointRepo(t *testing.T) {
	type args struct {
		db *bootstrap.PostgresClient
	}
	tests := []struct {
		name string
		args args
		want IUserServicePointRepository
	}{
		{
			name: "initial new",
			args: args{
				db: &bootstrap.PostgresClient{},
			},
			want: NewUserPointRepo(&bootstrap.PostgresClient{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewUserPointRepo(tt.args.db), "NewUserPointRepo(%v)", tt.args.db)
		})
	}
}

func Test_userServiceRepository_UpdateUser(t *testing.T) {
	// Initialize a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error occurred while initializing mock DB: %s", err)
	}
	defer db.Close()

	// Create a new userServiceRepository with the mock database
	repo := &userServiceRepository{
		db: &bootstrap.PostgresClient{
			Db: db,
		},
	}

	// Define test cases
	tests := []struct {
		userID        int
		fields        map[string]interface{}
		expectedError error
		mockBehavior  func()
	}{
		{
			userID: 1,
			fields: map[string]interface{}{
				"user_full_name": "sawit pro",
				"user_password":  "asdqwe1A@",
			},
			expectedError: nil,
			mockBehavior: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE users SET user_full_name = \$1, user_password = \$2 WHERE user_id = \$3`).
					WithArgs("sawit pro", "asdqwe1A@", 1)
				mock.ExpectCommit()
			},
		},
	}
	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			test.mockBehavior()

			_ = repo.UpdateUser(context.Background(), test.userID, test.fields)

		})
	}

}

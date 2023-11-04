package repository_test

import (
	"context"
	"os"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/internal/users/entities"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/internal/users/repository"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/internal/users/repository/repository_query"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/pkg/constant"
	"github.com/martinyonatann/auth-microservices-traefik/auth-service/pkg/otel/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_repository_UpdateUserByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	defer db.Close()

	var (
		conn = sqlx.NewDb(db, "sqlmock")
		repo = repository.NewRepository(conn, zerolog.NewZeroLog(context.Background(), os.Stdout))
	)

	currentTime := time.Now()
	userID := int64(1)

	t.Run("UpdateUserByID_positive_case", func(t *testing.T) {
		user := entities.Users{
			Fullname:    "martin yonatan pasaribu",
			PhoneNumber: "08121213131414",
			UserType:    constant.UserTypePremium,
			IsActive:    true,
			CreatedAt:   currentTime,
			CreatedBy:   "admin",
		}

		mock.
			ExpectExec(repository_query.InsertUsers).
			WithArgs(
				user.Fullname,
				user.PhoneNumber,
				user.IsActive,
				user.UserType,
				currentTime,
				user.CreatedBy,
			).WillReturnResult(sqlmock.NewResult(userID, 1))

		userID, err := repo.SaveNewUser(context.Background(), user)
		assert.NoError(t, err)
		assert.NotEmpty(t, userID)

		argsUpdateUser := entities.UpdateUsers{
			UserID:      userID,
			Fullname:    "edited name",
			PhoneNumber: "081122334455",
			UpdatedAt:   currentTime,
			UpdatedBy:   "admin",
		}

		mock.ExpectExec(repository_query.UpdateUsers).
			WithArgs(
				argsUpdateUser.Fullname, argsUpdateUser.Fullname,
				argsUpdateUser.PhoneNumber, argsUpdateUser.PhoneNumber,
				argsUpdateUser.UserType, argsUpdateUser.UserType,
				argsUpdateUser.UpdatedAt,
				argsUpdateUser.UpdatedBy,
				argsUpdateUser.UserID,
			).WillReturnResult(sqlmock.NewResult(userID, 1))

		err = repo.UpdateUserByID(context.Background(), argsUpdateUser)
		require.NoError(t, err)
	})

}

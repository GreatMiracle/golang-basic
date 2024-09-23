package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"simpleBank/util"
	"testing"
)

func createRandomUser(t *testing.T) User {
	generatePwd := util.RandomString(6)
	fmt.Println(generatePwd)
	hashedPassword, err := util.HashPassword(generatePwd)

	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       generatePwd,
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	//user1 := createRandomUser(t)
	//username := user1.Username
	user2, err := testQueries.GetUser(context.Background(), "irzwbs")
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, "irzwbs", user2.Username)
	require.Equal(t, "$2a$10$2r.bm5T9O54.kGCfN5mKgOFkIrRxdcTbJLBpfb5HTch0gf68mvCZm", user2.HashedPassword)
	require.Equal(t, "rkuizk", user2.FullName)
	require.Equal(t, "taqhrh@email.com", user2.Email)
	//require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	//require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestPassword(t *testing.T) {
	password := "secret"
	hashedPassword1, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	// Kiểm tra trường hợp mật khẩu đúng
	err = util.CheckPassword(password, hashedPassword1)
	require.NoError(t, err)

	// Kiểm tra trường hợp mật khẩu sai
	wrongPassword := "wrongpassword"
	err = util.CheckPassword(wrongPassword, hashedPassword1)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	// Kiểm tra mật khẩu băm khác nhau với cùng một mật khẩu
	hashedPassword2, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}

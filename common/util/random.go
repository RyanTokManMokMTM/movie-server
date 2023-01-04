package util

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ryantokmanmokmtm/movie-server/common/crytox"
	"github.com/ryantokmanmokmtm/movie-server/internal/models"
	"math/rand"
)

func RandomUser() models.User {
	gofakeit.New(7557)
	return models.User{
		ID:       uint(rand.Intn(10000)),
		Name:     gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: gofakeit.Password(true, true, true, false, false, 16),
		Avatar:   "/defaultAvatar.jpeg",
		Cover:    "/defaultCover.jpeg",
	}
}

func RandomUserEncryptPassword(pw, salt string) models.User {
	gofakeit.New(7557)
	return models.User{
		ID:       uint(rand.Intn(10000)),
		Name:     gofakeit.Name(),
		Email:    gofakeit.Email(),
		Password: crytox.PasswordEncrypt(pw, salt),
		Avatar:   "/defaultAvatar.jpeg",
		Cover:    "/defaultCover.jpeg",
	}
}

package crawle

import (
	"github.com/earlgray283/material-gakujo/api/db/model"
	auth "github.com/earlgray283/material-gakujo/api/server/libauth"
	"github.com/pkg/errors"
)

func AuthInfoFromUser(user *model.User, cryptoKey []byte) (string, string, error) {
	decryptedPassword, err := auth.Decrypt(user.GakujoEncryptedPassword, cryptoKey)
	if err != nil {
		return "", "", errors.WithStack(err)
	}

	return user.GakujoUsername, string(decryptedPassword), nil
}

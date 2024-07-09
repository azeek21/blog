package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/azeek21/blog/models"
)

var USER_NOT_FOUND_IN_CONTEXT_ERROR = errors.New("User not found in context")

func GetUser(ctx context.Context) (*models.User, error) {
	user := ctx.Value(models.USER_MODEL_NAME)
	if user != nil {
		return user.(*models.User), nil
	}
	return nil, USER_NOT_FOUND_IN_CONTEXT_ERROR
}

func IsAuthed(ctx context.Context) bool {
	user := ctx.Value(models.USER_MODEL_NAME)
	isAuthed := user != nil
	fmt.Printf("USER: %v -  %v+\n", isAuthed, user)
	return isAuthed
}

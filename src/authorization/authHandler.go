package authorization

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lunasky-hy/dialy-note-app/src/model"
	"github.com/lunasky-hy/dialy-note-app/src/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
    repos repository.DiaryRepository
}

func (h AuthHandler) AuthUser(user *model.User) (error) {
	registerdUser, db_err := h.repos.UserGet(user.Name)

	if db_err != nil {
		return db_err
	}

	pw_err := bcrypt.CompareHashAndPassword(registerdUser.Password, user.Password)

	if pw_err != nil {
		return pw_err
	}
    
    return nil
}

func (h AuthHandler) VerifyJwt(authHeader string, user *model.User) error {
    if authHeader == "" {
        return errors.New("bad request: Authorization Header does not exist")
    }

    splitStr := strings.Split(authHeader, " ")

    if len(splitStr) != 2 {
        return errors.New("bad request: JWT format is invalid")
    }
    if splitStr[0] != "Bearer" {
        return errors.New("bad request: Bearer is not included")
    }

    token, jwt_err := jwt.Parse(splitStr[1], func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        secret := os.Getenv("AUTH_SECRET")
        return []byte(secret), nil
    })

    if jwt_err != nil {
        return fmt.Errorf("jwt parse error : %s", jwt_err)
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        username := string(claims["username"].(string))

        tokenUser, db_err := h.repos.UserGet(username)

        if db_err != nil {
            return fmt.Errorf("%s", db_err)
        }
        
        *user = tokenUser
    }
    return nil
}

func CreateAuthHandler(repos repository.DiaryRepository) AuthHandler {
    handler := AuthHandler{repos: repos}
    return handler
}
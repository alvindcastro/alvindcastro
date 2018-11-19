package handler

import (
	"net/http"
	"time"

	"github.com/alvindcastro/travel-echo-mongo/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (handler *Handler) SignUp(context echo.Context) (err error) {
	// Bind
	user := &model.User{ID: bson.NewObjectId()}
	if err = context.Bind(user); err != nil {
		return
	}

	// Validate
	if user.Email == "" || user.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// Save user
	db := handler.DB.Clone()
	defer db.Close()
	if err = db.DB("mgo").C("users").Insert(user); err != nil {
		return
	}

	return context.JSON(http.StatusCreated, user)
}

func (handler *Handler) Login(context echo.Context) (err error) {
	// Bind
	user := new(model.User)
	if err = context.Bind(user); err != nil {
		return
	}

	// Find user
	db := handler.DB.Clone()
	defer db.Close()
	if err = db.DB("mgo").C("users").
		Find(bson.M{"email": user.Email, "password": user.Password}).One(user); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
		}
		return
	}

	//-----
	// JWT Implementation
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	user.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	user.Password = ""
	return context.JSON(http.StatusOK, user)
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}

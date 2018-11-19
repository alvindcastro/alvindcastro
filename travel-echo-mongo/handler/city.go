package handler

import (
	"github.com/alvindcastro/travel-echo-mongo/model"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

func (handler *Handler) CreateCity(context echo.Context) (err error) {
	user := &model.User{
		ID: bson.ObjectIdHex(userIDFromToken(context)),
	}
	city := &model.City{
		ID: bson.NewObjectId(),
	}
	if err = context.Bind(city); err != nil {
		return
	}

	// Validation
	if city.Name == "" || city.Desc == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid fields"}
	}

	// Find user from database
	db := handler.DB.Clone()
	defer db.Close()
	if err = db.DB("mgo").C("users").FindId(user.ID).One(user); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return
	}

	// Save post in database
	if err = db.DB("mgo").C("city").Insert(city); err != nil {
		return
	}
	return context.JSON(http.StatusCreated, city)
}

func (handler *Handler) FetchCities(c echo.Context) (err error) {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	// Retrieve cities from database
	var cities []*model.City
	db := handler.DB.Clone()
	if err = db.DB("mgo").C("city").
		Find(bson.M{}).
		Skip((page - 1) * limit).
		Limit(limit).
		All(&cities); err != nil {
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, cities)
}

func (handler *Handler) FetchCity(c echo.Context) (err error) {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	// Retrieve cities from database
	var city []*model.City
	var name string
	name = c.Param("name")
	db := handler.DB.Clone()
	if err = db.DB("mgo").C("city").
		Find(bson.M{"name": name}).All(&city); err != nil {
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, city)
}

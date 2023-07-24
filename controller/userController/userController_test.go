package userController

import (
	"bytes"
	"encoding/json"
	"go-blog/controller/publicController"
	"go-blog/dto"
	"go-blog/model"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var TestUser = model.User{
	Id:       1,
	Username: "test",
	Tasks:    []model.Task{},
}

var ExpectedResponse = "{\"ID\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\"" +
	",\"DeletedAt\":null,\"Id\":1,\"Username\":\"test\",\"Tasks\":[]}"

var TestUserCreateDTO = dto.UserCreateDTO{
	Username: "test1",
}

var ExpectedResponseCreate = "{\"ID\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\"," +
	"\"DeletedAt\":null,\"Id\":0,\"Username\":\"test1\",\"Tasks\":null}"

var ExpectedResponseCreateError = "{\"message\":\"User cannot be created!\"}"

var ExpectedResponseUpdate = "{\"ID\":0,\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\"," +
	"\"DeletedAt\":null,\"Id\":1,\"Username\":\"test1\",\"Tasks\":[]}"

var ExpectedResponseDelete = "{\"message\":\"User deleted successfully!\"}"

func TestPublicMethod(t *testing.T) {
	response := "{\"message\":\"Hello world!\"}"
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	publicController.HelloWorld(ctx)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, response, recorder.Body.String())
}

func TestGetUserById(t *testing.T) {
	oldMethod := GetUserByIdHelper
	defer func() { GetUserByIdHelper = oldMethod }()
	GetUserByIdHelper = func(ctx *gin.Context) model.User {
		return TestUser
	}
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = []gin.Param{
		{Key: "id", Value: "1"},
	}
	GetUserByID(ctx)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, ExpectedResponse, recorder.Body.String())
}

func TestGetUserByUsername(t *testing.T) {
	oldMethod := GetUserByUsernameHelper
	defer func() {
		GetUserByUsernameHelper = oldMethod
	}()
	GetUserByUsernameHelper = func(username string) model.User {
		return TestUser
	}
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = []gin.Param{
		{Key: "username", Value: "test"},
	}
	GetUserByUsername(ctx)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, ExpectedResponse, recorder.Body.String())
}

func TestCreateUser(t *testing.T) {
	oldMethod := CreateUserHelper
	defer func() { CreateUserHelper = oldMethod }()
	CreateUserHelper = func(user model.User) bool {
		return true
	}
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Request.Method = "POST"
	ctx.Request.Header.Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(TestUserCreateDTO)
	if err != nil {
		panic(err)
	}
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
	CreateUser(ctx)
	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, ExpectedResponseCreate, recorder.Body.String())
}

func TestCreateUserError(t *testing.T) {
	oldMethod := CreateUserHelper
	defer func() { CreateUserHelper = oldMethod }()
	CreateUserHelper = func(user model.User) bool {
		return false
	}
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Request.Method = "POST"
	ctx.Request.Header.Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(TestUserCreateDTO)
	if err != nil {
		panic(err)
	}
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
	CreateUser(ctx)
	assert.Equal(t, 400, recorder.Code)
	assert.Equal(t, ExpectedResponseCreateError, recorder.Body.String())
}

func TestUpdateUser(t *testing.T) {
	oldMethod := GetUserByIdHelper
	oldUpdateMethod := UpdateUserHelper
	defer func() {
		GetUserByIdHelper = oldMethod
		UpdateUserHelper = oldUpdateMethod
	}()
	GetUserByIdHelper = func(ctx *gin.Context) model.User {
		return TestUser
	}
	UpdateUserHelper = func(user model.User) {}
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Request.Method = "POST"
	ctx.Request.Header.Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(TestUserCreateDTO)
	if err != nil {
		panic(err)
	}
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
	ctx.Params = []gin.Param{
		{Key: "id", Value: "1"},
	}
	UpdateUser(ctx)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, ExpectedResponseUpdate, recorder.Body.String())
}

func TestDeleteUser(t *testing.T) {
	oldMethod := GetUserByIdHelper
	oldDeleteMethod := DeleteUserHelper
	defer func() {
		GetUserByIdHelper = oldMethod
		DeleteUserHelper = oldDeleteMethod
	}()
	GetUserByIdHelper = func(ctx *gin.Context) model.User {
		return TestUser
	}
	DeleteUserHelper = func(user model.User) {}
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = []gin.Param{
		{Key: "id", Value: "1"},
	}
	DeleteUserByID(ctx)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, ExpectedResponseDelete, recorder.Body.String())
}

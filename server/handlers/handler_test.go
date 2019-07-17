package handler

import (
	"database/sql"
	"fmt"
	"testing"
	model "trainingApi/server/models"

	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type fakeRepo struct {
}

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	handler  Handler
	fakeRepo fakeRepo
	// currency *model.Currency
}

func (fakeRepo *fakeRepo) GetByID() (*model.Currency, error) {
	m := new(model.Currency)
	return m, nil
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.handler = InitializeHandler(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_handler_GetValueByID() {

	handler := s.handler.GetValueByID()
	router := gin.New()
	router.GET("/api/currency/:id", handler)

	req, _ := http.NewRequest("GET", "/api/currency/:id", nil)
	resp := httptest.NewRecorder()
	fmt.Println("req:", req)
	fmt.Println("res:", resp)
	router.ServeHTTP(resp, req)

	assert.Equal(s.T(), resp.Body.String(), "")
}

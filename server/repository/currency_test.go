package repository

import (
	"database/sql"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	// "github.com/Rosaniline/gorm-ut/pkg/model"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	model "trainingApi/server/models"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository Repository
	person     *model.Currency
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

	s.repository = CreateRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_GetByID() {
	var (
		id   = int64(1)
		name = "test-name"
		bank = "Frances"
		buy  = float64(3.33)
		sell = float64(5.44)
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "person" WHERE (id = $1)`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(id, name))

	res, err := s.repository.GetByID(strconv.FormatInt(id, 10))

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&model.Currency{ID: id, Name: name, Bank: bank, Buy: buy, Sell: sell}, res))
}

// func (s *Suite) Test_repository_Create() {
// 	var (
// 		id   = uuid.NewV4()
// 		name = "test-name"
// 	)

// 	s.mock.ExpectQuery(regexp.QuoteMeta(
// 		`INSERT INTO "person" ("id","name")
// 			VALUES ($1,$2) RETURNING "person"."id"`)).
// 		WithArgs(id, name).
// 		WillReturnRows(
// 			sqlmock.NewRows([]string{"id"}).AddRow(id.String()))

// 	err := s.repository.Create(id, name)

// 	require.NoError(s.T(), err)
// }

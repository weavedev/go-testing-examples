package models

import (
	"context"

	uuid "github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"lab.weave.nl/weave/presentation/packages/testhelpers"
)

type PostgresTestSuite struct {
	testhelpers.BaseSuite
	db          *gorm.DB
	transaction *gorm.DB
	DBManager   DBManager
	Ctx         context.Context
}

type DBManager struct {
	UserDB *UserDB
}

func (s *PostgresTestSuite) SetupTest() {
	s.BaseSuite.SetupTest()
	s.transaction = s.db.Begin()
	s.Ctx = context.Background()
	s.DBManager = DBManager{
		UserDB: NewUserDB(s.transaction),
	}
}

func (s *PostgresTestSuite) TearDownTest() {
	s.transaction.Rollback()
	s.BaseSuite.TearDownTest()

}

func (s *PostgresTestSuite) SetDB(db *gorm.DB) {
	s.db = db
}

func (s *PostgresTestSuite) CreateDummyUsers(amount int) []uuid.UUID {
	var res []uuid.UUID
	for i := 0; i < amount; i++ {
		user := User{FirstName: "John", FamilyName: "Doe"}
		s.NoErrorWithFail(s.DBManager.UserDB.Add(s.Ctx, &user))
		res = append(res, user.ID)
	}
	return res
}

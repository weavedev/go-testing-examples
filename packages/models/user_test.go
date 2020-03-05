package models

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	PostgresTestSuite
}

func (s *UserTestSuite) TestAddUser() {
	s.NoErrorWithFail(s.DBManager.UserDB.Add(s.Ctx, &User{
		FirstName:  "John",
		FamilyName: "Doe",
		Email:      "john@doe.com",
		Password:   "someHashedSecret",
	}))
	users, err := s.DBManager.UserDB.List(s.Ctx)
	s.NoErrorWithFail(err)
	s.Equal(1, len(users))
	s.Equal("John", users[0].FirstName)
	s.Equal("Doe", users[0].FamilyName)
	s.Equal("john@doe.com", users[0].Email)
	s.Equal("someHashedSecret", users[0].Password)
}

func (s *UserTestSuite) TestAddUserPanic() {
	s.Panics(func() {
		err := s.DBManager.UserDB.Add(s.Ctx, nil)
		s.Error(err)
	})
}

func (s *UserTestSuite) TestListUser() {
	amountOfUsers := 3
	s.CreateDummyUsers(amountOfUsers)
	users, err := s.DBManager.UserDB.List(s.Ctx)
	s.NoErrorWithFail(err)
	s.Equal(amountOfUsers, len(users))
}

func (s *UserTestSuite) TestGetUser() {
	amountOfUsers := 1
	userIDs := s.CreateDummyUsers(amountOfUsers)
	_, err := s.DBManager.UserDB.Get(s.Ctx, userIDs[0])
	s.NoErrorWithFail(err)
}

func TestUserTestSuite(t *testing.T) {
	db := NewMustDBConnection()
	userTestSuite := &UserTestSuite{}
	userTestSuite.SetDB(db)
	suite.Run(t, userTestSuite)
}

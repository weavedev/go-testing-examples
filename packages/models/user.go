package models

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"lab.weave.nl/weave/presentation/packages/db"
)

// User model
type User struct {
	ID          uuid.UUID `gorm:"primary_key" sql:"type:uuid;default:uuid_generate_v4()"` // primary key
	CreatedAt   time.Time
	DeletedAt   *time.Time
	Email       string
	FamilyName  string
	FirstName   string
	MiddleName  string
	Password    string
	PhoneNumber string
	UpdatedAt   time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m User) TableName() string {
	return "users"

}

// UserDB is the implementation of the storage interface for
// User.
type UserDB struct {
	Db *gorm.DB
}

// NewUserDB creates a new storage type.
func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{Db: db}
}

// DB returns the underlying database.
func (m *UserDB) DB() interface{} {
	return m.Db
}

// UserStorage represents the storage interface.
type UserStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*User, error)
	Get(ctx context.Context, id uuid.UUID) (*User, error)
	Add(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserDB) TableName() string {
	return "users"
}

func NewMustDBConnection() *gorm.DB {
	m := []interface{}{
		User{},
	}
	db, err := db.ConnectTest(m)
	if err != nil {
		panic(fmt.Sprintf("unable to connect to db: %s", err))
	}
	return db
}

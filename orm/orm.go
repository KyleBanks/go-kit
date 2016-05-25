// Manages access to the database, including ORM
package orm

import (
	"errors"

	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/KyleBanks/go-kit/log"
	"reflect"
	"time"
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type ORM struct {
	Username string
	Password string
	Database string

	conn *gorm.DB
}

// Model is a type added to domain models that will provide ORM functionality.
type Model struct {
	gorm.Model
}

// open creates a database connection, or returns an existing one if present.
func (orm *ORM) Open() *gorm.DB {
	if orm.conn != nil {
		return orm.conn
	}

	// TODO: Use mysql
	dbName := fmt.Sprintf("/tmp/gorm-%d.db", time.Now().Unix())
	log.Info("Using DB:", dbName)

	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic(err)
	} else if db == nil {
		panic(errors.New("Database handle is nil!"))
	}

	log.Info("Database connection established:", orm.Database)

	// Enable logging
	db.SetLogger(log.Logger)
	db.LogMode(true)

	orm.conn = db
	return orm.conn
}

// AutoMigrate performs database migration for all Model types provided.
func (orm ORM) AutoMigrate(models []interface{}) error {
	for _, model := range models {
		modelName := reflect.Indirect(reflect.ValueOf(model)).Type()
		log.Info("Migrating model:", modelName)

		if err := orm.conn.AutoMigrate(model).GetErrors(); len(err) > 0 {
			log.Error("AutoMigrate failed for model", modelName)
			return err[0]
		}
		log.Info("Model migrated:", modelName)
	}

	return nil
}

// Exec performs a raw SQL query against the underlying database.
func (orm ORM) Exec(query string, output interface{}) {
	orm.conn.Raw(query).Scan(output)
}

// Begin starts a new database transaction.
func (orm ORM) Begin() *gorm.DB {
	return orm.conn.Begin()
}

// Where performs a query with "Where" parameters.
func (orm ORM) Where(query interface{}, args ...interface{}) *gorm.DB {
	return orm.conn.Where(query, args...)
}

// Create inserts a new model instance into the database.
func (orm ORM) Create(model interface{}) *gorm.DB {
	return orm.conn.Create(model)
}

// Save updates a model with the given attributes.
func (orm ORM) Save(value interface{}) *gorm.DB {
	return orm.conn.Save(value)
}

// Model specifies the domain model that subsequent queries will be run against.
func (orm ORM) Model(model interface{}) *gorm.DB {
	return orm.conn.Model(model)
}

// First returns the first model (ordered by ID) that matches the specified query.
func (orm ORM) First(model interface{}, where ...interface{}) *gorm.DB {
	return orm.conn.First(model, where...)
}

// ModelWithId returns an instance of the specified model with the given ID.
func (orm ORM) ModelWithId(model interface{}, id uint) error {
	// First check if the Model exists.
	// We do this so that we can avoid an error returned by the ORM
	// when a query returns no results.
	if exists, err := orm.ModelExistsWithId(model, id); err != nil {
		return err
	} else if !exists {
		return nil
	}

	// It exists, so let's load it
	if err := orm.First(model, id).Error; err != nil {
		return err
	}

	return nil
}

// ModelExistsWithId returns a boolean indicating if an instance of the
// specified model exists with a given ID.
func (orm ORM) ModelExistsWithId(model interface{}, id uint) (bool, error) {
	var count int64

	err := orm.Model(model).Where(id).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

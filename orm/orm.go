// Package orm manages access to a database, including ORM-like functionality.
package orm

import (
	"errors"
	"reflect"

	"github.com/KyleBanks/go-kit/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type ORM struct {
	conn *gorm.DB
}

// Model is a type added to domain models that will provide ORM functionality.
type Model struct {
	gorm.Model
}

// Open creates a database connection, or returns an existing one if present.
func (orm *ORM) Open(dialect, connectionString string) *gorm.DB {
	if orm.conn != nil {
		return orm.conn
	}

	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		log.Errorf("Error opening database connection: %v", err)
		panic(err)
	} else if db == nil {
		err := errors.New("Database handle is nil!")
		log.Errorf("Error opening database connection: %v", err)
		panic(err)
	}

	log.Infof("Database connection established: {Dialect: %v, ConnectionString: %v}", dialect, connectionString)

	// Enable logging
	db.SetLogger(log.Logger)
	db.LogMode(true)

	// Configure
	// TODO: Accept options as a param to Open
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(0) // Unlimited

	orm.conn = db
	return orm.conn
}

// AutoMigrate performs database migration for all Model types provided.
func (orm ORM) AutoMigrate(models []interface{}) error {
	for _, model := range models {
		modelName := reflect.Indirect(reflect.ValueOf(model)).Type()
		log.Info("Migrating model:", modelName)

		if err := orm.conn.AutoMigrate(model).Error; err != nil {
			log.Error("AutoMigrate failed for model", modelName)
			return err
		}
		log.Info("Model migrated:", modelName)
	}

	return nil
}

// Exec performs a raw SQL query against the underlying database.
func (orm ORM) Exec(query string, output interface{}) *gorm.DB {
	return orm.conn.Raw(query).Scan(output)
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

// Last returns the last model (ordered by ID) that matches the specified query.
func (orm ORM) Last(model interface{}, where ...interface{}) *gorm.DB {
	return orm.conn.Last(model, where...)
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

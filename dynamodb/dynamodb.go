// Package "dynamodb" provides a wrapper for the AWS DynamoDB SDK.
package dynamodb

//import (
//	"github.com/goamz/goamz/dynamodb"
//	"github.com/goamz/goamz/aws"
//	"time"
//	"errors"
//)
//
//const (
//	HashKey = "HASH"
//	RangeKey = "RANGE"
//
//	AttributeTypeString = "S"
//	AttributeTypeNumber = "N"
//	AttributeTypeBinary = "B"
//)
//
//var (
//	ErrTableNotFound = errors.New("ResourceNotFoundException: Cannot do operations on a non-existent table")
//)
//
//type Dynamo struct {
//	*dynamodb.Server
//}
//
//type DynamoConfiguration struct {
//	// The Region to use.
//	Region aws.Region
//
//	// Optional: Set a specific endpoint to use as the DynamoDB host.
//	Endpoint string
//}
//
//// New initializes and returns an new Dynamo connection.
//func New(c DynamoConfiguration) (*Dynamo, error) {
//	auth, err := aws.GetAuth("", "", "", time.Now().Add(time.Hour * 24 * 3))
//	if err != nil {
//		return nil, err
//	}
//
//	d := &dynamodb.Server{
//		Auth: auth,
//		Region: c.Region,
//	}
//
//	if len(c.Endpoint) > 0 {
//		d.Region.DynamoDBEndpoint = c.Endpoint
//	}
//
//	return &Dynamo{d}, nil
//}
//
//// TableExists returns a boolean indicating if a particular table exists.
//func (d Dynamo) TableExists(tableName string) (bool, error) {
//	if _, err := d.DescribeTable(tableName); err != nil && err.Error() != ErrTableNotFound.Error() {
//		return false, err
//	} else if err != nil && err.Error() == ErrTableNotFound.Error() {
//		return false, nil
//	}
//
//	return true, nil
//}
//
//// CreateTableIfNotExists creates a new table with the specified definition, if it doesn't already exists. A boolean
//// is returned that indicates if a new table was created.
//func (d Dynamo) CreateTableIfNotExists(table dynamodb.TableDescriptionT) (bool, error) {
//	if exists, err := d.TableExists(table.TableName); err != nil {
//		return false, err
//	} else if exists {
//		return false, nil
//	}
//
//	if _, err := d.CreateTable(table); err != nil {
//		return false, err
//	}
//
//	return true, nil
//}
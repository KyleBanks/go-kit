package dynamodb

//import (
//	"testing"
//	"github.com/goamz/goamz/aws"
//	"fmt"
//	"time"
//	"github.com/goamz/goamz/dynamodb"
//)
//
//func TestNew(t *testing.T) {
//	// Works with minimal configuration
//	_, err := New(DynamoConfiguration{
//		Region: aws.USEast,
//	})
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Works with local endpoint
//	newLocal(t)
//}
//
//
//func TestDynamo_TableExists(t *testing.T) {
//	// Ensure false returned for an unknown table
//	tableName := fmt.Sprintf("unknown%v", time.Now().Unix())
//	d := newLocal(t)
//
//	if exists, err := d.TableExists(tableName); err != nil {
//		t.Fatal(err)
//	} else if exists {
//		t.Fatal("Expected false to be returned for non-existent table; %v", tableName)
//	}
//
//	// Ensure a table that exists return true
//	table := dynamodb.TableDescriptionT{
//		TableName: tableName,
//		KeySchema: []dynamodb.KeySchemaT{
//			dynamodb.KeySchemaT{
//				AttributeName: "test",
//				KeyType: HashKey,
//			},
//		},
//		AttributeDefinitions: []dynamodb.AttributeDefinitionT{
//			dynamodb.AttributeDefinitionT{
//				Name: "test",
//				Type: AttributeTypeString,
//			},
//		},
//		ProvisionedThroughput: dynamodb.ProvisionedThroughputT{
//			WriteCapacityUnits: 1,
//			ReadCapacityUnits: 1,
//		},
//	}
//	if _, err := d.CreateTable(table); err != nil {
//		t.Fatal(err)
//	}
//
//	if exists, err := d.TableExists(tableName); err != nil {
//		t.Fatal(err)
//	} else if !exists {
//		t.Fatalf("Expected true to be returned for existent table: %v", table.TableName)
//	}
//}
//
//func TestDynamo_CreateTableIfNotExists(t *testing.T) {
//	table := dynamodb.TableDescriptionT{
//		TableName: fmt.Sprintf("testcreate%v", time.Now().Unix()),
//		KeySchema: []dynamodb.KeySchemaT{
//			dynamodb.KeySchemaT{
//				AttributeName: "test",
//				KeyType: HashKey,
//			},
//		},
//		AttributeDefinitions: []dynamodb.AttributeDefinitionT{
//			dynamodb.AttributeDefinitionT{
//				Name: "test",
//				Type: AttributeTypeString,
//			},
//		},
//		ProvisionedThroughput: dynamodb.ProvisionedThroughputT{
//			WriteCapacityUnits: 1,
//			ReadCapacityUnits: 1,
//		},
//	}
//	d := newLocal(t)
//
//	// Test with a new table
//	if created, err := d.CreateTableIfNotExists(table); err != nil {
//		t.Fatal(err)
//	} else if !created {
//		t.Fatal("Expected true to be returned for CreateTableIfNotExists when passing an unexisting table.")
//	}
//
//	// Test with an existing table
//	if created, err := d.CreateTableIfNotExists(table); err != nil {
//		t.Fatal(err)
//	} else if created {
//		t.Fatal("Excepted false to be returned for CreateTableIfNotExists when passing an existing table.")
//	}
//}
//
//func newLocal(t *testing.T) *Dynamo {
//	d, err := New(DynamoConfiguration{
//		Region: aws.USEast,
//		Endpoint: "http://localhost:8999",
//	})
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	return d
//}
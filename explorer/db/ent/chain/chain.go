// Code generated by ent, DO NOT EDIT.

package chain

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the chain type in the database.
	Label = "chain"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldChainID holds the string denoting the chain_id field in the database.
	FieldChainID = "chain_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the chain in the database.
	Table = "chains"
)

// Columns holds all SQL columns for chain fields.
var Columns = []string{
	FieldID,
	FieldChainID,
	FieldCreatedAt,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)

// OrderOption defines the ordering options for the Chain queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByChainID orders the results by the chain_id field.
func ByChainID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChainID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

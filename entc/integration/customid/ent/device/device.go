// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package device

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/schema"
)

const (
	// Label holds the string label denoting the device type in the database.
	Label = "device"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeActiveSession holds the string denoting the active_session edge name in mutations.
	EdgeActiveSession = "active_session"
	// EdgeSessions holds the string denoting the sessions edge name in mutations.
	EdgeSessions = "sessions"
	// Table holds the table name of the device in the database.
	Table = "devices"
	// ActiveSessionTable is the table that holds the active_session relation/edge.
	ActiveSessionTable = "devices"
	// ActiveSessionInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	ActiveSessionInverseTable = "sessions"
	// ActiveSessionColumn is the table column denoting the active_session relation/edge.
	ActiveSessionColumn = "device_active_session"
	// SessionsTable is the table that holds the sessions relation/edge.
	SessionsTable = "sessions"
	// SessionsInverseTable is the table name for the Session entity.
	// It exists in this package in order to avoid circular dependency with the "session" package.
	SessionsInverseTable = "sessions"
	// SessionsColumn is the table column denoting the sessions relation/edge.
	SessionsColumn = "device_sessions"
)

// Columns holds all SQL columns for device fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "devices"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"device_active_session",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() schema.ID
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func([]byte) error
)

// Order defines the ordering method for the Device queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByActiveSessionField orders the results by active_session field.
func ByActiveSessionField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newActiveSessionStep(), sql.OrderByField(field, opts...))
	}
}

// BySessionsCount orders the results by sessions count.
func BySessionsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSessionsStep(), opts...)
	}
}

// BySessions orders the results by sessions terms.
func BySessions(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSessionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newActiveSessionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ActiveSessionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ActiveSessionTable, ActiveSessionColumn),
	)
}
func newSessionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SessionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SessionsTable, SessionsColumn),
	)
}

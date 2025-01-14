// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package tweettag

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the tweettag type in the database.
	Label = "tweet_tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAddedAt holds the string denoting the added_at field in the database.
	FieldAddedAt = "added_at"
	// FieldTagID holds the string denoting the tag_id field in the database.
	FieldTagID = "tag_id"
	// FieldTweetID holds the string denoting the tweet_id field in the database.
	FieldTweetID = "tweet_id"
	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"
	// EdgeTweet holds the string denoting the tweet edge name in mutations.
	EdgeTweet = "tweet"
	// Table holds the table name of the tweettag in the database.
	Table = "tweet_tags"
	// TagTable is the table that holds the tag relation/edge.
	TagTable = "tweet_tags"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "tag_id"
	// TweetTable is the table that holds the tweet relation/edge.
	TweetTable = "tweet_tags"
	// TweetInverseTable is the table name for the Tweet entity.
	// It exists in this package in order to avoid circular dependency with the "tweet" package.
	TweetInverseTable = "tweets"
	// TweetColumn is the table column denoting the tweet relation/edge.
	TweetColumn = "tweet_id"
)

// Columns holds all SQL columns for tweettag fields.
var Columns = []string{
	FieldID,
	FieldAddedAt,
	FieldTagID,
	FieldTweetID,
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
	// DefaultAddedAt holds the default value on creation for the "added_at" field.
	DefaultAddedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Order defines the ordering method for the TweetTag queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAddedAt orders the results by the added_at field.
func ByAddedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldAddedAt, opts...).ToFunc()
}

// ByTagID orders the results by the tag_id field.
func ByTagID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTagID, opts...).ToFunc()
}

// ByTweetID orders the results by the tweet_id field.
func ByTweetID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTweetID, opts...).ToFunc()
}

// ByTagField orders the results by tag field.
func ByTagField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagStep(), sql.OrderByField(field, opts...))
	}
}

// ByTweetField orders the results by tweet field.
func ByTweetField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTweetStep(), sql.OrderByField(field, opts...))
	}
}
func newTagStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
	)
}
func newTweetStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TweetInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, TweetTable, TweetColumn),
	)
}

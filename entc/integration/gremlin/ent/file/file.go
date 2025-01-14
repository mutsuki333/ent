// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package file

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "fsize"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUser holds the string denoting the user field in the database.
	FieldUser = "user"
	// FieldGroup holds the string denoting the group field in the database.
	FieldGroup = "group"
	// FieldOp holds the string denoting the op field in the database.
	FieldOp = "op"
	// FieldFieldID holds the string denoting the field_id field in the database.
	FieldFieldID = "field_id"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeType holds the string denoting the type edge name in mutations.
	EdgeType = "type"
	// EdgeField holds the string denoting the field edge name in mutations.
	EdgeField = "field"
	// OwnerInverseLabel holds the string label denoting the owner inverse edge type in the database.
	OwnerInverseLabel = "user_files"
	// TypeInverseLabel holds the string label denoting the type inverse edge type in the database.
	TypeInverseLabel = "file_type_files"
	// FieldLabel holds the string label denoting the field edge type in the database.
	FieldLabel = "file_field"
)

var (
	// DefaultSize holds the default value on creation for the "size" field.
	DefaultSize int
	// SizeValidator is a validator for the "size" field. It is called by the builders before save.
	SizeValidator func(int) error
)

// Order defines the ordering method for the File queries.
type Order func(*dsl.Traversal)

// comment from another template.

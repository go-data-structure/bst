package bst

import "errors"

var (
	// ErrKeyExists key already exists
	ErrKeyExists = errors.New("key already exists")
	// ErrKeyNotExists key not exists
	ErrKeyNotExists = errors.New("key not exists")
)

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package postgresql

import ()

type Entry struct {
	ID          int64
	Uuid        string
	Title       string
	Content     string
	ContentType string
	IsEncrypted string
	InsertDate  interface{}
}

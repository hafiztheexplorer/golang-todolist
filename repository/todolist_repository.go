package repository

import (
	"context"
	"database/sql"
	"golang-todolist/model/domain"
)

// ini adalah
type TodolistRepository interface {
	Create(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist // action yang dibutuhkan todolist semua, return data nya juga semua dari domain.Todolist
	Update(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist // action yang dibutuhkan todolist semua, return data nya juga semua dari domain.Todolist
	Delete(ctx context.Context, tx *sql.Tx, todolist domain.Todolist)                 // action yang dibutuhkan todolist semua
	FindById(ctx context.Context, tx *sql.Tx, todolistID int64) domain.Todolist       // action yang dibutuhkan id todolist semua, return data juga semua
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Todolist                        // karena ini findAll ya otomatis return semua data
}

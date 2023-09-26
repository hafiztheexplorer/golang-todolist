package repository

import (
	"context"
	"database/sql"
	"golang-todolist/model/domain"
)

type TodolistRepositoryImplementation struct{}

func (repository *TodolistRepositoryImplementation) Create(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist {
	SQLScript := "INSERT INTO `todolist` (`todolisttitle`) values (?) "
	tx.ExecContext(ctx, SQLScript, todolist.TodolistTitle)
	panic("not implemented") // TODO: Implement
}

func (repository *TodolistRepositoryImplementation) Update(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) domain.Todolist {
	panic("not implemented") // TODO: Implement
}

func (repository *TodolistRepositoryImplementation) Delete(ctx context.Context, tx *sql.Tx, todolist domain.Todolist) {
	panic("not implemented") // TODO: Implement
}

func (repository *TodolistRepositoryImplementation) FindById(ctx context.Context, tx *sql.Tx, todolistID int64) domain.Todolist {
	panic("not implemented") // TODO: Implement
}

func (repository *TodolistRepositoryImplementation) FindAll(ctx context.Context, tx *sql.Tx) []domain.Todolist {
	panic("not implemented") // TODO: Implement
}

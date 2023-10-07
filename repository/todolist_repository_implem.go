package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang-todolist/model/domain"
)

type TodolistRepositoryImplementation struct{}

func NewTodolistRepository() TodolistRepository {
	return &TodolistRepositoryImplementation{}
}

func (repository *TodolistRepositoryImplementation) Create(ctx context.Context, tx *sql.Tx, Todolist domain.Todolist) domain.Todolist {
	// tujuan akhir yaitu kita kirim perintah sqlnya seperti apa dengan menggunakan golang context, lalu kita return valuenya ke Todolist domain.Todolist
	// kita set script sqlnya
	SQLScript := "INSERT INTO `todolist` (`todolisttitle`) values (?) "

	// kita gunakan tx.ExecContext untuk mengirimkan perintah ke database berupa context yang, dengan script seperti di SQLscript (mengirim nama todolist di todolisttitle pada tabel todolist), lalu tujuan databasenya di package repository di file domain.Todolist (digantikan dengan todolist) tepatnya pada kolom TodolistTitle
	hasil, _ := tx.ExecContext(ctx, SQLScript, Todolist.TodolistTitle)
	fmt.Println(hasil)

	// karena baris data ini ada auto_increment nya, maka sebelum kita return agar dieksekusi, kita masukkan ke function LastInsertId, agar ditempatkan di baris setelah id terakhir pada database,
	id, _ := hasil.LastInsertId()
	fmt.Println(id)

	// kita kirim menjadi return value dengan cara menggabungkan semuanya, menjadi Todolist, karena di dalamnya ada "hasil" yang digabung dengan informasi last insert id "id"
	Todolist.Id = id
	fmt.Println(Todolist.Id)
	return Todolist

	// ingat return statement harus sesuai domain.Todolist either domain.Todolist.Id or domain.Todolist.TodolistTitle
}

func (repository *TodolistRepositoryImplementation) Update(ctx context.Context, tx *sql.Tx, Todolist domain.Todolist) domain.Todolist {
	// tujuan akhir yaitu kita kirim perintah sqlnya seperti apa dengan menggunakan golang context, lalu kita return valuenya ke Todolist domain.Todolist
	// kita set script sqlnya
	SQLScript := "update `todolist` set `todolisttitle` = `?` where `id` = `?`"

	// kita gunakan tx.ExecContext untuk mengirimkan perintah ke database berupa context yang, dengan script seperti di SQLscript (mengirim nama todolist di todolisttitle pada tabel todolist), lalu tujuan databasenya di package repository di file domain.Todolist (digantikan dengan todolist) tepatnya pada kolom TodolistTitle
	hasil, _ := tx.ExecContext(ctx, SQLScript, Todolist.TodolistTitle)
	fmt.Println(hasil)

	// karena ini update data existing, tidak ada logic seperti pada create untuk memberi tahu next id yang tersedia, hanya edit data yang sudah ada
	return Todolist

	// ingat return statement harus sesuai domain.Todolist either domain.Todolist.Id or domain.Todolist.TodolistTitle
}

func (repository *TodolistRepositoryImplementation) Delete(ctx context.Context, tx *sql.Tx, Todolist domain.Todolist) {
	// tujuan akhir yaitu kita kirim perintah sqlnya seperti apa dengan menggunakan golang context, lalu kita return valuenya ke Todolist domain.Todolist
	// kita set script sqlnya
	SQLScript := "delete from `todolist` where id=`?`;"

	// kita gunakan tx.ExecContext untuk mengirimkan perintah ke database berupa context yang, dengan script seperti di SQLscript (mengirim nama todolist di todolisttitle pada tabel todolist), lalu tujuan databasenya di package repository di file domain.Todolist (digantikan dengan todolist) tepatnya pada kolom TodolistTitle
	hasil, _ := tx.ExecContext(ctx, SQLScript, Todolist.TodolistTitle)
	fmt.Println(hasil)

	// karena ini delete data existing, tidak perlu return data apapun
}

func (repository *TodolistRepositoryImplementation) FindById(ctx context.Context, tx *sql.Tx, TodolistID int64) domain.Todolist {
	panic("not implemented") // TODO: Implement
}

func (repository *TodolistRepositoryImplementation) FindAll(ctx context.Context, tx *sql.Tx) []domain.Todolist {
	panic("not implemented") // TODO: Implement
}

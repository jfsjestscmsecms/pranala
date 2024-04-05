package numfive

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type Worker struct {
	ID   string
	Name string
}

const conn_string = "postgres://username:passwd@localhost:5432/database_name"

type DB struct {
	conn_string string
}

func (x *DB) Create(d Worker) error {
	conn, err := pgx.Connect(context.Background(), x.conn_string)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "insert into worker (id,name) values($1,$2)", d.ID, d.Name)
	if err != nil {
		return err
	}

	return nil
}

func (x *DB) Read(id string) (Worker, error) {
	var w Worker
	conn, err := pgx.Connect(context.Background(), x.conn_string)
	if err != nil {
		return w, err
	}

	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(), "select id,name from worker where id=$1").Scan(&w.ID, &w.Name)
	if err != nil {
		return w, err
	}

	return w, nil
}

func (x *DB) Update(w Worker, id string) error {
	conn, err := pgx.Connect(context.Background(), x.conn_string)
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "update worker set id=$1,name=$2 where id=$3", w.ID, w.Name, id)
	if err != nil {
		return err
	}

	return nil
}

func (x *DB) Delete(id string) error {
	conn, err := pgx.Connect(context.Background(), x.conn_string)
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), "delete from worker where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

func NewDB() *DB {
	return &DB{conn_string: conn_string}
}

func NumFive() {
	p1 := Worker{ID: "1", Name: "Albert"}

	db := NewDB()

	err := db.Create(p1)
	if err != nil {
		log.Fatalln()
	}

	d, err := db.Read(p1.ID)
	if err != nil {
		log.Fatalln()
	}

	d.Name = "smith"
	err = db.Update(d, p1.ID)
	if err != nil {
		log.Fatalln()
	}

	err = db.Delete(p1.ID)
	if err != nil {
		log.Fatalln()
	}
}

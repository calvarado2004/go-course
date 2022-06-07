package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	//connect to a database
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=postgres user=postgres password=")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	defer conn.Close()

	log.Println("Connected to the Database")

	//test connection

	err = conn.Ping()
	if err != nil {
		log.Fatal("can not ping database")
	}

	log.Println("Database pinged!")

	//get rows from table

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	//insert a row

	query := `insert into todos.users (id,first_name,last_name) values ($1,$2,$3)`
	_, err = conn.Exec(query, 3, "Cristina", "Fernandez")
	if err != nil {
		log.Fatal("row not inserted due to an error", err)
	}

	log.Println("Inserted a row")

	//get rows from a table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	//update a row
	stmt := `update todos.users set last_name = $1 where last_name = $2`
	_, err = conn.Exec(stmt, "Gutierrez", "Fernandez")
	if err != nil {
		log.Fatal("row not updated due to an error", err)
	}

	log.Println("Updated a row")

	//get rows from a table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	//get one row by id
	query = `select id, first_name, last_name from todos.users where id = $1`

	var id int
	var firstName, lastName string

	row := conn.QueryRow(query, 1)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal("row not selected due to an error", err)
	}

	log.Println("QueryRow returns", id)

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	//delete a row
	query = `delete from todos.users where last_name = $1`
	_, err = conn.Exec(query, "Gutierrez")
	if err != nil {
		log.Fatal("row not deleted due to an error", err)
	}

	log.Println("Rows deleted")

	//get rows from a table
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

}

func getAllRows(conn *sql.DB) error {

	rows, err := conn.Query("select * from todos.users")
	if err != nil {
		log.Println(err)
		return err
	}

	defer rows.Close()

	var id int
	var first_name string
	var last_name string

	for rows.Next() {
		err := rows.Scan(&id, &first_name, &last_name)
		if err != nil {
			log.Println(err)
		}

		fmt.Println("Record is", id, first_name, last_name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("error scanning rows", err)
	}

	fmt.Println("-----------------------------")

	return nil
}

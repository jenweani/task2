package models

import (
	"database/sql"
	"errors"
)

// Person struct
type Person struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

// Person struct methods
// Insert a new Person field 
func (person *Person) CreatePerson(db *sql.DB) error {
	query := `
		INSERT INTO persons (name) VALUES ($1)
	`
	_, err := db.Exec(query, person.Name)
	person.ReadPerson(db, person.Id, person.Name)
	if err != nil {return err}
	return nil
}

// retrieve person row from db using id or name
func (person *Person) ReadPerson(db *sql.DB, Id int, name string) error {
	if Id == 0 && name != "" {
		err := db.QueryRow(`SELECT * FROM persons WHERE persons.name = $1`, name).Scan(&person.Id, &person.Name)
		if err != nil {return err}
	}else if (name == "" && Id != 0) || (name != "" && Id != 0) {
		err := db.QueryRow(`SELECT * FROM persons WHERE persons.id = $1`, Id).Scan(&person.Id, &person.Name)
		if err != nil {return err}
	}else {
		return errors.New("an id or name variable is needed to query db")
	}
	return nil
}

// update a person row in db using id 
func (person *Person) UpdatePerson(db *sql.DB, Id int, name string) error {
	if Id != 0{
		_, err := db.Exec(`UPDATE persons SET name = $1 WHERE id = $2;`, name, Id)
		if err != nil {return err}
	}else {
		return errors.New("an id variable is needed to query db")
	}
	person.Id = Id
	person.Name = name
	return nil
}

// delete a person row from db using id
func (person *Person) DeletePerson(db *sql.DB, Id int) error {
	if Id != 0{
		_, err := db.Exec(`DELETE FROM persons WHERE persons.id = $1`, Id)
		if err != nil {return err}
	}else {
		return errors.New("an id or name variable is needed to query db")
	}
	return nil
}
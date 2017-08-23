package main

import (
	"database/sql"
)

var conexion *sql.DB

func conectarBaseDatos() {
	db, err := sql.Open("sqlite3", "registro.personas.sqlite")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE  TABLE  IF NOT EXISTS persona(nombre text, ciudad text, genero text)")
	if err != nil {
		panic(err)
	}
	conexion = db
}
func cerrarBaseDatos() {
	conexion.Close()
}
func registrarPersona(persona *Persona) (int64, error) {
	stmt, err := conexion.Prepare("INSERT INTO persona(nombre, ciudad, genero) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(persona.Nombre, persona.Ciudad, persona.Genero)
	if err != nil {
		return 0, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}
func personas() ([]Persona, error) {
	personas := []Persona{}
	rows, err := conexion.Query("SELECT * FROM persona")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := Persona{}
		err := rows.Scan(&p.Nombre, &p.Ciudad, &p.Genero)
		if err != nil {
			return nil, err
		}
		personas = append(personas, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return personas, nil
}

package service

import "database/sql"

type Gokyo struct {
	ID         int
	Name       string
	Serie      int
	Difficulty string
}

func NewGokyoService(db *sql.DB) *GokyoService {
	return &GokyoService{db: db}
}

type GokyoService struct {
	db *sql.DB
}

func (s *GokyoService) GetStatus() (string, error) {
	return "API is Up", nil
}

func (s *GokyoService) GetGokyo() ([]Gokyo, error) {
	query := "Select id, name, serie, difficulty from gokyo "
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	var gokyos []Gokyo
	for rows.Next() {
		var gokyo Gokyo
		err := rows.Scan(&gokyo.ID, &gokyo.Name, &gokyo.Serie, &gokyo.Difficulty)
		if err != nil {
			return nil, err
		}
		gokyos = append(gokyos, gokyo)
	}
	return gokyos, nil
}

func (s *GokyoService) CreateGokyo(gokyo *Gokyo) error {
	query := "Insert into gokyo (name,serie,difficulty) values (?,?,?) "
	result, err := s.db.Exec(query, gokyo.Name, gokyo.Serie, gokyo.Difficulty)
	if err != nil {
		return err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	gokyo.ID = int(lastInsertId)
	return nil
}

func (s *GokyoService) DeleteGokyo(id int) error {
	query := "delete from gokyo where id = ? "
	_, err := s.db.Exec(query, id)
	return err
}

func (s *GokyoService) UpdateGokyo(gokyo *Gokyo) error {
	query := "Update Gokyo set Name = ?, Serie = ? , Difficulty = ? where id = ?"
	_, err := s.db.Exec(query, gokyo.Name, gokyo.Serie, gokyo.Difficulty, gokyo.ID)
	return err
}

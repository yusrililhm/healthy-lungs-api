package symtomp_pg

import (
	"database/sql"
	"expert_systems_api/entity"
	"expert_systems_api/pkg/exception"
	"expert_systems_api/symtomp/symtomp_repo"
	"log"
)

type symtompPg struct {
	db *sql.DB
}

const (
	fetchSymtompQuery = `
		select id, name, description, created_at, updated_at from symtomp
	`
)

func NewSymtompPg(db *sql.DB) symtomp_repo.SymtompRepo {
	return &symtompPg{
		db: db,
	}
}

// Fetch implements symtomp_repo.SymtompRepo.
func (s *symtompPg) Fetch() ([]*entity.Symtomp, exception.Exception) {

	symtomps := []*entity.Symtomp{}

	rows, err := s.db.Query(fetchSymtompQuery)

	if err != nil {
		log.Println(err.Error())
		return nil, exception.NewInternalServerError("something went wrong")
	}

	for rows.Next() {
		symtomp := entity.Symtomp{}

		if err := rows.Scan(
			&symtomp.Id,
			&symtomp.Name,
			&symtomp.Description,
			&symtomp.CreatedAt,
			&symtomp.UpdatedAt,
		); err != nil {
			log.Println(err.Error())
			return nil, exception.NewInternalServerError("something went wrong")
		}

		symtomps = append(symtomps, &symtomp)
	}

	defer rows.Close()

	return symtomps, nil
}

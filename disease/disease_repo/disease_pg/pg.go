package disease_pg

import (
	"database/sql"
	"expert_systems_api/disease/disease_repo"
	"expert_systems_api/entity"
	"expert_systems_api/pkg/exception"
	"log"
)

type diseasePg struct {
	db *sql.DB
}

const (
	fetchDiseaseQuery = ``
)

func NewDiseasePg(db *sql.DB) disease_repo.DiseaseRepo {
	return &diseasePg{
		db: db,
	}
}

// Fetch implements disease_repo.DiseaseRepo.
func (d *diseasePg) Fetch() ([]*entity.Disease, exception.Exception) {
	diseases := []*entity.Disease{}

	rows, err := d.db.Query(fetchDiseaseQuery)

	if err != nil {
		log.Println(err.Error())
		return nil, exception.NewInternalServerError("something went wrong")
	}

	for rows.Next() {
		disease := entity.Disease{}

		if err := rows.Scan(); err != nil {
			log.Println(err.Error())
			return nil, exception.NewInternalServerError("something went wrong")
		}

		diseases = append(diseases, &disease)
	}

	return diseases, nil
}

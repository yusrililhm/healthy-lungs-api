package user_pg

import (
	"database/sql"
	"log"

	"expert_systems_api/entity"
	"expert_systems_api/pkg/exception"
	"expert_systems_api/user/user_repo"
)

type userPg struct {
	db *sql.DB
}

const (
	addUserQuery = `INSERT INTO "user" (full_name, email, password, role) VALUES($1, $2, $3, $4)`

	fetchByEmailQuery = `SELECT id, full_name, email, password, role, created_at, updated_at FROM "user" WHERE email = $1`

	fetchByIdQuery = `SELECT id, full_name, email, password, role, created_at, updated_at FROM "user" WHERE id = $1`

	modifyUserQuery = `UPDATE FROM "user" SET full_name = $2, email = $3 WHERE id = $1`

	changePasswordQuery = `UPDATE FROM "user" SET password = $2 WHERE`
)

func NewUserRepo(db *sql.DB) user_repo.UserRepo {
	return &userPg{
		db: db,
	}
}

// Add implements user_repo.UserRepo.
func (pg *userPg) Add(user *entity.User) exception.Exception {

	tx, err := pg.db.Begin()

	if err != nil {
		tx.Rollback()
		log.Println("[Warn]", err.Error())
		return exception.NewInternalServerError("something went wrong")
	}

	stmt, err := tx.Prepare(addUserQuery)

	if err != nil {
		tx.Rollback()
		log.Println("[Warn]", err.Error())
		return exception.NewInternalServerError("something went wrong")
	}

	_, err = stmt.Exec(user.FullName, user.Email, user.Password, user.Role)

	if err != nil {
		// if err ==  {
		// 	tx.Rollback()
		// 	log.Println("[Warn]", err.Error())
		// 	return exception.NewConflictError("email has been used")
		// }

		tx.Rollback()
		log.Println("[Warn]", err.Error())
		return exception.NewInternalServerError("something went wrong")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return exception.NewInternalServerError("something went wrong")
	}

	return nil
}

// ChangePassword implements user_repo.UserRepo.
func (pg *userPg) ChangePassword(id int, newPassword string) exception.Exception {
	tx, err := pg.db.Begin()

	if err != nil {
		tx.Rollback()
		log.Println("[Warn]", err.Error())
		return exception.NewInternalServerError("something went wrong")
	}

	stmt, err := tx.Prepare(changePasswordQuery)

	if err != nil {
		tx.Rollback()
		log.Println("[Warn]", err.Error())
		return exception.NewInternalServerError("something went wrong")
	}

	_, err = stmt.Exec(id, newPassword)

	if err != nil {
		tx.Rollback()
		log.Println("[Warn]", err.Error())
		return exception.NewInternalServerError("something went wrong")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return exception.NewInternalServerError("something went wrong")
	}

	return nil
}

// FetchByEmail implements user_repo.UserRepo.
func (pg *userPg) FetchByEmail(email string) (*entity.User, exception.Exception) {

	user := entity.User{}

	row := pg.db.QueryRow(fetchByEmailQuery, email)

	if err := row.Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NewNotFoundError("user not found")
		}
		return nil, exception.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

// FetchById implements user_repo.UserRepo.
func (pg *userPg) FetchById(id int) (*entity.User, exception.Exception) {

	user := entity.User{}

	row := pg.db.QueryRow(fetchByIdQuery, id)

	if err := row.Scan(
		&user.Id,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NewNotFoundError("user not found")
		}
		return nil, exception.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

// Modify implements user_repo.UserRepo.
func (pg *userPg) Modify(id int, user *entity.User) exception.Exception {

	tx, err := pg.db.Begin()

	if err != nil {
		tx.Rollback()
		log.Println("[Warn]", err.Error())
		return exception.NewInternalServerError("something went wrong")
	}

	stmt, err := tx.Prepare(modifyUserQuery)

	if err != nil {
		tx.Rollback()
		log.Println("[Warn]", err.Error())
		return exception.NewInternalServerError("something went wrong")
	}

	_, err = stmt.Exec(id, user.FullName, user.Email)

	if err != nil {
		tx.Rollback()
		log.Println("[Warn]", err.Error())
		return exception.NewInternalServerError("something went wrong")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return exception.NewInternalServerError("something went wrong")
	}

	return nil
}

package repositories

import (
	"database/sql"
	"github.com/pcavezzan/sample-apiserver/app/models"
	"github.com/pcavezzan/sample-apiserver/log"
)

type ParkingRepository interface {
	GetAll() []models.Parking
}

func NewInMemoryParkingRepository() ParkingRepository {
	r := InMemoryParkingRepository{}
	r.Data = []models.Parking{
		{Number: 1, Owner: "Foo Bar", Car: "Dodge Viper"},
		{Number: 2, Owner: "James Bond", Car: "Aston Martin"},
	}
	return r
}

type InMemoryParkingRepository struct {
	Data []models.Parking
}

func (r InMemoryParkingRepository) GetAll() []models.Parking {
	return r.Data
}

type SqlParkingRepository struct {
	*sql.DB
}

func NewSqlParkingRepository(db *sql.DB) SqlParkingRepository {
	return SqlParkingRepository{
		DB: db,
	}
}

func (r SqlParkingRepository) GetAll() []models.Parking {
	res := make([]models.Parking, 0)
	tx, err := r.DB.Begin()
	if err != nil {
		log.Error(err)
		return res
	}

	stmt, err := tx.Prepare("SELECT number, owner, car FROM parking")
	if err != nil {
		log.Error(err)
		return rollbackTx(tx)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Error(err)
		return rollbackTx(tx)
	}
	defer rows.Close()

	for rows.Next() {
		d := models.Parking{}
		err = rows.Scan(
			&d.Number,
			&d.Owner,
			&d.Car,
		)
		res = append(res, d)
	}

	if err = rows.Err(); err != nil {
		log.Error(err)
		return []models.Parking{}
	}

	return res

}

func rollbackTx(tx *sql.Tx) []models.Parking {
	if ok := tx.Rollback(); ok != nil {
		log.Error("Could not rollback transaction", ok)
	}
	return []models.Parking{}
}
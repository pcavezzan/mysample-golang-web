package cli

import "github.com/pcavezzan/sample-apiserver/log"

const (
	create_table_parking = `
CREATE TABLE IF NOT EXISTS parking (
 id serial PRIMARY KEY,
 number SMALLINT UNIQUE,
 owner VARCHAR,
 car VARCHAR
);
`
	create_sample_parking = `
INSERT INTO parking (number, owner, car)
VALUES (29, 'James Bond', 'Aston Martin'),
(30, 'Foo Bar', 'Porsche 911');
`
)

func initializeDb() {
	db := factory.DataSource()

	// Should assemble our statement into one
	q := create_table_parking + create_sample_parking

	result, err := db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}
	nbRowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}
	log.Info("Init completed", nbRowsAffected)
}
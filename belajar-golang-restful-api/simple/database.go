package simple

type Database struct {
	Name string
}

func NewDatabasPostgreSQL() *DatabasePostgreSQL {
	return &DatabasePostgreSQL{Name: "PostgreSQL"}
}

func NewDatabasMySQL() *DatabaseMySQL {
	return &DatabaseMySQL{Name: "MySQL"}
}

type DatabasePostgreSQL Database
type DatabaseMySQL Database

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMySQL      *DatabaseMySQL
}

func NewDatabaseRepository(databasePostgreSQL *DatabasePostgreSQL, databaseMySQL *DatabaseMySQL) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: databasePostgreSQL, DatabaseMySQL: databaseMySQL}
}

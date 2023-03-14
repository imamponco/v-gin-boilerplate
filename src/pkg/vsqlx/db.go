package vsqlx

import (
	"fmt"
	"github.com/imamponco/v-gin-boilerplate/src/svc/contract"
	"github.com/jmoiron/sqlx"
	"log"
)

func InitDatabase(config *contract.Config) (*sqlx.DB, error) {

	// Set database default config
	setDefault(config)

	// Get datasource name
	dsn, err := getDSN(config)
	if err != nil {
		log.Fatalln(err)
	}

	// Connect to datasource
	conn, err := sqlx.Connect(config.DatabaseDriver, dsn)
	if err != nil {
		log.Fatalln(err)
	}

	return conn, err
}

func getDSN(c *contract.Config) (dsn string, err error) {
	switch c.DatabaseDriver {
	case DriverMySQL:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.DatabaseUsername, c.DatabasePassword, c.DatabaseHost, c.DatabasePort,
			c.DatabaseName)
	case DriverPostgreSQL:
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.DatabaseHost, c.DatabasePort,
			c.DatabaseUsername, c.DatabasePassword, c.DatabaseName)
	default:
		err = fmt.Errorf("vsqlx: unsupported database driver %s", c.DatabaseDriver)
	}
	return
}

func setDefault(c *contract.Config) {
	// If max idle connection is unset, set to 10
	if c.DatabaseMaxIdleConn == "" {
		c.DatabaseMaxIdleConn = "10"
	}
	// If max open connection is unset, set to 10
	if c.DatabaseMaxOpenConn == "" {
		c.DatabaseMaxOpenConn = "10"
	}
	// If max idle connection is unset, set to 1 second
	if c.DatabaseMaxConnLifetime == "" {
		c.DatabaseMaxConnLifetime = "1"
	}
}

// PrepareFmtRebind prepare sql statements from string format and rebind variable or exit cmd if fails or error
func PrepareStmtRebind(db *sqlx.DB, queryFmt string, args ...interface{}) *sqlx.Stmt {
	query := fmt.Sprintf(queryFmt, args...)
	query = db.Rebind(query)
	return PrepareStmt(db, query)
}

// PrepareNamed prepare sql statements with named bindvars or exit cmd if fails or error
func PrepareStmt(db *sqlx.DB, query string) *sqlx.Stmt {
	stmt, err := db.Preparex(query)
	if err != nil {
		panic(fmt.Errorf("vsqlx: error while preparing statment [%s] (%s)", query, err))
	}
	return stmt
}

// PrepareNamed prepare sql statements with named bindvars or exit cmd if fails or error
func PrepareNamed(db *sqlx.DB, query string) *sqlx.NamedStmt {
	stmt, err := db.PrepareNamed(query)
	if err != nil {
		panic(fmt.Errorf("vsqlx: error while preparing statment [%s] (%s)", query, err))
	}
	return stmt
}

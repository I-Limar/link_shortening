package db

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"text/template"
)

var Tables = struct {
	Links string
}{
	Links: "links",
}

// DBService represents a connection to database
type DBService struct {
	config     DBConfig
	connection *sql.DB
}

// DBConfig represent a config required to connect to database
type DBConfig struct {
	Host     string `required:"true"`
	Port     string `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
	Dbname   string `required:"true"`
}

func BuildConnectionString(cfg DBConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.Dbname, cfg.Port)
}

// NewDBService Creates new connection to database
func NewDBService(cfg DBConfig) (DBService, error) {
	var db DBService

	db.config = cfg
	dbInfo := BuildConnectionString(cfg)
	connection, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return db, fmt.Errorf("[ ERROR ] Couldn't open connection to Database: %s", err)
	}

	db.connection = connection

	return db, nil
}

type Executor struct {
	db       *DBService
	template *template.Template
}

func (e *Executor) GetRow(resp interface{}, name string, source interface{}, Args ...interface{}) error {
	b := bytes.NewBuffer(nil)

	if err := e.template.ExecuteTemplate(b, name, source); err != nil {
		return fmt.Errorf("sql template: %w", err)
	}
	row := e.db.connection.QueryRow(b.String(), Args...)

	err := row.Scan(resp)
	if err != nil {
		return fmt.Errorf("error: sql: %w", err)
	}

	return nil
}

func (e *Executor) SetRow(name string, source interface{}, Args ...interface{}) error {
	b := bytes.NewBuffer(nil)

	if err := e.template.ExecuteTemplate(b, name, source); err != nil {
		return fmt.Errorf("sql template: %w", err)
	}
	_, err := e.db.connection.Exec(b.String(), Args...)

	if err != nil {
		return fmt.Errorf("error: sql: %w", err)
	}

	return nil
}

// NewExecutor return new instance of Executor
func NewExecutor(db *DBService, queries string) *Executor {
	var err error

	t := template.New("default").Option("missingkey=error")

	if t, err = t.Parse(queries); err != nil {
		panic(err)
	}

	return &Executor{
		db:       db,
		template: t,
	}
}

package db

import (
	"fmt"
	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
	"github.com/vsafonkin/involta/internal/config"
)

type DB struct {
	Conn *reindexer.Reindexer
}

func NewConnect() (DB, error) {
	dns := fmt.Sprintf("cproto://%s:%s/%s", config.RServerHost(), config.RServerPort(), config.DBName())
	db := reindexer.NewReindex(dns, reindexer.WithCreateDBIfMissing())
	if err := db.Ping(); err != nil {
		return DB{}, err
	}
	return DB{Conn: db}, nil
}

func (db DB) Upsert(namespace string, s interface{}) error {
	return db.Conn.Upsert(namespace, s)
}

func (db DB) GetById(id int, namespace string) (interface{}, bool) {
	return db.Conn.Query(namespace).Where("id", reindexer.EQ, id).Get()
}

func (db DB) List(namespace string) ([]interface{}, error) {
	return db.Conn.Query(namespace).Exec().FetchAll()
}

func OpenNamespace(conn *reindexer.Reindexer, namespace string, s interface{}) error {
	return conn.OpenNamespace(namespace, reindexer.DefaultNamespaceOptions(), s)
}

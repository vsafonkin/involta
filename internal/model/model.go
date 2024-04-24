package model

import (
	"github.com/vsafonkin/involta/internal/cache"
	"github.com/vsafonkin/involta/internal/config"
	"github.com/vsafonkin/involta/internal/db"
)

var conn db.DB

type InnerDoc struct {
	Content string `reindex:"content"`
}

type Doc struct {
	Id      int        `reindex:"id,,pk"`
	Sort    int        `reindex:"sort"`
	Content []InnerDoc `reindex:"content"`
}

func NewDocModel(d db.DB) error {
	conn = d
	cache.NewCache()
	return db.OpenNamespace(d.Conn, config.Namespace(), Doc{})
}

func List() ([]interface{}, error) {
	return conn.List(config.Namespace())
}

func GetById(id int) (interface{}, bool) {
	doc, ok := cache.GetFromCache(id)
	if !ok {
		out, ok := conn.GetById(id, config.Namespace())
		if ok {
			cache.AddToCache(id, out)
		}
		return out, ok
	}
	return doc, ok
}

func Upsert(doc Doc) error {
	if err := conn.Upsert(config.Namespace(), doc); err != nil {
		return err
	}
	cache.InvalidateCache(doc.Id)
	return nil
}

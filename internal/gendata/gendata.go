package gendata

import (
	"math/rand"

	"github.com/vsafonkin/involta/internal/model"
)

func GenRandomData(n int) error {
	var doc model.Doc
	doc.Content = []model.InnerDoc{{Content: "Alice"}, {Content: "Bob"}}
	for i := range n {
		doc.Id = i
		doc.Sort = rand.Intn(10000)
		if err := model.Upsert(doc); err != nil {
			return err
		}
	}
	return nil
}

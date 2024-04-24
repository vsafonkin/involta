package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vsafonkin/involta/internal/config"
	"github.com/vsafonkin/involta/internal/model"
)

type response struct {
	Message string
}

func RunAPIServer() error {
	router := gin.Default()

	dbname := config.DBName()
	namespace := config.Namespace()
	url := fmt.Sprintf("/%s/%s", dbname, namespace)
	router.GET(url, list)
	router.POST(url, update)

	apiServerUrl := fmt.Sprintf("%s:%s", config.APIServerHost(), config.APIServerPort())
	if err := router.Run(apiServerUrl); err != nil {
		return err
	}
	return nil
}

func list(c *gin.Context) {
	values := c.Request.URL.Query()
	if v, ok := values["id"]; ok {
		id, err := strconv.Atoi(v[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, response{Message: fmt.Sprintf("Cannot convert id '%s' to integer", v[0])})
			return
		}
		doc, ok := model.GetById(id)
		if !ok {
			c.JSON(http.StatusBadRequest, response{Message: fmt.Sprintf("Document with id '%d' not found", id)})
			return
		}
		c.JSON(http.StatusOK, doc)
		return
	}

	docs, err := model.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, docs)
}

func update(c *gin.Context) {
	var doc model.Doc
	if err := c.BindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	if err := model.Upsert(doc); err != nil {
		c.JSON(http.StatusBadRequest, response{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response{Message: "Successful"})
}

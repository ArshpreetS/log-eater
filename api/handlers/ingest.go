package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ArshpreetS/log-ingester/models"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gin-gonic/gin"
)

func Ingest(g *gin.Context, es *elasticsearch.Client) {
	body := new(models.Log)

	if err := g.ShouldBindJSON(&body); err != nil {
		g.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
			Status  int    `json:"status"`
		}{
			Status:  0,
			Message: "Bad request: Doesn't contain all field values",
		})
		return
	}

	// Addition logic

	// Convert the document to JSON bytes
	docBytes, err := json.Marshal(body)
	println(string(docBytes))
	if err != nil {
		fmt.Println(err)
		g.JSON(http.StatusInternalServerError, struct {
			Message string `json:"message"`
			Status  int    `json:"status"`
		}{
			Status:  0,
			Message: "Server Error: Error while marshalling",
		})
		return
	}

	req := esapi.IndexRequest{
		Index:      "search-log_data",
		DocumentID: time.Now().Format("2006-01-02 15:04:05"),
		Body:       bytes.NewReader(docBytes),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), es)

	if err != nil || res.IsError() {
		fmt.Println(err)
		fmt.Println(res.String())
		g.JSON(http.StatusInternalServerError, struct {
			Message string `json:"message"`
			Status  int    `json:"status"`
		}{
			Status:  0,
			Message: "Server Error: Error while Adding log",
		})
		return
	}
	defer res.Body.Close()

	g.JSON(http.StatusOK, struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}{
		Status:  1,
		Message: "Success: Log added",
	})
}

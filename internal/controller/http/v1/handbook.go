package v1

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"handbooks_backend/internal/policy/dto"
	"log"
	"net/http"
)

func (h *Handler) createHandbook(ctx *gin.Context) {
	var input dto.CreateHandbookInput
	// generate and validate structure
	if err := ctx.BindJSON(&input); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	handbook, err := h.policy.CreateHandbook(ctx, input)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, handbook)
}

func (h *Handler) getHandbooks(ctx *gin.Context) {
	projectCode := ctx.Query("project_code")
	search := ctx.Query("search")
	handbooks, err := h.policy.GetHandbooks(ctx, projectCode, search)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, handbooks)
}

func (h *Handler) updateHandbook(ctx *gin.Context) {
	var input dto.UpdateHandbookInput
	// generate and validate structure
	if err := ctx.BindJSON(&input); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	handbook, err := h.policy.UpdateHandbook(ctx, ctx.Param("id"), input)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	filter := make(map[string]interface{})
	result, err := h.policy.GetRowsHandbook(ctx, handbook, filter)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (h *Handler) getHandbook(ctx *gin.Context) {
	filter := make(map[string]interface{})
	filter["search"] = ctx.Query("search")
	rows, err := h.policy.GetHandbook(ctx, ctx.Param("id"), filter)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, rows)
}

func (h *Handler) updateRowHandbook(ctx *gin.Context) {
	var input []map[string]interface{}
	if err := ctx.BindJSON(&input); err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	rows, err := h.policy.UpdateRowsHandbook(ctx, ctx.Param("id"), input)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, rows)
}

func (h *Handler) testELK(ctx *gin.Context) {
	type LogData struct {
		Message string `json:"message"`
	}

	elasticsearchURL := "https://10.154.0.148:9200" // Замените на URL вашего Elasticsearch
	indexName := "handbook_er"                      // Замените на имя вашего индекса
	logData := LogData{
		Message: "Пример сообщения для отправки в ELK систему",
	}
	jsonData, err := json.Marshal(logData)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", elasticsearchURL+"/"+indexName, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("Ошибка при отправке данных. Код ответа: %d", resp.StatusCode)
	} else {
		log.Println("Данные успешно отправлены в ELK систему")
	}

	ctx.JSON(http.StatusOK, "ok")
}

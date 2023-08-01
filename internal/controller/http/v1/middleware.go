package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	authorizationHeaders = "Authorization"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeaders)
	if header == "" {
		errorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		errorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	client := &http.Client{}
	req, _ := http.NewRequest(
		"GET", "https://profile.mrm-etagi.com/api/permissions", nil,
	)
	// добавляем заголовки
	req.Header.Add("Authorization", "Bearer "+headerParts[1])
	resp, err := client.Do(req)
	if err != nil {
		errorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	if resp.StatusCode != 200 {
		errorResponse(ctx, http.StatusUnauthorized, resp.Status)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var b []map[string]interface{}
	err = json.Unmarshal(body, &b)
	if err != nil {
		errorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	for _, val := range b {
		if val["code"] == "handbooks" {
			ctx.Set("projects", val["modules"])
		}
	}

}

package dto

import (
	"github.com/hetiansu5/urlquery"
	"github.com/imamponco/v-gin-boilerplate/src/pkg/vhttp"
	"github.com/imamponco/v-gin-boilerplate/src/svc/constant"
	"net/http"
	"strings"
	"time"
)

type Subject struct {
	ID          string
	RefID       int64
	Role        int64
	FullName    string
	SubjectType constant.SubjectType
	SessionXID  string
	Version     int64
	Metadata    map[string]string
}

type BaseField struct {
	CreatedAt  int64     `json:"createdAt" time_format:"unix" example:"1662803912"`
	UpdatedAt  int64     `json:"updatedAt" time_format:"unix" example:"1662803912"`
	ModifiedBy *Modifier `json:"modifiedBy"`
	Version    int64     `json:"version" example:"1"`
}

type Modifier struct {
	ID       string                `json:"-"`
	Role     constant.ModifierRole `json:"role" example:"USER"`
	FullName string                `json:"fullName" example:"USER"`
}

type GetHealth_Result struct {
	AppVersion     string    `json:"appVersion" example:"v0.1.0"`
	BuildSignature string    `json:"buildSignature" example:"2b38f457-577f-423b-a7c0-16c50a86398c"`
	Uptime         string    `json:"uptime" example:"1m39.5398474s"`
	ServerTime     time.Time `json:"serverTime" example:"2023-03-14T22:59:10.155009095+07:00"`
}

type GetDetail_Payload struct {
	XID     string   `json:"XID"`
	Subject *Subject `json:"subject"`
}

type GetDetailById_Payload struct {
	Id      int64    `json:"XID"`
	Subject *Subject `json:"subject"`
}

type List_Payload struct {
	Limit   int64             `json:"limit" query:"limit"`
	Skip    int64             `json:"skip" query:"skip"`
	SortBy  string            `json:"sortBy" query:"sortBy"`
	Filters map[string]string `json:"-" query:"filters"`
}

func GetListPayload(rq *http.Request, payload *List_Payload) error {
	// Parse query
	queryParam := ValidateQueryParams(rq.URL.RawQuery)
	err := urlquery.Unmarshal([]byte(queryParam), payload)
	if err != nil {
		return vhttp.ListQueryParamError
	}

	if payload.Filters == nil {
		payload.Filters = map[string]string{}
	}

	// Normalize Limit
	if payload.Limit <= 0 {
		payload.Limit = 10
	}

	// Normalize Skip
	if payload.Skip < 0 {
		payload.Skip = 0
	}

	return nil
}

func ValidateQueryParams(query string) string {
	if query == "" {
		return ""
	}

	getVariable := strings.Split(query, "&")

	for _, row := range getVariable {
		value := strings.Split(row, "=")
		if value[1] == "" {
			query = strings.Replace(query, value[0], "", -1)
		}
	}

	return query
}

type ListMetadata struct {
	Count  int64  `json:"count" example:"1"`
	Limit  int64  `json:"limit" example:"10"`
	Skip   int64  `json:"skip" example:"0"`
	SortBy string `json:"sortBy" example:"createdAt DESC"`
}

func ToListMetadata(p *List_Payload, count int64) *ListMetadata {
	return &ListMetadata{
		Count:  count,
		Limit:  p.Limit,
		Skip:   p.Skip,
		SortBy: p.SortBy,
	}
}

type Version_Payload struct {
	Version int64 `json:"version" example:"1"`
}

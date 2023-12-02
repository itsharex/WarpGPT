package common

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type RequestUrl interface {
	Generate(path string, rawquery string) string
}
type OfficialApiRequestUrl struct {
}
type UnOfficialApiRequestUrl struct {
}
type ReverseApiRequestUrl struct {
}
type ReverseBackendRequestUrl struct {
}
type ReversePublicApiRequestUrl struct {
}

func (u OfficialApiRequestUrl) Generate(path string, rawquery string) string {
	if rawquery == "" {
		return "https://api.openai.com/v1" + path
	}
	return "https://api.openai.com/v1" + path + "?" + rawquery
}
func (u UnOfficialApiRequestUrl) Generate(path string, rawquery string) string {
	if rawquery == "" {
		return "https://" + Env.OpenAI_HOST + "/backend-api/" + path
	}
	return "https://" + Env.OpenAI_HOST + "/backend-api/" + path + "?" + rawquery
}
func (u ReverseApiRequestUrl) Generate(path string, rawquery string) string {
	if rawquery == "" {
		return "https://" + Env.OpenAI_HOST + "/api/" + path
	}
	return "https://" + Env.OpenAI_HOST + "/api/" + path + "?" + rawquery
}
func (u ReverseBackendRequestUrl) Generate(path string, rawquery string) string {
	if rawquery == "" {
		return "https://" + Env.OpenAI_HOST + "/backend-api/" + path
	}
	return "https://" + Env.OpenAI_HOST + "/backend-api/" + path + "?" + rawquery
}
func (u ReversePublicApiRequestUrl) Generate(path string, rawquery string) string {
	if rawquery == "" {
		return "https://" + Env.OpenAI_HOST + "/public-api/" + path
	}
	return "https://" + Env.OpenAI_HOST + "/public-api/" + path + "?" + rawquery
}

func CheckRequest(c *gin.Context) RequestUrl {
	path := c.Request.URL.Path
	if strings.HasPrefix(path, "/backend-api") {
		return ReverseBackendRequestUrl{}
	}
	if strings.HasPrefix(path, "/api") {
		return ReverseApiRequestUrl{}
	}
	if strings.HasPrefix(path, "/public-api") {
		return ReversePublicApiRequestUrl{}
	}
	if strings.HasPrefix(path, "/v1") {
		return OfficialApiRequestUrl{}
	}
	if strings.HasPrefix(path, "/r") {
		return UnOfficialApiRequestUrl{}
	}
	return nil
}

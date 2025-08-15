package response

import (
	"encoding/json"
	"net/http"
	"os"
)

type ApiResponseStatus struct {
	Code int    `json:"code"`
	Type string `json:"type"`
}

type ApiResponse struct {
	ProcessId string            `json:"process_id"`
	Path      string            `json:"path"`
	Status    ApiResponseStatus `json:"status"`
	Request   interface{}       `json:"request"`
	Errors    []error           `json:"errors"`
	Data      interface{}       `json:"data"`
}

func (response *ApiResponse) ToJSON() string {
	jsonout, err := json.Marshal(response)
	if err != nil {
		return ""
	}

	return string(jsonout)
}

func NewApiResponse(path string) *ApiResponse {
	respStatus := ApiResponseStatus{
		Code: http.StatusOK,
		Type: http.StatusText(http.StatusOK),
	}

	respError := []error{}
	respData := make([]interface{}, 0)

	return &ApiResponse{
		ProcessId: os.Getenv("PROCESS_ID"),
		Status:    respStatus,
		Errors:    respError,
		Data:      respData,
		Path:      path,
	}
}

package api

import (
	"encoding/json"
	"net/http"
)

type ApiResult struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
	Error   string      `json:"error"`
}

func SetResult(statusCode int, result interface{}, errorInfo error, w http.ResponseWriter) {
	apiResult := &ApiResult{}

	if statusCode == http.StatusOK {
		apiResult.Success = true
	} else {
		apiResult.Success = false
	}

	apiResult.Result = result
	if errorInfo != nil {
		apiResult.Error = errorInfo.Error()
	}

	jsonResponse, err := json.Marshal(apiResult)

	if err != nil {
		apiResult.Success = false
		apiResult.Error = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("custom", "omrani")

	w.WriteHeader(statusCode)
	_, err = w.Write(jsonResponse)
	if err != nil {
		panic(err)
	}
	// fmt.Fprintf(w, "%s", string(jsonResponse))
}

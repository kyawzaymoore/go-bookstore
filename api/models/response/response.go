package response

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var ErrorCodeList map[string]struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseModelList[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []T    `json:"data"`
}

type ResponseModel[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func init() {
	errorPath := os.Getenv("ERROR_PATH")  // Get the errorcode path from the environment variable

    if errorPath == "" {
        errorPath = "helper/errorcode.json"
    }

	errorCodesJSON, err := ioutil.ReadFile(errorPath)
	if err != nil {
		fmt.Println("Error reading errorcode.json:", err)
		return
	}

	if err := json.Unmarshal(errorCodesJSON, &ErrorCodeList); err != nil {
		fmt.Println("Error decoding errorcode.json:", err)
		return
	}
}

func GetResponseModel[T any](errCode string, data T) *ResponseModel[T] {
	err, ok := ErrorCodeList[errCode]
	if ok {
		return &ResponseModel[T]{
			Code:    err.Code,
			Message: err.Message,
			Data:    data,
		}
	}
	return &ResponseModel[T]{
		Code:    -1,
		Message: "Unknown Error",
		Data:    data,
	}
}
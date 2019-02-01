package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const ApiUrl = "https://aws.random.cat/meow"

type ApiResponse struct {
	Url string `json:"file"`
	Error error
}

func getCatPic() ApiResponse {
	res, err := http.Get(ApiUrl)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return ApiResponse{
			Error: err,
		}
	}

	contents, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return ApiResponse{
			Error: err,
		}
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(contents, &apiResponse)

	if err != nil {
		return ApiResponse{
			Error: err,
		}
	}

	apiResponse.Url = strings.Replace(apiResponse.Url, "\\", "", 1)
	return apiResponse
}

package httpapi

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ridhozhr10/ottojunior/internal/model"
	"github.com/ridhozhr10/ottojunior/internal/repository"
)

type productHttpapiRepository struct {
	baseURL    string
	httpClient *http.Client
}

// NewProductHttpapiRepository create implementation for repository.Product
func NewProductHttpapiRepository(baseURL string) repository.Product {
	return &productHttpapiRepository{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (r *productHttpapiRepository) Get() ([]model.Product, error) {
	result := []model.Product{}
	request, err := http.NewRequest("GET", r.baseURL+"/interview/biller/v1/list", nil)
	if err != nil {
		return result, err
	}

	response, err := r.httpClient.Do(request)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	type rawResponseT struct {
		Code    int             `json:"code"`
		Status  string          `json:"status"`
		Message string          `json:"message"`
		Data    []model.Product `json:"data"`
	}
	var responseRaw rawResponseT
	err = json.NewDecoder(response.Body).Decode(&responseRaw)
	if err != nil {
		return result, err
	}
	result = responseRaw.Data
	return result, nil
}

func (r *productHttpapiRepository) GetByID(productID int) (model.Product, error) {
	result := model.Product{}
	pidstr := strconv.Itoa(productID)
	request, err := http.NewRequest("GET", r.baseURL+"/interview/biller/v1/detail?billerId="+pidstr, nil)
	if err != nil {
		return result, err
	}

	response, err := r.httpClient.Do(request)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()
	type rawResponseT struct {
		Code    int           `json:"code"`
		Status  string        `json:"status"`
		Message string        `json:"message"`
		Data    model.Product `json:"data"`
	}
	var responseRaw rawResponseT
	err = json.NewDecoder(response.Body).Decode(&responseRaw)
	if err != nil {
		return result, err
	}
	result = responseRaw.Data
	return result, nil
}

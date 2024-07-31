package course

import (
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/zchelalo/go_microservices_domain/domain"
)

type (
	DataResponse struct {
		Message string      `json:"message"`
		Status  int         `json:"status"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta"`
	}

	Transport interface {
		Get(id string) (*domain.Course, error)
	}

	clientHTTP struct {
		client resty.Client
	}
)

func NewHTTPClient(baseURL, token string) Transport {
	var authToken string

	if token != "" {
		authToken = fmt.Sprintf("Bearer %s", token)
	}

	return &clientHTTP{
		client: *resty.New().SetBaseURL(baseURL).SetAuthToken(authToken).SetTimeout(5 * time.Second),
	}
}

func (c *clientHTTP) Get(id string) (*domain.Course, error) {
	dataResponse := DataResponse{
		Data: &domain.Course{},
	}

	u := url.URL{}
	u.Path = fmt.Sprintf("/courses/%s", id)
	response, err := c.client.R().SetResult(&dataResponse).Get(u.String())
	if err != nil {
		return nil, err
	}

	if !response.IsSuccess() {
		if response.StatusCode() == 404 {
			return nil, ErrNotFound{response.String()}
		}

		return nil, fmt.Errorf("Error: %s", response.String())
	}

	return dataResponse.Data.(*domain.Course), nil
}

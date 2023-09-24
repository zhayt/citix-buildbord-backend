package news

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"innovatex-app/internal/config"
	"innovatex-app/internal/models"
	"net/http"
	"time"
)

type viaHTTP struct {
	url     string
	apiKey  string
	timeout time.Duration
}

func newViaHTTP(source *config.Source) *viaHTTP {
	return &viaHTTP{
		url:     source.NewsAPI,
		apiKey:  source.NewsAPIKey,
		timeout: source.Timeout,
	}
}

func (r *viaHTTP) GetAll(ctx context.Context) (*models.News, error) {
	client := http.Client{
		Timeout: r.timeout,
	}

	request, err := http.NewRequest(http.MethodGet, r.url, nil)
	if err != nil {
		zap.S().Error("Making request error: %s", err.Error())
		return nil, err
	}

	query := request.URL.Query()
	query.Add("apiKey", r.apiKey)
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		zap.S().Errorf("Doing request error: %s", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	var results models.News
	if err = json.NewDecoder(response.Body).Decode(&results); err != nil {
		zap.S().Errorf("Parsing response body error: %s", err.Error())
		return nil, err
	}

	return &results, nil
}

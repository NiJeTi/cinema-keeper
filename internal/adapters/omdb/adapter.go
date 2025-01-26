package omdb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nijeti/cinema-keeper/internal/models"
	"github.com/nijeti/cinema-keeper/internal/pkg/ptr"
)

type Config struct {
	Key string `conf:"key"`
}

type Adapter struct {
	config Config
	client *http.Client
}

func New(config Config) *Adapter {
	return &Adapter{
		config: config,
		client: &http.Client{},
	}
}

func (a *Adapter) MoviesByTitle(
	ctx context.Context, title string,
) ([]models.MovieShort, error) {
	var dto movieSearch
	if err := a.makeRequest(ctx, a.baseURL()+"&s="+title, &dto); err != nil {
		return nil, err
	}

	return dto.toModel(), nil
}

func (a *Adapter) MovieByID(
	ctx context.Context, id string,
) (*models.MovieShort, error) {
	var dto movie
	if err := a.makeRequest(ctx, a.baseURL()+"&i="+id, &dto); err != nil {
		return nil, err
	}

	return ptr.To(dto.toModel()), nil
}

func (a *Adapter) makeRequest(
	ctx context.Context, url string, result any,
) error {
	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet, url, http.NoBody,
	)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to get response: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return nil
}

package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Torebekov/shop-back/internal/models"
	"io"
	"net/http"
)

type Interface interface {
	GetAuth(ctx context.Context, authHeader string) (models.Auth, error)
}

func New(addr string) Interface {
	return &service{
		addr: addr,
	}
}

type service struct {
	addr string
}

func (s service) GetAuth(_ context.Context, authHeader string) (models.Auth, error) {
	req, err := http.NewRequest(http.MethodPost, s.addr+"/check", nil)
	if err != nil {
		return models.Auth{}, fmt.Errorf("creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return models.Auth{}, fmt.Errorf("client.Do: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Auth{}, fmt.Errorf("io.ReadAll: %v", err)
	}

	var result models.Auth
	if err := json.Unmarshal(respBody, &result); err != nil {
		return models.Auth{}, fmt.Errorf("faild to parse: %v", err)
	}

	return result, nil
}

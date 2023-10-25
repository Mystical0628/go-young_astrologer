package nasa

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"go-young_astrologer/package-core"
	"go-young_astrologer/service-nasa/domain/apod/entity"
	"io"
	"net/http"
	"time"
)

type ApodGateway struct {
	config *package_core.Config
}

func NewApodGateway(config *package_core.Config) *ApodGateway {
	return &ApodGateway{
		config: config,
	}
}

type ApodFetchImageDTO struct {
	Date      *time.Time `url:"date,omitempty" layout:"2006-01-02"`
	StartDate *time.Time `url:"start_date,omitempty" layout:"2006-01-02"`
	EndDate   *time.Time `url:"end_date,omitempty" layout:"2006-01-02"`
	Count     *int       `url:"count,omitempty"`
	Thumbs    *bool      `url:"thumbs,omitempty"`
}

func (g *ApodGateway) FetchImage(dto ApodFetchImageDTO) (*entity.Apod, error) {
	vals, err := query.Values(dto)
	if err != nil {
		return nil, err
	}
	vals.Add("api_key", g.config.NasaApiKey)

	resp, err := http.Get("https://api.nasa.gov/planetary/apod?" + vals.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	apod := new(entity.Apod)
	err = json.Unmarshal(jsonBytes, apod)
	if err != nil {
		return nil, err
	}

	return apod, nil
}

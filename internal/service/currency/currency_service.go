package currency

import (
	"encoding/json"
	"errors"
	"fintrack/internal/domain/model"
	"fintrack/internal/repo/currency"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// APIConfig holds the configuration for the external API.
type APIConfig struct {
	BaseURL string
	Version string
	APIKey  string
}

// API is the implementation of CurrencyConverterAPI.
type API struct {
	config APIConfig
}

// NewAPI creates and returns an API instance.
func NewAPI(config APIConfig) *API {
	return &API{config: config}
}

type CurrencyService interface {
	Convert(req model.ConvertRequest) (*model.Convert, error)
	ConvertCompact(req model.ConvertRequest) (model.ConvertCompact, error)
	ConvertHistorical(req model.ConvertHistoricalRequest) (*model.ConvertHistorical, error)
	ConvertHistoricalCompact(req model.ConvertHistoricalRequest) (model.ConvertHistoricalCompact, error)
	Currencies() (*model.Currency, error)
	Countries() (*model.Country, error)
	Usage() (*model.Usage, error)
}

type currencyServiceImpl struct {
	repo currency.CurrencyRepository
	api  *API
}

func NewCurrencyService(repo currency.CurrencyRepository, api *API) CurrencyService {
	return &currencyServiceImpl{repo: repo, api: api}
}

func call[T model.Response](a *API, shouldPrefixAPIPath bool, path string, handler func(q url.Values) error) (result *T, err error) {
	u, err := url.Parse(a.config.BaseURL)
	if err != nil {
		return nil, err
	}

	if shouldPrefixAPIPath {
		u = u.JoinPath("api").JoinPath(a.config.Version)
	}

	u = u.JoinPath(path)

	query := u.Query()
	query.Add("apiKey", a.config.APIKey)

	err = handler(query)
	if err != nil {
		return nil, err
	}

	u.RawQuery = query.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, parseError(resp)
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// parseError uses json.Unmarshal to returns Error whenever it is possible.
func parseError(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	e := model.Error{}
	err = json.Unmarshal(body, &e)
	if err != nil {
		return errors.New(string(body))
	}

	return errors.New(e.Error)
}

func (s *currencyServiceImpl) Convert(req model.ConvertRequest) (*model.Convert, error) {
	return call[model.Convert](s.api, true, "convert", func(q url.Values) error {
		if len(req.Q) == 0 {
			return errors.New("`Q` require at least one currency conversion")
		}
		q.Add("q", strings.Join(req.Q, ","))
		return nil
	})
}

func (s *currencyServiceImpl) ConvertCompact(req model.ConvertRequest) (model.ConvertCompact, error) {
	r, err := call[model.ConvertCompact](s.api, true, "convert", func(q url.Values) error {
		if len(req.Q) == 0 {
			return errors.New("`Q` require at least one currency conversion")
		}
		q.Add("compact", "ultra")
		q.Add("q", strings.Join(req.Q, ","))
		return nil
	})
	if err != nil {
		return model.ConvertCompact{}, err
	}

	return *r, nil
}

func (s *currencyServiceImpl) ConvertHistorical(req model.ConvertHistoricalRequest) (*model.ConvertHistorical, error) {
	return call[model.ConvertHistorical](s.api, true, "convert", func(q url.Values) error {
		if len(req.Q) == 0 {
			return errors.New("`Q` require at least one currency conversion")
		}
		if req.Date.IsZero() {
			return errors.New("`Date` is required")
		}
		q.Add("q", strings.Join(req.Q, ","))
		q.Add("date", req.Date.Format("2006-01-02"))
		if !req.EndDate.IsZero() {
			q.Add("endDate", req.EndDate.Format("2006-01-02"))
		}
		return nil
	})
}

func (s *currencyServiceImpl) ConvertHistoricalCompact(req model.ConvertHistoricalRequest) (model.ConvertHistoricalCompact, error) {
	r, err := call[model.ConvertHistoricalCompact](s.api, true, "convert", func(q url.Values) error {
		if len(req.Q) == 0 {
			return errors.New("`Q` require at least one currency conversion")
		}
		if req.Date.IsZero() {
			return errors.New("`Date` is required")
		}
		q.Add("compact", "ultra")
		q.Add("q", strings.Join(req.Q, ","))
		q.Add("date", req.Date.Format("2006-01-02"))
		if !req.EndDate.IsZero() {
			q.Add("endDate", req.EndDate.Format("2006-01-02"))
		}
		return nil
	})
	if err != nil {
		return model.ConvertHistoricalCompact{}, err
	}

	return *r, nil
}

func (s *currencyServiceImpl) Currencies() (*model.Currency, error) {
	return call[model.Currency](s.api, true, "currencies", func(q url.Values) error { return nil })
}

func (s *currencyServiceImpl) Countries() (*model.Country, error) {
	return call[model.Country](s.api, true, "countries", func(q url.Values) error { return nil })
}

func (s *currencyServiceImpl) Usage() (*model.Usage, error) {
	return call[model.Usage](s.api, false, "others/usage", func(q url.Values) error { return nil })
}

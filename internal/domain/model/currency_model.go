package model

import (
	"time"
)

// Error represents the API error response.
type Error struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

// ConvertRequest contains request fields for the Convert API.
type ConvertRequest struct {
	Q []string `json:"q"` // e.g., MYR_USD
}

// ConvertHistoricalRequest contains request fields for ConvertHistorical API.
type ConvertHistoricalRequest struct {
	Q       []string  `json:"q"`
	Date    time.Time `json:"date"`
	EndDate time.Time `json:"endDate,omitempty"`
}

// Convert is the result of the Convert API.
type Convert struct {
	Query struct {
		Count int `json:"count"`
	} `json:"query"`
	Results map[string]struct {
		ID  string  `json:"id"`
		Val float32 `json:"val"`
		To  string  `json:"to"`
		Fr  string  `json:"fr"`
	} `json:"results"`
}

// ConvertCompact is the compact result of the Convert API.
type ConvertCompact map[string]float32

// ConvertHistorical is the result of ConvertHistorical API.
type ConvertHistorical struct {
	Query struct {
		Count int `json:"count"`
	} `json:"query"`
	Date    string `json:"date"`
	EndDate string `json:"endDate,omitempty"`
	Results map[string]struct {
		ID  string             `json:"id"`
		To  string             `json:"to"`
		Fr  string             `json:"fr"`
		Val map[string]float32 `json:"val"`
	} `json:"results"`
}

// ConvertHistoricalCompact is the compact result of ConvertHistorical API.
type ConvertHistoricalCompact map[string]map[string]float32

// Currency is the result of the Currencies API.
type Currency struct {
	Results map[string]struct {
		ID             string `json:"id"`
		CurrencyName   string `json:"currencyName"`
		CurrencySymbol string `json:"currencySymbol"`
	} `json:"results"`
}

// Country is the result of the Countries API.
type Country struct {
	Results map[string]struct {
		ID             string `json:"id"`
		Alpha3         string `json:"alpha3"`
		CurrencyID     string `json:"currencyId"`
		CurrencyName   string `json:"currencyName"`
		CurrencySymbol string `json:"currencySymbol"`
		Name           string `json:"name"`
	} `json:"results"`
}

// Usage is the result of the Usage API.
type Usage struct {
	Timestamp time.Time `json:"timestamp"`
	Usage     int       `json:"usage"`
}

// Response is the generic response interface used by the call function.
type Response interface {
	Convert | ConvertCompact | ConvertHistorical | ConvertHistoricalCompact | Currency | Country | Usage
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"os"

	"github.com/mavryk-network/mvgo/mavryk"
)

type cmcResponse struct {
	Data map[string]struct {
		Slug  string `json:"slug"`
		Quote map[string]struct {
			Price float64 `json:"price"`
		} `json:"quote"`
	} `json:"data"`
}

func get_cmc_exchange_rate(token_slug string, apiKey string) (mavryk.Z, error) {
	if token_slug == "" || token_slug == "mavryk" {
		return mavryk.Zero, errors.New("Invalid token slug")
	}

	if apiKey == "" {
		return mavryk.Zero, errors.New("Invalid API key")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("slug", fmt.Sprintf("mavryk,%s", token_slug))
	q.Add("aux", "num_market_pairs")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}

	respBody, _ := io.ReadAll(resp.Body)

	var cmcResp cmcResponse
	if err := json.Unmarshal(respBody, &cmcResp); err != nil {
		return mavryk.Zero, err
	}

	if cmcResp.Data == nil {
		return mavryk.Zero, errors.New("Invalid response from CMC")
	}

	var mavrykPrice float64
	var tokenPrice float64
	for _, data := range cmcResp.Data {
		for _, quote := range data.Quote {
			switch data.Slug {
			case "mavryk":
				mavrykPrice = quote.Price
			case token_slug:
				tokenPrice = quote.Price
			}
		}
	}

	if mavrykPrice == 0 || tokenPrice == 0 {
		return mavryk.Zero, errors.New("Invalid price data")
	}

	exchangeRate := mavrykPrice / tokenPrice
	return mavryk.NewZ(int64(exchangeRate * float64(PRECISION))), nil
}

type CMCExchangeRateProvider struct {
	slug    string
	api_key string
	fee     mavryk.Z
	token   TokenConfiguration

	rate mavryk.Z
}

func NewCMCExchangeRateProvider(slug, api_key string, fee float64, token TokenConfiguration) *CMCExchangeRateProvider {
	return &CMCExchangeRateProvider{
		slug:    slug,
		api_key: api_key,
		fee:     mavryk.NewZ(int64(fee * float64(PRECISION))),
		token:   token,
	}
}

func (p *CMCExchangeRateProvider) RefreshExchangeRate() error {
	rate, err := get_cmc_exchange_rate(p.slug, p.api_key)
	if err != nil {
		return err
	}
	slog.Info("Exchange rate updated", "rate", rate)
	p.rate = rate
	return nil
}

func (p *CMCExchangeRateProvider) ExchangeMavToToken(mumav mavryk.Z) mavryk.Z {
	decimalsMultiplier := int64(math.Pow10(p.token.Decimals))

	token_amount := mumav.Mul(p.rate).Mul64(decimalsMultiplier).Div(mavryk.NewZ(PRECISION).Sub(p.fee)).Div64(MUMAV_FACTOR)
	slog.Debug("Exchanging amount", "amount", mumav, "token_amount", token_amount, "rate", p.rate, "fee", p.fee)
	return token_amount
}

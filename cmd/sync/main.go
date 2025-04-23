package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/juandavidaa/stocks-api/internal/dto"
)

type apiResp struct {
	Items    []dto.CreateStock `json:"items"`
	NextPage string            `json:"next_page"`
}

var (
	outputFile = "data.json"
	baseURL    string
	tokenType  string
	token      string
	client     = &http.Client{Timeout: 10 * time.Second}
	nextPage   = ""
	page       = 1
	RequestUrl = &url.URL{}
)

func main() {
	flag.Parse()
	if err := godotenv.Load(); err != nil {
		fmt.Println("❌ .env file not found (please a new one)")
		os.Exit(1)
	}

	token = os.Getenv("API_TOKEN")
	baseURL = os.Getenv("API_URL")
	tokenType = os.Getenv("API_TOKEN_TYPE")

	if token == "" {
		fmt.Println("❌ API_TOKEN not found (please add it to .env)")
		os.Exit(1)
	}

	Init()

	if err := sync(); err != nil {
		fmt.Println("❌ ", err)
		os.Exit(1)
	}
}

func sync() error {

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	RequestUrl, err = buildURL()
	if err != nil {
		return err
	}

	client = &http.Client{}

	for {
		setNextPage()
		items, cursor, err := fetchPage()
		nextPage = cursor
		if err != nil {
			return fmt.Errorf("page %d: %w", page, err)
		}

		for _, it := range items {
			if err := writeJSON(f, &it); err != nil {
				return err
			}
		}

		fmt.Printf("✔  page %d: %d items (cursor=%q)\n", page, len(items), cursor)

		if cursor == "" {
			break
		}
		page++
		nextPage = cursor
	}

	fmt.Printf("✅ syncronization complete \n")
	return nil
}

func buildURL() (*url.URL, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func setNextPage() {
	query := RequestUrl.Query()
	if nextPage != "" {
		query.Set("next_page", nextPage)
	}
	RequestUrl.RawQuery = query.Encode()
}

func fetchPage() ([]dto.CreateStock, string, error) {

	response, err := makeRequest()
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return nil, "", fmt.Errorf("status %d: %s", response.StatusCode, body)
	}

	var data apiResp
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, "", err
	}
	return data.Items, data.NextPage, nil
}

func makeRequest() (*http.Response, error) {
	req, err := http.NewRequest("GET", RequestUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", tokenType+" "+token)
	return client.Do(req)
}

func writeJSON(w *os.File, v *dto.CreateStock) error {
	cleanCurrency(v)
	err := GetCurrentPrice(v)

	if err != nil {
		fmt.Println("❌ ", err)
	}

	writer := bufio.NewWriter(w)
	defer writer.Flush()

	enc := json.NewEncoder(writer)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "")

	if err := enc.Encode(v); err != nil {
		return err
	}

	return nil
}

func cleanCurrency(v *dto.CreateStock) {
	v.TargetFrom = strings.ReplaceAll(v.TargetFrom, "$", "")
	v.TargetFrom = strings.ReplaceAll(v.TargetFrom, ",", "")
	v.TargetTo = strings.ReplaceAll(v.TargetTo, "$", "")
	v.TargetTo = strings.ReplaceAll(v.TargetTo, ",", "")
}

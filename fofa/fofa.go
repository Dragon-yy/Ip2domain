package fofa

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func FetchDomains(apiKey, ip string) ([][]string, error) {
	query := fmt.Sprintf(`ip="%s"`, ip)
	url := fmt.Sprintf("https://fofa.info/api/v1/search/all?key=%s&qbase64=%s", apiKey, base64.StdEncoding.EncodeToString([]byte(query)))

	resp, err := http.Get(url)
	time.Sleep(1 * time.Second)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data struct {
		Error   bool       `json:"error"`
		Results [][]string `json:"results"`
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if data.Error {
		return nil, fmt.Errorf("API returned an error")
	}

	return data.Results, nil
}

func SaveToCSV(data [][]string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}

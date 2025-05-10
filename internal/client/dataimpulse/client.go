package dataimpulse

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"stats/internal/model"
	"strconv"
)

type authTransport struct {
	apiKey string
	base   http.RoundTripper
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Basic "+t.apiKey)
	//fmt.Println(t.apiKey)
	return t.base.RoundTrip(req)
}

type DataImpulse struct {
	client *http.Client
}

func New(ApiKey string) *DataImpulse {
	c := &DataImpulse{}
	c.client = &http.Client{
		Transport: &authTransport{
			apiKey: ApiKey,
			base:   http.DefaultTransport,
		},
	}
	//c.Auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(c.login+":"+c.pass))
	return c

}

func (c *DataImpulse) Traffic() (TrafficModel, error) {
	const url = "https://gw.dataimpulse.com:777/api/stats"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return TrafficModel{}, fmt.Errorf("Clouvider Get Balance Error request: %w", err)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return TrafficModel{}, fmt.Errorf("Clouvider Get Balance Error request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return TrafficModel{}, fmt.Errorf(
			"unexpected status %d: %s",
			res.StatusCode, string(body),
		)

	}
	//resp, err := io.ReadAll(res.Body)
	//fmt.Println(string(resp))
	var balResp TrafficModel
	if err := json.NewDecoder(res.Body).Decode(&balResp); err != nil {
		return TrafficModel{}, fmt.Errorf("decoding JSON: %w", err)
	}
	return balResp, nil
}

func (c *DataImpulse) Stat() model.DataImpulseStats {
	cvBalResp, err := c.Traffic()
	if err != nil {
		fmt.Println(err)
		return model.DataImpulseStats{Error: err.Error()}
	} else {
		return model.DataImpulseStats{
			TrafficLeft: strconv.FormatFloat(
				float64(cvBalResp.TrafficLeft)/(1024*1024*1024),
				'f', 2, 64,
			),
		}
	}
}

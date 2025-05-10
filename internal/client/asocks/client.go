package asocks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"stats/internal/model"
	"strconv"
)

//func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
//	req.Header.Set("Authorization", "Basic "+t.apiKey)
//	//fmt.Println(t.apiKey)
//	return t.base.RoundTrip(req)
//}

type Asocks struct {
	client  *http.Client
	authUrl url.URL
}

func New(ApiKey string, proxyUrl string) *Asocks {
	c := &Asocks{}
	//c.apiKey = ApiKey
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		panic(err)
	}
	urlparse, err := url.Parse(fmt.Sprintf("https://api.asocks.com/?apiKey=%s", ApiKey))
	if err != nil {
		panic(err)
	}
	c.authUrl = *urlparse
	c.client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
	//c.url = fmt.Sprintf("https://api.sms-activate.ae/stubs/handler_api.php?api_key=%s", ApiKey)
	//c.Auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(c.login+":"+c.pass))
	return c

}

func (c *Asocks) Balance() (BalanceModel, error) {
	balurl := c.authUrl
	balurl.Path = "/v2/user/balance"
	requrl := balurl.String()
	req, err := http.NewRequest(http.MethodGet, requrl, nil)
	if err != nil {
		return BalanceModel{}, fmt.Errorf("Clouvider Get Balance Error request: %w", err)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return BalanceModel{}, fmt.Errorf("Clouvider Get Balance Error request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return BalanceModel{}, fmt.Errorf(
			"unexpected status %d: %s",
			res.StatusCode, string(body),
		)

	}
	//resp, err := io.ReadAll(res.Body)
	//fmt.Println(string(resp))
	var balResp BalanceModel
	if err := json.NewDecoder(res.Body).Decode(&balResp); err != nil {

		return BalanceModel{}, fmt.Errorf("decoding JSON: %w", err)
	}
	return balResp, nil
}

func (c *Asocks) Stat() model.AsocksStats {
	cvBalResp, err := c.Balance()
	if err != nil {
		fmt.Println(err)
		return model.AsocksStats{Error: err.Error()}
	} else {
		return model.AsocksStats{
			Balance: strconv.FormatFloat(cvBalResp.Balance, 'f', 2, 64),
		}
	}
}

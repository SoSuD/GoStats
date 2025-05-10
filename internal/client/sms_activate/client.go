package sms_activate

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"stats/internal/model"
	"strings"
)

type SmsActivate struct {
	client *http.Client
	//apiKey string
	url string
	//proxyUrl string
}

func New(ApiKey string, proxyUrl string) *SmsActivate {
	c := &SmsActivate{}
	//c.apiKey = ApiKey
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		panic(err)
	}
	c.client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
	c.url = fmt.Sprintf("https://api.sms-activate.ae/stubs/handler_api.php?api_key=%s", ApiKey)
	//c.Auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(c.login+":"+c.pass))
	return c

}

func (c *SmsActivate) Balance() (string, error) {
	req, err := http.NewRequest(http.MethodGet, c.url+"&action=getBalance", nil)
	if err != nil {
		return "", fmt.Errorf("SmsActivate Get Balance Error request: %w", err)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("SmsActivate Get Balance Error request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Printf("SmsActivate Get Balance Error status: %d\n", res.StatusCode)
		return "", fmt.Errorf("SmsActivate Get Balance Error status: %d", res.StatusCode)
	} else {
		body, _ := io.ReadAll(res.Body)
		parts := strings.SplitN(string(body), ":", 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("SmsActivate Get Balance Error response: %s", string(body))
		}
		balanceStr := strings.TrimSpace(parts[1])
		return balanceStr, nil

	}
	//resp, err := io.ReadAll(res.Body)
	//fmt.Println(string(resp))
	//return "", nil

}

func (c *SmsActivate) Stat() model.SmsActivateStats {
	cvBalResp, err := c.Balance()
	//fmt.Println(cvBalResp)
	if err != nil {
		fmt.Println(err)
		return model.SmsActivateStats{
			Error: err.Error(),
		}
	} else {
		return model.SmsActivateStats{
			Balance: cvBalResp,
		}
	}
}

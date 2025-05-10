package clouvider

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"stats/internal/model"
	"time"
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

type Cloud struct {
	client *http.Client
}

func New(ApiKey string) *Cloud {
	c := &Cloud{}
	c.client = &http.Client{
		Transport: &authTransport{
			apiKey: ApiKey,
			base:   http.DefaultTransport,
		},
	}
	//c.Auth = "Basic " + base64.StdEncoding.EncodeToString([]byte(c.login+":"+c.pass))
	return c

}

func (c *Cloud) ListServices() (ListResponse, error) {
	const url = "https://console.clouvider.co.uk/api/service"

	// 1) Создаём запрос
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ListResponse{}, fmt.Errorf("creating request: %w", err)
	}

	// 2) Отправляем
	res, err := c.client.Do(req)
	if err != nil {
		return ListResponse{}, fmt.Errorf("making request: %w", err)
	}
	defer res.Body.Close()

	// 3) Проверяем HTTP-статус
	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return ListResponse{}, fmt.Errorf(
			"unexpected status %d: %s",
			res.StatusCode, string(body),
		)
	}

	// 4) Декодируем JSON напрямую из Body в модель
	var listResp ListResponse
	if err := json.NewDecoder(res.Body).Decode(&listResp); err != nil {
		return ListResponse{}, fmt.Errorf("decoding JSON: %w", err)
	}

	return listResp, nil
}

func (c *Cloud) Balance() (BalanceModel, error) {
	const url = "https://console.clouvider.co.uk/api/balance"
	req, err := http.NewRequest(http.MethodGet, url, nil)
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

func (c *Cloud) Stat() model.ClouviderStats {
	stat := model.ClouviderStats{}

	balCh := make(chan string, 1)
	srvCh := make(chan int8, 1)
	errorCh := make(chan string, 2)
	go func() {
		cvBalResp, err := c.Balance()
		if err != nil {
			fmt.Println(err)
			//stat.Error = err.Error()
			errorCh <- err.Error()
		} else {
			//stat.Balance = cvBalResp.Details.AccCredit
			balCh <- cvBalResp.Details.AccCredit
		}
	}()
	go func() {
		cvServersResp, err := c.ListServices()
		if err != nil {
			fmt.Println(err)
			errorCh <- err.Error()
			//stat.Error = err.Error()
		} else {
			srvCh <- int8(len(cvServersResp.Services))
			//stat.ServersCount = int8(len(cvServersResp.Services))
		}

	}()
	for i := 0; i < 2; i++ {
		select {
		case svc := <-balCh:
			stat.Balance = svc
		case svc := <-srvCh:
			stat.ServersCount = svc
		case svc := <-errorCh:
			stat.Error = svc
			//return stat
		case <-time.After(2 * time.Second):
			// если ни один канал не ответил за 2 секунды
			stat.Error = fmt.Sprintf("operation timed out after %v", 2*time.Second)
			//return stat
		}
	}
	return stat

}

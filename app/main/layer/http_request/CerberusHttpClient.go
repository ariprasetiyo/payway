package httprequest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"payway/app/main/layer/model"
	"payway/app/main/layer/server"
	"payway/app/main/layer/status"
	"payway/app/main/layer/util"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	TOKO_CERBERUS_HOST                   string
	TOKO_CERBERUS_ENABLE_LOG             bool
	TOKO_CERBERUS_ENABLE_BODY_HEADER_LOG bool
	TOKO_CERBERUS_HTTP_TIMEOUT_IN_MS     int64
)

func NewCerberusHttpClient() *CerberusHttpClient {
	TOKO_CERBERUS_HOST = os.Getenv(util.CONFIG_TOKO_CERBERUS_HOST)
	TOKO_CERBERUS_ENABLE_LOG, _ = strconv.ParseBool(os.Getenv(util.CONFIG_TOKO_CERBERUS_ENABLE_LOG))
	TOKO_CERBERUS_ENABLE_BODY_HEADER_LOG, _ = strconv.ParseBool(os.Getenv(util.CONFIG_TOKO_CERBERUS_ENABLE_BODY_HEADER_LOG))
	TOKO_CERBERUS_HTTP_TIMEOUT_IN_MS, _ = strconv.ParseInt(os.Getenv(util.CONFIG_TOKO_CERBERUS_HTTP_TIMEOUT_IN_MS), 10, 32)
	return &CerberusHttpClient{httpClient: httpClientCerberus()}
}

type CerberusHttpClient struct {
	httpClient *http.Client
}

const (
	cerberusAuthUrl = "/oauth/merchant/validate/"
)

func httpClientCerberus() *http.Client {
	httpTimeout := time.Duration(rand.Int63n(TOKO_CERBERUS_HTTP_TIMEOUT_IN_MS))
	httpClientLogger := server.NewHttpClientLogger(TOKO_CERBERUS_ENABLE_LOG, TOKO_CERBERUS_ENABLE_BODY_HEADER_LOG)
	client := &http.Client{Timeout: httpTimeout * time.Millisecond, Transport: httpClientLogger}
	return client
}

func (hc CerberusHttpClient) DoRequestToCerberus(ceberusAggregatorModel model.CeberusAggregatorModel) model.CeberusAggregatorResponse {

	req, err := http.NewRequest(util.POST, buildAuthPathUrl(ceberusAggregatorModel.ClientId), bytes.NewReader([]byte(ceberusAggregatorModel.JsonBodyRequest)))
	if err != nil {
		logrus.Error("Error Occurred clientid", ceberusAggregatorModel.ClientId, err)
	}

	req.Header.Set(util.HEADER_AUTHORIZATION, ceberusAggregatorModel.Authorization)
	req.Header.Set(util.HEADER_SIGNATURE, ceberusAggregatorModel.Signature)
	req.Header.Set(util.HEADER_REQUEST_TIME, ceberusAggregatorModel.RequestTime)
	req.Header.Set(util.HEADER_SOURCE_URL, ceberusAggregatorModel.SourceUrl)
	req.Header.Set(util.HEADER_METHOD, ceberusAggregatorModel.Method)
	req.Header.Set(util.HEADER_CONTENT_TYPE, util.CONTENT_TYPE_APPLICATION_JSON)

	response, err := hc.httpClient.Do(req)
	if err != nil {
		logrus.Error("Error sending request to API endpoint clientid", ceberusAggregatorModel.ClientId, err)
	}
	defer response.Body.Close()

	var cerberusResponse model.CeberusAggregatorResponse
	cerberusResponse.Status = status.CERBERUS_UNAUTHORIZED
	if response.StatusCode != http.StatusOK {
		return cerberusResponse
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Error("Couldn't parse response body clientid", ceberusAggregatorModel.ClientId, err)
		return cerberusResponse
	}

	error := json.Unmarshal(body, &cerberusResponse)
	if error != nil {
		logrus.Warn("error unmarshal json from cerberus client id", ceberusAggregatorModel.ClientId, error.Error())
		return cerberusResponse
	}
	return cerberusResponse
}

func buildAuthPathUrl(username string) string {
	return TOKO_CERBERUS_HOST + cerberusAuthUrl + username
}

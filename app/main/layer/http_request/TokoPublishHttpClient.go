package httprequest

import (
	"bytes"
	"encoding/base64"
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
	TOKO_PUBLISH_HOST               string
	TOKO_PUBLISH_ENABLE_LOG         bool
	TOKO_PUBLISH_USERNAME           string
	TOKO_PUBLISH_PASSWORD           string
	TOKO_PUBLISH_HTTP_TIMEOUT_IN_MS int64
)

func NewTokoPublishttpClient() *TokoPublishHttpClient {
	TOKO_PUBLISH_HOST = os.Getenv(util.CONFIG_TOKO_PUBLISH_HOST)
	TOKO_PUBLISH_ENABLE_LOG, _ = strconv.ParseBool(os.Getenv(util.CONFIG_TOKO_PUBLISH_ENABLE_LOG))
	TOKO_PUBLISH_USERNAME = os.Getenv(util.CONFIG_TOKO_PUBLISH_USERNAME)
	TOKO_PUBLISH_PASSWORD = os.Getenv(util.CONFIG_TOKO_PUBLISH_PASSWORD)
	TOKO_PUBLISH_HTTP_TIMEOUT_IN_MS, _ = strconv.ParseInt(os.Getenv(util.CONFIG_TOKO_PUBLISH_HTTP_TIMEOUT_IN_MS), 10, 32)
	return &TokoPublishHttpClient{httpClient: httpClientTokoPublish()}
}

type TokoPublishHttpClient struct {
	httpClient *http.Client
}

const (
	createReciptUrl = "/api/tokopublish/print-receipt/"
)

func httpClientTokoPublish() *http.Client {
	httpTimeout := time.Duration(rand.Int63n(TOKO_PUBLISH_HTTP_TIMEOUT_IN_MS))
	httpClientLogger := server.NewHttpClientLogger(TOKO_PUBLISH_ENABLE_LOG, false)
	client := &http.Client{Timeout: httpTimeout * time.Millisecond, Transport: httpClientLogger}
	return client
}

func (hc TokoPublishHttpClient) DoRequestCreateReceipt(receiptRequest model.CreateReciptRequest) model.ResponseModel {

	var responseModel model.ResponseModel
	responseModel.RequestId = receiptRequest.RequestId
	responseModel.Type = receiptRequest.Type
	responseModel.Body = model.CreateReciptResponseBody{}
	responseModel.Status = status.CREATE_RECEIPT_FAILED
	responseModel.StatusMessage = status.CREATE_RECEIPT_STATUS_DESCRIPTION(status.CREATE_RECEIPT_FAILED)

	requestJsonBody, _ := json.Marshal(receiptRequest)
	req, err := http.NewRequest(util.POST, buildCreateReceiptUrl(receiptRequest.Body.MerchantId), bytes.NewReader(requestJsonBody))
	if err != nil {
		logrus.Error("Error Occurred clientid", receiptRequest.RequestId, err.Error())

	}

	req.Header.Set(util.HEADER_AUTHORIZATION, buildBasicAuth())
	req.Header.Set(util.HEADER_CONTENT_TYPE, util.CONTENT_TYPE_APPLICATION_JSON)

	response, err := hc.httpClient.Do(req)
	if err != nil {
		logrus.Error("Error sending request to API endpoint ", receiptRequest.RequestId, err.Error())
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Error("Couldn't parse response body ", receiptRequest.RequestId, err.Error())
	}

	errorUnmarshal := json.Unmarshal(body, &responseModel)
	if errorUnmarshal != nil {
		logrus.Error("error unmarshal json from cerberus ", receiptRequest.RequestId, errorUnmarshal.Error())
	}
	return responseModel

}

func buildCreateReceiptUrl(merchantId string) string {
	return TOKO_PUBLISH_HOST + createReciptUrl + merchantId
}

func buildBasicAuth() string {
	credential := TOKO_PUBLISH_USERNAME + ":" + TOKO_PUBLISH_PASSWORD
	credentialBase64 := base64.StdEncoding.EncodeToString([]byte(credential))
	return "Basic " + credentialBase64
}

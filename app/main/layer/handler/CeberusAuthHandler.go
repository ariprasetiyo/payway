package handler

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	httprequest "payway/app/main/layer/http_request"
	"payway/app/main/layer/interfaces"
	"payway/app/main/layer/model"
	"payway/app/main/layer/status"
	"payway/app/main/layer/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewCerberusAuthorization(hcCerberus *httprequest.CerberusHttpClient) interfaces.Cerberus {
	return &cerberusAuthorization{httpClientCerberus: hcCerberus}
}

type cerberusAuthorization struct {
	httpClientCerberus *httprequest.CerberusHttpClient
}

func (cerberus cerberusAuthorization) Execute() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientId := c.GetHeader(util.HEADER_CLIENT_ID)
		signature := c.GetHeader(util.HEADER_SIGNATURE)
		requestTime := c.GetHeader(util.HEADER_REQUEST_TIME)
		authorization := c.GetHeader(util.HEADER_AUTHORIZATION)
		httpMethod := c.Request.Method
		sourceUrl := c.Request.URL.String()
		jsonRequestBody := getBodyRequest(c, clientId)

		if util.IsEmptyString(clientId) && util.IsEmptyString(signature) &&
			util.IsEmptyString(requestTime) && util.IsEmptyString(authorization) &&
			util.IsEmptyString(httpMethod) && util.IsEmptyString(sourceUrl) &&
			util.IsEmptyString(jsonRequestBody) {
			unauthorized(c)
			logrus.Info("invalid request", clientId, " signature:", signature, " requestTime:", requestTime, "authorization:", authorization, " httpMethod:", httpMethod, " sourceUrl:", sourceUrl, " request body:", jsonRequestBody)
			return
		}

		var cerberusModel model.CeberusAggregatorModel
		cerberusModel.Authorization = authorization
		cerberusModel.ClientId = clientId
		cerberusModel.RequestTime = requestTime
		cerberusModel.Signature = signature
		cerberusModel.SourceUrl = sourceUrl
		cerberusModel.JsonBodyRequest = `{"map":` + jsonRequestBody + `}`
		cerberusModel.Method = httpMethod

		if !isValidAuth(cerberus, cerberusModel) {
			unauthorized(c)
			return
		}
	}
}

func getBodyRequest(c *gin.Context, clientId string) string {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Error("error read request body", clientId, err.Error())
		unauthorized(c)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))
	jsonRequestBody := string(jsonData)
	if len(strings.TrimSpace(jsonRequestBody)) == 0 {
		jsonRequestBody = util.EMPTY_JSON_OBJECT
	}
	return jsonRequestBody
}

func unauthorized(c *gin.Context) {
	c.Header("WWW-Authenticate", "Unauthorized")
	c.AbortWithStatus(http.StatusUnauthorized)
}

func isValidAuth(cerberus cerberusAuthorization, ceberusAggregatorModel model.CeberusAggregatorModel) bool {
	cerberusResponse := cerberus.httpClientCerberus.DoRequestToCerberus(ceberusAggregatorModel)
	if cerberusResponse.Status == status.CERBERUS_SUCCEES {
		return true
	}
	logrus.Infoln("invalid auth client id", ceberusAggregatorModel.ClientId)
	return false
}

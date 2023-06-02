package handler

import (
	"net/http"
	httprequest "payway/app/main/layer/http_request"
	"payway/app/main/layer/interfaces"
	"payway/app/main/layer/model"
	"payway/app/main/layer/repository"
	"payway/app/main/layer/status"
	"payway/app/main/layer/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

var (
	databaseImpl          repository.Database
	tokoPublishHttpClient *httprequest.TokoPublishHttpClient
)

func NewCreateReciptHandler(httpClient *httprequest.TokoPublishHttpClient, db *repository.Database) interfaces.Handler {
	databaseImpl = *db
	tokoPublishHttpClient = httpClient
	return &createReciptHandler{}
}

type createReciptHandler struct {
}

func (handler *createReciptHandler) Execute(c *gin.Context) {

	var merchantId string = c.Params.ByName(util.PARAM_MERCHANT_ID)
	var createReceiptRequest *model.CreateReciptRequest
	var responseModel model.ResponseModel
	var response = model.Response{http.StatusOK, responseModel}

	if err := c.ShouldBindBodyWith(&createReceiptRequest, binding.JSON); err != nil {
		logrus.Errorln("invalid request body merchant id", merchantId, err)
		response = model.Response{http.StatusBadRequest, responseModel}
	} else if util.IsEmptyString(createReceiptRequest.Body.ReceiptSendType) &&
		util.IsEmptyString(createReceiptRequest.Body.BuyerName) &&
		util.IsEmptyString(createReceiptRequest.Body.BuyerPaymentStatus) {
		response = model.Response{http.StatusBadRequest, responseModel}
	} else {
		responseModel.RequestId = createReceiptRequest.RequestId
		responseModel.Type = createReceiptRequest.Type
		responseModel.Body = model.CreateReciptResponseBody{}
		if createReceiptRequest.Body.BuyerPayment < 0 {
			responseModel.Status = status.CREATE_RECEIPT_INVALID_BUYER_AMOUNT
			responseModel.StatusMessage = status.CREATE_RECEIPT_STATUS_DESCRIPTION(status.CREATE_RECEIPT_INVALID_BUYER_AMOUNT)
		} else if merchantId != createReceiptRequest.Body.MerchantId {
			responseModel.Status = status.CREATE_RECEIPT_INVALID_MERCHANT_ID_PATH_URL
			responseModel.StatusMessage = status.CREATE_RECEIPT_STATUS_DESCRIPTION(status.CREATE_RECEIPT_INVALID_MERCHANT_ID_PATH_URL)
		} else {
			responseModel = doProcess(c, *createReceiptRequest)
		}
		response = model.Response{http.StatusOK, responseModel}
	}
	c.JSON(response.HttpStatus, response.Response)
}

func doProcess(c *gin.Context, createReceiptRequest model.CreateReciptRequest) model.ResponseModel {
	if createReceiptRequest.Body.MerchantIdType == status.USER_ID_TYPE_CHATAT {
		createReceiptRequest.Body.MerchantId = databaseImpl.GetMerchantIdByUserIdChatat(c, createReceiptRequest.Body.MerchantId)
	}
	return tokoPublishHttpClient.DoRequestCreateReceipt(createReceiptRequest)
}

package handler

import (
	"net/http"
	securityModel "payway/app/main/domain/security/model"
	"payway/app/main/layer/interfaces"
	"payway/app/main/layer/util"
	model "payway/app/main/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func NewGenerateTokenHandler() interfaces.Handler {
	return &generateTokenHandler{}
}

type generateTokenHandler struct {
}

func (gth *generateTokenHandler) Execute(ctx *gin.Context) {
	xTimestamp := ctx.GetHeader(util.HEADER_X_TIMESTAMP)
	xClientKey := ctx.GetHeader(util.HEADER_X_CLIENT_KEY)
	xSignature := ctx.GetHeader(util.HEADER_X_SIGNATURE)
	contentType := ctx.GetHeader(util.HEADER_X_CONTENT_TYPE)
	xPartnerId := ctx.GetHeader(util.HEADER_X_PARTNER_ID)
	tokenRequest := securityModel.TokenRequest{}
	err := ctx.ShouldBindBodyWith(tokenRequest, binding.JSON)
	util.IsErrorDoPrintWithMessage(&xPartnerId, &err)

	gth.isValid(&xTimestamp, &xClientKey, &xSignature, &contentType, &xPartnerId, &tokenRequest)
}

func (gth *generateTokenHandler) isValid(xTimestamp *string, xClientKey *string, xSignature *string, contentType *string, xPartnerId *string, tokenRequest *securityModel.TokenRequest) *securityModel.TokenResponse {

	if util.IsEmptyString(*xTimestamp) {
		responseCode := gth.responseCode("400", "00", "00")
		responseCodeModel := model.Response{ResponseCode: responseCode, ResponseMessage: util.HEADER_X_TIMESTAMP + " is empty"}
		return &securityModel.TokenResponse{Response: responseCodeModel}
	} else if util.IsEmptyString(*xClientKey) {
		responseCode := gth.responseCode("400", "00", "00")
		responseCodeModel := model.Response{ResponseCode: responseCode, ResponseMessage: util.HEADER_X_CLIENT_KEY + " is empty"}
		return &securityModel.TokenResponse{Response: responseCodeModel}
	} else if util.IsEmptyString(*xSignature) {
		responseCode := gth.responseCode("400", "00", "00")
		responseCodeModel := model.Response{ResponseCode: responseCode, ResponseMessage: util.HEADER_X_SIGNATURE + " is empty"}
		return &securityModel.TokenResponse{Response: responseCodeModel}
	} else if util.IsEmptyString(*contentType) {
		responseCode := gth.responseCode("400", "00", "00")
		responseCodeModel := model.Response{ResponseCode: responseCode, ResponseMessage: util.HEADER_X_CONTENT_TYPE + " is empty"}
		return &securityModel.TokenResponse{Response: responseCodeModel}
	} else if util.IsEmptyString(*xPartnerId) {
		responseCode := gth.responseCode("400", "00", "00")
		responseCodeModel := model.Response{ResponseCode: responseCode, ResponseMessage: util.HEADER_X_PARTNER_ID + " is empty"}
		return &securityModel.TokenResponse{Response: responseCodeModel}
	} else if tokenRequest == nil {
		responseCode := gth.responseCode("400", "00", "00")
		responseCodeModel := model.Response{ResponseCode: responseCode, ResponseMessage: util.HEADER_X_CLIENT_KEY + " is empty"}
		return &securityModel.TokenResponse{Response: responseCodeModel}
	} else if util.IsEmptyString(tokenRequest.GrantType) {
		responseCode := gth.responseCode("400", "00", "00")
		responseCodeModel := model.Response{ResponseCode: responseCode, ResponseMessage: util.HEADER_X_CLIENT_KEY + " is empty"}
		return &securityModel.TokenResponse{Response: responseCodeModel}
	}
	return nil
}

func (gth *generateTokenHandler) responseCode(httpStatusCode string, serviceCode string, caseCode string) string {
	return httpStatusCode + serviceCode + caseCode
}

func unauthorized(c *gin.Context) {
	c.Header("WWW-Authenticate", "Unauthorized")
	c.AbortWithStatus(http.StatusUnauthorized)
}

func badRequest(c *gin.Context) {
	c.JSON(400, "")
}

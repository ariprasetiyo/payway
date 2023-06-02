package status

const (
	CREATE_RECEIPT_SUCCEES                      int = 100
	CREATE_RECEIPT_FAILED                       int = 300
	CREATE_RECEIPT_INVALID_MERCHANT_ID_PATH_URL int = 301
	CREATE_RECEIPT_INVALID_BUYER_AMOUNT         int = 302
	CREATE_RECEIPT_INVALID_RECEIPT_TYPE         int = 303
)

func CREATE_RECEIPT_STATUS_DESCRIPTION(value int) string {
	switch value {
	case CREATE_RECEIPT_SUCCEES:
		return "success"
	case CREATE_RECEIPT_INVALID_RECEIPT_TYPE:
		return "invalid receipt type"
	case CREATE_RECEIPT_INVALID_MERCHANT_ID_PATH_URL:
		return "not match merchant id with path url"
	case CREATE_RECEIPT_INVALID_BUYER_AMOUNT:
		return "invalid buyer amount"
	default:
		return "create receipt failed"
	}
}

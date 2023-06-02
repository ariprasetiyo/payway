package model

type CreateReciptRequest struct {
	RequestModel
	Body CreateReciptRequestBody `json:"body"`
}

type CreateReciptRequestBody struct {
	RequestModel
	Language           string       `json:"language"`
	LogoUrl            string       `json:"logo_url"`
	ReceiptSendType    string       `json:"receipt_send_type"`
	CashierName        string       `json:"cashier_name"`
	BuyerName          string       `json:"buyer_name"`
	BuyerEmail         string       `json:"buyer_email"`
	BuyerPhoneNumber   string       `json:"buyer_phone_number"`
	OrderNumber        string       `json:"order_number"`
	InvoiceNumber      string       `json:"invoice_number"`
	BuyerPaymentStatus string       `json:"buyer_payment_status"`
	BuyerPayment       int64        `json:"buyer_payment"`
	BuyerChange        int64        `json:"buyer_change"`
	CreatedAt          string       `json:"created_at"`
	MerchantIdType     string       `json:"merchant_id_type"`
	MerchantId         string       `json:"merchant_id"`
	OrderItems         []OrderItems `json:"order_items"`
}

type OrderItems struct {
	ProductName  string `json:"product_name"`
	Quantity     int32  `json:"quantity"`
	ProductPrice int64  `json:"product_price"`
}

package types

type Request struct {
	OrderID           string `json:"order_id"`
	OrderDesc         string `json:"order_desc"`
	Currency          string `json:"currency"`
	Amount            string `json:"amount"`
	Signature         string `json:"signature"`
	MerchantID        string `json:"merchant_id"`
	ServerCallbackURL string `json:"server_callback_url"`
}

type APIResponse struct {
	Response Response `json:"response"`
}

type Response struct {
	OrderID                 string `json:"order_id"`
	MerchantID              int    `json:"merchant_id"`
	OrderDesc               string `json:"order_desc"`
	Signature               string `json:"signature"`
	Amount                  int    `json:"amount"`
	Currency                string `json:"currency"`
	Version                 string `json:"version"`
	ResponseURL             string `json:"response_url"`
	ServerCallbackURL       string `json:"server_callback_url"`
	PaymentSystems          string `json:"payment_systems"`
	PaymentMethod           string `json:"payment_method"`
	DefaultPaymentSystem    string `json:"default_payment_system"`
	Lifetime                int    `json:"lifetime"`
	MerchantData            string `json:"merchant_data"`
	Preauth                 string `json:"preauth"`
	SenderEmail             string `json:"sender_email"`
	Delayed                 string `json:"delayed"`
	Lang                    string `json:"lang"`
	ProductID               string `json:"product_id"`
	RequiredRectoken        string `json:"required_rectoken"`
	Verification            string `json:"verification"`
	VerificationType        string `json:"verification_type"`
	Rectoken                string `json:"rectoken"`
	ReceiverRectoken        string `json:"receiver_rectoken"`
	DesignID                int    `json:"design_id"`
	Subscription            string `json:"subscription"`
	SubscriptionCallbackURL string `json:"subscription_callback_url"`
}

type ErrorResponse struct {
	Status  string `json:"response_status"`
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}

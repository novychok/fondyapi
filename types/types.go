package types

type APIRequest struct {
	OrderID    string `json:"order_id"`
	OrderDesc  string `json:"order_desc"`
	Currency   string `json:"currency"`
	Amount     string `json:"amount"`
	Signature  string `json:"signature"`
	MerchantID string `json:"merchant_id"`
}

type Request struct {
	OrderID             string `json:"order_id"`
	MerchantID          int    `json:"merchant_id"`
	Amount              int    `json:"amount"`
	Currency            string `json:"currency"`
	OrderStatus         string `json:"order_status"`
	TranType            string `json:"tran_type"`
	SenderCellPhone     string `json:"sender_cell_phone"`
	SenderAccount       string `json:"sender_account"`
	MaskedCard          string `json:"masked_card"`
	CardBin             int    `json:"card_bin"`
	CardType            string `json:"card_type"`
	RRN                 string `json:"rrn"`
	ApprovalCode        string `json:"approval_code"`
	ResponseCode        int    `json:"response_code"`
	ResponseDescription string `json:"response_description"`
	ReversalAmount      int    `json:"reversal_amount"`
	SettlementAmount    int    `json:"settlement_amount"`
	SettlementCurrency  string `json:"settlement_currency"`
	OrderTime           string `json:"order_time"`
	SettlementDate      string `json:"settlement_date"`
	ECI                 int    `json:"eci"`
	PaymentSystem       string `json:"payment_system"`
	SenderEmail         string `json:"sender_email"`
	PaymentID           int    `json:"payment_id"`
	ActualAmount        int    `json:"actual_amount"`
	ActualCurrency      string `json:"actual_currency"`
	ProductID           string `json:"product_id"`
	MerchantData        string `json:"merchant_data"`
	VerificationStatus  string `json:"verification_status"`
	Rectoken            string `json:"rectoken"`
	RectokenLifetime    string `json:"rectoken_lifetime"`
	AdditionalInfo      string `json:"additional_info"`
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

type MidResponse struct {
	Status    string `json:"response_status"`
	Url       string `json:"checkout_code"`
	PaymentId int    `json:"payment_id"`
}

type ErrorResponse struct {
	Status  string `json:"response_status"`
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}

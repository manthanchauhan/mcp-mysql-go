package loan

// RenewRebookSpinWheelOfferPieResponseDto represents a single offer pie item
// We're not using this currently, but keeping for reference
type RenewRebookSpinWheelOfferPieResponseDto struct {
	OfferId              int64   `json:"offerId"`
	InterestRate         float64 `json:"interestRate"`
	ProcessingFee        float64 `json:"processingFee"`
	TopUpAmount          float64 `json:"topUpAmount"`
	TopUpInCurrencyValue string  `json:"topUpInCurrencyValue"`
	ColorCode            string  `json:"colorCode"`
	IsActive             bool    `json:"isActive"`
	Description          string  `json:"description"`
	IsSelected           bool    `json:"isSelected"`
}

// RenewRebookSpinWheelConsumerAppOfferListResponseDto represents the response from the API
type RenewRebookSpinWheelConsumerAppOfferListResponseDto struct {
	ActualTopUpInCurrency string `json:"actualTopUpInCurrency"`
	BestOfferId           int64  `json:"bestOfferId"`
	// Ignoring offerList as requested
}

// RenewalOfferResponse represents the API response structure
type RenewalOfferResponse struct {
	Result RenewRebookSpinWheelConsumerAppOfferListResponseDto `json:"result"`
	Error  *string                                             `json:"error"`
	Status int                                                 `json:"status"`
}

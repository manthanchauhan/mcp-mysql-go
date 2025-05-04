package loan

type Loan struct {
	UserId              int     `json:"userId"`
	LoanId              int     `json:"loanId"`
	Name                string  `json:"name"`
	Mobile              string  `json:"mobile"`
	Status              string  `json:"status"`
	TotalLoanAmount     float64 `json:"totalLoanAmount"`
	IsGoldLoanTaken     bool    `json:"isGoldLoanTaken"`
	IsPersonalLoanTaken bool    `json:"isPersonalLoanTaken"`
}

type LoanListResponse struct {
	Result []Loan  `json:"result"`
	Error  *string `json:"error"`
	Status int     `json:"status"`
}

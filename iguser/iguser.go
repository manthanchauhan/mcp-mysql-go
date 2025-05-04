package iguser

type IgUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Username string `json:"username"`
	PhotoURL string `json:"photo_url"`
}

type IgUserResponse struct {
	Success bool    `json:"success"`
	Data    *IgUser `json:"data"`
	Error   string  `json:"error,omitempty"`
}

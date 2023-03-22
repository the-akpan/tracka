package schemas

type Controller struct {
	SecretKey  string `json:"secret_key"`
	BlockKey   string `json:"block_key"`
	CookieName string `json:"cookie_name"`
}

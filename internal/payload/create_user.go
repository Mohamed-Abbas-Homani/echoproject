package payload

type CreateUserPayload struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	OrganizationId uint64 `json:"organization_id"`
}

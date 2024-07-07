package payload

type UserUpdatePayload struct {
	Id             uint64 `param:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	OrganizationId uint64 `json:"organization_id"`
}

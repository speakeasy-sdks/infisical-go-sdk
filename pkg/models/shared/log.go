package shared

type Log struct {
	ID          *string  `json:"_id,omitempty"`
	ActionNames []string `json:"actionNames,omitempty"`
	Actions     []Action `json:"actions,omitempty"`
	Channel     *string  `json:"channel,omitempty"`
	CreatedAt   *string  `json:"createdAt,omitempty"`
	IPAddress   *string  `json:"ipAddress,omitempty"`
	UpdatedAt   *string  `json:"updatedAt,omitempty"`
	User        *string  `json:"user,omitempty"`
	Workspace   *string  `json:"workspace,omitempty"`
}

package shared

type Membership struct {
	Role      *string `json:"role,omitempty"`
	Status    *string `json:"status,omitempty"`
	User      *User   `json:"user,omitempty"`
	Workspace *string `json:"workspace,omitempty"`
}

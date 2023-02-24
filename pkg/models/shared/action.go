package shared

type Action struct {
	Name      *string  `json:"name,omitempty"`
	Payload   *Payload `json:"payload,omitempty"`
	User      *string  `json:"user,omitempty"`
	Workspace *string  `json:"workspace,omitempty"`
}

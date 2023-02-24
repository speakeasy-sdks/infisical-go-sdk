package shared

type Workspace struct {
	ID           *string       `json:"_id,omitempty"`
	Environments []Environment `json:"environments,omitempty"`
	Name         *string       `json:"name,omitempty"`
	Organization *string       `json:"organization,omitempty"`
}

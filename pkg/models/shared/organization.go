package shared

type Organization struct {
	ID         *string `json:"_id,omitempty"`
	CustomerID *string `json:"customerId,omitempty"`
	Name       *string `json:"name,omitempty"`
}

package shared

type Payload struct {
	NewSecretVersion *string `json:"newSecretVersion,omitempty"`
	OldSecretVersion *string `json:"oldSecretVersion,omitempty"`
}

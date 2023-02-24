package shared

type Key struct {
	Encryptedkey *string `json:"encryptedkey,omitempty"`
	Nonce        *string `json:"nonce,omitempty"`
	Receiver     *string `json:"receiver,omitempty"`
	Sender       *Sender `json:"sender,omitempty"`
	Workspace    *string `json:"workspace,omitempty"`
}

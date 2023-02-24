package shared

type User struct {
	ID                  *string `json:"_id,omitempty"`
	CreatedAt           *string `json:"createdAt,omitempty"`
	Email               *string `json:"email,omitempty"`
	EncryptedPrivateKey *string `json:"encryptedPrivateKey,omitempty"`
	FirstName           *string `json:"firstName,omitempty"`
	Iv                  *string `json:"iv,omitempty"`
	LastName            *string `json:"lastName,omitempty"`
	PublicKey           *string `json:"publicKey,omitempty"`
	Tag                 *string `json:"tag,omitempty"`
	UpdatedAt           *string `json:"updatedAt,omitempty"`
}

package shared

type SecretVersion struct {
	ID                    *string `json:"_id,omitempty"`
	CreatedAt             *string `json:"createdAt,omitempty"`
	Environment           *string `json:"environment,omitempty"`
	IsDeleted             *string `json:"isDeleted,omitempty"`
	Secret                *string `json:"secret,omitempty"`
	SecretKeyCiphertext   *string `json:"secretKeyCiphertext,omitempty"`
	SecretKeyIV           *string `json:"secretKeyIV,omitempty"`
	SecretKeyTag          *string `json:"secretKeyTag,omitempty"`
	SecretValueCiphertext *string `json:"secretValueCiphertext,omitempty"`
	SecretValueIV         *string `json:"secretValueIV,omitempty"`
	SecretValueTag        *string `json:"secretValueTag,omitempty"`
	Type                  *string `json:"type,omitempty"`
	UpdatedAt             *string `json:"updatedAt,omitempty"`
	User                  *string `json:"user,omitempty"`
	Version               *int64  `json:"version,omitempty"`
	Workspace             *string `json:"workspace,omitempty"`
}

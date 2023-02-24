package shared

type Secret struct {
	ID                      *string  `json:"_id,omitempty"`
	CreatedAt               *string  `json:"createdAt,omitempty"`
	SecretCommentCiphertext *string  `json:"secretCommentCiphertext,omitempty"`
	SecretCommentIV         *string  `json:"secretCommentIV,omitempty"`
	SecretCommentTag        *string  `json:"secretCommentTag,omitempty"`
	SecretKeyCiphertext     *string  `json:"secretKeyCiphertext,omitempty"`
	SecretKeyIV             *string  `json:"secretKeyIV,omitempty"`
	SecretKeyTag            *string  `json:"secretKeyTag,omitempty"`
	SecretValueCiphertext   *string  `json:"secretValueCiphertext,omitempty"`
	SecretValueIV           *string  `json:"secretValueIV,omitempty"`
	SecretValueTag          *string  `json:"secretValueTag,omitempty"`
	Type                    *string  `json:"type,omitempty"`
	UpdatedAt               *string  `json:"updatedAt,omitempty"`
	User                    *string  `json:"user,omitempty"`
	Version                 *float64 `json:"version,omitempty"`
	Workspace               *string  `json:"workspace,omitempty"`
}

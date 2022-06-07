package entity

type Authentication struct {
	Token string `gorm:"not null"`
}

type RefreshTokenPayload struct {
	RefreshToken string `json:"refreshToken" validate:"required,jwt"`
}

type NewLogin struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type NewAccessToken struct {
	AccessToken string `json:"accessToken"`
}

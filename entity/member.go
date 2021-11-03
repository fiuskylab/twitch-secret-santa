package entity

// Member is the secret santa participant
type Member struct {
	// Name is the member's Twitch's username
	Name string `json:"name"`
	// Secretis the secret friend Twitch's username
	Secret string `json:"secret"`
}

package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	ImagerURL  string `json:"image_url"`
}

func FormatUser(user User, token string) UserFormatter {
	formattedUser := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
		ImagerURL:  user.AvatarFileName,
	}

	return formattedUser
}

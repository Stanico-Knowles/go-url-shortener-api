package userenums

const (
	// Error messages
	EMAIL_IS_REQUIRED              string = "Email is required"
	EMAIL_IS_INVALID               string = "Invalid email"
	PASSWORD_IS_REQUIRED           string = "Password is required"
	FIRST_NAME_IS_REQUIRED         string = "First name is required"
	LAST_NAME_IS_REQUIRED          string = "Last name is required"
	PASSWORD_LENGTH_INVALID        string = "Password must be between 8 and 20 characters"
	USER_NOT_FOUND                 string = "User not found"
	USER_WITH_EMAIL_ALREADY_EXISTS string = "A user with this email already exists"
	SOMETHING_WENT_WRONG           string = "Unable to create user at this time"
)

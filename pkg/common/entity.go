package common

type Config struct {
	Port string `mapstructure:"PORT"`
}

type UserData struct {
	UserID    int32
	FirstName string
	City      string
	Phone     string
	Height    float32
	Married   bool
}

type NotFoundList struct {
	UsersNotFound []int32
}

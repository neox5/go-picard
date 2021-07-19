package info

type Author struct {
	Name  string
	Alias string
	Email string
}

func Neox5() *Author {
	return &Author{
		Name:  "Christian Faustmann",
		Alias: "neox5",
		Email: "faustmannchr@gmail.com",
	}
}

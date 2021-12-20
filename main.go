package main

func main() {
	app := App{
		Members: []Member{
			{
				Name:             "Default Member",
				Email:            "default@example.com",
				RegistrationDate: "Tue Nov 10 10:00:00 2009",
			},
		},
	}

	app.run()
}

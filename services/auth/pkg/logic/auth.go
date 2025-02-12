package logic

func Auth(email, password *string) bool {
	if *email == "1" && *password == "1" {
		return true
	} else {
		return false
	}
}

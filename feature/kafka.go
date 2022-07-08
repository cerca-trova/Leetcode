package feature

import "os"

func Return_user_id() int {
	user_id := os.Geteuid()
	return user_id
}

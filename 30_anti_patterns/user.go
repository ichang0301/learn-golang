// https://quii.gitbook.io/learn-go-with-tests/meta/anti-patterns

// naming test doubles: https://quii.dev/Start_naming_your_test_doubles_correctly

// simple made easy: https://www.infoq.com/presentations/Simple-Made-Easy

package anti_patterns

import "net/http"

type User struct {
	// Some user fields
}
type UserService interface {
	Register(newUser User) error
}

func NewRegistrationHandler(userService UserService) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// parse user
		// register user
		// check error, send response
	}
}

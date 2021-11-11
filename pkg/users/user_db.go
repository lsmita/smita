/*
	Package users provides a mock of user database which has username and passwword map.
*/
package users

var USER_DB = make(map[string]string)

func init() {
	USER_DB["foo"] = "bar"
	USER_DB["good"] = "stuff"
}

/* Check whether the user/passwd are matching the one in user database */
func Valid(user, pswd string) bool {
	if p, ok := USER_DB[user]; ok {
		if (len(p) == len(pswd)) && (p == pswd) {
			return true
		}
	}

	return false
}

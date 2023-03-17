package consts

const (
	UserStateNormal = 1

	UserStateFreeze = 2

	UserStateClose = 3
)

func IsUserNormal(state int) bool {
	return UserStateNormal == state
}

func IsUserFreeze(state int) bool {
	return UserStateFreeze == state
}

func IsUserClose(state int) bool {
	return UserStateClose == state
}

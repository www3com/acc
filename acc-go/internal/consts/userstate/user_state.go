package userstate

const (
	UserStateNormal = 1

	UserStateFreeze = 2

	UserStateClose = 3
)

func IsNormal(state int) bool {
	return UserStateNormal == state
}

func IsFreeze(state int) bool {
	return UserStateFreeze == state
}

func IsClose(state int) bool {
	return UserStateClose == state
}

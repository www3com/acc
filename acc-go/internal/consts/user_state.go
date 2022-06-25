package consts

const (
	UserStateNormal = 1

	UserStateFreeze = 2

	UserStateClose = 3
)

func IsUserNormal(state int) bool {
	if UserStateNormal == state {
		return true
	}
	return false
}

func IsUserFreeze(state int) bool {
	if UserStateFreeze == state {
		return true
	}
	return false
}

func IsUserClose(state int) bool {
	if UserStateClose == state {
		return true
	}
	return false
}

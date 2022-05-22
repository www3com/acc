package consts

const (
	USER_STATE_NORMAL = 1

	USER_STATE_FREEZE = 2

	USER_STATE_CLOSE = 3
)

func IsUserNormal(state int) bool {
	if USER_STATE_NORMAL == state {
		return true
	}
	return false
}

func IsUserFreeze(state int) bool {
	if USER_STATE_FREEZE == state {
		return true
	}
	return false
}

func IsUserClose(state int) bool {
	if USER_STATE_CLOSE == state {
		return true
	}
	return false
}

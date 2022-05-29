package e

func GetMessage(lang string, key int) string {
	switch lang {
	case "zh":
		return zh[key]
	case "en":
		return en[key]
	}
	return ""
}

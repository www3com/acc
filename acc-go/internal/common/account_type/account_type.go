package account_type

var accountType = map[int]string{
	1: "初始账户",
	2: "余额调整",
	3: "收入",
	4: "支出",
	5: "借出",
	6: "借入",
	7: "收债",
	8: "还债",
	9: "转账",
}

func GetValue(key int) string {
	return accountType[key]
}

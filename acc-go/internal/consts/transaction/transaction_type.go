package transaction

const (
	// TypeExpenses 支出
	TypeExpenses = 1

	// TypeIncome 收入
	TypeIncome = 2

	// TypeLend 借出
	TypeLend = 3

	// TypeBorrow 借入
	TypeBorrow = 4

	// TypeDebtIn 收债
	TypeDebtIn = 5

	// TypeDebtOut 还债
	TypeDebtOut = 6

	// TypeTransfer 转账
	TypeTransfer = 7
)

var Mapper = map[int]string{
	TypeExpenses: "支出",
	TypeIncome:   "收入",
	TypeLend:     "借出",
	TypeBorrow:   "借入",
	TypeDebtIn:   "收债",
	TypeDebtOut:  "还债",
	TypeTransfer: "转账",
}

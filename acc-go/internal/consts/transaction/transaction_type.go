package transaction

const (
	// TypeExpenses 支出
	TypeExpenses = 1

	// TypeIncome 收入
	TypeIncome = 2

	// TypeTransfer 转账
	TypeTransfer = 3

	// TypeLend 借出
	TypeLend = 4

	// TypeBorrow 借入
	TypeBorrow = 5

	// TypeDebtIn 收债
	TypeDebtIn = 6

	// TypeDebtOut 还债
	TypeDebtOut = 7
)

var Mapper = map[int]string{
	TypeExpenses: "支出",
	TypeIncome:   "收入",
	TypeTransfer: "转账",
	TypeLend:     "借出",
	TypeBorrow:   "借入",
	TypeDebtIn:   "收债",
	TypeDebtOut:  "还债",
}

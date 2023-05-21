package transaction

const (
	// TypeExpenses 支出
	TypeExpenses = iota + 1

	// TypeIncome 收入
	TypeIncome

	// TypeTransfer 转账
	TypeTransfer

	// TypeLend 借出
	TypeLend

	// TypeBorrow 借入
	TypeBorrow

	// TypeDebtIn 收债
	TypeDebtIn

	// TypeDebtOut 还债
	TypeDebtOut

	TypeBalanceChange

	TypeReceivablesChange

	TypeDebtChange
)

var Mapper = map[int]string{
	TypeExpenses:          "支出",
	TypeIncome:            "收入",
	TypeTransfer:          "转账",
	TypeLend:              "借出",
	TypeBorrow:            "借入",
	TypeDebtIn:            "收债",
	TypeDebtOut:           "还债",
	TypeBalanceChange:     "余额变更",
	TypeReceivablesChange: "债权变更",
	TypeDebtChange:        "负债变更",
}

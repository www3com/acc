package account

const (

	// TypeAsset 账户类型 - 资产
	TypeAsset = iota + 1

	// TypeReceivables 账户类型 - 应收款（债权）
	TypeReceivables

	// TypeDebt 账户类型 - 负债
	TypeDebt

	// TypeEquity 账户类型 - 权益
	TypeEquity

	// TypeIncome 账户类型 - 收入
	TypeIncome

	// TypeExpenses 账户类型 - 支出
	TypeExpenses
)

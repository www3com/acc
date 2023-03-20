package model

import (
	"github.com/shopspring/decimal"
	"github.com/upbos/go-saber/db"
	"gorm.io/gorm"
)

type Account struct {
	ID         int64           `json:"id" gorm:"primary_key"`
	CreateTime int64           `json:"-"`
	UpdateTime int64           `json:"-"`
	LedgerId   int64           `json:"-"`
	Type       int             `json:"-"`
	Name       string          `json:"name"`
	Code       string          `json:"code"`
	Level      int             `json:"level"`
	Debit      decimal.Decimal `json:"debit"`
	Credit     decimal.Decimal `json:"credit"`
	Balance    decimal.Decimal `json:"balance"`
	Icon       string          `json:"icon"`
	CurrencyId int             `json:"currencyId"`
	Remark     string          `json:"remark"`
	Children   []*Account      `json:"children" gorm:"-"`
}

func (Account) TableName() string {
	return "acc_account"
}

type max struct {
	Value string
}

type AccountDao struct{}

func (d *AccountDao) Insert(account *Account) error {
	return db.DB.Create(account).Error
}

func (d *AccountDao) Update(account *Account) error {
	sql := "update acc_account set name = ?, remark = ?, update_time = ? where id = ? and ledger_id = ?"
	return db.DB.Exec(sql, account.Name, account.Remark, account.UpdateTime, account.ID, account.LedgerId).Error
}

func (d *AccountDao) BatchInsert(tx *gorm.DB, ledgerId, tLedgerId int64, now int64) (err error) {
	sql := `insert into acc_account (ledger_id, type, name, code, level, debit, credit, balance, icon, currency_id, remark, create_time, update_time)
			select  ?, type, name, code, level, 0, 0, 0, icon, currency_id, remark, ?, ? from tpl_account where ledger_id = ?`
	return tx.Exec(sql, ledgerId, now, now, tLedgerId).Error
}

func (d *AccountDao) Get(ledgerId, accountId int64) (*Account, error) {
	var account Account
	err := db.DB.Where("ledger_id = ? and id = ?", ledgerId, accountId).First(&account).Error
	return &account, err
}

func (d *AccountDao) List(ledgerId int64) (accounts []*Account, err error) {
	err = db.DB.Model(Account{}).Where("ledger_id = ? and type in (1,2)", ledgerId).Order("id asc").Find(&accounts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return accounts, nil
}

func (d *AccountDao) MaxCode(ledgerId int64, code string, childLevel int) (string, error) {
	var m max
	err := db.DB.Raw("select max(code) value from acc_account where ledger_id = ? and code like ? and level = ?", ledgerId, code+"%", childLevel).Scan(&m).Error
	return m.Value, err
}

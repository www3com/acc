package service

import (
	"acc/internal/consts"
	"acc/internal/dao"
	"errors"
	"github.com/upbos/go-saber/log"
	"gorm.io/gorm"
	"time"
)

type Ledger struct {
	LedgerId int64  `form:"ledgerId" binding:"required"`
	Name     string `form:"name"`
	OwnerId  int64  `form:"ownerId"`
}

type LedgerService struct{}

func (s *LedgerService) New(tx *gorm.DB, tLedgerId int64, ownerId int64) error {
	tLedger, err := tplLedgerDao.Get(tLedgerId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return consts.ErrTemplateNotFound
	}
	now := time.Now().UnixMilli()
	ledger := dao.Ledger{
		OwnerId:    ownerId,
		TLedgerId:  tLedger.ID,
		Name:       tLedger.Name,
		Icon:       tLedger.Icon,
		Remark:     tLedger.Remark,
		Selected:   1,
		CreateTime: now,
		UpdateTime: now,
	}
	if err := ledgerDao.CancelDefault(ownerId); err != nil {
		return err
	}
	ledgerId, err := ledgerDao.Insert(tx, &ledger)
	if err != nil {
		return err
	}
	if err := accountDao.BatchInsert(tx, ledgerId, tLedgerId, now); err != nil {
		return err
	}
	if err := projectDao.BatchInsert(tx, ledgerId, tLedgerId, now); err != nil {
		return err
	}
	if err := memberDao.BatchInsert(tx, ledgerId, tLedgerId, now); err != nil {
		return err
	}
	if err := supplierDao.BatchInsert(tx, ledgerId, tLedgerId, now); err != nil {
		return err
	}
	return nil
}

func (s *LedgerService) List(ownerId int64) (*[]dao.Ledger, error) {
	ledgers, err := ledgerDao.List(ownerId)
	if err != nil {
		log.Error(err, "list ledger")
		return nil, err
	}
	return ledgers, nil
}

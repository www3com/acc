import { makeAutoObservable } from 'mobx';
import { listLedgers } from '@/services/ledger';
import { getCurrentLedger, setCurrentLedger } from '@/components/session';


export class LayoutStore {
  visible = false;
  ledgers: [] = [];

  constructor() {
    makeAutoObservable(this);
    this.list();
  }

  setVisible(visible: boolean) {
    this.visible = visible;
  }

  * list(): any {
    let ledgers = yield listLedgers();
    if (ledgers == null || ledgers.length == 0) {
      return;
    }
    this.ledgers = ledgers;

    // 本地缓存中不存在账本
    let ledger = getCurrentLedger();
    if (ledger == null) {
      setCurrentLedger(ledgers[0]);
    }

    // 本地缓存中存在账本，但是属于其他人的账本
    for (const l of ledgers) {
      if (l.id == ledger.id) {
        return;
      }
    }

    setCurrentLedger(ledger[0]);
  }
}
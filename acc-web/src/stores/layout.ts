import { makeAutoObservable } from 'mobx';
import { listLedgers } from '@/services/ledger';

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
  }
}
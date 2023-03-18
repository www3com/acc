import { makeAutoObservable } from 'mobx';

import { listAccount, saveAccount } from '@/services/account';


interface AccountDO {
  id?: number,
  name?: string,
  remark?: string,
  parentId: number
}

export class Account {
  accounts: [] = [];

  dialogProps = {
    item: {},
    open: false,
  };

  constructor() {
    makeAutoObservable(this);
  }

  * list(): any {
    const r = yield listAccount();
    this.accounts = r.data;
  }

  showDialog(account: AccountDO) {
    this.dialogProps = { open: true, item: account };
  }

  closeDialog() {
    this.dialogProps.open = false;
  }

  * saveAccount(account: AccountDO) {
    yield saveAccount(account);
    this.closeDialog();
  }

}
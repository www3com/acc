import {makeAutoObservable} from 'mobx';

import {listAccounts, saveAccount} from '@/services/account';


interface AccountDO {
    id?: number,
    name?: string,
    remark?: string,
    parentId: number
}

export class Account {
    accountDetails = {
        total: 0.00,
        debt: 0.00,
        netAmount: 0.00,
        details: []
    };

    dialogProps = {
        item: {},
        open: false,
    };

    constructor() {
        makeAutoObservable(this);
    }

    * list(): any {
        const r = yield listAccounts();
        this.accountDetails = r.data;
    }

    showDialog(account: AccountDO) {
        this.dialogProps = {open: true, item: account};
    }

    closeDialog() {
        this.dialogProps.open = false;
    }

    * saveAccount(account: AccountDO) {
        yield saveAccount(account);
        this.closeDialog();
    }

}
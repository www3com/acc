import {makeAutoObservable} from 'mobx';

import {deleteAccount, listAccounts, saveAccount} from '@/services/account';
import request, {OK} from "@/components/Request";
import {notification} from "antd";


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
        loading: false,
    };

    constructor() {
        makeAutoObservable(this);
    }

    * list(): any {
        const r = yield listAccounts();
        this.accountDetails = r.data;
    }

    showDialog(account: AccountDO) {
        this.dialogProps = {...this.dialogProps, open: true, item: account};
        console.log(account)
    }

    closeDialog() {
        this.dialogProps.open = false;
    }

    * saveAccount(account: AccountDO) {
        try {
            this.dialogProps.loading = true
            yield saveAccount(account);
            this.closeDialog();
            this.list()
        } finally {
            this.dialogProps.loading = false
        }
    }

    * deleteAccount(code: string): any {
        let r = yield deleteAccount(code);
        if (r.code == OK) {
            this.list()
            return
        }
        notification.warning({
            message: r.message,
        })
    }
}
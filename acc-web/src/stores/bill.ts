import {makeAutoObservable} from "mobx";
import {listAccounts, listAccountsByTypes} from "@/services/account";
import {listProjects} from "@/services/project";
import {listMembers} from "@/services/member";
import {listSuppliers} from "@/services/supplier";
import {listTotalTransaction, listTransactions, saveTransaction} from "@/services/bill";
import {Dayjs} from "dayjs";

interface TransactionParams {
    startTime: Dayjs | null,
    endTime: Dayjs | null,
    accounts: [],
    cpAccounts: [],
    projects: [],
    members: [],
    suppliers: []
}

export class BillStore {
    params: TransactionParams = {
        startTime: null,
        endTime: null,
        accounts: [],
        cpAccounts: [],
        projects: [],
        members: [],
        suppliers: []
    };

    accounts = [];
    cpAccounts = [];
    projects = [];
    members = [];
    suppliers = [];
    incomes = [];
    expenses = [];
    debts = [];

    transactions = [];
    totalTransaction = {
        income: '0.00',
        expense: '0.00',
        balance: '0.00'
    }

    loading = false;

    constructor() {
        makeAutoObservable(this);
        this.listTransactions();
        this.listTotalTransaction();
        this.listAccounts();
        this.listProjects();
        this.listMembers();
        this.listSuppliers();
    }

    setParams(params: any) {
        this.params = {...this.params, ...params}
        this.listTransactions();
        this.listTotalTransaction();
    }

    * listAccounts(): any {
        this.incomes = yield listAccountsByTypes([5]);
        this.expenses = yield listAccountsByTypes([6]);
        this.debts = yield listAccountsByTypes([3]);
        this.accounts = yield listAccountsByTypes([5, 6]);
        this.cpAccounts = yield listAccountsByTypes([1, 2, 3]);
    }

    * listProjects(): any {
        const r = yield listProjects();
        this.projects = r.data;
    }

    * listMembers(): any {
        const r = yield listMembers();
        this.members = r.data;
    }

    * listSuppliers(): any {
        const r = yield listSuppliers();
        this.suppliers = r.data;
    }

    * listTransactions(): any {
        const r = yield listTransactions(this.params)
        this.transactions = r.data;
    }

    * listTotalTransaction(): any {
        const r = yield listTotalTransaction(this.params)
        this.totalTransaction = r.data;
    }

    * saveTransaction(transaction: any): any {
        try {
            this.loading = true
            yield saveTransaction(transaction);
            this.listTransactions()
            this.listTotalTransaction()
        } catch (e) {
            throw e;
        } finally {
            this.loading = false
        }
    }
}
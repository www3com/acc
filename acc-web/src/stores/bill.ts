import {makeAutoObservable} from "mobx";
import {listAccounts, listAccountsByTypes} from "@/services/account";
import {listProjects} from "@/services/project";
import {listMembers} from "@/services/member";
import {listSuppliers} from "@/services/supplier";
import {listTotalTransaction, listTransactions, saveTransaction} from "@/services/bill";
import {Dayjs} from "dayjs";

interface TransactionParams {
    startDate: Dayjs | null,
    endDate: Dayjs | null,
    accounts: number[],
    cpAccounts: number[],
    projects: number[],
    members: number[],
    suppliers: number[]
}

export class BillStore {

    accounts = [];
    cpAccounts = [];
    projects = [];
    members = [];
    suppliers = [];
    incomes = [];
    expenses = [];
    debts = [];

    params: TransactionParams = {
        startDate: null,
        endDate: null,
        accounts: [],
        cpAccounts: [],
        projects: [],
        members: [],
        suppliers: []
    };

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
        const p = {
            ...this.params,
            startDate: this.params.startDate?.valueOf(),
            endDate: this.params.endDate?.valueOf(),
        }
        const r = yield listTransactions(p)
        this.transactions = r.data;
    }

    * listTotalTransaction(): any {
        const p = {
            ...this.params,
            startDate: this.params.startDate?.valueOf(),
            endDate: this.params.endDate?.valueOf(),
        }
        const r = yield listTotalTransaction(p)
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
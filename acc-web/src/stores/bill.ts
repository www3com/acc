import {makeAutoObservable} from "mobx";
import {listAccounts, listExpenses, listIncomeExpenses, listIncomes, saveAccount} from "@/services/account";
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
        this.listIncomes();
        this.listExpenses();
        this.listIncomeExpenses();
        this.listCpAccounts();
        this.listProjects();
        this.listMembers();
        this.listSuppliers();
    }

    setParams(params: any) {
        this.params = {...this.params, ...params}
    }

    * listIncomes(): any {
        const r = yield listIncomes();
        this.incomes = r.data;
    }

    * listExpenses(): any {
        const r = yield listExpenses();
        this.expenses = r.data;
    }

    * listIncomeExpenses(): any {
        const r = yield listIncomeExpenses();
        this.accounts = r.data;
    }

    * listCpAccounts(): any {
        const r = yield listAccounts();
        this.cpAccounts = r.data;
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
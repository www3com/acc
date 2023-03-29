import {makeAutoObservable} from "mobx";
import {listAccounts, listExpenses, listIncomeExpenses, listIncomes} from "@/services/account";
import {listProjects} from "@/services/project";
import {listMembers} from "@/services/member";
import {listSuppliers} from "@/services/supplier";
import {listTransactions} from "@/services/bill";

export class BillStore {

    accounts = [];
    cpAccounts = [];
    projects = [];
    members = [];
    suppliers = [];
    incomes = [];
    expenses = [];

    params = {
        startDate: null,
        endDate: null,
        accounts: [],
        cpAccounts: [],
        projects: [],
        members: [],
        suppliers: []
    };

    transactions = [];

    constructor() {
        makeAutoObservable(this);
        this.listTransactions();
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
        const p = {}
        const r = yield listTransactions(p)
        this.transactions = r.data;
    }

}
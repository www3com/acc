import {makeAutoObservable} from "mobx";
import {listAccounts, ListIncomeExpenses} from "@/services/account";
import {listProjects} from "@/services/project";
import {listMembers} from "@/services/member";
import {listSuppliers} from "@/services/supplier";

export class BillStore {

    accounts = [];
    cpAccounts = [];
    projects = [];
    members = [];
    suppliers = [];

    params = {
        startDate: null,
        endDate: null,
        accounts: [],
        cpAccounts: [],
        projects: [],
        members: [],
        suppliers: []
    }

    constructor() {
        makeAutoObservable(this);
        this.ListIncomeExpenses();
        this.listCpAccounts();
        this.listProjects();
        this.listMembers();
        this.listSuppliers();
    }

    * ListIncomeExpenses(): any {
        const r = yield ListIncomeExpenses();
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

    selectAccounts(items: any) {
        console.log(items)
        // this.selectedAccounts = items;
    }

}
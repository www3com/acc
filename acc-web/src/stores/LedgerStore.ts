import {makeAutoObservable} from "mobx";

export class LedgerStore {

  constructor() {
    makeAutoObservable(this)
  }


}

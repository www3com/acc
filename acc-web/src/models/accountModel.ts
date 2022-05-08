import {useCallback, useState} from "react";
import {listAccount} from '../services/account';

export default function accountModel() {
  const [accounts, setAccounts] = useState<any>();

  const list = async (ledgerId: number) => {
    let accounts = await listAccount(ledgerId)
    setAccounts(accounts.data)
  }

  return {accounts, list}
}

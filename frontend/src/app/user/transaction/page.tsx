// NOTE: pages display transaction
//TODO: Create transaction page

"use client";

import { useEffect, useState } from "react";
import style from '../../styles/componentstyle/defaultbuttonstyle.module.css';
import Button from '../../components/button/button'
import axios from "axios";

type Transaction = {
  id: string;
  type: string;
  amount: number;
  date: string;
};

axios.defaults.baseURL = 'http://localhost:8080';

export default function Transaction() {
  const [amount, setAmount] = useState("");
  const [accountNumber, setAccountNumber] = useState<string | null>(null);
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [result, setResult] = useState("Processing...");

  const reloadTransactionsHistory = async () => {
    try {
      const res = await axios.get(`/api/account/transaction/${accountNumber}/history`, {
        headers: {
          "Authorization": `Bearer ${localStorage.getItem("authToken")}`,
        }
      }
      );
      console.log(res)
      setTransactions(res.data.data.transactions); // gán dữ liệu thẳng từ API
    } catch (err) {
      console.error(err);
    }
  }

  useEffect(() => {
    setAccountNumber(localStorage.getItem("accountNumber"));
    if (accountNumber) {
      reloadTransactionsHistory();
    }
  }, [accountNumber]);

  const handleTransaction = async (type: "dep" | "with") => {
    try {
      const res = await axios.post(`/api/account/transaction/${accountNumber}`,
        {
          "amount": parseInt(amount),
          "type": type,
          "time": ""
        }, {
        headers: {
          "Authorization": `Bearer ${localStorage.getItem("authToken")}`,
        }
      });
      if (res.status === 200) {
        setResult("Transaction successful!. Your new balance is" + res.data.newBalance);
      }
      reloadTransactionsHistory()
    } catch (err) {
      console.error(err);
    };
  }

  return (
    <div className="flex gap-8">
      {/* Cột bên trái */}
      <div className="w-1/3">
        <input
          type="text"
          placeholder=""
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          className="border p-2 rounded mb-2 w-full"
        />

        {/* Các nút giao dịch */}
        <div className="mt-4 space-y-2">
          <Button onClick={() => handleTransaction("with")} label="Withdraw" className={style.greenButton} />
          <Button onClick={() => handleTransaction("dep")} label="Deposit" className={style.redButton} />
          <p>Result: {result} </p>
        </div>
      </div>

      {/* Cột bên phải */}
      <div className="w-2/3">
        <h2 className="text-lg font-bold mb-4">Transaction History</h2>
        <table className="table-auto border-collapse border border-gray-300 w-full text-left">
          <thead>
            <tr className="bg-gray-100">
              <th className="border border-gray-300 text-gray-800 px-4 py-2">ID</th>
              <th className="border border-gray-300 text-gray-800 px-4 py-2">Type</th>
              <th className="border border-gray-300 text-gray-800 px-4 py-2">Amount</th>
              <th className="border border-gray-300 text-gray-800 px-4 py-2">Date</th>
            </tr>
          </thead>
          <tbody>
            {transactions.map((t: Transaction) => (
              <tr key={t.id}>
                <td className="bg-gray-100 border border-gray-300 text-gray-800 px-4 py-2">{t.id}</td>
                <td className="bg-gray-100 border border-gray-300 text-gray-800 px-4 py-2">{t.type}</td>
                <td className="bg-gray-100 border border-gray-300 text-gray-800 px-4 py-2">{t.amount}</td>
                <td className="bg-gray-100 border border-gray-300 text-gray-800 px-4 py-2">{new Date(t.date).toLocaleString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
// NOTE: show all transaction table history 

// TODO: display as table for view only

"use client";

import style from '../../styles/componentstyle/defaultbuttonstyle.module.css';
import Button from '../../components/button/button';

type Transaction = {
  id: string;
  type: string;
  amount: number;
  date: string;
};

export default function TransHistory() {
  const [transactions, setTransactions] = useState<Transaction[]>([]);

  return (
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
                <td className="border border-gray-300 text-gray-800 px-4 py-2">{t.id}</td>
                <td className="border border-gray-300 text-gray-800 px-4 py-2">{t.type}</td>
                <td className="border border-gray-300 text-gray-800 px-4 py-2">{t.amount}</td>
                <td className="border border-gray-300 text-gray-800 px-4 py-2">{new Date(t.date).toLocaleString()}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
  );
}


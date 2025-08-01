// NOTE: pages display transaction
//TODO: Create transaction page

"use client";

import { useState } from "react";
import style from '../../styles/componentstyle/defaultbuttonstyle.module.css';
import Button from '../../components/button/button'
import axios from "axios";

export default function Transaction() {
  const [amount, setAmount] = useState("");
  let accountNumber = localStorage.getItem("accountNumber") // get account number from local storage
  const handleWithDraw = async () => {

    const res = await axios.post(`/api/transaction/${accountNumber}`, { amount: 300 });
    console.log(res.data);
    console.log("withdraw from account:", amount);
  };

  const handleDeposit = () => {
    console.log("deposit into account:", amount);
  };

  return (
    <div>
      <input
        type="text"
        placeholder=""
        value={amount}
        onChange={(e) => setAmount(e.target.value)}
        className="border p-2 rounded mb-2"
      />

      {/* Các nút giao dịch */}
      <div className="mt-4">
        <Button onClick={handleWithDraw} label="Withdraw" className={style.customButton} />
        <Button onClick={handleDeposit} label="Deposit" className={style.customButton} />
      </div>
    </div>
  );
}

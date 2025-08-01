// NOTE: user for transfer between account
// TODO: create transfer page
// FIX: event handler cannot passed to client component props

'use client';

import React, { useState } from 'react';
import style from '../../styles/componentstyle/defaultbuttonstyle.module.css';
import Button from '../../components/button/button';
import { useRouter} from "next/navigation";

export default function Transfer() {
  const router = useRouter();
  const [amount, setAmount] = useState<string>('');
  const [destinationAccount, setDestinationAccount] = useState<string>('');
  const [bank, setBank] = useState<string>('');

  const sendAmount = async () =>{
    console.log("sending to account...")
    const body = {
      to_account_number : destinationAccount,
      amount: amount,
      bank: bank
    };

    try {
      const res = await fetch("http://localhost:8080/user/transfer", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(body), // Gửi username thay vì email
      });

      const data = await res.json();

      if (res.ok) {
        console.log("Tranfer success:", body);
        router.push("/dashboard");
      } else {
        console.error("Login failed:", data.message);
        alert("Login failed: " + (data.message || "Unknown error"));
      }
    } catch (error) {
      console.error("Request error:", error);
      alert("Network error or backend not reachable.");
    }

  }

  return (
    <div>
        <span>Welcome To World Bank</span>
        <div style={{ marginBottom: '1rem' }}>
        <label>Số tiền:</label><br />
        <input
          type="number"
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
          placeholder="Nhập số tiền"
        />
      </div>

      <div style={{ marginBottom: '1rem' }}>
        <label>Tài khoản đích:</label><br />
        <input
          type="text"
          value={destinationAccount}
          onChange={(e) => setDestinationAccount(e.target.value)}
          placeholder="Nhập số tài khoản"
        />
      </div>

      <div style={{ marginBottom: '1rem' }}>
        <label>Ngân hàng:</label><br />
        <select
          value={bank}
          onChange={(e) => setBank(e.target.value)}
        >
          <option value="">-- Chọn ngân hàng --</option>
          <option value="VCB">Vietcombank</option>
          <option value="TCB">Techcombank</option>
          <option value="BIDV">BIDV</option>
          <option value="ACB">ACB</option>
          <option value="MB">MB Bank</option>
        </select>
      </div>
        <Button onClick={sendAmount} label="Confirm" className={style.customButton} />
    </div>
  );
}

// NOTE: pages display transaction
//TODO: Create transaction page

"use client";

import style from '../../styles/componentstyle/defaultbuttonstyle.module.css';
import Button from '../../components/button/button'

export default function Transaction() {
  const handleWithDraw = () =>{
    console.log("withdraw")
  }
  const handleDeposit = () =>{
    console.log("deposit")
  }
  return (
    <div>
        <span>Welcome To World Bank</span>
        <Button onClick={handleWithDraw} label="WithDraw" className={style.customButton} />
        <Button onClick={handleDeposit} label="Deposit" className={style.customButton} />
    </div>
  );
}

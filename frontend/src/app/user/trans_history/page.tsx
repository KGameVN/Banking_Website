//NOTE: show all transaction table history 

// TODO: display as table for view only

"use client";

import style from '../../styles/componentstyle/defaultbuttonstyle.module.css';
import Button from '../../components/button/button';

export default function TransHistory() {
  const k = () =>{
    console.log("withdraw")
  }
  const l = () =>{
    console.log("deposit")
  }
  return (
    <div>
        <span>Welcome To World Bank</span>
        <Button onClick={k} label="WithDraw" className={style.customButton} />
        <Button onClick={l} label="Deposit" className={style.customButton} />
    </div>
  );
}

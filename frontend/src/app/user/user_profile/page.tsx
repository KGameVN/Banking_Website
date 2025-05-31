// NOTE: user for transfer between account
// TODO: create transfer page
// FIX: event handler cannot passed to client component props

'use client';

import style from '../../styles/componentstyle/defaultbuttonstyle.module.css';
import Button from '../../components/button/button';

export default function Transfer() {
  const sendAmount = () =>{
    console.log("sending to account...")
  }
  return (
    <div>
        <span>Welcome To World Bank</span>
        <Button onClick={sendAmount} label="Confirm" className={style.customButton} />
    </div>
  );
}

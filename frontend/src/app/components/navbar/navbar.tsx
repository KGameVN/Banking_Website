//
"use client"
import React from "react";
import {useRouter} from "next/navigation";

type NavbarProps = {
  className: string;
}

// TODO: should fixed number or not?

export default function NavBar({className}: NavbarProps){
  const router = useRouter();
  return (
    <nav className={`bg-gray-800 text-white px-6 py-3 flex justify-between items-center ${className}`}>
      <div className="text-lg font-bold">🏦 MyBank</div>
      <div className="flex gap-4">
        <button onClick={() => router.push("/dashboard")} className="hover:underline">Dashboard</button>
        <button onClick={() => router.push("/user/transaction")} className="hover:underline">Transaction</button>
        <button onClick={() => router.push("/user/transfer")} className="hover:underline">Transfer</button>
        <button onClick={() => router.push("/login")} className="hover:underline text-red-300">Logout</button>
      </div>
    </nav>
  );
}

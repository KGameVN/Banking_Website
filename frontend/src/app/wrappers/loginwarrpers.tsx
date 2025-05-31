"use client";

import dynamic from "next/dynamic";

const Login = dynamic(() => import("../user/login"));

export default function LoginWrapper() {
  return <Login />;
}

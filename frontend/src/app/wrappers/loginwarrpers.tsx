"use client";

import dynamic from "next/dynamic";

const Login = dynamic(() => import("../components/login"));

export default function LoginWrapper() {
  return <Login />;
}

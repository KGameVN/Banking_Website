"use client";

import { useState } from "react";

export default function Login() {
  const [open, setOpen] = useState(false);
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const res = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      const data = await res.json();

      if (res.ok) {
        console.log("Login success:", data);
        setOpen(false);
      } else {
        console.error("Login failed:", data.message);
        alert("Login failed: " + (data.message || "Unknown error"));
      }
    } catch (error) {
      console.error("Request error:", error);
      alert("Network error or backend not reachable.");
    }
  };


  return (
    <>
      <button
        onClick={() => setOpen(true)}
        className="mt-4 px-4 py-2 bg-blue-500 text-white rounded"
      >
        Login
      </button>

      {open && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white p-6 rounded shadow-lg w-96 relative">
            <button
              onClick={() => setOpen(false)}
              className="absolute top-2 right-2 text-gray-500"
            >
              ✕
            </button>
            <h2 className="text-xl font-bold mb-4 text-center">Login</h2>
            <form onSubmit={handleSubmit} className="space-y-4">
              <input
                type="email"
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="w-full p-2 border rounded"
              />
              <input
                type="password"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="w-full p-2 border rounded"
              />
              <button
                type="submit"
                className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600"
              >
                Submit
              </button>
            </form>
          </div>
        </div>
      )}
    </>
  );
}

"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import Image from "next/image";

export default function Login() {
  const [open, setOpen] = useState(false);
  const [username, setUsername] = useState(""); // Đổi email thành username
  const [password, setPassword] = useState("");
  const [rememberMe, setRememberMe] = useState(false);
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    try {
      const token = localStorage.getItem("authToken")
      const res = await fetch("http://localhost:8080/api/user/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`
        },
        body: JSON.stringify({ username, password, rememberMe }), // Gửi username thay vì email
      });

      const data = await res.json();

      if (res.ok) {
        console.log("Login success:", data);
        localStorage.setItem("accountNumber", data.user.accountnumber) // stored account number in local storage
        setOpen(false);
        router.push("/dashboard");
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
        className="mt-4 px-4 py-2 bg-gradient-to-r from-blue-500 to-indigo-500 text-white rounded-md shadow-md hover:scale-105 transform transition duration-300"
      >
        Login
      </button>

      {open && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-gradient-to-br from-white via-blue-50 to-blue-100 p-8 rounded-2xl shadow-2xl w-full max-w-md relative border border-blue-200">
            <button
              onClick={() => setOpen(false)}
              className="absolute top-2 right-2 text-gray-500"
            >
              ✕
            </button>
            <Image
              src="/images/icon/banking.png"
              alt="Logo"
              className="w-16 h-16 mx-auto mb-4"
            />
            <h2 className="text-2xl font-bold mb-6 text-center text-blue-700">
              Welcome Back
            </h2>
            <form onSubmit={handleSubmit} className="space-y-5">
              <input
                placeholder="Username"
                value={username} // Bind với username
                onChange={(e) => setUsername(e.target.value)} // Cập nhật giá trị username
                className="w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-400 text-black"
              />
              <input
                type="password"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-400 text-black"
              />

              <div className="flex items-center justify-between text-sm">
                <label className="flex items-center space-x-2">
                  <input
                    type="checkbox"
                    checked={rememberMe}
                    onChange={() => setRememberMe(!rememberMe)}
                    className="accent-blue-500 text-black"
                  />
                  <span className="text-black">Remember me</span>
                </label>
                <button
                  type="button"
                  onClick={() => alert("Forgot password clicked!")}
                  className="text-blue-500 hover:underline text-black"
                >
                  Forgot Password?
                </button>
              </div>

              <button
                type="submit"
                className="w-full bg-gradient-to-r from-blue-500 to-indigo-500 text-white py-2 rounded-md hover:scale-105 transform transition duration-300 shadow-lg"
              >
                Submit
              </button>

              <div className="text-center text-black">
                <span>Dont have an account?</span>{" "}
                <button
                  type="button"
                  onClick={() => alert("Register clicked!")}
                  className="text-blue-500 hover:underline"
                >
                  Register
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </>
  );
}

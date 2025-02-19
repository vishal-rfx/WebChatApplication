"use client";
import axios from "axios";
import { useRouter } from "next/navigation";
import React, { useState } from "react";
import { useAuthStore } from "./zustand/useAuthStore";

const Auth = () => {
  const [username, setUserName] = useState("");
  const [password, setPassword] = useState("");
  const {updateAuthName} = useAuthStore();
  const router = useRouter();

  const signUpFunc = async (event) => {
    event.preventDefault();

    try {
      const res = await axios.post(
        "http://localhost:8000/auth/signup",
        {
          username: username,
          password: password,
        },
        {
          withCredentials: true,
        }
      );

      console.log(res);
      if (res.status != 201) {
        alert("Username not created");
      } else if (res.data.message === "Username already exists") {
        alert(res.data.message);
      } else {
        alert("Username created. Please sign in");
      }
    } catch (error) {
      console.log(error);
    }
  };

  const loginFunc = async (event) => {
    event.preventDefault();
    try {
      const res = await axios.post(
        "http://localhost:8000/auth/login",
        {
          username: username,
          password: password,
        },
        {
          withCredentials: true,
        }
      );

      console.log(res);
      if (res.status != 200) {
        alert("Login unsuccessful");
      } else {
        updateAuthName(username);
        router.push("/chat");
      }
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div>
      <div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
        <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
          <form action="#" method="POST" className="space-y-6">
            <div>
              <label
                htmlFor="username"
                className="block text-sm/6 font-medium text-gray-900"
              >
                Username
              </label>
              <div className="mt-2">
                <input
                  id="username"
                  name="username"
                  type="text"
                  required
                  onChange={(e) => setUserName(e.target.value)}
                  autoComplete="username"
                  className="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
                />
              </div>
            </div>

            <div>
              <div className="flex items-center justify-between">
                <label
                  htmlFor="password"
                  className="block text-sm/6 font-medium text-gray-900"
                >
                  Password
                </label>
              </div>
              <div className="mt-2">
                <input
                  id="password"
                  name="password"
                  type="password"
                  required
                  onChange={(e) => setPassword(e.target.value)}
                  autoComplete="current-password"
                  className="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-900 outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600 sm:text-sm/6"
                />
              </div>
            </div>

            <div className="flex">
              <button
                onClick={signUpFunc}
                type="submit"
                className="flex w-1/2 justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 mr-2"
              >
                Sign Up
              </button>
              <button
                type="submit"
                onClick={loginFunc}
                className="flex w-1/2 justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-xs hover:bg-indigo-500 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
              >
                Sign in
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Auth;

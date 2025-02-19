"use client";
import React, { use, useEffect, useState } from "react";
import { useAuthStore } from "../zustand/useAuthStore";

function Chat() {
  const [msg, setMsg] = useState("");
  const [socket, setSocket] = useState(null);
  const [msgs, setMsgs] = useState([]);
  const { authName } = useAuthStore();

  useEffect(() => {
    // Establish a websocket connection
    console.log("Establishing websocket connection");
    const queryParams = new URLSearchParams({ authName }).toString();
    const wsUrl = `ws://localhost:4000/ws?${queryParams}`;
    const newSocket = new WebSocket(wsUrl);
    setSocket(newSocket);
    newSocket.onmessage = (evt) => {
      console.log(evt);
      setMsgs((prevMsgs) => [
        ...prevMsgs,
        { text: evt.data, sentByCurrentUser: false },
      ]);
    };

    // When the page closes, this cleanup function is called.
    return () => newSocket.close();
  }, []);

  const sendMsg = (e) => {
    e.preventDefault();
    if (socket && socket.readyState == WebSocket.OPEN) {
      socket.send(msg);
      setMsgs((prevMsgs) => [
        ...prevMsgs,
        { text: msg, sentByCurrentUser: true },
      ]);
      setMsg("");
    }
  };

  return (
    <div>
      <div className="msg-container">
        {msgs.map((msg, index) => (
          <div
            className={`m-5 ${
              msg.sentByCurrentUser ? "text-right" : "text-left"
            }`}
            key={index}
          >
            <span
              className={`p-3 rounded-lg ${
                msg.sentByCurrentUser ? "bg-blue-200" : "bg-green-200"
              }`}
            >
              {msg.text}
            </span>
          </div>
        ))}
      </div>

      <form
        onSubmit={sendMsg}
        className="max-w-md mx-auto my-10 fixed bottom-0 left-0 right-0 p-4 bg-white"
      >
        <div className="relative">
          <input
            type="text"
            value={msg}
            onChange={(e) => setMsg(e.target.value)}
            placeholder="Type your text here"
            required
            className="block w-full p-4 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
          />
          <button
            type="submit"
            className="text-white absolute end-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Send
          </button>
        </div>
      </form>
    </div>
  );
}

export default Chat;

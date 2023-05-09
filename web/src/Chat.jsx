import React, { useEffect, useState, useContext } from "react";
import { useNavigate } from "react-router-dom";
import UserContext from "./UserContext";

import axios from "axios";

const Chat = () => {
  const navigate = useNavigate();
  const [websocket, setWebsocket] = useState(null);
  const [message, setMessage] = useState("");
  const [users, setUsers] = useState([]);
  const [selectedUserId, setSelectedUserId] = useState(null);
  const [onlineUsers, setOnlineUsers] = useState([]);
  const { username, ID, setUsername, setID } = useContext(UserContext);

  useEffect(() => {
    const checkJwtToken = async () => {
      const jwtToken = localStorage.getItem("jwt");
      if (!jwtToken) {
        clearStorageAndNavigate();
        return;
      }
      try {
        const response = await axios.get("http://localhost:8080/jwtauth", {
          headers: {
            Authorization: `Bearer ${jwtToken}`,
          },
        });
        if (response.status != 200) {
          clearStorageAndNavigate();
        }
        setUsername(response.data.jwtAuthResponse.username);
        setID(response.data.jwtAuthResponse.id);
      } catch (error) {
        clearStorageAndNavigate();
      }
    };
    const clearStorageAndNavigate = () => {
      localStorage.removeItem("jwt");
      navigate("/");
    };

    checkJwtToken();
  }, []);

  const handleLogout = async () => {
    localStorage.removeItem("jwt");
    navigate("/");
  };

  useEffect(() => {
    const getUsers = async () => {
      const response = await axios.get("http://localhost:8080/users");
      if (response.status === 200) {
        setUsers(response.data.users);
      }
    };
    getUsers();
  }, []);

  useEffect(() => {
    const websocket = new WebSocket("ws://localhost:8080/ws");
    websocket.onopen = () => {
      // Once the connection is open, send the JWT to the server
      const jwtToken = localStorage.getItem("jwt");
      websocket.send(JSON.stringify({ type: "auth", token: jwtToken }));
    };
    setWebsocket(websocket);

    return () => {
      if (websocket) {
        websocket.close();
      }
    };
  }, []);

  const sendMessage = () => {
    if (websocket && message) {
      websocket.send(message);
      setMessage("");
    }
  };

  return (
    <div>
      <div className="flex h-screen">
        <div className="bg-blue-50 w-1/5 flex flex-col">
          <div className="flex-grow">
            <div className="text-blue-600 font-bold flex gap-2 mb-4">
              GO React Chat
            </div>
            {users.map((user) => (
              <div
                key={user.id}
                className={`border-b border-gray-300 py-2 px-2 ${
                  selectedUserId === user.id
                    ? "bg-blue-300 font-semibold pl-4"
                    : ""
                }`}
                onClick={() => setSelectedUserId(user.id)}
              >
                {user.username}
              </div>
            ))}
          </div>
          <div className="p-2 text-center flex items-center justify-center">
            <span className="text-gray-600 font-semibold mr-2 flex">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                fill="currentColor"
                className="w-6 h-6 "
              >
                <path
                  fillRule="evenodd"
                  d="M7.5 6a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM3.751 20.105a8.25 8.25 0 0116.498 0 .75.75 0 01-.437.695A18.683 18.683 0 0112 22.5c-2.786 0-5.433-.608-7.812-1.7a.75.75 0 01-.437-.695z"
                  clipRule="evenodd"
                />
              </svg>
              {username}
            </span>
            <button
              className="text-blue-500 font-bold py-1 px-2 rounded-md border bg-blue-100"
              onClick={handleLogout}
            >
              Logout
            </button>
          </div>
        </div>
        <div className="bg-blue-300 w-4/5 p-1 flex flex-col">
          <div className="flex-grow">messages</div>
          <div className="flex gap-2">
            <input
              type="text"
              placeholder="Message to"
              className="bg-white border p-2 rounded-lg flex-grow"
              value={message}
              onChange={(e) => setMessage(e.target.value)}
            />
            <button onClick={sendMessage}>Send</button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Chat;

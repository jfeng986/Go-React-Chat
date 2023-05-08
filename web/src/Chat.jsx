import React, { useEffect, useState, useContext } from "react";
import { useNavigate } from "react-router-dom";
import UserContext from "./UserContext";

import axios from "axios";

const Chat = () => {
  const navigate = useNavigate();
  const [websocket, setWebsocket] = useState(null);
  const [message, setMessage] = useState("");
  const [onlinePeople, setOnlinePeople] = useState({});
  //const [messages, setMessages] = useState([]);

  const { username, ID, setUsername, setID } = useContext(UserContext);

  useEffect(() => {
    const checkJwtToken = async () => {
      const jwtToken = localStorage.getItem("jwt");
      const storedUsername = localStorage.getItem("username");
      const storedUserID = localStorage.getItem("id");

      if (storedUsername) {
        setUsername(storedUsername);
      }
      if (storedUserID) {
        setID(storedUserID);
      }

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
      } catch (error) {
        clearStorageAndNavigate();
      }
    };

    const clearStorageAndNavigate = () => {
      localStorage.removeItem("jwt");
      localStorage.removeItem("username");
      localStorage.removeItem("id");
      navigate("/");
    };

    checkJwtToken();
  }, []);

  const handleLogout = async () => {
    localStorage.removeItem("jwt");
    navigate("/");
  };

  useEffect(() => {
    const websocket = new WebSocket("ws://localhost:8080/ws");
    setWebsocket(websocket);
    /*
    return () => {
      if (websocket) {
        websocket.close();
      }
    };
    */
    websocket.addEventListener("message", handleMessage);
  }, []);

  function handleMessage(event) {
    const messageData = JSON.parse(event.data);
    if ("online" in messageData) {
      showOnlineUser(messageData.online);
    } else {
      console.log(messageData);
    }
  }

  const sendMessage = () => {
    if (websocket && message) {
      websocket.send(message);
      setMessage("");
    }
  };

  /*
  useEffect(() => {
    if (ws) {
      ws.onmessage = (event) => {
        setMessages((prevMessages) => [...prevMessages, event.data]);
      };
    }
  }, [ws]);
  */

  function showOnlinePeople(onlinePeople) {
    const onlinePeopleSet = {};
    onlinePeople.forEach(({ userId, username }) => {
      onlinePeopleSet[userId] = username;
    });
    setOnlinePeople(onlinePeopleSet);
  }

  return (
    <div>
      <h1>Welcome {username}!</h1>
      <p>Your ID is: {ID}</p>
      <div className="flex h-screen">
        <div className="bg-blue-100 w-1/5 pl-4 pt-4">
          <div className="text-blue-600 font-bold flex gap-2 mb-4">
            GO React Chat
          </div>
          {Object.keys(onlinePeople).map((userId) => (
            <div key={userId} className="border-b border-gray-100 py-2">
              {onlinePeople[userId]}
            </div>
          ))}
          <button onClick={handleLogout}>Logout</button>
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

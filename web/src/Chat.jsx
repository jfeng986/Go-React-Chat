import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

const Chat = () => {
  const navigate = useNavigate();

  useEffect(() => {
    const checkJwtToken = async () => {
      const jwtToken = localStorage.getItem("jwt");
      if (jwtToken) {
        try {
          const response = await axios.get("http://localhost:8080/profile", {
            headers: {
              Authorization: `Bearer ${jwtToken}`,
            },
          });
          if (response.status != 200) {
            localStorage.removeItem("jwt");
            navigate("/");
          }
        } catch (error) {
          localStorage.removeItem("jwt");
          navigate("/");
        }
      } else {
        navigate("/");
      }
    };
    checkJwtToken();
  }, []);

  const handleLogout = async () => {
    localStorage.removeItem("jwt");
    navigate("/");
  };

  const [ws, setWs] = useState(null);
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState([]);

  useEffect(() => {
    const websocket = new WebSocket("ws://localhost:8080/ws");
    setWs(websocket);

    return () => {
      if (websocket) {
        websocket.close();
      }
    };
  }, []);

  useEffect(() => {
    if (ws) {
      ws.onmessage = (event) => {
        setMessages((prevMessages) => [...prevMessages, event.data]);
      };
    }
  }, [ws]);

  const sendMessage = () => {
    if (ws && message) {
      ws.send(message);
      setMessage("");
    }
  };

  return (
    <div>
      <h1>Welcome to Chat</h1>
      <button onClick={handleLogout}>Logout</button>
      <br />

      <h1>Chat</h1>

      <div>
        {messages.map((msg, index) => (
          <p key={index}>{msg}</p>
        ))}
      </div>

      <input
        type="text"
        value={message}
        onChange={(e) => setMessage(e.target.value)}
      />
      <button onClick={sendMessage}>Send</button>
    </div>
  );
};

export default Chat;

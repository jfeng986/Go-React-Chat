import React, { useEffect } from "react";
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

  return (
    <div>
      <h1>Welcome to Chat</h1>
      <button onClick={handleLogout}>Logout</button>
    </div>
  );
};

export default Chat;

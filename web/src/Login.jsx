import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

import axios from "axios";

function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
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
          if (response.status === 200) {
            navigate("/chat");
          }
        } catch (error) {
          localStorage.removeItem("jwt");
        }
      }
    };
    checkJwtToken();
  }, []);

  const Login = async () => {
    try {
      const response = await axios.post("http://localhost:8080/login", {
        username,
        password,
      });
      if (response.status == 200) {
        const token = response.headers["authorization"].split(" ")[1];
        localStorage.setItem("jwt", token);
        navigate("/chat");
      } else {
        alert("Invalid username or password");
      }
    } catch (error) {
      alert("Invalid username or password");
      console.error(error);
    }
  };

  const Register = async () => {
    try {
      const response = await axios.post("http://localhost:8080/register", {
        username,
        password,
      });
      if (response.status == 200) {
        navigate("/chat");
      } else {
        alert("error");
      }
    } catch (error) {
      alert("catch error");
      console.error(error);
    }
  };

  return (
    <div className="container mx-auto mt-10">
      <form
        onSubmit={(e) => {
          e.preventDefault();
        }}
        className="w-full max-w-sm mx-auto"
      >
        <div className="mb-4">
          <label className="block text-gray-700 text-sm font-bold mb-2">
            Username:
          </label>
          <input
            className="shadow border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="username"
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div className="mb-6">
          <label className="block text-gray-700 text-sm font-bold mb-2">
            Password:
          </label>
          <input
            className="shadow border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
            id="password"
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <div className="flex items-center justify-between">
          <button
            className="w-28 bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            onClick={Login}
          >
            Login
          </button>
          <button
            className="w-28 bg-green-500 hover:bg-green-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            onClick={Register}
          >
            Register
          </button>
        </div>
      </form>
    </div>
  );
}

export default Login;

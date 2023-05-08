import { useState } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import LoginAndRegister from "./LoginAndRegister";
import Chat from "./Chat";
import UserContext from "./UserContext";

function App() {
  const [username, setUsername] = useState("");
  const [ID, setID] = useState("");
  return (
    <UserContext.Provider value={{ username, setUsername, ID, setID }}>
      <Router>
        <Routes>
          <Route path="/" element={<LoginAndRegister />} />
          <Route path="/chat" element={<Chat />} />
        </Routes>
      </Router>
    </UserContext.Provider>
  );
}

export default App;

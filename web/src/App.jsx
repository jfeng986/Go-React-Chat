import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import LoginAndRegister from "./LoginAndRegister";
import Chat from "./Chat";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LoginAndRegister />} />
        <Route path="/chat" element={<Chat />} />
      </Routes>
    </Router>
  );
}

export default App;

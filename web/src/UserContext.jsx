import { createContext } from "react";

const UserContext = createContext({
  username: "",
  ID: "",
  setUsername: () => {},
  setID: () => {},
});

export default UserContext;

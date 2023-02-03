import { createContext, useState } from "react";

const UserContext = createContext();

const UserProvider = ({ children }) => {
    const [userStatus, setUserStatus] = useState({
        connected: false,
        name: "",
        token: "",
        error: false,
        errorMessage: "",
    });

    return (
        <UserContext.Provider value={{ userStatus, setUserStatus }}>
            {children}
        </UserContext.Provider>
    );
};

export { UserContext, UserProvider };

import { Navigate, useLocation } from "react-router-dom";
import { useContext } from "react";
import { UserContext } from "./UserContext";

export default function RequireAuth({ children }) {
    const { userStatus } = useContext(UserContext);
    const location = useLocation();

    if (!userStatus.connected) {
        // Redirect them to the /login page, but save the current location they were
        // trying to go to when they were redirected. This allows us to send them
        // along to that page after they login, which is a nicer user experience
        // than dropping them off on the home page.
        return <Navigate to="/" state={{ from: location }} replace />;
    }

    return children;
}

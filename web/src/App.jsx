import { Outlet } from "react-router-dom";
import BNavbar from "./components/BNavbar";

function App() {
    return (
        <div className="App">
            <BNavbar></BNavbar>
            <div className="bContainer">
                <Outlet />
            </div>
        </div>
    );
}

export default App;

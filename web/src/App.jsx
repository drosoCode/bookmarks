import { Outlet } from "react-router-dom";
import BNavbar from "./components/BNavbar";
import Footer from "./components/Footer";

function App() {
    return (
        <div className="App">
            <BNavbar></BNavbar>
            <div className="bContainer">
                <Outlet />
            </div>
            <Footer></Footer>
        </div>
    );
}

export default App;

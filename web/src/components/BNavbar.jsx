import Container from "react-bootstrap/Container";
import Nav from "react-bootstrap/Nav";
import Navbar from "react-bootstrap/Navbar";
import { LinkContainer } from "react-router-bootstrap";
import { useAPI } from "../Api";
import { UserContext } from "../components/UserContext";
import { useContext } from "react";
import { useNavigate } from "react-router-dom";

export default function BNavbar() {
    const { basePath } = useAPI();
    const { userStatus, setUserStatus } = useContext(UserContext);
    const navigate = useNavigate();
    const logout = () => {
        const data = { connected: false, name: "", token: "" };
        setUserStatus(data);
        localStorage.setItem("bookmarks", JSON.stringify(data));
        document.cookie = "bookmarktoken=";
        navigate("/");
    };

    return (
        <div className="mb-4">
            <Navbar bg="dark" variant="dark">
                <Container>
                    <Navbar.Brand href="#/home">
                        <img
                            src="./public/icon.png"
                            width="25"
                            height="25"
                            className="d-inline-block align-top"
                            href="#/home"
                            style={{ marginTop: "4px", marginRight: "5px" }}
                        />
                        Bookmarks
                    </Navbar.Brand>
                    <Nav className="me-auto">
                        <LinkContainer to="/home">
                            <Nav.Link>Home</Nav.Link>
                        </LinkContainer>
                        <LinkContainer to="/add">
                            <Nav.Link>Add</Nav.Link>
                        </LinkContainer>
                        <LinkContainer to="/tag">
                            <Nav.Link>Tags</Nav.Link>
                        </LinkContainer>
                        <LinkContainer to="/token">
                            <Nav.Link>Tokens</Nav.Link>
                        </LinkContainer>
                        <Nav.Link href={basePath + "swagger"}>API</Nav.Link>
                        {userStatus.connected ? (
                            <Nav.Link onClick={logout}>Logout</Nav.Link>
                        ) : (
                            ""
                        )}
                    </Nav>
                </Container>
            </Navbar>
        </div>
    );
}

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
            <Navbar bg="dark" variant="dark" expand="sm">
                <Container>
                    <Navbar.Brand href="#/home">
                        <img
                            src="./icons/icon256.png"
                            width="25"
                            height="25"
                            className="d-inline-block align-top"
                            href="#/home"
                            style={{ marginTop: "4px", marginRight: "5px" }}
                        />
                        Bookmarks
                    </Navbar.Brand>
                    <Navbar.Toggle aria-controls="responsive-navbar-nav" />
                    <Navbar.Collapse id="responsive-navbar-nav">
                        <Nav className="me-auto">
                            <LinkContainer to="/home">
                                <Nav.Link>
                                    <i className="fa-solid fa-home fa-sm"></i>
                                    &nbsp; Home
                                </Nav.Link>
                            </LinkContainer>
                            <LinkContainer to="/add">
                                <Nav.Link>
                                    <i className="fa-solid fa-circle-plus fa-sm"></i>
                                    &nbsp;Add
                                </Nav.Link>
                            </LinkContainer>
                            <LinkContainer to="/tag">
                                <Nav.Link>
                                    <i className="fa-solid fa-tag fa-sm"></i>
                                    &nbsp;Tags
                                </Nav.Link>
                            </LinkContainer>
                            <LinkContainer to="/token">
                                <Nav.Link>
                                    <i className="fa-solid fa-key fa-sm"></i>
                                    &nbsp;Tokens
                                </Nav.Link>
                            </LinkContainer>
                            <Nav.Link href={basePath + "swagger"}>
                                <i className="fa-solid fa-book fa-sm"></i>
                                &nbsp;API
                            </Nav.Link>
                            {userStatus.connected ? (
                                <Nav.Link onClick={logout}>
                                    <i class="fa-solid fa-right-from-bracket fa-sm"></i>
                                    &nbsp;Logout
                                </Nav.Link>
                            ) : (
                                ""
                            )}
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
        </div>
    );
}

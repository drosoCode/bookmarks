import { useNavigate } from "react-router-dom";
import { useContext, useEffect, useRef } from "react";
import { UserContext } from "../components/UserContext";
import { useAPI } from "../Api";
import Alert from "react-bootstrap/Alert";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";

export default function LoginPage(props) {
    const { userStatus, setUserStatus } = useContext(UserContext);
    const { api } = useAPI();
    const navigate = useNavigate();

    const token = useRef(null);
    const loginToken = () => {
        const data = {
            connected: true,
            name: "Test",
            token: token.current.value,
            clearLS: false,
        };
        setUserStatus(data);
        localStorage.setItem("bookmarks", JSON.stringify(data));
        document.cookie = "bookmarktoken=" + token.current.value;
        navigate("/home");
    };

    useEffect(() => {
        if (userStatus.connected) {
            navigate("/home");
        } else {
            const data = localStorage.getItem("bookmarks");
            if (
                data !== "" &&
                data !== undefined &&
                data !== null &&
                !userStatus.clearLS
            ) {
                // login using stored data in localStorage
                const user = JSON.parse(data);
                if (user !== null && user.connected) {
                    setUserStatus(user);
                    document.cookie = "bookmarktoken=" + user.token;
                    navigate("/home");
                    return;
                }
            }
            // else login using api
            api("user/login", "GET").then((d) => {
                if (d != null) {
                    const data = {
                        connected: true,
                        name: d.name,
                        token: d.token,
                        clearLS: false,
                    };
                    setUserStatus(data);
                    localStorage.setItem("bookmarks", JSON.stringify(data));
                    document.cookie = "bookmarktoken=" + d.token;
                    navigate("/home");
                }
            });
        }
    }, []);

    return (
        <div>
            {!userStatus.connected ? (
                <div>
                    <Alert key="danger" variant="danger">
                        Login Failed {userStatus.errorMessage}
                    </Alert>
                    <InputGroup className="mb-3">
                        <Form.Control placeholder="Token" ref={token} />
                        <Button variant="outline-primary" onClick={loginToken}>
                            <i className="fa-solid fa-key fa-sm"></i>
                            &nbsp; Login with Token
                        </Button>
                    </InputGroup>
                </div>
            ) : (
                ""
            )}
        </div>
    );
}

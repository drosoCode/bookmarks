import Alert from "react-bootstrap/Alert";
import Table from "react-bootstrap/Table";
import Button from "react-bootstrap/Button";
import CloseButton from "react-bootstrap/CloseButton";
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";
import { useEffect, useState, useRef } from "react";
import { useAPI } from "../Api";

export default function TokenPage(props) {
    const { api } = useAPI();
    const [data, setData] = useState([]);
    const [token, setToken] = useState(null);
    const tokenName = useRef();

    useEffect(() => {
        api("token", "GET").then((data) => {
            if (data !== null) {
                setData(data);
            }
        });
    }, [token]);

    const deleteToken = (id) => {
        api("token/" + id, "DELETE").then((d) => {
            setData(data.filter((x) => x.id != id));
        });
    };

    const addToken = () => {
        api("token", "POST", { name: tokenName.current.value }).then((data) => {
            if (data !== null) {
                setToken(data.token);
            }
        });
    };

    return (
        <div>
            {token != null ? (
                <Alert key="success" variant="success">
                    Your token is: {token}
                    &nbsp;
                    <CloseButton
                        onClick={() => {
                            setToken(null);
                        }}
                    />
                </Alert>
            ) : (
                ""
            )}
            <InputGroup className="mb-3">
                <Form.Control
                    placeholder="Token Name"
                    ref={tokenName}
                    className="bg-dark text-white"
                />
                <Button variant="success" onClick={addToken}>
                    Add Token
                </Button>
            </InputGroup>
            <br />
            <Table striped bordered hover variant="dark">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Creation Date</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {data.map((x) => (
                        <tr key={x.id}>
                            <td>{x.id}</td>
                            <td>{x.name}</td>
                            <td>
                                {new Date(x.addDate * 1000)
                                    .toISOString()
                                    .replace("T", " ")
                                    .substring(0, 19)}
                            </td>
                            <td>
                                <Button
                                    variant="danger"
                                    onClick={() => {
                                        deleteToken(x.id);
                                    }}
                                >
                                    Delete
                                </Button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </Table>
        </div>
    );
}

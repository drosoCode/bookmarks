import Table from "react-bootstrap/Table";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";
import Tag from "../components/Tag";
import { useEffect, useState, useRef } from "react";
import { useAPI } from "../Api";

export default function TagPage(props) {
    const { api } = useAPI();
    const [data, setData] = useState([]);
    const tagName = useRef();
    const tagColor = useRef();

    const update = () => {
        api("tag", "GET").then((data) => {
            if (data !== null) {
                setData(data);
            }
        });
    };

    useEffect(() => {
        update();
    }, []);

    const deleteTag = (id) => {
        api("tag/" + id, "DELETE").then((d) => {
            setData(data.filter((x) => x.id != id));
        });
    };

    const addTag = () => {
        api("tag", "POST", {
            name: tagName.current.value,
            color: tagColor.current.value,
        }).then((data) => {
            update();
        });
    };

    return (
        <div>
            <InputGroup className="mb-3">
                <Form.Control
                    placeholder="Tag Name"
                    ref={tagName}
                    className="bg-dark text-white"
                />
                <input
                    className="form-control-color bg-dark text-white"
                    type="color"
                    ref={tagColor}
                />
                <Button variant="success" onClick={addTag}>
                    <i className="fa-solid fa-plus fa-sm"></i>
                    &nbsp; Add Tag
                </Button>
            </InputGroup>
            <br />
            <Table striped bordered hover variant="dark">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Tag</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {data.map((x) => (
                        <tr key={x.id}>
                            <td>{x.id}</td>
                            <td>
                                <Tag name={x.name} color={x.color} />
                            </td>
                            <td>
                                <Button
                                    variant="danger"
                                    onClick={() => {
                                        deleteTag(x.id);
                                    }}
                                >
                                    <i className="fa-solid fa-trash fa-sm"></i>
                                    &nbsp; Delete
                                </Button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </Table>
        </div>
    );
}

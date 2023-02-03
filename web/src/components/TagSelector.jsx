import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import ListGroup from "react-bootstrap/ListGroup";
import Form from "react-bootstrap/Form";
import InputGroup from "react-bootstrap/InputGroup";
import { useState } from "react";
import Tag from "./Tag";

export default function TagSelector(props) {
    const [show, setShow] = useState(false);

    const update = () => {
        console.log("change");
        setShow(false);
    };

    return (
        <div>
            <InputGroup className="mb-3">
                <div
                    className="form-control"
                    style={{
                        backgroundColor: "#212529",
                        borderColor: "#313539",
                    }}
                >
                    <Tag name="tst" color="#ff0000" />
                </div>
                <Button
                    variant="success"
                    onClick={() => {
                        setShow(true);
                    }}
                >
                    Add Tag
                </Button>
            </InputGroup>

            <Modal
                variant="dark"
                show={show}
                onHide={() => {
                    setShow(false);
                }}
            >
                <Modal.Header closeButton>
                    <Modal.Title>Select Tags</Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    <input type="text" className="form-control" />
                    <br />
                    <ListGroup variant="dark">
                        <ListGroup.Item>
                            <Tag name="test" color="#00ff00" />
                        </ListGroup.Item>
                        <ListGroup.Item active={true}>
                            <Tag name="test" color="#00ff00" />
                        </ListGroup.Item>
                        <ListGroup.Item active={true}>
                            <Tag name="test" color="#00ff00" />
                        </ListGroup.Item>
                    </ListGroup>
                </Modal.Body>

                <Modal.Footer>
                    <Button
                        variant="danger"
                        onClick={() => {
                            setShow(false);
                        }}
                    >
                        Cancel
                    </Button>
                    <Button variant="success" onClick={update}>
                        OK
                    </Button>
                </Modal.Footer>
            </Modal>
        </div>
    );
}

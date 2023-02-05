import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import ListGroup from "react-bootstrap/ListGroup";
import InputGroup from "react-bootstrap/InputGroup";
import { useState, useEffect, useRef } from "react";
import { useAPI } from "../Api";
import Tag from "./Tag";

export default function TagSelector(props) {
    const [show, setShow] = useState(false);
    const [data, setData] = useState([]);
    const [selected, setSelected] = useState([]);
    const [visibleList, setVisibleList] = useState([]);
    const search = useRef("");
    const { api } = useAPI();

    useEffect(() => {
        api("tag", "GET").then((data) => {
            if (data !== null) {
                setData(data);
                setVisibleList(data);
            }
        });
    }, []);

    useEffect(() => {
        setSelected([]);
    }, [props.clear]);

    const filter = () => {
        const arr = [];
        data.forEach((e) => {
            if (
                search.current.value == "" ||
                e.name.indexOf(search.current.value) >= 0
            )
                arr.push(e);
        });
        setVisibleList([...arr]);
    };

    const select = (id) => {
        if (selected.includes(id)) selected.splice(selected.indexOf(id), 1);
        else selected.push(id);
        setSelected([...selected]);
    };

    return (
        <div>
            <InputGroup className="mb-3">
                <div className="form-control bg-dark text-white">
                    {data
                        .filter((t) => selected.includes(t.id))
                        .map((x) => (
                            <Tag name={x.name} color={x.color} key={x.id} />
                        ))}
                </div>
                <Button
                    variant="success"
                    onClick={() => {
                        setShow(true);
                    }}
                >
                    <i className="fa-solid fa-plus fa-sm"></i>
                    &nbsp; Add Tag
                </Button>
                <Button
                    variant="danger"
                    onClick={() => {
                        props.onChange([]);
                        setSelected([]);
                    }}
                >
                    <i className="fa-solid fa-xmark fa-sm"></i>
                    &nbsp; Clear
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
                    <Modal.Title>
                        &nbsp;
                        <i className="fa-solid fa-tag"></i>
                        &nbsp; Select Tags
                    </Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    <input
                        type="text"
                        className="form-control"
                        onChange={filter}
                        ref={search}
                    />
                    <br />
                    <ListGroup variant="dark">
                        {visibleList.map((x) => (
                            <ListGroup.Item
                                active={selected.includes(x.id)}
                                onClick={() => {
                                    select(x.id);
                                }}
                                key={x.id}
                            >
                                <Tag name={x.name} color={x.color} />
                            </ListGroup.Item>
                        ))}
                    </ListGroup>
                </Modal.Body>

                <Modal.Footer>
                    <Button
                        variant="success"
                        onClick={() => {
                            props.onChange(selected);
                            setShow(false);
                        }}
                    >
                        <i className="fa-solid fa-check fa-sm"></i>
                        &nbsp; OK
                    </Button>
                </Modal.Footer>
            </Modal>
        </div>
    );
}

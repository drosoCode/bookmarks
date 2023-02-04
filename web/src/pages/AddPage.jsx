import TagSelector from "../components/TagSelector";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { useState, useRef } from "react";
import { useAPI } from "../Api";

export default function AddPage(props) {
    const [clear, setClear] = useState(0);
    const [tags, setTags] = useState([]);
    const url = useRef();
    const checkbox = useRef();
    const { api } = useAPI();

    const add = () => {
        api("bookmark", "POST", {
            link: url.current.value,
            tags: tags,
            save: checkbox.current.checked,
        }).then((data) => {
            if (data !== null) {
                url.current.value = "";
                setClear(clear + 1);
                setTags([]);
            }
        });
    };

    return (
        <div style={{ textAlign: "left" }} className="card bg-dark">
            <div className="card-body">
                <Form>
                    <Form.Group className="mb-3">
                        <Form.Label className="text-white">Link</Form.Label>
                        <Form.Control
                            placeholder="Bookmark URL"
                            ref={url}
                            className="bg-dark text-white"
                        />
                    </Form.Group>

                    <Form.Group className="mb-3">
                        <Form.Label className="text-white">Tags</Form.Label>
                        <TagSelector
                            clear={clear}
                            onChange={(data) => {
                                setTags(data);
                            }}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formBasicCheckbox">
                        <Form.Check
                            type="checkbox"
                            label="Save HTML"
                            className="text-white"
                            ref={checkbox}
                        />
                    </Form.Group>
                    <Button variant="primary" onClick={add}>
                        ADD
                    </Button>
                </Form>
            </div>
        </div>
    );
}

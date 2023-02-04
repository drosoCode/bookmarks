import Tag from "./Tag";
import { useAPI } from "../Api";
import OverlayTrigger from "react-bootstrap/OverlayTrigger";
import Tooltip from "react-bootstrap/Tooltip";

export default function BookmarkCard(props) {
    const { basePath } = useAPI();

    return (
        <div
            className="card bg-dark text-white border-secondary h-100 mx-auto"
            style={{ maxWidth: "25rem" }}
        >
            {props.description != "" ? (
                <OverlayTrigger
                    key="bottom"
                    placement="bottom"
                    overlay={
                        <Tooltip id={`tooltip-bottom`}>
                            {props.description}
                        </Tooltip>
                    }
                >
                    <div className="card-header border-secondary">
                        {props.name}
                    </div>
                </OverlayTrigger>
            ) : (
                <div className="card-header border-secondary">{props.name}</div>
            )}
            <a href={props.link} target="_blank">
                <img
                    src={basePath + "cache/" + props.id + "/image.png"}
                    className="card-img-top"
                />
            </a>
            <div className="card-body border-secondary">
                <a
                    href={props.link}
                    className="btn btn-sm btn-primary"
                    target="_blank"
                >
                    Open
                </a>
                &nbsp;
                {props.save ? (
                    <>
                        <a
                            href={basePath + "cache/" + props.id + "/html"}
                            className="btn btn-sm btn-primary"
                            target="_blank"
                        >
                            Save
                        </a>
                        &nbsp;
                    </>
                ) : (
                    ""
                )}
                <a
                    className="btn btn-sm btn-danger"
                    onClick={() => {
                        props.onDelete(props.id);
                    }}
                >
                    Delete
                </a>
            </div>
            <div className="card-footer text-muted border-secondary">
                <span style={{ float: "left", marginTop: "3px" }}>
                    {new Date(props.addDate * 1000)
                        .toISOString()
                        .replace("T", " ")
                        .substring(0, 19)}
                </span>
                <span style={{ float: "right" }}>
                    {props.tags.map((x) => (
                        <Tag name={x.name} color={x.color} key={x.id} />
                    ))}
                </span>
            </div>
        </div>
    );
}

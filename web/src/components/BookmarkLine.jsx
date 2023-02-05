import Tag from "./Tag";
import { useAPI } from "../Api";

export default function BookmarkLine(props) {
    const { basePath } = useAPI();

    return (
        <tr>
            <td>{props.id}</td>
            <td>
                <a href={props.link} target="_blank">
                    {props.name}
                </a>
            </td>
            <td>{props.description}</td>
            <td>
                {props.tags.map((x) => (
                    <Tag name={x.name} color={x.color} key={x.id} />
                ))}
            </td>
            <td>
                {new Date(props.addDate * 1000)
                    .toISOString()
                    .replace("T", " ")
                    .substring(0, 19)}
            </td>
            <td>
                <a
                    href={props.link}
                    className="btn btn-sm btn-primary my-1"
                    target="_blank"
                >
                    <i className="fa-solid fa-arrow-up-right-from-square fa-sm"></i>
                    &nbsp; Open
                </a>
                &nbsp;
                {props.save ? (
                    <>
                        <a
                            href={basePath + "cache/" + props.id + "/html"}
                            className="btn btn-sm btn-primary my-1"
                            target="_blank"
                        >
                            <i className="fa-solid fa-floppy-disk fa-sm"></i>
                            &nbsp; Save
                        </a>
                        &nbsp;
                    </>
                ) : (
                    ""
                )}
                <a
                    className="btn btn-sm btn-danger my-1"
                    onClick={() => {
                        props.onDelete(props.id);
                    }}
                >
                    <i className="fa-solid fa-trash fa-sm"></i>&nbsp; Delete
                </a>
            </td>
        </tr>
    );
}

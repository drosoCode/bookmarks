import TagSelector from "../components/TagSelector";
import BookmarkCard from "../components/BookmarkCard";
import BookmarkLine from "../components/BookmarkLine";
import { useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import Table from "react-bootstrap/Table";
import { useAPI } from "../Api";

export default function HomePage(props) {
    const [displayedBookmarks, setDisplayedBookmarks] = useState([]);
    const [selectedTags, setSelectedTags] = useState([]);
    const [tags, setTags] = useState([]);
    const [andMode, setAndMode] = useState(false);
    const [gridMode, setGridMode] = useState(true);
    const [bookmarks, setBookmarks] = useState([]);
    const navigate = useNavigate();
    const { api } = useAPI();

    useEffect(() => {
        if (window.location.search.includes("?title=")) {
            navigate({
                pathname: "/add",
                search: window.location.search,
            });
        }
        api("tag", "GET").then((data) => {
            if (data !== null) {
                setTags(data);
            }
        });
        api("bookmark", "GET").then((data) => {
            if (data !== null) {
                setBookmarks(data);
                setDisplayedBookmarks(data);
            }
        });
    }, []);

    useEffect(() => {
        if (selectedTags.length == 0) {
            setDisplayedBookmarks(bookmarks);
            return;
        }
        if (andMode) {
            setDisplayedBookmarks(
                bookmarks.filter((b) =>
                    selectedTags.every((v) => b.tags.includes(v))
                )
            );
        } else {
            setDisplayedBookmarks(
                bookmarks.filter((b) =>
                    selectedTags.some((v) => b.tags.includes(v))
                )
            );
        }
    }, [selectedTags, andMode, bookmarks]);

    const del = (id) => {
        api("bookmark/" + id, "DELETE").then((d) => {
            setBookmarks(bookmarks.filter((x) => x.id != id));
        });
    };

    return (
        <div>
            <div className="row bg-secondary pt-3 mb-4 rounded mx-2">
                <div className="col-md-4 col-lg-4 col-xl-3 mb-3 mb-md-0">
                    <button
                        style={{ display: "inline-block" }}
                        type="button"
                        className="btn btn-primary"
                        onClick={() => {
                            setAndMode(!andMode);
                        }}
                    >
                        <i className="fa-solid fa-filter fa-sm"></i>
                        &nbsp;
                        {andMode ? "Filter AND" : "Filter OR"}
                    </button>
                    &nbsp;
                    <button
                        style={{ display: "inline-block" }}
                        type="button"
                        className="btn btn-primary"
                        onClick={() => {
                            setGridMode(!gridMode);
                        }}
                    >
                        {gridMode ? (
                            <>
                                <i className="fa-solid fa-grip fa-sm"></i>&nbsp;
                                Display Grid
                            </>
                        ) : (
                            <>
                                <i className="fa-solid fa-bars fa-sm"></i>&nbsp;
                                Display List
                            </>
                        )}
                    </button>
                </div>
                <div className="col-md-8 col-lg-8 col-xl-9">
                    <TagSelector
                        onChange={(data) => {
                            setSelectedTags(data);
                        }}
                    />
                </div>
            </div>
            {gridMode ? (
                <div className="row card-group">
                    {displayedBookmarks.map((x) => (
                        <div
                            className="col-xs-12 col-sm-6 col-md-6 col-lg-4 col-xl-2 mb-3"
                            key={x.id}
                        >
                            <BookmarkCard
                                id={x.id}
                                name={x.name}
                                addDate={x.addDate}
                                description={x.description}
                                save={x.save}
                                link={x.link}
                                tags={tags.filter((t) => x.tags.includes(t.id))}
                                onDelete={del}
                            />
                        </div>
                    ))}
                </div>
            ) : (
                <Table striped bordered hover variant="dark">
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Name</th>
                            <th>Description</th>
                            <th>Tags</th>
                            <th>Date</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {displayedBookmarks.map((x) => (
                            <BookmarkLine
                                id={x.id}
                                name={x.name}
                                addDate={x.addDate}
                                description={x.description}
                                save={x.save}
                                link={x.link}
                                tags={tags.filter((t) => x.tags.includes(t.id))}
                                onDelete={del}
                                key={x.id}
                            />
                        ))}
                    </tbody>
                </Table>
            )}
        </div>
    );
}

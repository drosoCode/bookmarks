export default function Tag(props) {
    const getColor = (color) => {
        const x =
            parseInt(color.substring(1, 2), 16) * 0.299 +
            parseInt(color.substring(3, 4), 16) * 0.587 +
            parseInt(color.substring(5, 6), 16) * 0.114;
        if (x > 10) return "#000000";
        else "#ffffff";
    };

    return (
        <h6>
            <button
                disabled
                className="badge"
                style={{
                    backgroundColor: props.color,
                    color: getColor(props.color),
                    borderWidth: 0,
                    padding: "8px",
                }}
            >
                {props.name}
            </button>
        </h6>
    );
}

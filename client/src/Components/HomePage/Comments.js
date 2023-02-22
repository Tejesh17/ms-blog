import React from "react";

const Comments = (props) => {
	return (
		<>
			<div className="AllComments mt-1">
				<ul className="list-disc list-outside">
					{props.comments &&
						props.comments.map((c) => {
							return <li key={c.commentid}>- {c.content}</li>;
						})}
				</ul>
			</div>
		</>
	);
};

export default Comments;

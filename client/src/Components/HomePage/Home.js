import React, { useEffect, useState } from "react";
import axios from "axios";
import Heading from "./Heading";
import PostCard from "./PostCard";

const Home = () => {
	const [postTitle, setPostTitle] = useState("");
	const [postCards, SetPostCards] = useState([]);

	const GetCards = async () => {
		try {
			const result = await axios.get("http://localhost:8082/posts");
			if (result.data) {
				let allposts = [];
				for (const postcard in result.data) {
					allposts.push(result.data[postcard]);
				}
				SetPostCards(allposts);
			}
		} catch (e) {
			console.log(e);
		}
	};

	const CreateCard = async () => {
		try {
			if (postTitle === "") return;
			const result = await axios.post("http://localhost:8080/posts", {
				title: postTitle,
			});
			if (result.data) {
				alert("Post Created!");
				setPostTitle("");
			}
		} catch (e) {
			console.log(e);
		}
	};

	useEffect(() => {
		GetCards();
	}, []);

	return (
		<>
			<Heading />
			<div className="CreatePost m-4">
				To create a post, enter the post title and hit create!
				<div className="flex flex-col mt-3">
					<input
						className="bg-white focus:outline-none focus:shadow-outline border border-gray-5 rounded py-1 px-4 block w-54"
						type="text"
						placeholder="Title"
						value={postTitle}
						onChange={(e) => {
							setPostTitle(e.target.value);
						}}
					/>
					<button
						className="bg-blue-500 text-white py-1 px-2 rounded mt-1 hover:bg-blue-600"
						onClick={CreateCard}
					>
						Create
					</button>
				</div>
			</div>
			<div className="AllCards flex flex-row flex-wrap justify-center">
				{postCards.map((post) => (
					<PostCard key={post.id} post={post} />
				))}
			</div>
		</>
	);
};

export default Home;

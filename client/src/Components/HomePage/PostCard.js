import React, { useEffect, useState } from 'react';
import Comments from './Comments';
import axios from 'axios';

const PostCard = (props) =>{

	const [comments, setComments]= useState([])
	const [newComment, setNewComment]= useState("")
	
	const GetComments = async (id)=>{
		try{
			let result = await axios.get(`http://localhost:8081/comments/${id}`)
			if(result.data ){
				setComments(result.data.comments)
			}
		}catch(e){
			console.log(e)
		}
	} 

	const CreateComment = async ()=>{
		try{
			if(newComment === "")return
			let result = await axios.post(`http://localhost:8081/comments`,{
				postid: props.id,
				comment: newComment
			})
			if(result.data ){
				alert("Comment added!")
				setNewComment("")
				GetComments(props.id)
			}
		}catch(e){
			console.log(e)
		}
	}

	useEffect(()=>{
		GetComments(props.id)
	},[])

	return(
		<>
		<div className='PostCard basis-1/4 m-1  rounded box-border h-65 w-90 p-2 border-4'>
			<div className='Title align-middle text-center underline '>
				{props.title}
			</div>
			<div className='text-sm '>
				Comments:
			</div>
			<div className='text-sm h-20 overflow-y-auto'>
					<Comments  comments= {comments}/>
			</div>
			<div className="flex flex-col items-center">
				<input className="bg-white focus:outline-none focus:shadow-outline border border-gray-5 rounded py-1 px-4 block w-54" type="text" placeholder="Enter Your Comment" onChange={(e)=>{setNewComment(e.target.value)}} value={newComment}/>
				<button className="bg-blue-500 text-white py-1 px-2 rounded mt-1 hover:bg-blue-600" onClick={CreateComment}>Submit Comment</button>
			</div>


		</div>
		</>
	)
}

export default PostCard
import React from "react";
import ReactDOM from "react-dom";

export const Message = ({ onEditSaveClick, onDeleteClick, message }) => {
	const [isEditing, setIsEditing] = React.useState(false);
	const [editableMessage, setEditableMessage] = React.useState(message.message);

	const handleEditSaveClick = (id, newMessage) => {
		onEditSaveClick(id, newMessage);
		setIsEditing(false);
	};

	return (
		<div style={{ margin: "25px" }}>
			<p>ID: {message.id}</p>
			{isEditing && (<input type="text" value={editableMessage} onChange={e => setEditableMessage(e.target.value)} /> )}
			{isEditing && (<button onClick={() => handleEditSaveClick(message.id, editableMessage)}>Save</button>)}
			{!isEditing && (<strong style={{ display: "block", marginBottom: "10px" }}>{message.message}</strong>)}
			{!isEditing && (<button onClick={() => setIsEditing(!isEditing)}>Edit</button>)}
			<button onClick={() => onDeleteClick(message.id)}>Delete</button>
		</div>
	);
};
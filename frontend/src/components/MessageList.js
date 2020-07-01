import React from "react";
import ReactDOM from "react-dom";

import { Message } from './Message';

export const MessageList = ({ fetchMessages, messages }) => {
	if (!messages.length) {
		return <div>No messages were found. Create one below!</div>;
	}
	const handleEditSaveClick = async (id, message) => {
		await fetch(`api/v1/messages/${id}`, {
			method: 'PUT',
			body: JSON.stringify({ message })
		});

		fetchMessages();
	};

	const handleDeleteClick = async (id) => {
		await fetch(`api/v1/messages/${id}`, {
			method: 'DELETE',
		});
		
		fetchMessages();
	};

	return (
		<>
			{messages.map(message => (
				<Message
					key={message.id}
					onEditSaveClick={handleEditSaveClick}
					onDeleteClick={handleDeleteClick}
					message={message}
				/>
			))}
		</>
	);
};
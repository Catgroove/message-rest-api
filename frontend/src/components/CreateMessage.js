import React from "react";
import ReactDOM from "react-dom";

export const CreateMessage = ({ fetchMessages }) => {
	const [message, setMessage] = React.useState("");

	const handleCreateClick = async () => {
		if (!message) return;

		await fetch('api/v1/messages', {
			method: 'POST',
			body: JSON.stringify({ message })
		});

		setMessage("");
		fetchMessages();
	};

	return (
		<>
			<input type="text" value={message} onChange={e => setMessage(e.target.value)} placeholder="Message" />
			<button onClick={handleCreateClick}>Create</button>
		</>
	);
};

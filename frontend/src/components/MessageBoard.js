import React from "react";
import ReactDOM from "react-dom";

import { MessageList } from './MessageList';
import { CreateMessage } from './CreateMessage';

export const MessageBoard = () => {
	const [messages, setMessages] = React.useState([]);

	const fetchMessages = async () => {
		const response = await fetch('api/v1/messages').then(response => response.json());
		setMessages(response || []);
	};

	React.useEffect(() => {
		fetchMessages();
	}, []);

	return (
		<>
			<MessageList fetchMessages={fetchMessages} messages={messages} />
			<CreateMessage fetchMessages={fetchMessages} />
		</>
	);
};

import React from "react";
import ReactDOM from "react-dom";

import { MessageBoard } from './components/MessageBoard'

const App = () => {
	return (
		<MessageBoard />
	);
};

const node = document.getElementById("app");
ReactDOM.render(<App />, node);

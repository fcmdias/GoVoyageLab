import React, { useState, useEffect } from "react";

function Chat() {
    const [ws, setWs] = useState(null);
    const [messages, setMessages] = useState([]);

    useEffect(() => {
        const websocket = new WebSocket("ws://localhost:8080/ws");
        websocket.onmessage = (event) => {
            setMessages((prevMessages) => [...prevMessages, event.data]);
        };
        setWs(websocket);
        return () => websocket.close();
    }, []);

    const sendMessage = (message) => {
        ws && ws.send(message);
    };

    return (
        <div>
            <div>
                {messages.map((msg, index) => (
                    <p key={index}>{msg}</p>
                ))}
            </div>
            <button onClick={() => sendMessage("hello" + new Date().toString())}>
                Send Message
            </button>
        </div>
    );
}

export default Chat;

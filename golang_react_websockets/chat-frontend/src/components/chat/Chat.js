import React, { useState, useEffect } from "react";

function Chat() {
    const [ws, setWs] = useState(null);
    const [messages, setMessages] = useState([]);
    const [usersList, setUsersList] = useState([]);


    const group = "teachers";
    const username = "sky" + Math.random().toString(36).substring(2, 7);

    
    useEffect(() => {
        const websocket = new WebSocket(`ws://localhost:8080/ws?group=${group}&username=${username}`);

        websocket.onmessage = (event) => {
            const message = JSON.parse(event.data);

            switch (message.type) {
                case "message":
                    console.log("setting message: ", message);
                    setMessages((prevMessages) => [...prevMessages, message]);
                    console.log("messages: ", messages);
                    break;
                case "users_list":
                    setUsersList(message.users);
                    break;
                // ... handle other message types if needed
            }
        };

        setWs(websocket);
        return () => websocket.close();
    }, []);

    const sendMessage = (message) => {
        console.log("sending message: ", message);
        ws && ws.send(JSON.stringify({
            type: "message",
            content: message,
        }));
    };
return (
    <div>
        <h1>Chat</h1>
        <span>Username: {username}</span>   
        <h3>Online Users:</h3>
        <ul>
            {usersList.map((user) => (
                <li key={user}>{user}</li>
            ))}
        </ul>
        <div>
            {messages.map((msg, index) => (
                <p key={index}>{msg.content}</p>
            ))}
        </div>
        <button onClick={() => sendMessage("hello " + new Date().toString())}>
            Send Message
        </button>
    </div>
);
            }

export default Chat;

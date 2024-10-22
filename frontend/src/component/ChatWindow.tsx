import React, { useState } from "react";

interface ChatWindowProps {
  currentGroup: string;
}

interface Message {
  text: string;
  group: string;
}

const ChatWindow = ({ currentGroup }: Readonly<ChatWindowProps>) => {
  const [messages, setMessages] = useState<Message[]>([]);
  const [newMessage, setNewMessage] = useState<string>("");

  const handleSendMessage = () => {
    if (newMessage.trim() !== "") {
      setMessages([...messages, { text: newMessage, group: currentGroup }]);
      setNewMessage("");
    }
  };

  const handleInputChange = (e: InputChangeEvent) => {
    setNewMessage(e.target.value);
  };

  return (
    <div style={{ flex: 1, padding: "20px" }}>
      <h2>Chat ({currentGroup})</h2>
      <div style={{ height: "calc(100vh - 200px)", overflowY: "auto", border: "1px solid #ccc", padding: "10px" }}>
        {messages
          .filter((message) => message.group === currentGroup)
          .map((message, index) => (
            <div key={index} style={{ marginBottom: "10px" }}>
              {message.text}
            </div>
          ))}
      </div>
      <div style={{ marginTop: "20px" }}>
        <input type="text" value={newMessage} onChange={handleInputChange} style={{ width: "80%", padding: "10px" }} />
        <button onClick={handleSendMessage} style={{ padding: "10px", marginLeft: "10px" }}>
          Send
        </button>
      </div>
    </div>
  );
};

export default ChatWindow;

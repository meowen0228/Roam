import { useEffect, useState } from "react";
import useWebSocket from "react-use-websocket";

interface ChatWindowProps {
  currentGroup: string;
}

interface Message {
  id: string;
  content: string;
}

const ChatWindow = ({ currentGroup }: Readonly<ChatWindowProps>) => {
  const id = localStorage.getItem("id") || "";
  const { sendMessage, lastJsonMessage } = useWebSocket<Message>("ws://localhost:8080/ws", {
    queryParams: { id }
  });
  const [messages, setMessages] = useState<Message[]>([]);
  const [newMessage, setNewMessage] = useState<string>("");

  const handleSendMessage = () => {
    if (newMessage.trim() !== "") {
      setNewMessage("");
      sendMessage(newMessage);
    }
  };

  const handleInputChange = (e: InputChangeEvent) => {
    setNewMessage(e.target.value);
  };

  useEffect(() => {
    if (lastJsonMessage !== null) {
      setMessages((prev) =>
        prev.concat({
          id: lastJsonMessage.id,
          content: lastJsonMessage.content
        })
      );
    }
  }, [lastJsonMessage]);

  return (
    <div style={{ flex: 1, padding: "20px" }}>
      <h2>({currentGroup})</h2>
      <div style={{ height: "calc(100vh - 200px)", overflowY: "auto", border: "1px solid #ccc", padding: "10px" }}>
        {messages.map((v, index) => {
          if (v.id === id) {
            return (
              <div key={index} style={{ marginBottom: "10px", textAlign: "right" }}>
                <div>
                  <span style={{ fontSize: "12px", color: "#888" }}>You</span>
                </div>
                {v.content}
              </div>
            );
          }
          return (
            <div key={index} style={{ marginBottom: "10px", textAlign: "left" }}>
              <div>
                <span style={{ fontSize: "12px", color: "#888" }}>{v.id}</span>
              </div>
              {v.content}
            </div>
          );
        })}
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

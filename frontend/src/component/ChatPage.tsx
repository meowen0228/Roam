import { useState } from "react";
import ChatWindow from "./ChatWindow";
import GroupList from "./GroupList";

const ChatPage = () => {
  const [currentGroup, setCurrentGroup] = useState<string>("general");
  return (
    <div style={{ width: "100%", display: "flex" }}>
      <GroupList
        groups={[]}
        onGroupChange={(group) => {
          setCurrentGroup(group);
        }}
      />
      <ChatWindow currentGroup={currentGroup} />
    </div>
  );
};

export default ChatPage;

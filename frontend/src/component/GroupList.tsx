interface GroupListProps {
  groups: string[];
  onGroupChange: (group: string) => void;
}

const GroupList = ({ groups, onGroupChange }: Readonly<GroupListProps>) => {
  return (
    <div style={{ minWidth: 100, maxWidth: "250px", borderRight: "1px solid #ccc" }}>
      <h2>Groups</h2>
      <ul style={{ listStyleType: "none", padding: 0 }}>
        {groups.map((group) => (
          <li key={group} onClick={() => onGroupChange(group)} style={{ cursor: "pointer", padding: "10px 20px" }}>
            {group}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default GroupList;

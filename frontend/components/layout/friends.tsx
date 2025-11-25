function Friend(){
  const friends = ["Alice", "Bob", "Charlie", "David", "Eve"];

  return(
    <>
      {friends.map((name) => (
        <li key={name} className="w-full h-20 bg-white border-b flex items-center px-4 hover:bg-gray-200 cursor-pointer font-semibold">
          <span className="rounded-full bg-red-300 w-10 h-10"/>
          <div className="p-4">{name}</div>
        </li>
      ))}
    </>
  )
}

export default function FriendList() {
  return (
    <div className="w-80 h-full bg-gray-100 shadow-md">
      <ul>
        <Friend />
      </ul>
    </div>
  );
}
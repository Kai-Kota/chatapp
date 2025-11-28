import Room from "../ui/room";

function Friend(){
  const friends = ["Alice", "Bob"];

  return(
    <>
      {friends.map((name, idx) => (
        <Room 
          key={idx}
          name={name}
        />
      ))}
    </>
  )
}

export default function FriendList() {
  return (
    <div className="w-80 h-full bg-gray-100 flex flex-col shadow-md">
      <ul className="flex-auto">
        <Friend />
      </ul>

      <div className="w-full h-30 bg-cyan-800 p-3">
        <p className="text-white">フレンド追加</p>
        <div className="flex justify-between">
          <input 
            type="text"
            className="mt-2 p-2 rounded-md border bg-white border-gray-300 text-sm"
          />
          <input 
            type="submit"
            value="追加"
            className=" mt-2 p-2 rounded-md bg-blue-500 text-white text-sm hover:bg-blue-600 cursor-pointer"
           />
         </div>
      </div>
    </div>
  );
}
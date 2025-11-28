import Message from "../ui/message";

function Chats(){
  let messages = [
    { userid: 1, text: "こんにちは！" },
    { userid: 2, text: "元気ですか？" },
    { userid: 1, text: "はい、元気です！ありがとう。" },
  ]

  return (
    <div className="bg-blue-100 flex-1 h-130 p-4 space-y-3">
      {messages.map((msg, index) => (
        <Message
          key={msg.userid + "_" + index}
          userid={msg.userid}
          text={msg.text}  
        />
        ))}
    </div>
  )
}

function MessageInputField(){
    return (
      <div className="p-4 bg-white flex items-end gap-2">
        <textarea
          className="flex-1 min-h-[100px] max-h-40 resize-none rounded-md px-3 py-2 text-sm shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-300"
          placeholder="メッセージを入力..."
          aria-label="メッセージ入力"
        />
        <button className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 focus:outline-none">
          送信
        </button>
      </div>
    )
}

export default function ChatField() {
  return (
    <>
      <div className="flex flex-col w-full h-full">
        <Chats />
        <MessageInputField />
      </div>
    </>
  );
}
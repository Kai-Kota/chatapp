function Chats(){
  return (
    <div className="bg-blue-100 flex-1 h-130 p-4 space-y-3">
      {/* 受信メッセージ */}
      <div className="flex items-start gap-3">
        <div className="w-8 h-8 rounded-full bg-gray-300 flex-shrink-0" />
        <div className="bg-white px-3 py-2 rounded-lg shadow-sm max-w-[70%] text-sm">
          元気？
        </div>
      </div>
      {/* 送信メッセージ（右寄せ） */}
      <div className="flex items-start gap-3 justify-end">
        <div className="bg-blue-500 text-white px-3 py-2 rounded-lg shadow-sm max-w-[70%] text-sm">
          元気だよ
        </div>
        <div className="w-8 h-8 rounded-full bg-blue-400 flex-shrink-0" />
      </div>
      {/* 受信メッセージ*/}
      <div className="flex items-start gap-3">
        <div className="w-8 h-8 rounded-full bg-gray-300 flex-shrink-0" />
        <div className="bg-white px-3 py-2 rounded-lg shadow-sm max-w-[70%] text-sm">
          そうでっか
        </div>
      </div>
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
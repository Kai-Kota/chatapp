import ChatField from "@/components/layout/chats";
import Friends from "@/components/layout/friends";
import Header from "@/components/layout/header";

export default function Home() {
  return (
    <div className="w-full h-160">
      <Header/>
      <div className="flex h-full">
        <Friends/>
        <ChatField/>
      </div>
    </div>
  );
}
  
import Friends from "@/components/layout/friends";
import Header from "@/components/layout/header";
import Image from "next/image";

export default function Home() {
  return (
    <div className="w-full h-screen bg-blue-100">
      <Header/>
      <Friends/>
    </div>
  );
}

"use client";

import Broadcast from "@/components/Broadcast";
import Blog from "@/components/Blog";

export default function Home() {
  return (
    <div className="min-h-screen p-20">
      <h1 className="text-5xl font-bold">Broadcast</h1>
      <div className="grid place-content-center p-10">
        <div className="flex flex-wrap gap-10">
          <div className="flex-2">
          <Blog />
          </div>
          <div className="flex-1">
            <Broadcast />
          </div>
        </div>
      </div>
    </div>
  );
}

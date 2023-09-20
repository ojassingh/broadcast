"use client";

import { useRouter } from "next/navigation";
import del from "@/functions/postHandlers/delete";

export default function Post(props: any) {
  const router = useRouter();
  const blog = props.blog;

  console.log(blog);

  async function handleClick(id: any) {
    try {
      const deleted = await (await del(id))?.data;
      router.refresh();
    } catch (error) {
      console.log(error);
    }
  }

  return (
    <div className="relative bg-blue-500/20 p-6 rounded-2xl">
      <button
        onClick={() => {
          handleClick(blog.id);
        }}
        className="text-sm right-0 absolute top-2 right-2"
      >
        🗑️
      </button>
      <h1 className="font-medium text-xl">{blog.title}</h1>
      <p className="pb-4">{blog.content}</p>
      <p className="text-sm right-0 absolute bottom-2 right-2">{blog.date}</p>
    </div>
  );
}

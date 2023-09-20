"use client";
import del from "@/functions/postHandlers/delete";
import { useRouter } from "next/navigation";
import { readAll } from "@/functions/postHandlers/read";
import { useState, useEffect} from "react";
import { Button } from "../components/ui/button";
import Post from "./Post";
import { Dialogue } from "./Dialogue";

export default function Blog() {

  const [blogs, setBlogs] = useState<any>([]);

  async function getData(){
    try{
      const allBlogs = (await readAll()).data;
      setBlogs(allBlogs)
      console.log(allBlogs)
    }
    catch(error){
      console.log(error)
    }
  }

  useEffect(()=>{
    getData();
  }, [])

  async function handleClick(id: any){
    
  }

  const router = useRouter();

  return (
    <div className="">
      <h1 className="text-xl font-bold">Welcome to your braodcasts!</h1>
      {blogs && <div className="flex flex-wrap gap-10">
        {blogs.map((blog: any, i:number) => {
          return <Post blog={blog} key={i}/>;
        })}
      </div>}
      {!blogs && <div className="flex flex-wrap gap-10">
        <h1>You don't have any blogs yet! Post one!</h1>
      </div>}
      <Dialogue/>
    </div>
  );
}

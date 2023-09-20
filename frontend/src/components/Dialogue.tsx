import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "./ui/textarea";
import { useState } from "react";
import create from "@/functions/postHandlers/create";
import { useRouter } from "next/navigation";

export function Dialogue() {

  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [loading, setLoading] = useState(false)
  const router = useRouter()
  
  async function handleClick(id: any){

    

    try{
      setLoading(true);
      const data: any = {
        Title: title,
        Content: content,
        Date: (new Date()).toDateString()
      }
      // console.log("Initiating...")
      const creation = await create(data)
      setLoading(false)
    }
    catch(error){
      console.log(error)
      alert("failed")
    }
  }

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button className="rounded-3xl" variant="outline">New Post</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Make a new post!</DialogTitle>
          <DialogDescription>
           Create a new post by submitting info here!
          </DialogDescription>
        </DialogHeader>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="name" className="text-right">
              Title
            </Label>
            <Input value={title} onChange={(e)=>{setTitle(e.target.value)}} placeholder="My day.." id="name" className="col-span-3 rounded-xl" />
          </div>
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="content" className="text-right">
              Content
            </Label>
            <Textarea value={content} onChange={(e)=>{setContent(e.target.value)}} id="content" placeholder="input your content here..." className="p-2 rounded-xl col-span-3 bg-black outline outline-1" />
          </div>
        </div>
        <DialogFooter className="">
          <Button onClick={handleClick} type="submit">{!loading ? "Submit" : "Loading..."}</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

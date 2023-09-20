'use client'

import {connect, sendMsg} from '../functions/index'
import { useEffect, useState} from 'react'
import ChatHistory from '@/components/ChatHistory'

export default function Broadcast() {

    const [state, setState] = useState<any>([])
    const [value, setValue] = useState<string>("");
  
  
    useEffect(()=>{
      connect((msg: any)=>{
        setState((prevChatHistory: any)=> [...prevChatHistory, msg]);
      });
    }, [])
  
    function send(){
      sendMsg(value)
    }
  
    return (
      <div className=''>
        <h1 className='text-xl font-bold'>Social</h1>
        <ChatHistory chatHistory={state} />
  
        <input onKeyDown={(e)=>{
          if(e.key === 'Enter'){
            send()
            setValue('')
          }
        }}  onChange={(e)=>{
          setValue(e.target.value)
        }} value={value} className='rounded-3xl p-2 text-black' placeholder='write your message here'/>
        <button onClick={()=>{
          send()
          setValue('')
        }} className='rounded-full outline bg-blue-500 text-lg font-medium p-2 m-4'>Send!</button>
      </div>
    )
  }
  
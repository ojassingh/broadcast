'use client'
import Image from 'next/image'
import {connect, sendMsg} from '../functions/index'
import { useEffect } from 'react'

export default function Home() {


  useEffect(()=>{
    connect();
  }, [])

  function send(){
    sendMsg("test")
  }

  return (
    <div className='grid place-content-center min-h-screen'>
      <h1 className='text-4xl font-bold'>Hi there!</h1>
      <button onClick={send} className='rounded-full outline bg-blue-500 text-lg font-medium p-2 m-4'>Click me</button>
    </div>
  )
}

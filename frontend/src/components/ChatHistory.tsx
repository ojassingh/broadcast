'use client'

export default function ChatHistory(props: any){

    // const messages = props.chatHistory.map((msg: any, index: any) => {
    //     <p key={index}>{msg.data}</p>
    // })

    const messageHistory = props.chatHistory;
    const messages = messageHistory.map((message: any, i: any)=>{
        return <p key={i}>{message.data}</p>
    })

    return(<div className="ChatHistory">
    <h2>Chat History</h2>
        {messages}
  </div>)
}
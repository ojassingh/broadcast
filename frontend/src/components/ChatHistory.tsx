"use client";

export default function ChatHistory(props: any) {

  const messageHistory = props.chatHistory;
  const messages = messageHistory.map((message: any, i: any) => {

    const text = JSON.parse(message.data).Body;
    const user = JSON.parse(message.data).Type;
    // console.log(text)

    return (
      <p className="p-2 rounded-xl bg-gray-900" key={i}>
        <span className="font-medium">User {user}</span>: {text}
      </p>
    );
  });

  return (
    <div className="ChatHistory">
      <h2>Chat History</h2>
      <div className="grid gap-2">{messages}</div>
    </div>
  );
}

# Broadcast - Social Media App

Broadcast is a social media application that allows users to post and broadcast messages to their social circle in real-time. This repository contains the source code for the Broadcast app, which is built using various technologies including Go, Gin, Websockets, MongoDB, and Next.js with Tailwind CSS for the frontend.

## Features

- Create and post messages similar to "tweets."
- Real-time broadcasting of messages to your social circle.
- User authentication and authorization.
- MongoDB database for storing user data and messages.
- Modern and responsive user interface designed with Tailwind CSS.
- Integration with Next.js for frontend rendering.

## Technologies Used

- **Backend:**
  - Go - Programming language used for the backend server.
  - Gin - Web framework for building RESTful APIs.
  - Websockets - Real-time communication for broadcasting messages.
  - MongoDB - NoSQL database for storing user data and messages.

- **Frontend:**
  - Next.js - React framework for building the frontend.
  - Tailwind CSS - Utility-first CSS framework for designing the user interface.

## Getting Started

To run the Broadcast app locally, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/broadcast.git
   cd broadcast
   ```
2. Configure the environment variables. Create a .env file in the root directory and set the following variables:

```bash
# Backend Configuration
DB_CONNECTION_STRING=mongodb://localhost:27017/broadcast
JWT_SECRET=your-secret-key

# Frontend Configuration
NEXT_PUBLIC_API_URL=http://localhost:8080/api
```
3. Install dependencies for both the backend and frontend:

```bash
cd backend
go mod tidy

# Install frontend dependencies
cd ../frontend
npm install
```

4. Start the backend server

```bash
go run main.go
```

5. Start the frontend development server:
```bash
cd frontend
npm run dev
```

6. Access the app in your web browser at http://localhost:3000.





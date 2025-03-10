# Backend part of Salyte app 

Backend consists of 3 services and an API gateway.

## Installation

Prequisites:

- docker
- docker-compose
- git

Also it is necessary to open port 3000 for main API access

Clone repoisotry in a new folder

```
git clone http://github.com/p4elkab35t/salyte_backend.git
```

Then you should switch to the directoy and initialize docker-compose

```
cd salyte_backend
```
```
docker-compose up --build [-d for running in th background]
```

Now you can access to services (depends on open ports) and main API gateway on port 3000.

## API Reference

Below is an API for each of the service (only API for main API gateway)

### Auth service

#### /lifecheck
- **Method**: GET
- **Description**: Check if the Auth service is alive.
- **Response**: `{ message: "Auth Service Alive" }`

#### /verify
- **Method**: GET
- **Description**: Verify a token.
- **Headers**: `Authorization: Bearer <token>`
- **Query Params**: `user_id=<user_id>`
- **Response**: `{ status: 0, ... }` or `{ error: "Verification failed" }`

#### /signin
- **Method**: GET, POST
- **Description**: Sign in a user.
- **Headers**: `Authorization: Bearer <token>` (for GET)
- **Body**: `{ email: <email>, password: <password> }` (for POST)
- **Response**: `{ status: 0, ... }` or `{ error: "Sign in failed" }`

#### /signup
- **Method**: POST
- **Description**: Register a new user.
- **Body**: `{ email: <email>, password: <password> }`
- **Response**: `{ status: 0, ... }` or `{ error: "Sign up failed" }`

#### /signout
- **Method**: GET
- **Description**: Sign out a user.
- **Headers**: `Authorization: Bearer <token>`
- **Response**: `{ status: 0, ... }` or `{ error: "Sign out failed" }`

### Social service

#### /profile
- **Methods**: GET, PUT
- **Description**: Manage user profiles.

#### /follow
- **Methods**: POST, DELETE
- **Description**: Follow or unfollow users.

#### /friends
- **Methods**: GET, POST
- **Description**: Manage friends list.

#### /community
- **Methods**: GET, POST, PUT
- **Description**: Manage communities.

#### /post
- **Methods**: GET, POST, PUT, DELETE
- **Description**: Manage posts.

#### /post/comment
- **Methods**: POST, PUT, DELETE
- **Description**: Manage comments on posts.

#### /post/likes
- **Methods**: GET
- **Description**: Get likes on posts.

#### /post/like
- **Methods**: POST, DELETE
- **Description**: Like or unlike posts.

### Message service

#### /getMessagesByChatID
- **Method**: GET
- **Description**: Get messages by chat ID.
- **Query Params**: `chat_id=<chat_id>&user_id=<user_id>`
- **Response**: `{ status: 0, messages: [...] }` or `{ error: "Failed to get messages" }`

#### /getUnreadMessages
- **Method**: GET
- **Description**: Get unread messages.
- **Query Params**: `user_id=<user_id>`
- **Response**: `{ status: 0, messages: [...] }` or `{ error: "Failed to get unread messages" }`

#### /deleteAllMessagesByChatID
- **Method**: POST
- **Description**: Delete all messages by chat ID.
- **Query Params**: `chat_id=<chat_id>&user_id=<user_id>`
- **Body**: `{ ... }`
- **Response**: `{ status: 0 }` or `{ error: "Failed to delete messages" }`

#### /getChat
- **Method**: GET
- **Description**: Get chat details.
- **Query Params**: `chat_id=<chat_id>`
- **Response**: `{ status: 0, chat: { ... } }` or `{ error: "Failed to get chat" }`

#### /createChat
- **Method**: POST
- **Description**: Create a new chat.
- **Body**: `{ ... }`
- **Response**: `{ status: 0, chat: { ... } }` or `{ error: "Failed to create chat" }`

#### /getAllChats
- **Method**: GET
- **Description**: Get all chats for a user.
- **Query Params**: `user_id=<user_id>`
- **Response**: `{ status: 0, chats: [...] }` or `{ error: "Failed to get chats" }`

#### /addUserToChat
- **Method**: POST
- **Description**: Add a user to a chat.
- **Query Params**: `chat_id=<chat_id>&user_id=<user_id>&added_user_id=<added_user_id>`
- **Body**: `{ ... }`
- **Response**: `{ status: 0 }` or `{ error: "Failed to add user to chat" }`

#### /removeUserFromChat
- **Method**: POST
- **Description**: Remove a user from a chat.
- **Query Params**: `chat_id=<chat_id>&user_id=<user_id>&removed_user_id=<removed_user_id>`
- **Body**: `{ ... }`
- **Response**: `{ status: 0 }` or `{ error: "Failed to remove user from chat" }`

#### /getChatMembers
- **Method**: GET
- **Description**: Get members of a chat.
- **Query Params**: `chat_id=<chat_id>&user_id=<user_id>`
- **Response**: `{ status: 0, members: [...] }` or `{ error: "Failed to get chat members" }`

#### /getChatByID
- **Method**: GET
- **Description**: Get chat by ID.
- **Query Params**: `chat_id=<chat_id>&user_id=<user_id>`
- **Response**: `{ status: 0, chat: { ... } }` or `{ error: "Failed to get chat" }`

#### /getReactions
- **Method**: GET
- **Description**: Get reactions for a message.
- **Query Params**: `message_id=<message_id>&user_id=<user_id>`
- **Response**: `{ status: 0, reactions: [...] }` or `{ error: "Failed to get reactions" }`

### Message realtime reference

#### WebSocket Connection
- **URL**: `ws://localhost:3000/api/message`
- **Description**: Establish a WebSocket connection for real-time messaging.

#### Incoming Requests

##### open
- **Description**: Open a WebSocket connection.
- **Data**: `{ chatID: <chat_id>, token: <token> }`
- **Response**: `{ type: "connected", message: "Chat connected" }` or `{ type: "error", message: "Unauthorized" }`

##### message
- **Description**: Handle incoming messages.
- **Data**: 
  - **get**: `{ type: "get", message_id: <message_id> }`
  - **send**: `{ type: "send", chat_id: <chat_id>, message: <message> }`
  - **edit**: `{ type: "edit", message_id: <message_id>, message: <new_message> }`
  - **delete**: `{ type: "delete", message_id: <message_id> }`
  - **reactionApply**: `{ type: "reactionApply", message_id: <message_id>, reaction: <reaction> }`
  - **reactionRemove**: `{ type: "reactionRemove", message_id: <message_id>, reaction: <reaction> }`
  - **read**: `{ type: "read", message_id: <message_id> }`
- **Response**: 
  - **get**: `{ type: "success", message: { ... } }` or `{ type: "error", message: "Failed to get message" }`
  - **send**: `{ type: "success", message: "Message sent" }` or `{ type: "error", message: "Failed to send message" }`
  - **edit**: `{ type: "success", message: "Message edited" }` or `{ type: "error", message: "Failed to edit message" }`
  - **delete**: `{ type: "success", message: "Message deleted" }` or `{ type: "error", message: "Failed to delete message" }`
  - **reactionApply**: `{ type: "success", message: "Reaction applied" }` or `{ type: "error", message: "Failed to apply reaction" }`
  - **reactionRemove**: `{ type: "success", message: "Reaction removed" }` or `{ type: "error", message: "Failed to remove reaction" }`
  - **read**: `{ type: "success", message: "Message marked as read" }` or `{ type: "error", message: "Failed to mark message as read" }`

##### close
- **Description**: Close a WebSocket connection.
- **Response**: `{ type: "disconnected", message: "Chat disconnected" }`

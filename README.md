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
- **Response**: 
  ```json
  {
    "status": 0,
    "user_id": "string",
    "email": "string",
    "roles": ["string"]
  }
  ```
  or `{ error: "Verification failed" }`

#### /signin
- **Method**: GET, POST
- **Description**: Sign in a user.
- **Headers**: `Authorization: Bearer <token>` (for GET)
- **Body**: `{ email: <email>, password: <password> }` (for POST)
- **Response**: 
  ```json
  {
    "status": 0,
    "user_id": "string",
    "email": "string",
    "token": "string"
  }
  ```
  or `{ error: "Sign in failed" }`

#### /signup
- **Method**: POST
- **Description**: Register a new user.
- **Body**: `{ email: <email>, password: <password> }`
- **Response**: 
  ```json
  {
    "status": 0,
    "user_id": "string",
    "email": "string"
  }
  ```
  or `{ error: "Sign up failed" }`

#### /signout
- **Method**: GET
- **Description**: Sign out a user.
- **Headers**: `Authorization: Bearer <token>`
- **Response**: 
  ```json
  {
    "status": 0,
    "message": "Signed out successfully"
  }
  ```
  or `{ error: "Sign out failed" }`

### Social service

#### /profile
- **Methods**: GET, PUT
- **Description**: Manage user profiles.
- **GET Response**: 
  ```json
  {
    "profile_id": "string",
    "user_id": "string",
    "username": "string",
    "bio": "string",
    "profile_picture_url": "string",
    "visibility": "string",
    "created_at": "string",
    "updated_at": "string"
  }
  ```
  or `{ error: "Failed to get profile" }`
- **PUT Body**: 
  ```json
  {
    "username": "string",
    "bio": "string",
    "profile_picture_url": "string",
    "visibility": "string"
  }
  ```
- **PUT Response**: `{ message: "profile updated" }` or `{ error: "Failed to update profile" }`

#### /follow
- **Methods**: POST, DELETE
- **Description**: Follow or unfollow users.
- **POST Query Params**: `profileID=<profile_id>`
- **POST Response**: `{ message: "profile followed" }` or `{ error: "Failed to follow profile" }`
- **DELETE Query Params**: `profileID=<profile_id>`
- **DELETE Response**: `{ message: "profile unfollowed" }` or `{ error: "Failed to unfollow profile" }`

#### /friends
- **Methods**: GET, POST
- **Description**: Manage friends list.
- **GET Response**: 
  ```json
  [
    {
      "profile_id": "string",
      "user_id": "string",
      "username": "string",
      "bio": "string",
      "profile_picture_url": "string",
      "visibility": "string",
      "created_at": "string",
      "updated_at": "string"
    }
  ]
  ```
  or `{ error: "Failed to get friends" }`
- **POST Query Params**: `profileID=<profile_id>`
- **POST Response**: `{ message: "profile followed" }` or `{ error: "Failed to follow profile" }`

#### /community
- **Methods**: GET, POST, PUT
- **Description**: Manage communities.
- **GET Query Params**: `communityID=<community_id>`
- **GET Response**: 
  ```json
  {
    "community_id": "string",
    "name": "string",
    "description": "string",
    "profile_picture_url": "string",
    "visibility": "string",
    "created_at": "string",
    "updated_at": "string"
  }
  ```
  or `{ error: "Failed to get community" }`
- **POST Body**: 
  ```json
  {
    "name": "string",
    "description": "string",
    "profile_picture_url": "string",
    "visibility": "string"
  }
  ```
- **POST Response**: 
  ```json
  {
    "community_id": "string",
    "created_at": "string",
    "updated_at": "string"
  }
  ```
  or `{ error: "Failed to create community" }`
- **PUT Query Params**: `communityID=<community_id>`
- **PUT Body**: 
  ```json
  {
    "name": "string",
    "description": "string",
    "profile_picture_url": "string",
    "visibility": "string"
  }
  ```
- **PUT Response**: `{ message: "community updated" }` or `{ error: "Failed to update community" }`

#### /post
- **Methods**: GET, POST, PUT, DELETE
- **Description**: Manage posts.
- **GET Query Params**: `postID=<post_id>`
- **GET Response**: 
  ```json
  {
    "post_id": "string",
    "profile_id": "string",
    "community_id": "string",
    "content": "string",
    "media_url": "string",
    "visibility": "string",
    "created_at": "string",
    "updated_at": "string"
  }
  ```
  or `{ error: "Failed to get post" }`
- **POST Body**: 
  ```json
  {
    "profile_id": "string",
    "community_id": "string",
    "content": "string",
    "media_url": "string",
    "visibility": "string"
  }
  ```
- **POST Response**: 
  ```json
  {
    "post_id": "string",
    "created_at": "string",
    "updated_at": "string"
  }
  ```
  or `{ error: "Failed to create post" }`
- **PUT Query Params**: `postID=<post_id>`
- **PUT Body**: 
  ```json
  {
    "content": "string",
    "media_url": "string",
    "visibility": "string"
  }
  ```
- **PUT Response**: `{ message: "post updated" }` or `{ error: "Failed to update post" }`
- **DELETE Query Params**: `postID=<post_id>`
- **DELETE Response**: `{ message: "post deleted" }` or `{ error: "Failed to delete post" }`

#### /post/comment
- **Methods**: POST, PUT, DELETE
- **Description**: Manage comments on posts.
- **POST Body**: 
  ```json
  {
    "profile_id": "string",
    "post_id": "string",
    "content": "string"
  }
  ```
- **POST Response**: 
  ```json
  {
    "comment_id": "string",
    "created_at": "string",
    "updated_at": "string"
  }
  ```
  or `{ error: "Failed to create comment" }`
- **PUT Query Params**: `commentID=<comment_id>`
- **PUT Body**: 
  ```json
  {
    "content": "string"
  }
  ```
- **PUT Response**: `{ message: "comment updated" }` or `{ error: "Failed to update comment" }`
- **DELETE Query Params**: `commentID=<comment_id>`
- **DELETE Response**: `{ message: "comment deleted" }` or `{ error: "Failed to delete comment" }`

#### /post/likes
- **Methods**: GET
- **Description**: Get likes on posts.
- **GET Query Params**: `postID=<post_id>`
- **GET Response**: 
  ```json
  [
    {
      "profile_id": "string",
      "user_id": "string",
      "username": "string",
      "bio": "string",
      "profile_picture_url": "string",
      "visibility": "string",
      "created_at": "string",
      "updated_at": "string"
    }
  ]
  ```
  or `{ error: "Failed to get likes" }`

#### /post/like
- **Methods**: POST, DELETE
- **Description**: Like or unlike posts.
- **POST Query Params**: `postID=<post_id>`
- **POST Response**: `{ message: "post liked" }` or `{ error: "Failed to like post" }`
- **DELETE Query Params**: `postID=<post_id>`
- **DELETE Response**: `{ message: "post unliked" }` or `{ error: "Failed to unlike post" }`

### Message service

#### /getMessagesByChatID
- **Method**: GET
- **Description**: Get messages by chat ID.
- **Query Params**: `chat_id=<chat_id>&user_id=<user_id>`
- **Response**: 
  ```json
  {
    "status": 0,
    "messages": [
      {
        "message_id": "string",
        "chat_id": "string",
        "sender_id": "string",
        "content": "string",
        "created_at": "string",
        "updated_at": "string"
      }
    ]
  }
  ```
  or `{ error: "Failed to get messages" }`

#### /getUnreadMessages
- **Method**: GET
- **Description**: Get unread messages.
- **Query Params**: `user_id=<user_id>`
- **Response**: 
  ```json
  {
    "status": 0,
    "messages": [
      {
        "message_id": "string",
        "chat_id": "string",
        "sender_id": "string",
        "content": "string",
        "created_at": "string",
        "updated_at": "string"
      }
    ]
  }
  ```
  or `{ error: "Failed to get unread messages" }`

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
- **Response**: 
  ```json
  {
    "status": 0,
    "chat": {
      "chat_id": "string",
      "name": "string",
      "created_at": "string",
      "updated_at": "string"
    }
  }
  ```
  or `{ error: "Failed to get chat" }`

#### /createChat
- **Method**: POST
- **Description**: Create a new chat.
- **Body**: `{ ... }`
- **Response**: 
  ```json
  {
    "status": 0,
    "chat": {
      "chat_id": "string",
      "name": "string",
      "created_at": "string",
      "updated_at": "string"
    }
  }
  ```
  or `{ error: "Failed to create chat" }`

#### /getAllChats
- **Method**: GET
- **Description**: Get all chats for a user.
- **Query Params**: `user_id=<user_id>`
- **Response**: 
  ```json
  {
    "status": 0,
    "chats": [
      {
        "chat_id": "string",
        "name": "string",
        "created_at": "string",
        "updated_at": "string"
      }
    ]
  }
  ```
  or `{ error: "Failed to get chats" }`

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
- **Response**: 
  ```json
  {
    "status": 0,
    "members": [
      {
        "user_id": "string",
        "username": "string",
        "profile_picture_url": "string"
      }
    ]
  }
  ```
  or `{ error: "Failed to get chat members" }`

#### /getChatByID
- **Method**: GET
- **Description**: Get chat by ID.
- **Query Params**: `chat_id=<chat_id>&user_id=<user_id>`
- **Response**: 
  ```json
  {
    "status": 0,
    "chat": {
      "chat_id": "string",
      "name": "string",
      "created_at": "string",
      "updated_at": "string"
    }
  }
  ```
  or `{ error: "Failed to get chat" }`

#### /getReactions
- **Method**: GET
- **Description**: Get reactions for a message.
- **Query Params**: `message_id=<message_id>&user_id=<user_id>`
- **Response**: 
  ```json
  {
    "status": 0,
    "reactions": [
      {
        "user_id": "string",
        "reaction": "string"
      }
    ]
  }
  ```
  or `{ error: "Failed to get reactions" }`

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

#### Reactive Updates Interface

##### newMessage
- **Description**: Broadcasted when a new message is sent in the chat.
- **Data**:
  ```json
  {
    "type": "newMessage",
    "message": {
      "message_id": "string",
      "chat_id": "string",
      "sender_id": "string",
      "content": "string",
      "created_at": "string",
      "updated_at": "string"
    }
  }
  ```
- **Purpose**: Notify all clients in the chat room about the new message.

##### editedMessage
- **Description**: Broadcasted when a message is edited in the chat.
- **Data**:
  ```json
  {
    "type": "editedMessage",
    "message": {
      "message_id": "string",
      "chat_id": "string",
      "sender_id": "string",
      "content": "string",
      "created_at": "string",
      "updated_at": "string"
    }
  }
  ```
- **Purpose**: Notify all clients in the chat room about the edited message.

##### deletedMessage
- **Description**: Broadcasted when a message is deleted in the chat.
- **Data**:
  ```json
  {
    "type": "deletedMessage",
    "message_id": "string"
  }
  ```
- **Purpose**: Notify all clients in the chat room about the deleted message.

##### reactionApplied
- **Description**: Broadcasted when a reaction is applied to a message in the chat.
- **Data**:
  ```json
  {
    "type": "reactionApplied",
    "message_id": "string",
    "reaction": "string"
  }
  ```
- **Purpose**: Notify all clients in the chat room about the applied reaction.

##### reactionRemoved
- **Description**: Broadcasted when a reaction is removed from a message in the chat.
- **Data**:
  ```json
  {
    "type": "reactionRemoved",
    "message_id": "string",
    "reaction": "string"
  }
  ```
- **Purpose**: Notify all clients in the chat room about the removed reaction.

##### messageRead
- **Description**: Broadcasted when a message is marked as read in the chat.
- **Data**:
  ```json
  {
    "type": "messageRead",
    "message_id": "string"
  }
  ```
- **Purpose**: Notify all clients in the chat room about the read message.


##### close
- **Description**: Close a WebSocket connection.
- **Response**: `{ type: "disconnected", message: "Chat disconnected" }`

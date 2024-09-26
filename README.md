# TCPChat
## Overview
TCPChat is a command-line utility that recreates the functionality of NetCat within a server-client architecture. It enables a group chat system where multiple clients can connect to a server, send messages, and receive updates in real-time.

# Features
- TCP Connection: Establishes a server-client relationship, allowing multiple clients to connect.
- Client Naming: Each client must provide a unique name upon connection.
- Connection Control: Limits the number of simultaneous connections.
- Message Transmission: Clients can send messages, which are timestamped and tagged with the sender's name.
- Message History: New clients receive all previous messages upon joining.
- Join/Leave Notifications: Clients are informed when others join or leave the chat.
- Non-empty Message Requirement: Prevents the sending of empty messages.
- Default Port: If no port is specified, defaults to port 8989. Usage is as follows:

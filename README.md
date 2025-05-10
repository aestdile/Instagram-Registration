<div align="center">

# Instagram Model

<img src="https://img.shields.io/badge/💻%20C%23%20.NET-Gray?style=for-the-badge&logo=csharp&logoColor=white" />  
<img src="https://img.shields.io/badge/JSON20%FILESr-Purple?style=for-the-badge&logo=visualstudio&logoColor=white" />

</div>

<hr style="border: 1px solid gray; width: 100%;">

## 📋 Project Overview

A Golang-based Instagram-like application that provides user registration, authentication, post management, and interactive social features like likes and comments. The application tracks engagement metrics in real-time and offers a responsive, user-friendly interface.

## ✨ Features

### 👤 User Management
- **Registration**: Secure user signup with validation
- **Login**: Authentication system with session management
- **Profile Management**: Update user information and settings

### 📝 Content Features
- **Posts Creation**: Share content with descriptions
- **Media Upload**: Support for images
- **Feed Generation**: View posts from followed users

### 💫 Interaction
- **Likes System**: 
  - Like/unlike posts
  - Real-time like counter updates
  - Track user engagement metrics
- **Comments**: 
  - Add comments to posts
  - Comment counter and management

### 📊 Analytics
- View engagement statistics
- Track most popular content
- Monitor user activity

## 🚀 Getting Started

### Prerequisites
- Go 1.16+
- Git

### Installation

```bash
# Clone the repository
git clone https://github.com/aestdile/instagram-registration.git

# Navigate to project directory
cd instagram-registration

# Run the application
go run main.go
```

### Configuration
Create a `config.json` file in the project root (if not present):

```json
{
  "port": "8080",
  "db_path": "./data"
}
```

## 🗄️ Data Structure

The project uses JSON files for data storage:

- `users.json`: Stores user information and credentials
- `posts.json`: Contains post data, including likes and comments

## 📦 Project Structure

```
instagram-registration/
├── main.go              # Application entry point
├── users.json           # User data storage
├── posts.json           # Posts data storage
├── LICENSE.txt          # MIT License
├── .gitignore           # Git ignore rules
├── .gitattributes       # Git attributes
└── README.md            # Project documentation
```

## 🔍 API Endpoints

### User Management
- `POST /api/register` - Create new user account
- `POST /api/login` - Authenticate user
- `GET /api/users/:id` - Get user profile
- `PUT /api/users/:id` - Update user profile

### Posts
- `POST /api/posts` - Create new post
- `GET /api/posts` - Get all posts
- `GET /api/posts/:id` - Get specific post
- `DELETE /api/posts/:id` - Delete post

### Likes
- `POST /api/posts/:id/like` - Like a post
- `DELETE /api/posts/:id/like` - Unlike a post
- `GET /api/posts/:id/likes` - Get post likes count

## 🎯 Future Improvements

- Add followers/following system
- Implement direct messaging
- Add notifications system
- Create mobile application
- Implement caching for better performance

## 📜 License

This project is licensed under the MIT License - see the [LICENSE.txt](LICENSE.txt) file for details.

## ✍️ Author

**👤 Mukhtor Eshboyev🪄**\
🔗 GitHub: [@aestdile](https://github.com/aestdile)\
📌 *"When you finish this project, upload it to GitHub and send me the repository link, I'll wait for it!"*

---

<div align="center">
  <p>Made with ❤️ using Go</p>
</div>

# Connect with me!

## 🌐 Social Networks

<div align="center">
  <a href="https://t.me/aestdile"><img src="https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white" /></a>
  <a href="https://github.com/aestdile"><img src="https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white" /></a>
  <a href="https://leetcode.com/aestdile"><img src="https://img.shields.io/badge/LeetCode-FFA116?style=for-the-badge&logo=leetcode&logoColor=black" /></a>
  <a href="https://linkedin.com/in/aestdile"><img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white" /></a>
  <a href="https://youtube.com/@aestdile"><img src="https://img.shields.io/badge/YouTube-FF0000?style=for-the-badge&logo=youtube&logoColor=white" /></a>
  <a href="https://instagram.com/aestdile"><img src="https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white" /></a>
  <a href="https://facebook.com/aestdile"><img src="https://img.shields.io/badge/Facebook-1877F2?style=for-the-badge&logo=facebook&logoColor=white" /></a>
  <a href="mailto:aestdile@gmail.com"><img src="https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white" /></a>
</div>


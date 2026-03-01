# 📝 Personal Blog Platform

A lightweight, full-featured personal blog built with **Go** and pure **HTML/CSS**. Features JWT authentication, file-based storage, and a beautiful admin dashboard.

![Go Version](https://img.shields.io/badge/Go-1.25.0-blue) ![License](https://img.shields.io/badge/License-MIT-green)

---

## ✨ Features

### 👥 **Guest Features**
- 📖 View all published articles on the home page
- 🔍 **Search functionality** - search articles by title or content
- 📄 Read individual articles with full content and publication date
- 🎨 Responsive, modern design

### 🔐 **Admin Features**
- 📝 Create new articles with title, content, and publication date
- ✏️ Edit existing articles
- 🗑️ Delete articles with confirmation
- 📊 Dashboard with all articles overview
- 👤 User authentication with JWT tokens
- 🔒 Secure session management with HTTP-only cookies

---

## 🚀 Quick Start

### Prerequisites
- Go 1.25.0 or higher
- Git

### Installation

1. **Clone the repository**
```bash
git clone https://github.com/CrabRus/Web-Blog.git
cd web-blog
```

2. **Install dependencies**
```bash
go mod download
```

3. **Create `.env` file**
```bash
cp .env.example .env
```

4. **Configure environment variables**
```env
JWT_SECRET=your-secret-key-here
ADMIN_USERNAME=admin
ADMIN_PASSWORD=your-password
```

5. **Run the server**
```bash
go run main.go
```

The application will start at `http://localhost:8080`

---

## 📚 Project Structure

```
web-blog/
├── handlers/              # HTTP request handlers
│   ├── article_handlers.go    # Article CRUD operations
│   ├── auth.go                # Authentication handlers
│   └── middleware/
│       └── jwt_middleware.go  # JWT validation & auth middleware
├── model/                 # Data structures
│   └── article.go         # Article model
├── utils/                 # Utility functions
│   ├── server_utils.go    # Template parsing, file operations
│   └── validation.go      # Input validation
├── templates/             # HTML templates
│   ├── home.html          # Homepage with article list & search
│   ├── articlepage.html   # Single article view
│   ├── dashboard.html     # Admin dashboard
│   ├── newArticle.html    # Create article form
│   ├── updateArticle.html # Edit article form
│   ├── login.html         # Login page
│   ├── login_error.html   # Login error page
│   └── search_results.html # Search results page
├── articles/              # JSON article storage
├── main.go                # Application entry point
├── go.mod                 # Go module file
├── .env                   # Environment variables
└── README.md              # This file
```

---

## 🔗 Routes

### **Public Routes**

| Method | Route | Description |
|--------|-------|-------------|
| GET | `/` | Home page with article list |
| GET | `/articles/:id` | View single article |
| GET | `/search?q=query` | Search articles |
| GET | `/login` | Login page |
| POST | `/login` | Submit login credentials |
| GET | `/logout` | Logout and clear session |

### **Admin Routes** (Protected by JWT middleware)

| Method | Route | Description |
|--------|-------|-------------|
| GET | `/dashboard` | Admin dashboard |
| GET | `/new` | Create article form |
| POST | `/new` | Create new article |
| GET | `/edit/:id` | Edit article form |
| PUT | `/edit/:id` | Update article |
| DELETE | `/delete/:id` | Delete article |

---

## 🔐 Authentication

The blog uses **JWT (JSON Web Token)** authentication with HTTP-only cookies:

1. **Login** → Server validates credentials
2. **Token Generation** → JWT token created with 24-hour expiry
3. **Cookie Storage** → Token stored in secure HTTP-only cookie
4. **Protection** → Middleware validates token on each protected request
5. **Logout** → Cookie cleared, user session terminated

### Environment Variables
```env
JWT_SECRET=your-secure-secret-key      # Used to sign JWT tokens
ADMIN_USERNAME=admin                   # Admin login username
ADMIN_PASSWORD=your-secure-password    # Admin login password
```

---

## 📄 Data Storage

Articles are stored as **JSON files** in the `articles/` directory:

```json
{
  "id": 1,
  "title": "Getting Started with Go",
  "content": "Go is a powerful programming language...",
  "published": "2024-03-01",
  "author": "admin"
}
```

File naming convention: `article{id}.json`

---

## ✅ Input Validation

All user inputs are validated before processing:

| Field | Validation |
|-------|-----------|
| **Title** | Required, max 100 characters |
| **Content** | Required, no length limit |
| **Published Date** | Required, format: YYYY-MM-DD |

---

## 🛠️ Technologies Used

### Backend
- **Go** - Fast, compiled language for server-side logic
- **net/http** - Go's standard HTTP server package
- **golang-jwt** - JWT token generation and validation
- **encoding/json** - JSON marshaling/unmarshaling
- **godotenv** - Environment variable management

### Frontend
- **HTML5** - Semantic markup
- **CSS3** - Modern styling with flexbox
- **Vanilla JavaScript** - Client-side interactions (delete confirmation, forms)

### Data Storage
- **JSON Files** - Lightweight, file-based article storage

---

## 🧪 Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

Test files:
- `handlers/article_handler_test.go` - Article CRUD tests
- `handlers/auth_test.go` - Authentication tests
- `handlers/middleware/middleware_test.go` - Middleware tests
- `model/article_test.go` - Model validation tests
- `utils/validation_test.go` - Input validation tests
- `utils/server_utils_test.go` - Utility function tests

---

## 📈 Performance Features

- ✅ **HTTP-only Cookies** - Protection against XSS attacks
- ✅ **JWT Tokens** - Stateless authentication
- ✅ **Case-insensitive Search** - Flexible article discovery
- ✅ **Error Handling** - Graceful failure handling
- ✅ **Responsive Design** - Works on all devices

---

## 🚦 Getting Started Examples

### Create a New Article

1. Navigate to `http://localhost:8080/login`
2. Login with your admin credentials
3. Click "Dashboard"
4. Click "+ Add" button
5. Fill in the form and click "Publish"

### Search for Articles

1. On the home page, type in the search box
2. Click "Пошук" (Search) button
3. View filtered results

### Edit an Article

1. Go to Dashboard
2. Find the article
3. Click "Edit" button
4. Modify the content
5. Click "Publish" to save

---

## 📋 API Response Examples

### Get All Articles
**Request:** `GET /`
**Response:** HTML page with article list

### Get Single Article
**Request:** `GET /articles/1`
**Response:** HTML page with article content

### Search Articles
**Request:** `GET /search?q=golang`
**Response:** HTML page with search results

---


## 📦 Dependencies

```
github.com/golang-jwt/jwt/v5      # JWT token handling
github.com/joho/godotenv          # .env file parsing
golang.org/x/text                 # Text utilities
```

View `go.mod` and `go.sum` for detailed dependency information.

---

## 🐛 Troubleshooting

### Articles not showing?
- Check if `articles/` directory exists
- Ensure JSON files are properly formatted
- Check file permissions

### Login not working?
- Verify `.env` file exists and is properly configured
- Check if `JWT_SECRET` is set
- Ensure `ADMIN_USERNAME` and `ADMIN_PASSWORD` match

### Templates not loading?
- Ensure `templates/` directory exists
- Check file paths in template names
- Verify HTML file extensions are `.html`

---

## 📚 Future Enhancements

Planned features for future releases:

- [ ] 💬 **Comments System** - Allow readers to comment on articles
- [ ] 🏷️ **Categories & Tags** - Organize articles by topic
- [ ] 📊 **Article Sorting** - Sort by date, popularity, etc.
- [ ] 🗄️ **Database Support** - Migrate from files to SQL database
- [ ] 🔄 **RSS Feed** - Subscribe to new articles

---

## 📄 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

---


## 📊 Project Stats

- **Lines of Code:** ~2,000+
- **Test Coverage:** Comprehensive test suite
- **Supported Go Version:** 1.25.0+
- **Last Updated:** March 2026

---
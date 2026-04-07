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

#### For Local Development
- Go 1.25.0 or higher
- Git

#### For Docker
- Docker 20.10.0 or higher
- Docker Compose 1.29.0 or higher

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

## 🐳 Docker Installation (Recommended)

### Using Docker Compose (Simple & Recommended)

The easiest way to run the blog is with Docker Compose. All dependencies are automatically handled.

#### 1. **Clone the repository**
```bash
git clone https://github.com/CrabRus/Web-Blog.git
cd web-blog
```

#### 2. **Create `.env` file from example**
```bash
cp .env.example .env
```

#### 3. **Customize environment variables** (optional)
Edit `.env` file to change port or credentials:
```env
PORT=8080
ADMIN_USERNAME=admin
ADMIN_PASSWORD=your-secure-password
```

#### 4. **Build and run the container**
```bash
docker-compose up -d
```

#### 5. **Verify it's running**
```bash
docker-compose ps
```

You should see the blog service with `healthy` status. The blog will be available at `http://localhost:8080`

#### 6. **Stop the container**
```bash
docker-compose down
```

#### 7. **Stop and remove all data**
```bash
docker-compose down -v
```

### Using Docker only (without Compose)

#### 1. **Build the Docker image**
```bash
docker build -t personal-blog .
```

#### 2. **Create volume for articles**
```bash
docker volume create blog-articles
```

#### 3. **Run the container**
```bash
docker run -d \
  --name personal-blog \
  -p 8080:8080 \
  -v blog-articles:/app/articles \
  -e ADMIN_USERNAME=admin \
  -e ADMIN_PASSWORD=123 \
  personal-blog
```

#### 4. **Check if it's running**
```bash
docker ps
docker logs personal-blog
```

#### 5. **Stop the container**
```bash
docker stop personal-blog
docker rm personal-blog
```

---

## 📦 Docker Files Explained

### **Dockerfile** Features:
- **Multi-stage build**: Reduces final image size from ~500MB to ~25MB
- **Builder stage**: Compiles Go application with all dependencies
- **Runtime stage**: Alpine Linux based for minimal footprint
- **Non-root user**: Runs as `bloguser` for security
- **Health check**: Automatically monitors application availability
- **Environment variables**: Configurable port and credentials

### **docker-compose.yml** Features:
- **Service definition**: Automated build and configuration
- **Named volume**: Persists articles between restarts
- **Port mapping**: Accessible on configured PORT (default 8080)
- **Restart policy**: Automatically restarts on failure
- **Health checks**: Built-in container health monitoring
- **Network isolation**: Default bridge network for security

### **.dockerignore** Excludes:
- Git files (`.git`, `.gitignore`)
- IDE files (`.vscode`, `.idea`)
- Test files (`*_test.go`, `coverage/`)
- Build artifacts (`vendor/`, `dist/`)
- Environment files (`.env` files)
- Articles directory (populated at runtime)

---

## 🎯 Docker Commands Reference

### Docker Compose Commands

```bash
# Build and start services
docker-compose up -d

# View running services
docker-compose ps

# View service logs
docker-compose logs -f blog

# Rebuild images
docker-compose build

# Rebuild and restart
docker-compose up -d --build

# Stop services
docker-compose stop

# Stop and remove containers
docker-compose down

# Remove services and volumes
docker-compose down -v

# Execute command in running container
docker-compose exec blog /bin/sh

# View service health
docker-compose ps
```

### Docker Commands (without Compose)

```bash
# Build the image
docker build -t personal-blog .

# View images
docker images

# Run container
docker run -d --name blog -p 8080:8080 personal-blog

# View running containers
docker ps

# View container logs
docker logs blog
docker logs -f blog

# Stop container
docker stop blog

# Remove container
docker rm blog

# Remove image
docker rmi personal-blog

# Execute command in container
docker exec blog /bin/sh

# Get container stats
docker stats blog

# Inspect container
docker inspect blog
```

### Volume Commands

```bash
# List volumes
docker volume ls

# Create volume
docker volume create blog-articles

# Inspect volume
docker inspect blog-articles

# Remove volume
docker volume rm blog-articles

# List files in volume (using container)
docker run -v blog-articles:/data alpine ls -la /data
```

---

## 🔐 Environment Variables

Create a `.env` file based on `.env.example`:

```env
# Port Configuration
PORT=8080

# Admin Authentication
ADMIN_USERNAME=admin
ADMIN_PASSWORD=123

# Application Environment
APP_ENV=development

# Log Level
LOG_LEVEL=info
```

**Important Notes:**
- Never commit `.env` file to git (already in `.gitignore`)
- Use strong passwords in production
- Change default credentials before deployment
- Use `.env.example` as a template only

---

## ✅ Verification Checklist

After running `docker-compose up -d`, verify everything works:

- [ ] Container is running: `docker-compose ps` shows `healthy`
- [ ] Accessible via browser: `http://localhost:8080`
- [ ] Can view articles on home page
- [ ] Can login with credentials from `.env`
- [ ] Can create new article
- [ ] Articles persist after container restart: `docker-compose restart`
- [ ] Data persists after full recreation: `docker-compose down && docker-compose up -d`

---

## 📊 Docker Image Details

```
Image: personal-blog
Size: ~25MB (optimized)
Base Image: alpine:3.18
Go Version: 1.25.0
Runtime User: bloguser (uid: 1000)
Port: 8080
Health Check: Enabled
```

**Size Breakdown:**
- Alpine base: ~7MB
- Go runtime: ~10MB
- Application binary: ~8MB
- Total: ~25MB

---

## 🐛 Docker Troubleshooting

### Container won't start

```bash
# Check logs
docker-compose logs blog

# Common issues:
# 1. Port already in use
docker-compose down
docker-compose up -d

# 2. Previous container still running
docker ps -a
docker rm -f personal-blog

# 3. Volume issues
docker volume ls
docker volume prune
```

### Health check failing

```bash
# Check container status
docker-compose ps

# Manually test health
docker-compose exec blog wget -q http://localhost:8080 -O /dev/null && echo "OK"

# View detailed logs
docker-compose logs -f blog --tail 50
```

### Permission denied errors

```bash
# Fix permissions
docker-compose down
docker volume rm blog-articles
docker-compose up -d
```

### Articles not persisting

```bash
# Check volume is mounted
docker inspect personal-blog | grep -A 20 "Mounts"

# Verify articles directory
docker-compose exec blog ls -la /app/articles/

# Check volume contents
docker run -v blog-articles:/data alpine ls -la /data/
```

### Cannot login to admin

```bash
# Verify .env file
cat .env

# Check environment variables in container
docker-compose exec blog env | grep AUTH

# Restart container
docker-compose restart blog
```

---

## 🚀 Performance Tips

- **First run** takes longer as it builds the image
- **Subsequent runs** are fast due to layer caching
- **Layer caching** is optimized: dependencies cached before code
- Remove unused images: `docker image prune -a`
- Clean up volumes: `docker volume prune`
- Limit container resources if needed:

```yaml
# In docker-compose.yml services.blog
deploy:
  resources:
    limits:
      memory: 512M
      cpus: '0.5'
```

---

## 🔄 Development Workflow with Docker

### Hot Reload (Docker Compose)

For development with automatic rebuilds:

```bash
# Build and run in foreground (see logs)
docker-compose up --build

# Press Ctrl+C to stop
# Modify files
# Start again
docker-compose up --build
```

### Access Application Shell

```bash
docker-compose exec blog /bin/sh

# Now you can run commands inside container
ls -la /app
cat /app/articles/article1.json
```

### Debug Mode

```bash
# Run with interactive shell
docker-compose run --rm blog /bin/sh

# Inside container
ls -la
ps aux
```

---

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
# Student Management System - NextJS + Golang + PostgreSQL

Hệ thống quản lý sinh viên full-stack sử dụng NextJS cho frontend, Golang cho backend API, và PostgreSQL làm cơ sở dữ liệu.

## 🏗️ Cấu trúc dự án chi tiết

```
Project/
├── frontend/                    # NextJS Frontend Application
│   ├── src/
│   │   ├── app/                # App Router (NextJS 13+)
│   │   │   ├── layout.tsx      # Root layout component
│   │   │   ├── page.tsx        # Main dashboard page
│   │   │   └── globals.css     # Global styles
│   │   └── lib/                # Shared libraries và utilities
│   │       └── api.ts          # API client cho backend communication
│   ├── public/                 # Static assets
│   ├── package.json            # Dependencies và scripts
│   ├── tailwind.config.ts      # Tailwind CSS configuration
│   ├── next.config.js          # NextJS configuration
│   ├── .env.local              # Environment variables (local)
│   └── .env.local.example      # Environment template
├── backend/                    # Golang Backend API Server
│   ├── cmd/                    # Application entry points
│   │   └── main.go             # Main server application
│   ├── internal/               # Private application code
│   │   ├── handlers/           # HTTP request handlers
│   │   │   └── student.go      # Student CRUD operations
│   │   ├── models/             # Database models và structs
│   │   │   └── student.go      # Student model definition
│   │   └── database/           # Database connection và configuration
│   │       └── database.go     # PostgreSQL connection setup
│   ├── scripts/                # Utility scripts
│   │   ├── seed_students.go    # Database seeding script
│   │   └── utils.go            # Helper functions
│   ├── go.mod                  # Go module dependencies
│   ├── go.sum                  # Dependency checksums
│   ├── .env                    # Environment variables
│   ├── .env.example            # Environment template
│   ├── .gitignore              # Git ignore rules
│   └── Dockerfile              # Docker container configuration
├── docker-compose.yml          # Multi-container Docker setup
└── README.md                   # Project documentation
```

## � Mô tả chi tiết các thư mục và chức năng

### 🎨 Frontend (NextJS)

#### `frontend/src/app/`

- **Mục đích**: Chứa các pages và layouts của ứng dụng NextJS 13+ với App Router
- **Chức năng**:
  - `layout.tsx`: Root layout component, định nghĩa cấu trúc HTML cơ bản
  - `page.tsx`: Trang chính hiển thị dashboard quản lý sinh viên
  - `globals.css`: Styles toàn cục và Tailwind CSS imports

#### `frontend/src/lib/`

- **Mục đích**: Chứa các thư viện và utilities dùng chung
- **Chức năng**:
  - `api.ts`: API client class để giao tiếp với backend
  - Định nghĩa TypeScript interfaces cho Student, ApiResponse
  - Xử lý HTTP requests (GET, POST, PUT, DELETE)
  - Error handling và logging

#### `frontend/public/`

- **Mục đích**: Chứa static assets (images, icons, fonts)
- **Chức năng**: Serve các file tĩnh trực tiếp từ root URL

#### Configuration Files

- `package.json`: Quản lý dependencies, scripts, và metadata
- `tailwind.config.ts`: Cấu hình Tailwind CSS framework
- `next.config.js`: Cấu hình NextJS (routing, optimization, etc.)
- `.env.local`: Environment variables cho development

### 🔧 Backend (Golang)

#### `backend/cmd/`

- **Mục đích**: Entry points của ứng dụng
- **Chức năng**:
  - `main.go`: Khởi tạo server, database connection, routes
  - Setup middleware (CORS, logging)
  - Định nghĩa API endpoints và port listening

#### `backend/internal/handlers/`

- **Mục đích**: HTTP request handlers (Controllers trong MVC)
- **Chức năng**:
  - `student.go`: CRUD operations cho Student entity
  - Xử lý HTTP requests và responses
  - Validation và error handling
  - Business logic layer

#### `backend/internal/models/`

- **Mục đích**: Database models và data structures
- **Chức năng**:
  - `student.go`: Student struct với GORM tags
  - Định nghĩa database schema
  - Relationships và constraints
  - Table naming conventions

#### `backend/internal/database/`

- **Mục đích**: Database connection và configuration
- **Chức năng**:
  - `database.go`: PostgreSQL connection setup
  - GORM initialization
  - Auto-migration functions
  - Connection pooling

#### `backend/scripts/`

- **Mục đích**: Utility scripts và tools
- **Chức năng**:
  - `seed_students.go`: Tạo dữ liệu mẫu cho testing
  - `utils.go`: Helper functions (stringPtr, intPtr, etc.)
  - Database maintenance scripts

#### Configuration Files

- `go.mod`: Go module definition và dependencies
- `go.sum`: Dependency checksums cho security
- `.env`: Environment variables cho production/development
- `Dockerfile`: Container configuration cho deployment

### 🐳 Docker & Deployment

#### `docker-compose.yml`

- **Mục đích**: Multi-container orchestration
- **Chức năng**:
  - PostgreSQL database container
  - Backend API container
  - Frontend web container
  - Network configuration và volume mounting

## �🚀 Yêu cầu hệ thống

- **Node.js** >= 18.0.0
- **Go** >= 1.21
- **PostgreSQL** >= 13
- **Git**
- **Docker** (optional, cho containerized deployment)

## 🏛️ Kiến trúc hệ thống

### Architecture Overview

```
┌─────────────────┐    HTTP/REST API    ┌─────────────────┐    SQL Queries    ┌─────────────────┐
│                 │ ◄─────────────────► │                 │ ◄───────────────► │                 │
│   Frontend      │                     │   Backend       │                   │   Database      │
│   (NextJS)      │                     │   (Golang)      │                   │   (PostgreSQL)  │
│                 │                     │                 │                   │                 │
└─────────────────┘                     └─────────────────┘                   └─────────────────┘
│                                       │                                     │
├─ React Components                     ├─ Gin HTTP Server                   ├─ Students Table
├─ TypeScript                          ├─ GORM ORM                          ├─ Indexes
├─ Tailwind CSS                        ├─ RESTful APIs                      ├─ Constraints
├─ API Client                          ├─ Middleware                        └─ Relationships
└─ State Management                    └─ Error Handling
```

### Data Flow

1. **User Interaction**: User tương tác với React components
2. **API Call**: Frontend gọi API thông qua API client
3. **HTTP Request**: Request được gửi đến Golang backend
4. **Handler Processing**: Gin router chuyển request đến appropriate handler
5. **Database Operation**: Handler sử dụng GORM để thao tác với PostgreSQL
6. **Response**: Data được trả về qua JSON response
7. **UI Update**: Frontend cập nhật UI với data mới

### Technology Stack

#### Frontend Stack

- **NextJS 14**: React framework với App Router
- **TypeScript**: Type-safe JavaScript
- **Tailwind CSS**: Utility-first CSS framework
- **React Hooks**: State management và side effects

#### Backend Stack

- **Golang**: High-performance backend language
- **Gin**: HTTP web framework
- **GORM**: Object-Relational Mapping library
- **PostgreSQL Driver**: Database connectivity

#### Database Stack

- **PostgreSQL**: Relational database
- **GORM Migrations**: Schema management
- **Indexes**: Performance optimization
- **Constraints**: Data integrity

## 📦 Cài đặt và thiết lập

### 1. Clone dự án

```bash
git clone <repository-url>
cd Project
```

### 2. Thiết lập PostgreSQL

Tạo database mới:

```sql
CREATE DATABASE project_db;
CREATE USER project_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE project_db TO project_user;
```

### 3. Thiết lập Backend (Golang)

```bash
cd backend

# Copy file cấu hình
cp .env.example .env

# Chỉnh sửa file .env với thông tin database của bạn
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=project_user
# DB_PASSWORD=your_password
# DB_NAME=project_db
# DB_SSLMODE=disable
# PORT=8080

# Cài đặt dependencies
go mod tidy

# Chạy server
go run cmd/main.go

# Tùy chọn: Seed dữ liệu mẫu
go run cmd/seed.go
```

Server sẽ chạy tại: `http://localhost:8080`

### 4. Thiết lập Frontend (NextJS)

```bash
cd frontend

# Cài đặt dependencies
npm install

# Copy file cấu hình
cp .env.local.example .env.local

# Chỉnh sửa file .env.local nếu cần
# NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1

# Chạy development server
npm run dev
```

Frontend sẽ chạy tại: `http://localhost:3000`

## 🔧 API Endpoints

### Health Check

- `GET /health` - Kiểm tra trạng thái server

### Users

- `GET /api/v1/users` - Lấy danh sách users
- `POST /api/v1/users` - Tạo user mới
- `GET /api/v1/users/:id` - Lấy thông tin user theo ID

### Employees

- `GET /api/v1/employees` - Lấy danh sách nhân viên
- `POST /api/v1/employees` - Tạo nhân viên mới
- `GET /api/v1/employees/:id` - Lấy thông tin nhân viên theo ID
- `PUT /api/v1/employees/:id` - Cập nhật thông tin nhân viên
- `DELETE /api/v1/employees/:id` - Xóa nhân viên (soft delete)
- `GET /api/v1/employees/department/:departmentId` - Lấy nhân viên theo phòng ban
- `GET /api/v1/employees/status?status=active` - Lấy nhân viên theo trạng thái

### Departments

- `GET /api/v1/departments` - Lấy danh sách phòng ban
- `POST /api/v1/departments` - Tạo phòng ban mới

### Ví dụ tạo nhân viên mới:

```bash
curl -X POST http://localhost:8080/api/v1/employees \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@company.com",
    "phone": "+1234567890",
    "department_id": 1,
    "position": "Software Engineer",
    "status": "active",
    "employee_id": "EMP001"
  }'
```

## 🛠️ Development

### Backend Development

```bash
cd backend

# Chạy với hot reload (cần cài air)
go install github.com/cosmtrek/air@latest
air

# Hoặc chạy thông thường
go run cmd/main.go

# Build production
go build -o bin/server cmd/main.go
```

### Frontend Development

```bash
cd frontend

# Development mode
npm run dev

# Build production
npm run build

# Start production server
npm start

# Lint code
npm run lint
```

## 🗄️ Database Schema

### Users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
```

## 🔒 Environment Variables

### Backend (.env)

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=project_db
DB_SSLMODE=disable
PORT=8080
GIN_MODE=debug
```

### Frontend (.env.local)

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api/v1
```

## 🚀 Production Deployment

### Backend

```bash
# Build binary
go build -o bin/server cmd/main.go

# Run production server
GIN_MODE=release ./bin/server
```

### Frontend

```bash
# Build static files
npm run build

# Deploy to Vercel, Netlify, hoặc server khác
```

## 🧪 Testing

### Backend Testing

```bash
cd backend
go test ./...
```

### Frontend Testing

```bash
cd frontend
npm test
```

## 📝 Features và Chức năng

### 🎯 Core Features

- ✅ **Student Management**: CRUD operations cho sinh viên
- ✅ **Real-time UI**: Instant feedback và loading states
- ✅ **Data Validation**: Frontend và backend validation
- ✅ **Error Handling**: Comprehensive error management
- ✅ **Responsive Design**: Mobile-friendly interface
- ✅ **Type Safety**: Full TypeScript support

### 🔧 Technical Features

- ✅ **RESTful API**: Chuẩn REST với Golang + Gin
- ✅ **ORM Integration**: GORM cho database operations
- ✅ **Auto Migration**: Tự động tạo database schema
- ✅ **CORS Support**: Cross-origin resource sharing
- ✅ **Environment Config**: Flexible configuration management
- ✅ **Docker Support**: Containerized deployment
- ✅ **Soft Delete**: GORM soft delete functionality
- ✅ **Connection Pooling**: Optimized database connections

### 🎨 UI/UX Features

- ✅ **Modern Design**: Clean và professional interface
- ✅ **Form Validation**: Real-time input validation
- ✅ **Loading States**: Visual feedback cho user actions
- ✅ **Confirmation Dialogs**: Safe delete operations
- ✅ **Success Messages**: Clear action feedback
- ✅ **Error Display**: User-friendly error messages
- ✅ **Responsive Layout**: Works on all screen sizes

## 🔧 Development Guidelines

### Code Structure

- **Frontend**: Follow React best practices và NextJS conventions
- **Backend**: Implement clean architecture với separation of concerns
- **Database**: Use GORM conventions cho model definitions
- **API**: Follow RESTful principles và consistent naming

### Naming Conventions

- **Files**: snake_case cho Go files, camelCase cho TypeScript
- **Functions**: camelCase cho JavaScript/TypeScript, PascalCase cho Go exports
- **Variables**: camelCase cho local variables, UPPER_CASE cho constants
- **Database**: snake_case cho table và column names

### Error Handling

- **Frontend**: Use try-catch với meaningful error messages
- **Backend**: Return appropriate HTTP status codes
- **Logging**: Use console.log cho frontend, log package cho backend

## 🐛 Troubleshooting

### Common Issues

#### Backend không khởi động

```bash
# Kiểm tra Go installation
go version

# Kiểm tra dependencies
go mod tidy

# Kiểm tra database connection
psql -h localhost -U postgres -d test_db
```

#### Frontend không load data

```bash
# Kiểm tra backend server
curl http://localhost:8080/health

# Kiểm tra API endpoint
curl http://localhost:8080/api/v1/students

# Check browser console for errors
```

#### Database connection errors

- Kiểm tra PostgreSQL service đang chạy
- Verify database credentials trong .env file
- Ensure database `test_db` đã được tạo

#### CORS errors

- Backend đã có CORS middleware configured
- Check API_URL trong frontend .env.local

### Performance Tips

- Use database indexes cho frequently queried fields
- Implement pagination cho large datasets
- Use connection pooling cho database
- Optimize bundle size với NextJS

## 🤝 Contributing

1. Fork dự án
2. Tạo feature branch (`git checkout -b feature/AmazingFeature`)
3. Follow coding standards và conventions
4. Write tests cho new features
5. Commit changes (`git commit -m 'Add some AmazingFeature'`)
6. Push to branch (`git push origin feature/AmazingFeature`)
7. Tạo Pull Request

## 📚 Additional Resources

- [NextJS Documentation](https://nextjs.org/docs)
- [Golang Documentation](https://golang.org/doc/)
- [GORM Documentation](https://gorm.io/docs/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Tailwind CSS Documentation](https://tailwindcss.com/docs)

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

## 📞 Support

Nếu bạn gặp vấn đề, hãy tạo issue trên GitHub repository.

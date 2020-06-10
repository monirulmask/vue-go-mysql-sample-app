# vue-go-mysql-sample-app

This is sample project created with Go/Gin/Vue.js/MySQL. For ORM i have used gorm.
## How to run for development

### Go version 1.14.4 has beed used for development.

### Step 1: Create Database named 'schoolmanagement'

### Step 2: Install Gin(If already installed then no need to install)
```bash
go get github.com/gin-gonic/gin
```

### Step 3: Install Mysql Driver(If already installed then no need to install)
```bash
go get github.com/go-sql-driver/mysql
```

### Step 4: Install Gorm Driver(If already installed then no need to install)
```bash
go get github.com/jinzhu/gorm
```

### Step 5: Go to Source(vue-go-mysql-sample-app) Folder and type following Command
```bash
go run server.go
```

After running server, application can be browse in following URL http://localhost:8080/schoolapp/



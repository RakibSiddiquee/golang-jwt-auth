# Golang Fiber JWT Authentication

## Install Fiber Framework:  
go get github.com/gofiber/fiber/v2   

## Install GORM with MySQL: 
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql   

### We used bcrypt for password hashing  

## Run the project  
go run main.go  

### Registration URL: 
http://localhost:8000/api/register   

### Login URL: 
http://localhost:8000/api/login   

### Get logged in user after login: 
http://localhost:8000/api/user   

### Logout URL: 
http://localhost:8000/api/logout   




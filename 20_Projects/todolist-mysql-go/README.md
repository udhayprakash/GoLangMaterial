docker run -d -p 3307:3306 --name mysql -e MYSQL_ROOT_PASSWORD=root mysql
docker exec -it mysql mysql -uroot -proot -e 'CREATE DATABASE todolist'

go get -u github.com/gorilla/mux
go get -u github.com/sirupsen/logrus
go get -u github.com/jinzhu/gorm
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jinzhu/gorm/dialects/mysql

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	data := "Hello Go/Gin!!"
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{"data": data})
	})

	router.Run()
}

type Todo struct {
	gorm.Model
	Text   string
	Status string
}

//DB init
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("cant open db!（dbInit）")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

//DB add
func dbInsert(text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("cant open db! （dbInsert)")
	}
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

//DB update
func dbUpdate(id int, text string, status string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("cant open db! （dbUpdate)")
	}
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//DB delete
func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("cant open db! （dbDelete)")
	}
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//DB index
func dbGetAll() []Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("cant open db! (dbGetAll())")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//DB get one
func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("cant open db! (dbGetOne())")
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

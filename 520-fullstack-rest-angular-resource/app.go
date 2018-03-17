/// Przykład pokazujący proste RESTowe API
// aby było prościej: w jednym pliku
// dodamy także super szybki memory storage
// nie tam jakieś wolne Redisy
package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Definiujemy struktury danych
type Person struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	City  string `json:"city,omitempty"`
	Phone string `json:"phone"`
}

type People map[int]Person

// Memory storage
type Storage struct {
	MaxId int
	Data  People
}

func (storage *Storage) New(person Person) Person {
	storage.MaxId++
	person.Id = storage.MaxId
	storage.Data[person.Id] = person
	return person
}

func (storage *Storage) GetAll() []Person {
	result := []Person{}
	for _, person := range storage.Data {
		result = append(result, person)
	}
	return result
}

func (storage *Storage) Get(id int) Person {
	return storage.Data[id]
}

func (storage *Storage) Set(person Person) {
	storage.Data[person.Id] = person
}

func (storage *Storage) Delete(id int) {
	delete(storage.Data, id)
}

func idFromContext(context *gin.Context) int {
	val, _ := strconv.Atoi(context.Params.ByName("id"))
	return val
}

var storage Storage

func usersListHandler(c *gin.Context) {
	people := storage.GetAll()
	c.JSON(200, people)
}

func userGetHandler(c *gin.Context) {
	id := idFromContext(c)
	c.JSON(200, storage.Get(id))
}

func userDeleteHandler(c *gin.Context) {
	id := idFromContext(c)
	storage.Delete(id)
	c.JSON(200, gin.H{"deleted": 1})
}

func userSaveHandler(c *gin.Context) {
	id := idFromContext(c)
	person := storage.Get(id)

	err := c.Bind(&person)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	storage.Set(person)
	c.JSON(200, gin.H{"updated": 1, "person": person})
}

func userNewHandler(c *gin.Context) {
	person := Person{}
	err := c.Bind(&person)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	person = storage.New(person)

	c.JSON(200, gin.H{"inserted": 1, "person": person})
}

func init() {
	var person Person

	// nasz memory storage potrzebuje trochę danych
	storage = Storage{MaxId: 0, Data: People{}}

	person = Person{}
	person.Name = "Jacek Placek"
	person.Age = 18
	person.Phone = "+48 666 666 666"
	storage.New(person)

	person = Person{}
	person.Name = "Wacek Placek"
	person.Age = 120
	person.City = "Poznań"
	person.Phone = "+48 777 777 777"
	storage.New(person)

	person = Person{}
	person.Name = "Wicek Micek"
	person.Age = 23
	person.City = "Wa-wa"
	person.Phone = "+48 123 123 123"
	storage.New(person)

}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", "")
}

func main() {
	app := gin.New()

	app.LoadHTMLFiles("./static/index.html")
	app.Static("/static", "static")

	app.GET("/", indexHandler)

	// to be or not to be restfull - version your api
	v1 := app.Group("/v1")

	// mess with the best or die like the REST
	v1.GET("/users", usersListHandler)
	v1.GET("/users/:id", userGetHandler)
	v1.POST("/users", userNewHandler)
	v1.PUT("/users/:id", userSaveHandler)
	v1.DELETE("/users/:id", userDeleteHandler)

	app.Run(":8080")
}

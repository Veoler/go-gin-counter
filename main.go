package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	// "fmt"
)

var counter int

func main() {
	r := gin.Default()

	r.GET("/ping", func (c *gin.Context) {
		c.String(200, "All Okey")
	})

	// GET /current — вызывает Value() и возвращает JSON { "value": <число> };
	r.GET("/current", func (c *gin.Context) {
		c.JSON(200, gin.H{"value": Value()})
	})

	r.GET("/increment", func (c *gin.Context) {
		c.JSON(200, gin.H{"value": Inc1()})
	})

	r.GET("/decrement", func (c *gin.Context) {
		c.JSON(200, gin.H{"value": Dec1()})
	})

	r.GET("/reset", func (c *gin.Context) {
		c.JSON(200, gin.H{"value": Reset()})
	})

	r.GET("/add", func(c *gin.Context) {
		val := c.Query("n")

		n, err := strconv.Atoi(val)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "неверное значение параметра n",})
			return
		}

		newVal := Add(n)
		c.JSON(http.StatusOK, gin.H{"value": newVal})
	})

	r.GET("/subtract", func(c *gin.Context) {
		val := c.Query("n")

		n, err := strconv.Atoi(val)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "неверное значение параметра n",})
			return
		}

		newVal := Subtract(n)
		c.JSON(http.StatusOK, gin.H{"value": newVal})
	})

	r.Run(":8080")

	// counter = 1

	// Value()
	// fmt.Println(counter)
	// Inc1()
	// fmt.Println(counter)
	// Dec1()
	// fmt.Println(counter)
	// Reset()
	// fmt.Println(counter)
	
}

func Inc1() int {
  // +1 и вернуть новое значение
	counter++
	return counter
}
func Dec1() int   {
	// -1 и вернуть новое значение
	counter--
	return counter
}
func Reset() int  {
	// 0 и вернуть 0
	counter = 0 
	return counter
}
func Value() int  {
	// просто вернуть текущее значение
	return counter
}

func Add(n int) int {
	counter += n
	return counter
}

func Subtract(n int) int {
	counter -= n
	return counter
}
package main

import (
	"fmt"
	cache2 "github.com/ammorteza/redis_cache/cache"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Input struct {
	X 			int		`json:"x"`
	Y 			int		`json:"y"`
}

func main(){
	router := gin.Default()
	router.POST("/calc", func(context *gin.Context) {
		var input Input
		if err := context.BindJSON(&input); err != nil{
			context.JSON(http.StatusInternalServerError, nil)
			return
		}

		cache := cache2.NewRedisCache(1, 10)
		calcRes, err := cache.Get(strconv.Itoa(input.X) + strconv.Itoa(input.Y))
		fmt.Println(calcRes)
		if err != nil{
			result := 0
			for i := 1; i <= input.X; i++{
				for j := 1; j <= input.Y; j++{
					result += i * j
				}
			}

			if err := cache.Set(strconv.Itoa(input.X) + strconv.Itoa(input.Y), strconv.Itoa(result)); err != nil{
				context.JSON(http.StatusInternalServerError, err)
				return
			}

			context.JSON(http.StatusOK, strconv.Itoa(result))
		}else{
			context.JSON(http.StatusOK, calcRes)
		}

	})
	log.Fatal(router.Run(":8080"))
}
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Pessoa struct {
	Nome string `json:"nome"`
}

func main() {

	fmt.Println("Iniciando:")

	r := gin.Default()

	r.GET("/user/:usuario", func(c *gin.Context) {
		user := c.Params.ByName("usuario")

		if user == "Fabrício" {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": "Fabrício"})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "usuário não existe"})
		}
	})

	r.POST("user", func(c *gin.Context) {

		var pessoa Pessoa

		if c.Bind(&pessoa) == nil {

			encoder := json.NewEncoder(os.Stdout)
			encoder.Encode(pessoa)
			fmt.Println(" - ", pessoa.Nome)
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		} else {
			fmt.Println("erro")
		}
	})

	r.Run(":8080")

}

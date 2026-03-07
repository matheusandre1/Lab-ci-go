package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/matheusandre1/Lab-ci-go/config"
	"github.com/matheusandre1/Lab-ci-go/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/:nome", controllers.Saudacoes)
	r.GET("/alunos", controllers.TodosAlunos)
	r.GET("/alunos/:id", controllers.BuscarAlunoPorID)
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/alunos/", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)
	if err := r.Run(":" + config.GetAppPort()); err != nil {
		log.Panicf("Erro ao iniciar servidor: %v", err)
	}
}

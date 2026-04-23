package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/jonathantvrs/pismo/docs"
	"github.com/jonathantvrs/pismo/internal/handler"
	"github.com/jonathantvrs/pismo/internal/repository"
	"github.com/jonathantvrs/pismo/internal/usecase"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initDependencies(db *sql.DB) (accHandler *handler.AccountHandler, txHandler *handler.TransactionHandler) {
	//txRepo := repository.NewTransactionMemoryRepo()
	//accRepo := repository.NewAccountMemoryRepo()
	//opTypeRepo := repository.NewOperationMemoryRepo()

	accRepo := repository.NewSQLiteAccountRepo(db)
	txRepo := repository.NewSQLiteTransactionRepo(db)
	opTypeRepo := repository.NewSQLiteOperationRepo(db)

	accUseCase := usecase.NewAccountUseCase(accRepo)
	txUseCase := usecase.NewTransactionUseCase(txRepo, accRepo, opTypeRepo)

	accHandler = handler.NewAccountHandler(accUseCase)
	txHandler = handler.NewTransactionHandler(txUseCase)

	return
}

func main() {
	db, err := repository.InitSQLiteDB("./transactions.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	accHandler, txHandler := initDependencies(db)

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/accounts", accHandler.Create)
	r.GET("/accounts/:accountId", accHandler.Get)
	r.POST("/transactions", txHandler.Create)

	r.Run(":8080")
}

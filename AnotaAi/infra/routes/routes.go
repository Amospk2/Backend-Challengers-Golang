package routes

import (
	"api/infra/controllers"
	"api/infra/database"
	"api/infra/middleware"
	"api/infra/service"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func addRoutes(muxR *mux.Router, pool *pgxpool.Pool) {
	service := service.CreateSession()
	NewUserRouter(controllers.NewUserController(database.NewUserRepositoryImp(pool))).Load(muxR)
	NewProductRouter(controllers.NewProductController(database.NewProductRepository(pool), service)).Load(muxR)
	NewCategoryRouter(controllers.NewCategoryController(database.NewCategoryRepository(pool), service)).Load(muxR)
	NewAuthRouter(controllers.NewAuthController(database.NewUserRepositoryImp(pool))).Load(muxR)
	muxR.Use(mux.CORSMethodMiddleware(muxR))
}

func NewServer(env map[string]string) *mux.Router {
	mux := mux.NewRouter()

	connect := database.NewConnect(env["DATABASE_URL"])

	addRoutes(mux, connect)

	mux.Use(middleware.ApplicationTypeSet)

	return mux
}

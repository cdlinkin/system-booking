package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cdlinkin/system-booking/internal/config"
	"github.com/cdlinkin/system-booking/internal/handler"
	"github.com/cdlinkin/system-booking/internal/repo"
	"github.com/cdlinkin/system-booking/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cfg := config.DBConfigLoad()

	db := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, db)
	if err != nil {
		fmt.Printf("Ошибка при подключении БД: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		fmt.Printf("Произошла ошибка при выполнении пинг-запроса: %v", err)
	}
	fmt.Println("База данных подключена успешно")

	userRepo := repo.NewUserRepo(pool)
	resourceRepo := repo.NewResourceRepo(pool)
	bookingRepo := repo.NewBookingRepo(pool)

	userService := service.NewUserService(userRepo)
	resourceService := service.NewResourceService(resourceRepo)
	bookingService := service.NewBookingService(bookingRepo)

	userHandler := handler.NewUserHandler(userService, bookingService)
	resourceHandler := handler.NewResourceHandler(resourceService)
	bookingHandler := handler.NewBookingHandler(bookingService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/users", userHandler.Register)
	mux.HandleFunc("GET /api/resources", resourceHandler.GetResources)
	mux.HandleFunc("POST /api/bookings", bookingHandler.CreateBooking)
	mux.HandleFunc("GET /api/bookings/", bookingHandler.GetId)
	mux.HandleFunc("GET /api/users/", userHandler.GetBookingsByUserID)

	http.ListenAndServe(":9090", mux)
	fmt.Println("server starting :9090")
}

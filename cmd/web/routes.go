package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/loidinhm31/go-bookings-system/internal/config"
	"github.com/loidinhm31/go-bookings-system/internal/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)

	mux.Get("/about", handlers.Repo.About)

	mux.Get("/generals-quarters", handlers.Repo.Generals)

	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/book-room", handlers.Repo.BookRoom)

	mux.Get("/user/login", handlers.Repo.ShowLogin)
	mux.Post("/user/login", handlers.Repo.PostLogin)
	mux.Get("/user/logout", handlers.Repo.Logout)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Route("/admin", func(r chi.Router) {
		r.Use(Auth)
		r.Get("/dashboard", handlers.Repo.AdminDashboard)

		r.Get("/reservations-new", handlers.Repo.AdminNewReservations)
		r.Get("/reservations-all", handlers.Repo.AdminAllReservations)
		r.Get("/reservations-calendar", handlers.Repo.AdminReservationsCalendar)
		r.Post("/reservations-calendar", handlers.Repo.AdminPostReservationsCalendar)

		r.Get("/reservations/{src}/{id}/show", handlers.Repo.AdminShowReservation)
		r.Post("/reservations/{src}/{id}", handlers.Repo.AdminPostShowReservation)

		r.Get("/process-reservation/{src}/{id}/action", handlers.Repo.AdminProcessReservation)
		r.Get("/delete-reservation/{src}/{id}/action", handlers.Repo.AdminDeleteReservation)
	})

	return mux
}

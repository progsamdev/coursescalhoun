package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/csrf"
	"github.com/progsamdev/coursescalhoun/controllers"
	"github.com/progsamdev/coursescalhoun/models"
	"github.com/progsamdev/coursescalhoun/templates"
	"github.com/progsamdev/coursescalhoun/views"
)

func main() {
	r := chi.NewRouter() //chi->new Router
	r.Use(middleware.Logger)

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	config := models.DefaultPostgresConfig()
	fmt.Println(config.ToString())
	db, err := models.Open(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		panic(err)
	}
	defer db.Close()

	userSer := models.UserService{
		DB: db,
	}

	tokenManager := models.TokenManager{}

	sessionService := models.SessionService{
		DB:           db,
		TokenManager: tokenManager,
	}

	usersC := controllers.User{
		UserService:    &userSer,
		SessionService: &sessionService,
	}

	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", "tailwind.gohtml",
	))

	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", "tailwind.gohtml",
	))

	usersC.Templates.CurrentUser = views.Must(views.ParseFS(
		templates.FS,
		"currentuser.gohtml", "tailwind.gohtml",
	))

	r.Get("/signup", usersC.New)
	r.Post("/signup", usersC.Create)

	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/users/me", usersC.CurrentUser)
	r.Post("/signout", usersC.ProcessSignOut) //could be delete

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")

	csrfKey := "gMxYc0afnN5tQaFLLToTUKTIsRFR9AI"
	csrfMw := csrf.Protect([]byte(csrfKey),
		csrf.Secure(false),
	)
	http.ListenAndServe(":3000", csrfMw(r))
}

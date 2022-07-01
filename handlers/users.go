package handlers

// responsible for all API endpoints having the /users prefix as we specified in the main handler file

import (
    "context"
    "fmt"
    "net/http"
    "strconv"
    "github.com/go-chi/chi"
    "github.com/go-chi/render"
    "github.com/aicyp/escape-dan-app/controllers"
    "github.com/aicyp/escape-dan-app/models"
)

var userIdKey = "userId"

func users(router chi.Router) {
    router.Get("/", getAllUsers)
    router.Post("/", createUser)
    router.Route("/{userId}", func(router chi.Router) {
        router.Use(UserContext)
        router.Get("/", getUser)
        router.Put("/", updateUser)
        router.Delete("/", deleteUser)
    })
}

func UserContext(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userId := chi.URLParam(r, "userId")
        if userId == "" {
            render.Render(w, r, ErrorRenderer(fmt.Errorf("user ID is required")))
            return
        }
        id, err := strconv.Atoi(userId)
        if err != nil {
            render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid user ID")))
        }
        ctx := context.WithValue(r.Context(), userIdKey, id)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Handler for the AddUser() function
func createUser(w http.ResponseWriter, r *http.Request) {
    user := &models.User{}
    if err := render.Bind(r, user); err != nil {
        render.Render(w, r, ErrBadRequest)
        return
    }
    if err := dbInstance.AddUser(user); err != nil {
        render.Render(w, r, ErrorRenderer(err))
        return
    }
    if err := render.Render(w, r, user); err != nil {
        render.Render(w, r, ServerErrorRenderer(err))
        return
    }
}

// Handler for the GetAllUsers() function
func getAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := dbInstance.GetAllUsers()
    if err != nil {
        render.Render(w, r, ServerErrorRenderer(err))
        return
    }
    if err := render.Render(w, r, users); err != nil {
        render.Render(w, r, ErrorRenderer(err))
    }
}

// Handler for the GetUserById() function
func getUser(w http.ResponseWriter, r *http.Request) {
    userId := r.Context().Value(userIdKey).(int)
    user, err := dbInstance.GetUserById(userId)
    if err != nil {
        if err == db.ErrNoMatch {
            render.Render(w, r, ErrNotFound)
        } else {
            render.Render(w, r, ErrorRenderer(err))
        }
        return
    }
    if err := render.Render(w, r, &user); err != nil {
        render.Render(w, r, ServerErrorRenderer(err))
        return
    }
}

// Handler for the DeleteUser() function
func deleteUser(w http.ResponseWriter, r *http.Request) {
    userId := r.Context().Value(userIdKey).(int)
    err := dbInstance.DeleteUser(userId)
    if err != nil {
        if err == db.ErrNoMatch {
            render.Render(w, r, ErrNotFound)
        } else {
            render.Render(w, r, ServerErrorRenderer(err))
        }
        return
    }
}

// Handler for the UpdateUser() function
func updateUser(w http.ResponseWriter, r *http.Request) {
    userId := r.Context().Value(userIdKey).(int)
    userData := models.User{}
    if err := render.Bind(r, &userData); err != nil {
        render.Render(w, r, ErrBadRequest)
        return
    }
    user, err := dbInstance.UpdateUser(userId, userData)
    if err != nil {
        if err == db.ErrNoMatch {
            render.Render(w, r, ErrNotFound)
        } else {
            render.Render(w, r, ServerErrorRenderer(err))
        }
        return
    }
    if err := render.Render(w, r, &user); err != nil {
        render.Render(w, r, ServerErrorRenderer(err))
        return
    }
}
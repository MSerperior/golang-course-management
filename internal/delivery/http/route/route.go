package route

import (
	"golang-clean-architecture/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                   *fiber.App
	UserController        *http.UserController
	ContactController     *http.ContactController
	AddressController     *http.AddressController
	CategoryController    *http.CategoryController
	CourseController      *http.CourseController
	SectionController     *http.SectionController
	LessonController      *http.LessonController
	EnrollmentController  *http.EnrollmentController
	ReviewController      *http.ReviewController
	RoleController        *http.RoleController
	TransactionController *http.TransactionController
	UserRoleController    *http.UserRoleController
	AuthMiddleware        fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/_login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Delete("/api/users", c.UserController.Logout)
	c.App.Patch("/api/users/_current", c.UserController.Update)
	c.App.Get("/api/users/_current", c.UserController.Current)

	c.App.Get("/api/contacts", c.ContactController.List)
	c.App.Post("/api/contacts", c.ContactController.Create)
	c.App.Put("/api/contacts/:contactId", c.ContactController.Update)
	c.App.Get("/api/contacts/:contactId", c.ContactController.Get)
	c.App.Delete("/api/contacts/:contactId", c.ContactController.Delete)

	c.App.Get("/api/contacts/:contactId/addresses", c.AddressController.List)
	c.App.Post("/api/contacts/:contactId/addresses", c.AddressController.Create)
	c.App.Put("/api/contacts/:contactId/addresses/:addressId", c.AddressController.Update)
	c.App.Get("/api/contacts/:contactId/addresses/:addressId", c.AddressController.Get)
	c.App.Delete("/api/contacts/:contactId/addresses/:addressId", c.AddressController.Delete)

	c.App.Get("/api/categories", c.CategoryController.List)
	c.App.Post("/api/categories", c.CategoryController.Create)
	c.App.Put("/api/categories/:categoryId", c.CategoryController.Update)
	c.App.Get("/api/categories/:categoryId", c.CategoryController.Get)
	c.App.Delete("/api/categories/:categoryId", c.CategoryController.Delete)

	// Courses
	c.App.Get("/api/courses", c.CourseController.List)
	c.App.Post("/api/courses", c.CourseController.Create)
	c.App.Put("/api/courses/:courseId", c.CourseController.Update)
	c.App.Get("/api/courses/:courseId", c.CourseController.Get)
	c.App.Delete("/api/courses/:courseId", c.CourseController.Delete)

	// Sections
	c.App.Post("/api/sections", c.SectionController.Create)
	c.App.Get("/api/sections/:sectionId", c.SectionController.Get)
	c.App.Put("/api/sections/:sectionId", c.SectionController.Update)
	c.App.Delete("/api/sections/:sectionId", c.SectionController.Delete)

	// Lessons
	c.App.Post("/api/lessons", c.LessonController.Create)
	c.App.Get("/api/lessons/:lessonId", c.LessonController.Get)
	c.App.Put("/api/lessons/:lessonId", c.LessonController.Update)
	c.App.Delete("/api/lessons/:lessonId", c.LessonController.Delete)

	// Enrollments
	c.App.Post("/api/enrollments", c.EnrollmentController.Create)
	c.App.Get("/api/enrollments/:enrollmentId", c.EnrollmentController.Get)
	c.App.Put("/api/enrollments/:enrollmentId", c.EnrollmentController.Update)
	c.App.Delete("/api/enrollments/:enrollmentId", c.EnrollmentController.Delete)

	// Reviews
	c.App.Post("/api/reviews", c.ReviewController.Create)
	c.App.Get("/api/reviews/:reviewId", c.ReviewController.Get)
	c.App.Put("/api/reviews/:reviewId", c.ReviewController.Update)
	c.App.Delete("/api/reviews/:reviewId", c.ReviewController.Delete)

	// Roles
	c.App.Get("/api/roles", c.RoleController.List)
	c.App.Post("/api/roles", c.RoleController.Create)
	c.App.Put("/api/roles/:roleId", c.RoleController.Update)
	c.App.Get("/api/roles/:roleId", c.RoleController.Get)
	c.App.Delete("/api/roles/:roleId", c.RoleController.Delete)

	// Transactions
	c.App.Post("/api/transactions", c.TransactionController.Create)
	c.App.Get("/api/transactions/:transactionId", c.TransactionController.Get)
	c.App.Put("/api/transactions/:transactionId", c.TransactionController.Update)
	c.App.Delete("/api/transactions/:transactionId", c.TransactionController.Delete)

	// User roles
	c.App.Post("/api/user_roles", c.UserRoleController.Create)
	c.App.Get("/api/users/:userId/roles", c.UserRoleController.ListByUser)
	c.App.Delete("/api/user_roles/:userRoleId", c.UserRoleController.Delete)
}

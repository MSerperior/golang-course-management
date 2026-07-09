package config

import (
	"golang-clean-architecture/db/migrations"
	"golang-clean-architecture/internal/delivery/http"
	"golang-clean-architecture/internal/delivery/http/middleware"
	"golang-clean-architecture/internal/delivery/http/route"
	"golang-clean-architecture/internal/gateway/messaging"
	"golang-clean-architecture/internal/repository"
	"golang-clean-architecture/internal/usecase"

	"github.com/IBM/sarama"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
	Producer sarama.SyncProducer
}

func Bootstrap(config *BootstrapConfig) {
	migrations.Migrate(config.DB)

	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)
	contactRepository := repository.NewContactRepository(config.Log)
	addressRepository := repository.NewAddressRepository(config.Log)
	categoryRepository := repository.NewCategoryRepository(config.Log)
	courseRepository := repository.NewCourseRepository(config.Log)
	enrollmentRepository := repository.NewEnrollmentRepository(config.Log)
	lessonRepository := repository.NewLessonRepository(config.Log)
	reviewRepository := repository.NewReviewRepository(config.Log)
	roleRepository := repository.NewRoleRepository(config.Log)
	sectionRepository := repository.NewSectionRepository(config.Log)
	transactionRepository := repository.NewTransactionRepository(config.Log)
	userRoleRepository := repository.NewUserRoleRepository(config.Log)

	// setup producer
	var userProducer *messaging.UserProducer
	var contactProducer *messaging.ContactProducer
	var addressProducer *messaging.AddressProducer

	if config.Producer != nil {
		userProducer = messaging.NewUserProducer(config.Producer, config.Log)
		contactProducer = messaging.NewContactProducer(config.Producer, config.Log)
		addressProducer = messaging.NewAddressProducer(config.Producer, config.Log)
	}

	// setup use cases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository, userProducer)
	contactUseCase := usecase.NewContactUseCase(config.DB, config.Log, config.Validate, contactRepository, contactProducer)
	addressUseCase := usecase.NewAddressUseCase(config.DB, config.Log, config.Validate, contactRepository, addressRepository, addressProducer)
	categoryUseCase := usecase.NewCategoryUseCase(config.DB, config.Log, config.Validate, categoryRepository)
	courseUseCase := usecase.NewCourseUseCase(config.DB, config.Log, config.Validate, courseRepository)
	enrollmentUseCase := usecase.NewEnrollmentUseCase(config.DB, config.Log, config.Validate, enrollmentRepository)
	lessonUseCase := usecase.NewLessonUseCase(config.DB, config.Log, config.Validate, lessonRepository, sectionRepository)
	reviewUseCase := usecase.NewReviewUseCase(config.DB, config.Log, config.Validate, reviewRepository)
	roleUseCase := usecase.NewRoleUseCase(config.DB, config.Log, config.Validate, roleRepository)
	sectionUseCase := usecase.NewSectionUseCase(config.DB, config.Log, config.Validate, sectionRepository)
	transactionUseCase := usecase.NewTransactionUseCase(config.DB, config.Log, config.Validate, transactionRepository)
	userRoleUseCase := usecase.NewUserRoleUseCase(config.DB, config.Log, config.Validate, userRoleRepository, roleRepository)

	// setup controller
	userController := http.NewUserController(userUseCase, config.Log)
	contactController := http.NewContactController(contactUseCase, config.Log)
	addressController := http.NewAddressController(addressUseCase, config.Log)
	categoryController := http.NewCategoryController(categoryUseCase, config.Log)
	courseController := http.NewCourseController(courseUseCase, config.Log)
	enrollmentController := http.NewEnrollmentController(enrollmentUseCase, config.Log)
	sectionController := http.NewSectionController(sectionUseCase, config.Log)
	lessonController := http.NewLessonController(lessonUseCase, config.Log)
	reviewController := http.NewReviewController(reviewUseCase, config.Log)
	roleController := http.NewRoleController(roleUseCase, config.Log)
	transactionController := http.NewTransactionController(transactionUseCase, config.Log)
	userRoleController := http.NewUserRoleController(userRoleUseCase, config.Log)

	// setup middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	routeConfig := route.RouteConfig{
		App:                   config.App,
		UserController:        userController,
		ContactController:     contactController,
		AddressController:     addressController,
		CategoryController:    categoryController,
		CourseController:      courseController,
		SectionController:     sectionController,
		LessonController:      lessonController,
		EnrollmentController:  enrollmentController,
		ReviewController:      reviewController,
		RoleController:        roleController,
		TransactionController: transactionController,
		UserRoleController:    userRoleController,
		AuthMiddleware:        authMiddleware,
	}
	routeConfig.Setup()
}

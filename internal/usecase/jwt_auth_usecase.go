package usecase

import (
	"crypto/rand"
	"errors"
	"fmt"
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/repository"
	"math/big"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var JWTSecretKey = []byte("juara-coding-super-secret-key-2026-batch-1")

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type AuthService interface {
	Register(email, password, fullName string) (*entity.User, error)
	Login(email, password string) (string, error)
	Logout(tokenString string) error
	GetProfile(id uint) (*entity.User, error)
	VerifyOTP(email, otp string) error
}

type authService struct {
	DB            *gorm.DB
	Log           *logrus.Logger
	Validate      *validator.Validate
	userRepo      repository.UserRepository
	blacklistRepo repository.BlacklistedTokenRepository
	emailService  EmailService
}

// New Auth Service
func NewAuthService(
	userRepo repository.UserRepository,
	blacklistRepo repository.BlacklistedTokenRepository,
	emailService EmailService,
) *authService {
	return &authService{
		userRepo:      userRepo,
		blacklistRepo: blacklistRepo,
		emailService:  emailService,
	}
}

func generateOTP() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(900000))
	return fmt.Sprintf("%06d", n.Int64()+100000)
}

// Register
func (s *authService) Register(email, password, fullName string) (*entity.User, error) {
	// Check if email already exists
	existingUser, err := s.userRepo.FindByEmail(s.DB, email)
	if err == nil && existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	otpCode := generateOTP()

	// Create the new user
	user := entity.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     fullName,
	}

	// Save the user to the database
	err = s.userRepo.Create(s.DB, &user)
	if err != nil {
		return nil, err
	}

	go func() {
		_ = s.emailService.SendOTP(email, otpCode)
	}()

	return &user, nil
}

// Login
func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(s.DB, email)
	if err != nil {
		return "", errors.New("invalid email")
	}

	// Compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	claims := JWTClaims{
		UserID: 0,
		Email:  user.Email,
		Role:   "customer",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecretKey)
}

func (s *authService) Logout(tokenString string) error {
	return s.blacklistRepo.BlacklistToken(tokenString, time.Now().Add(24*time.Hour))
}

func (s *authService) GetProfile(id uint) (*entity.User, error) {
	var user entity.User
	err := s.userRepo.FindById(s.DB, &user, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *authService) VerifyOTP(email, otp string) error {
	_, err := s.userRepo.FindByEmail(s.DB, email)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	if otp == "" {
		return errors.New("kode OTP salah")
	}

	return nil
}

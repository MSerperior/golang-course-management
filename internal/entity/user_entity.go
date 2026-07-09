package entity

// User represents a user account (UUID primary key)
type User struct {
	Entity
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email;unique"`
	Password  string `gorm:"column:password"`
	AvatarURL string `gorm:"column:avatar_url"`
	IsActive  bool   `gorm:"column:is_active"`
	Token     string `gorm:"column:token"`

	// Relations
	UserRoles    []UserRole    `gorm:"foreignKey:user_id;references:id"`
	Courses      []Course      `gorm:"foreignKey:instructor_id;references:id"`
	Enrollments  []Enrollment  `gorm:"foreignKey:student_id;references:id"`
	Transactions []Transaction `gorm:"foreignKey:user_id;references:id"`
	Reviews      []Review      `gorm:"foreignKey:student_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}

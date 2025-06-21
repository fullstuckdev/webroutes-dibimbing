package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`

	// One to one with profile
	Profile *Profile `json:"profile,omitempty"`

	// One to many (1 users punya banyak postingan)
	Posts []Post `json:"posts,omitempty"`
}

// Profile model (one to one with users)
type Profile struct {
	gorm.Model
	UserID uint `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Bio string `json:"bio"`
}

// Post model (one to many with users)
type Post struct {
	gorm.Model
	UserID uint `json:"user_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	ExpiresAt *time.Time `json:"expires_at"`

	// Many to many
	Tags []Tag `json:"tags,omitempty" gorm:"many2many:post_tags"`
}

// Tag model (many to many with post)
type Tag struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type LoginRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

// Base Request
type CreateProfileRequest struct {
	FirstName string `json:"first_name" binding:"required"` // required artinya wajib dimasukan
	LastName string `json:"last_name" binding:"required"`  // required artinya wajib dimasukan
	Bio string `json:"bio"`
}

type CreatePostRequest struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CreateTagRequest struct {
	Name  string `json:"name" binding:"required"`
}

// Base Response
type APIResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}


func (p *Post) SetExpiration() {
	expirationDate := time.Now().AddDate(0, 0, 30) // 30 hari
	p.ExpiresAt = &expirationDate
}

func (p *Post) GetExpiration() *time.Time {
	return p.ExpiresAt
}
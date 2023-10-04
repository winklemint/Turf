package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"turf/config"
	"turf/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func AdminSignup(c *gin.Context) {
	var body struct {
		Name     string
		Contact  string
		Password string
		Email    string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "failed to read body",
			"data":   "null",
		})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "failed to hash password",
			"data":   "null",
		})
		return
	}
	bodys := models.Admin{
		Name:     body.Name,
		Contact:  body.Contact,
		Password: string(password),
		Email:    body.Email,
	}

	result := config.DB.Create(&bodys)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Admin Allready Exist",
			"data":   "null",
		})
		return
	}

	//Response
	c.JSON(http.StatusCreated, gin.H{
		"status": 201,

		"success": "Admin Successfully Created",
		"data":    bodys,
	})
}
func AdminLogin(c *gin.Context) {
	var body struct {
		Name     string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "failed to hash password",
			"data":   "null",
		})
		return
	}

	var admin models.Admin
	config.DB.Table("admins").Select("id", "name", "password").Where("name", body.Name).Scan(&admin)
	fmt.Println(admin)
	fmt.Println(admin.ID)
	if admin.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"Error":  "Admin Does Not Exist",
			"data":   "null",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Invalid Password",
			"data":   "null",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": admin.ID,
		"exp": time.Now().Add(time.Hour * 4).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Failed to create token",
			"data":   "null",
		})
		return
	}

	// send the generated jwt token back & set it in cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*4, "", "", false, true)
	admin.LastLogin = time.Now()
	config.DB.Save(&admin)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Admin Login Successfully",
	})
}

// func ChangePassword(c *gin.Context) {
// 	var body struct {
// 		NewPassword string
// 	}
// 	err := c.Bind(&body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": "400",
// 			"error":  "failed to hash password",
// 			"data":   "null",
// 		})
// 		return

// 	}
// 	password, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 14)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": "400",
// 			"error":  "failed to hash password",
// 			"data":   "null",
// 		})
// 		return
// 	}
// 	admins := models.Admin{

// 		Password: string(password),
// 	}

// 	result :=  config.DB.Model(&admins).Where("id = ?", shareholder.ID).Update("chip_wallet", amount)
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": "400",
// 			"error":  "Admin Allready Exist",
// 			"data":   "null",
// 		})
// 		return
// 	}

// 	//Response
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": 200,

// 		"success": "Admin Successfully Created",
// 		"data":    bodys,
// 	})
// }
func AddSlot(c *gin.Context) {
	var body struct {
		StartSlot string
		EndSlot   string
	}
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   "null",
		})
		return
	}
	// startSlot, err := time.Parse("15:04", body.StartSlot)
	// if err != nil {
	// 	fmt.Println("StartSlot parivartan mein error:", err)
	// 	return
	// }

	// endSlot, err := time.Parse("15:04", body.EndSlot)
	// if err != nil {
	// 	fmt.Println("EndSlot parivartan mein error:", err)
	// 	return
	// }
	slot := models.Slot{StartSlot: body.StartSlot, EndSlot: body.EndSlot}
	result := config.DB.Create(&slot)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Slot Allready Exist",
			"data":   "null",
		})
		return
	}

	//Response
	c.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"success": "Admin Successfully Created",
		"data":    slot,
	})

}

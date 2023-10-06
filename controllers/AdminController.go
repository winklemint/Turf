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
	slot := models.Time_Slot{Start_time: body.StartSlot, End_time: body.EndSlot}
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
		"success": "Slot Successfully Created",
		"data":    slot,
	})

}
func AddPackage(c *gin.Context) {
	var body struct {
		Name   string
		Price  float64
		Status bool
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
	packageModel := &models.Package{Name: body.Name, Price: body.Price, Status: body.Status}
	result := config.DB.Create(&packageModel)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "package Allready Exist",
			"data":   "null",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"success": "Package Successfully Created",
		"data":    body,
	})
}
func Package(c *gin.Context) {
	var body struct {
		Name   string
		Price  float64
		Status bool
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
	packageModel := &models.Package{Name: body.Name, Price: body.Price, Status: body.Status}
	result := config.DB.Create(&packageModel)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "package Allready Exist",
			"data":   "null",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"success": "Package Successfully Created",
		"data":    body,
	})
}
func UpdateAdmin(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
		Contact  string
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
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// decode & validate the same

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// find the user with token sub i.e user id
		var admin models.Admin
		config.DB.First(&admin, claims["sub"])

		if admin.ID == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		}

		result := config.DB.Find(&admin).Where("id = ?", claims["sub"])
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "400",
				"error":  "Admin Update UnSuccessfully",
				"data":   "null",
			})
			return
		}
		Hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "failed to hash password",
				"data":   "null",
			})
			return
		}

		admins := models.Admin{Name: body.Name, Email: body.Email, Contact: body.Contact, Password: string(Hash)}
		result = config.DB.Model(&admin).Where("id = ?", claims["sub"]).Updates(admins)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "400",
				"error":  "Admin Update UnSuccessfully",
				"data":   "null",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"success": "Admin Update Successfully",
			"data":    body,
		})

	}
}
func UpdateSlot(c *gin.Context) {
	Id := c.Param("id")
	var body struct {
		StartSlot string
		EndSlot   string
		Status    bool
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

	admin := models.Time_Slot{Start_time: body.StartSlot, End_time: body.EndSlot, Status: body.Status}
	result := config.DB.Model(&admin).Where("id = ?", Id).Updates(admin)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Slot Update UnSuccessfully",
			"data":   "null",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Slot Update Successfully",
		"data":    body,
	})

}
func GetAllSlot(c *gin.Context) {
	var slot []models.Time_Slot
	result := config.DB.Find(&slot)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 404,
			"error":  "failed to get all slot",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "slot details",
		"data":    slot,
	})

}

func UpdatePackage(c *gin.Context) {
	Id := c.Param("id")
	var body struct {
		Name   string ` grom:"unique"`
		Price  float64
		Status bool
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
	fmt.Println(body.Status)
	admin := models.Package{Name: body.Name, Price: body.Price, Status: body.Status}
	result := config.DB.Model(&admin).Where("id = ?", Id).Updates(admin)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Package Update UnSuccessfully",
			"data":   "null",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Package Update Successfully",
		"data":    body,
	})

}
func GetAllPackage(c *gin.Context) {
	var pkg []models.Package
	result := config.DB.Find(&pkg)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 404,
			"error":  "failed to get all slot",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "slot details",
		"data":    pkg,
	})
}

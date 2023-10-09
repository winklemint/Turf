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
		Role     int
		Status   string
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
	if body.Status == "Super Admin" {
		body.Role = 1
	} else {
		body.Role = 0
	}
	bodys := models.Admin{
		Name:     body.Name,
		Contact:  body.Contact,
		Password: string(password),
		Email:    body.Email,
		Role:     body.Role,
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

//			"success": "Admin Successfully Created",
//			"data":    bodys,
//		})
//	}
func Add_Branch(c *gin.Context) {
	var body struct {
		Turf_name             string
		Branch_name           string
		Branch_address        string
		Branch_email          string
		Branch_contact_number string
		GST_no                string
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
	branch := models.Branch_info_management{Turf_name: body.Turf_name, Branch_name: body.Branch_name, Branch_email: body.Branch_email, Branch_contact_number: body.Branch_contact_number, Branch_address: body.Branch_address, GST_no: body.GST_no}
	result := config.DB.Create(&branch)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Branch Already Exist",
			"data":   "null",
		})
		return
	}

	//Response
	c.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"success": "Branch Successfully Created",
		"data":    branch,
	})
}

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
		Status int
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

// func Package(c *gin.Context) {
// 	var body struct {
// 		Name   string
// 		Price  float64
// 		Status bool
// 	}
// 	err := c.Bind(&body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": 400,
// 			"error":  "failed to read body",
// 			"data":   "null",
// 		})
// 		return
// 	}
// 	packageModel := &models.Package{Name: body.Name, Price: body.Price, Status: body.Status}
// 	result := config.DB.Create(&packageModel)
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": "400",
// 			"error":  "package Allready Exist",
// 			"data":   "null",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{
// 		"status":  200,
// 		"success": "Package Successfully Created",
// 		"data":    body,
// 	})
// }
func UpdateAdmin(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
		Contact  string
		Role     string
		Status   int
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
		if body.Role == "Super Admin" {
			body.Status = 1
		} else {
			body.Status = 2
		}
		fmt.Println(admin.ID)
		admins := models.Admin{Name: body.Name, Email: body.Email, Contact: body.Contact, Password: string(Hash), Role: body.Status}
		result = config.DB.Model(&admin).Where("id = ?", admin.ID).Updates(admins)
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
		Status    int
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
		Status int
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
func GetConfirmBooking(c *gin.Context) {
	var Pkg []models.Confirm_Booking_Table
	var slot_id []int
	var Package []interface{}
	result := config.DB.Find(&Pkg)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 404,
			"error":  "failed to get all booking details",
		})
		return
	}

	for i := 0; i < len(Pkg); i++ {
		result := config.DB.Model(&models.Turf_Bookings{}).Where("order_id = ?", Pkg[i].Booking_order_id).Pluck("slot_id", &slot_id)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 404,
				"error":  "failed to get all detail",
			})
			return
		}

		var pkgSlots []interface{}

		for _, s := range slot_id {
			var slt models.Time_Slot
			result := config.DB.Where("id = ? ", s).Find(&slt)

			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Failed to find slot by start_slot",
				})
				return
			}

			// Create a map for slot details
			slotData := map[string]interface{}{
				"starttime": slt.Start_time,
				"endtime":   slt.End_time,
			}
			pkgSlots = append(pkgSlots, slotData)
		}
		Package = append(Package, Pkg[i], pkgSlots)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "confirmed booking details",
		"data":    Package,
	})
}
func UpdatecomfirmDetails(c *gin.Context) {
	Id := c.Param("id")
	var body struct {
		Paid_amount             float64
		Remaining_amount_to_pay float64
		Booking_status          string
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
	var Status int
	if body.Booking_status == "Confirm" {
		Status = 4

	} else {
		Status = 1
	}
	confirm_booking := models.Confirm_Booking_Table{Paid_amount: body.Paid_amount, Remaining_amount_to_pay: body.Remaining_amount_to_pay, Booking_status: Status}
	result := config.DB.Model(&models.Confirm_Booking_Table{}).Where("id = ?", Id).Updates(&confirm_booking)
	// result = config.DB.Exec(result)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "confirm table Update UnSuccessfully",
			"data":   "null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "confirm table Update Successfully",
		"data":    confirm_booking,
	})

}
func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := config.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to load user details",
			"data":   "null",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "success to load user details",
		"data":    users,
	})
}
func UpdateUserDetails(c *gin.Context) {
	Id := c.Param("id")
	var body struct {
		Full_Name      string
		Email          string
		Password       string
		Contact        string
		Account_Status string
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
	Hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})
		return
	}
	var Account_Status int
	switch {
	case body.Account_Status == "Active":
		Account_Status = 1
	case body.Account_Status == "Hold":
		Account_Status = 2
	case body.Account_Status == "Block":
		Account_Status = 3
	default:
		Account_Status = 0
	}

	users := models.User{Full_Name: body.Full_Name, Email: body.Email, Contact: body.Contact, Password: string(Hash), Account_Status: Account_Status}
	result := config.DB.Model(&users).Where("id = ?", Id).Updates(users)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "User Update UnSuccessfully",
			"data":   "null",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "User Update Successfully",
		"data":    body,
	})

}

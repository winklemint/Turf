package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"turf/config"
	"turf/models"

	"github.com/dariubs/percent"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func AdminSignup(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Name     string
		Contact  string
		Password string
		Email    string
		Role     int

		Branch_name   int
		Authorization string
	}
	if c.Bind(&body) != nil {
		logrus.Infoln("Failed t0 read b0dy f0r AdminSignup")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Read Body",
			"data":   nil,
		})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		logrus.Infoln("Admin Signup", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Hashed Password",
			"data":   nil,
		})
		return
	}

	// var branch models.Branch_info_management
	// result := config.DB.Find(&branch, "branch_name=?", body.Branch_name)
	// if result.Error != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status": 400,
	// 		"error":  "Error finding branch id",
	// 		"data":   nil,
	// 	})
	// 	return
	// }

	BranchID := uint(body.Branch_name)

	fmt.Println(BranchID)

	bodys := models.Admin{
		Name:           body.Name,
		Contact:        body.Contact,
		Password:       string(password),
		Email:          body.Email,
		Role:           body.Role,
		Turf_branch_id: BranchID,

		Authorization: body.Authorization,
	}

	result := config.DB.Create(&bodys)
	if result.Error != nil {
		logrus.Infoln("admin signup-", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Admin Allready Exist",
			"data":   nil,
		})
		return
	}

	//Response
	c.JSON(http.StatusCreated, gin.H{
		"status":  201,
		"success": "Admin Successfully Created",
		"data":    bodys,
	})
}
func AdminUpdateById(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	// Create an instance of the model outside of if-else scope
	var adminToUpdate models.Admin
	var body struct {
		Name          string
		Contact       string
		Password      string
		Email         string
		Role          int
		Branch_name   int
		Authorization string
	}
	if err := c.Bind(&body); err != nil {
		logrus.Infoln("failed t0 read b0dy AdminUpdateById")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Read Body",
			"data":   nil,
		})
		return
	}
	// Check if 'Password' is provided, and hash it if necessary
	if body.Password != "" {
		// Hash the password using bcrypt
		password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
		if err != nil {
			logrus.Infoln("AdminUpdateById-", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "failed to hash password",
				"data":   nil,
			})
			return
		}
		adminToUpdate.Password = string(password)
	}
	// Update the other fields
	adminToUpdate.Name = body.Name
	adminToUpdate.Contact = body.Contact
	adminToUpdate.Email = body.Email
	adminToUpdate.Role = body.Role
	adminToUpdate.Turf_branch_id = uint(body.Branch_name)
	adminToUpdate.Authorization = body.Authorization
	// Update the admin using the provided 'id'
	result := config.DB.Model(&models.Admin{}).Where("id = ?", id).Updates(&adminToUpdate)
	if result.Error != nil {
		logrus.Infoln("AdminUpdateById - ", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Error Finding Branch Id",
			"data":   nil,
		})
		return
	}
	BranchID := uint(body.Branch_name)
	fmt.Println(BranchID)
	bodys := models.Admin{
		Name:           body.Name,
		Contact:        body.Contact,
		Email:          body.Email,
		Role:           body.Role,
		Turf_branch_id: BranchID,
		Authorization:  body.Authorization,
	}
	result = config.DB.Model(&bodys).Where("id=?", id).Updates(&bodys)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Updated Admin Details",
			"data":   nil,
		})
		return
	}
	//Response
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Successfully Updated Admin Details",
		"data":    bodys,
	})
}
func AdminGetById(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var staff models.Admin
	result := config.DB.Find(&staff, "id=?", id)
	if result.Error != nil {
		logrus.Infoln("AdminGetById - ", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Get Admin Details",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"Success": "Successfully Get Admin Details",
		"Data":    staff,
	})

}
func AdminProfile(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	// var response struct {
	// 	Data []interface{}
	// }
	adminId, _ := c.Request.Cookie("AID")
	id, _ := strconv.Atoi(adminId.Value)

	// Initialize Admin and Branch models
	var admin []models.Admin

	// Create a database query to join Admin and Branch tables and select only branch_name
	result := config.DB.Debug().Raw(`SELECT admins.name, admins.contact, admins.email, admins.role, admins.turf_branch_id, admins.id, branch_info_managements.branch_name FROM admins JOIN branch_info_managements ON admins.turf_branch_id = branch_info_managements.id WHERE admins.id = ?`, id).Scan(&admin)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Get Admin Details",
			"data":   nil,
		})
		return
	}
	// for _, packageData := range admin {
	// 	response.Data = append(response.Data, packageData)
	// }

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"Success": "Successfully Get Admin Details",
		"Data":    admin,
	})
}
func UpdateProfile(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	adminId, _ := c.Request.Cookie("AID")
	id, _ := strconv.Atoi(adminId.Value)
	var body struct {
		Name    string
		Contact string
		Email   string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Read Body",
			"data":   nil,
		})
		return
	}

	bodys := models.Admin{
		Name:    body.Name,
		Contact: body.Contact,
		Email:   body.Email,
	}

	result := config.DB.Model(&bodys).Where("id=?", id).Updates(&bodys)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Admin Allready Exist",
			"data":   nil,
		})
		return
	}

	//Response
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Admin Successfully Updated",
		"data":    bodys,
	})

}
func AdminDelete(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var admin models.Admin
	result := config.DB.Model(&admin).Where("id=?", id).Delete(&admin)
	if result.Error != nil {
		logrus.Infoln("AdminDelete - ", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Unsuccessfully Deleted Staff",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Successfully Deleted Staff",
		"data":    nil,
	})
}

func GetConfirmBookingTop5(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var data []models.Confirm_Booking_Table
	// role, _ := c.Request.Cookie("Role")
	// Role, _ := strconv.Atoi(role.Value)
	branchID, _ := c.Request.Cookie("Branch_id")
	branchid, _ := strconv.Atoi(branchID.Value)

	// if Role != 1 {

	result := config.DB.Model(&models.Confirm_Booking_Table{}).Limit(5).Order("ID DESC").Where("branch_id=?", branchid).Find(&data)

	if result.Error != nil {
		logrus.Infoln("GetC0nfirmB00kingT0p5 - ", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 404,
			"error":  "failed to get all branch details ",
		})
		return

	}
	// } else {

	// 	result := config.DB.Model(&models.Confirm_Booking_Table{}).Limit(5).Order("ID DESC").Find(&data)
	// 	if result.Error != nil {
	// 		logrus.Infoln("GetC0nfirmB00kingT0p5 - ", result.Error)
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"status": 404,
	// 			"error":  "failed to get confirmed booking details",
	// 		})
	// 		return
	// 	}
	// }
	var responseData []interface{}
	for _, booking := range data {
		var user models.User
		result := config.DB.Find(&user, booking.User_id)
		if result.Error != nil {
			logrus.Infoln("GetC0nfirmB00kingT0p5 - ", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 404,
				"error":  "Failed To User Name",
				"data":   nil,
			})
			return
		}
		var branch models.Branch_info_management
		result = config.DB.Find(&branch, booking.Branch_id)
		if result.Error != nil {
			logrus.Infoln("GetC0nfirmB00kingT0p5 - ", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 404,
				"error":  "Failed To Fetch  Branch Name",
				"data":   nil,
			})
			return
		}
		bookingData := map[string]interface{}{
			"ID":                      booking.ID,
			"CreatedAt":               booking.CreatedAt,
			"User_id":                 booking.User_id,
			"User_name":               user.Full_Name,
			"Date":                    booking.Date,
			"Booking_order_id":        booking.Booking_order_id,
			"Total_price":             booking.Total_price,
			"Total_min_amount_to_pay": booking.Total_min_amount_to_pay,
			"Paid_amount":             booking.Paid_amount,
			"Remaining_amount_to_pay": booking.Remaining_amount_to_pay,
			"Booking_status":          booking.Booking_status,
			"Branch_name":             branch.Branch_name,
		}
		responseData = append(responseData, bookingData)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Confirmed Booking Details",
		"data":    responseData,
	})
}

// func GetConfirmBookingTop5Super(c *gin.Context) {
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
// 	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
// 	if c.Request.Method == "OPTIONS" {
// 		c.JSON(http.StatusOK, gin.H{})
// 		return
// 	}

// 	var data []models.Confirm_Booking_Table
// 	result := config.DB.Model(&models.Confirm_Booking_Table{}).Limit(5).Order("ID DESC").Find(&data)
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": 404,
// 			"error":  "failed to get confirmed booking details",
// 		})
// 		return
// 	}
// 	var responseData []interface{}
// 	for _, booking := range data {
// 		var user models.User
// 		result := config.DB.Find(&user, booking.User_id)
// 		if result.Error != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status": 404,
// 				"error":  "failed to user name",
// 			})
// 			return
// 		}
// 		var branch models.Branch_info_management
// 		result = config.DB.Find(&branch, booking.Branch_id)
// 		if result.Error != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status": 404,
// 				"error":  "failed to fetch  branch name",
// 			})
// 			return
// 		}
// 		bookingData := map[string]interface{}{
// 			"ID":                      booking.ID,
// 			"CreatedAt":               booking.CreatedAt,
// 			"User_id":                 booking.User_id,
// 			"User_name":               user.Full_Name,
// 			"Date":                    booking.Date,
// 			"Booking_order_id":        booking.Booking_order_id,
// 			"Total_price":             booking.Total_price,
// 			"Total_min_amount_to_pay": booking.Total_min_amount_to_pay,
// 			"Paid_amount":             booking.Paid_amount,
// 			"Remaining_amount_to_pay": booking.Remaining_amount_to_pay,
// 			"Booking_status":          booking.Booking_status,
// 			"Branch_name":             branch.Branch_name,
// 		}
// 		responseData = append(responseData, bookingData)
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  200,
// 		"success": "confirmed booking details",
// 		"data":    responseData,
// 	})
// }

func AdminLogin(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Name     string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		logrus.Infoln("failed t0 read b0dy admin l0gin")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Hashed  Password",
			"data":   nil,
		})
		return
	}

	var admin models.Admin

	result := config.DB.Find(&admin, "name=?", body.Name)
	if admin.ID == 0 || result.Error != nil {
		logrus.Infoln("Admin l0gin - ", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"Error":  "Admin Does Not Exist",
			"data":   nil,
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(body.Password))

	if err != nil {
		logrus.Infoln(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Invalid Password",
			"data":   nil,
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": admin.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		logrus.Infoln(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Create Token",
			"data":   nil,
		})
		return
	}

	adminIDString := strconv.FormatUint(uint64(admin.ID), 10)
	adminRole := strconv.Itoa(admin.Role)
	adminBranch := strconv.Itoa(int(admin.Turf_branch_id))

	// send the generated jwt token back & set it in cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 7200, "", "", true, true)
	c.SetCookie("AID", adminIDString, 7200, "", "", true, true)
	c.SetCookie("Role", adminRole, 7200, "", "", true, true)
	c.SetCookie("branch_id", adminBranch, 7200, "", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Admin Login Successfully",
		"data":    admin,
	})
}

// func ChangePassword(c *gin.Context) {
// 	var body struct {
// 		NewPassword string
// 	}
// 	err := c.Bind(&body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": 400,
// 			"error":  "failed to hash password",
// 			"data":   nil,
// 		})
// 		return

// 	}
// 	password, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 14)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": 400,
// 			"error":  "failed to hash password",
// 			"data":   nil,
// 		})
// 		return
// 	}
// 	admins := models.Admin{

// 		Password: string(password),
// 	}

// 	result :=  config.DB.Model(&admins).Where("id = ?", shareholder.ID).Update("chip_wallet", amount)
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": 400,
// 			"error":  "Admin Allready Exist",
// 			"data":   nil,
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
// func Select_branch(c *gin.Context) {
// 	var branch models.Branch_info_management
// 	result := config.DB.Select("branch_name").Find(&branch)
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": 400,
// 			"error":  "Failed to get branch",
// 			"data":   nil,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  200,
// 		"success": "Branch details",
// 		"data":    branch.Branch_name,
// 	})
// 	return

// }

func Add_Branch(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Turf_name             string
		Branch_name           string
		Branch_address        string
		Branch_email          string
		Branch_contact_number string
		GST_no                string
		Status                int
		Ground_Size           string
		Image                 string
	}

	err := c.Bind(&body)
	if err != nil {
		logrus.Infoln(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Read Body",
			"data":   nil,
		})
		return
	}
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filePath := filepath.Join("./uploads/branch", file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	if filepath.Ext(filePath) != ".jpg" && filepath.Ext(filePath) != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Upload The Right File Format (jpg or png)",
			"data":   nil,
		})
		return
	}
	branch := models.Branch_info_management{Turf_name: body.Turf_name, Branch_name: body.Branch_name, Branch_email: body.Branch_email, Branch_contact_number: body.Branch_contact_number, Branch_address: body.Branch_address, GST_no: body.GST_no, Status: body.Status, Ground_Size: body.Ground_Size, Image: filePath}
	result := config.DB.Create(&branch)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Branch Already Exist",
			"data":   nil,
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
func Update_Branch(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var body struct {
		Turf_name             string
		Branch_name           string
		Branch_address        string
		Branch_email          string
		Branch_contact_number string
		GST_no                string
		Status                int
		Image                 string
	}
	if c.Bind(&body) != nil {
		logrus.Infoln("failed t0 read b0dy Update_Branch")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Read Body",
			"data":   nil,
		})
		return

	}

	branch := models.Branch_info_management{Turf_name: body.Turf_name, Branch_name: body.Branch_name, Branch_email: body.Branch_email, Branch_contact_number: body.Branch_contact_number, Branch_address: body.Branch_address, GST_no: body.GST_no, Status: body.Status}
	result := config.DB.Model(&branch).Where("id=?", id).Updates(&branch)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Branch Updated Unsuccessfully",
			"data":   nil,
		})
		return
	}

	if body.Image != "" {

		file, err := c.FormFile("image")
		if err != nil {
			logrus.Infof("Failed to read image fr0m DB %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filePath := filepath.Join("./uploads/branch", file.Filename)

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			logrus.Infof("Failed to save file in DB %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}
		if filepath.Ext(filePath) != ".jpg" && filepath.Ext(filePath) != ".png" {
			logrus.Infoln("Upload the right file format (jpg or png) Update branch func")
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Upload The Right File Format (jpg or png)",
				"data":   nil,
			})
			return
		}

		fmt.Println(filePath)

		branch = models.Branch_info_management{Image: filePath}
		result = config.DB.Model(&branch).Where("id=?", id).Updates(&branch)
		if result.Error != nil {
			logrus.Infof("Failed to get data from DB %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Branch Update Unsuccessfully",
				"data":   nil,
			})
			return
		}
	} else {
		fmt.Println("n image")

		c.JSON(http.StatusCreated, gin.H{
			"status":  200,
			"success": "Branch Successfully Updated",
			"data":    branch,
		})
	}

}

func GET_All_Branch(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var branch []models.Branch_info_management
	result := config.DB.Find(&branch)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "No Branch Found",
			"data":   nil,
		})
		return
	}

	//Response
	c.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"success": "Get All Branch Successfully",
		"data":    branch,
	})

}
func ActiveBranch(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var branch []models.Branch_info_management
	result := config.DB.Find(&branch, "status=1")
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "No Branch Found",
			"data":   nil,
		})
		return
	}

	//Response
	c.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"success": "All Branch  Successfully",
		"data":    branch,
	})

}
func Get_IdBy_Branch_NAme(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Branch_Name string
	}
	if c.Bind(&body) != nil {
		logrus.Infoln("failed t0 read b0dy Get_IdBy_Branch_NAme")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Invalid Request Body ",
			"data":   nil,
		})
	}

	var branches models.Branch_info_management
	result := config.DB.Find(&branches, "branch_name=?", body.Branch_Name)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "Record Not Found",
			"data":   nil,
		})
	}

	Id := strconv.FormatUint(uint64(branches.ID), 10)
	c.SetCookie("Branch_id", Id, 7200, "", "", false, false)

}
func GET_All_Branch_Id(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	Id := c.Param("id")
	var branch models.Branch_info_management
	result := config.DB.Find(&branch, "id=?", Id)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "No Branch Found",
			"data":   nil,
		})
		return
	}

	//Response
	c.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"success": "Get Branch Details Successfully",
		"data":    branch,
	})

}
func Delete_Branch(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	Id := c.Param("id")
	var branch models.Branch_info_management
	result := config.DB.Model(&branch).Where("id=?", Id).Delete(&branch)
	if result.Error != nil {
		logrus.Infof("Failed to delete data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Unsuccessfully Deleted Branch",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Successfully Deleted Branch",
		"data":    nil,
	})
}
func GetBranchimagesById(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")

	var branch models.Branch_info_management
	result := config.DB.Find(&branch, "id=?", id)

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Fetch Branch Image",
			"data":   nil,
		})
		return
	}

	// Determine the file path based on the file format (you may need to store this information in your model)
	var filePath string
	if strings.HasSuffix(branch.Image, ".jpg") {
		filePath = branch.Image
		c.Header("Content-Type", "image/jpeg")
	} else if strings.HasSuffix(branch.Image, ".png") {
		filePath = branch.Image
		c.Header("Content-Type", "image/png")
	} else {
		// Handle unsupported image formats
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Unsupported Image Format",
			"data":   nil,
		})
		return
	}

	// Read the image file
	imageData, err := ioutil.ReadFile(filePath)
	if err != nil {
		logrus.Infof("Failed t0 read image %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"error":  "Internal Server Error",
			"data":   nil,
		})
		return
	}

	// Send the image data as the response
	c.Data(http.StatusOK, c.GetHeader("Content-Type"), imageData)
}
func AddSlot(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		StartSlot string
		EndSlot   string
		Day       string
		Branch_id int
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed to read b0dy %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Read Body",
			"data":   nil,
		})
		return
	}

	// Check if a slot with the same attributes already exists
	var existingSlot models.Time_Slot
	result := config.DB.Where("start_time = ? AND end_time = ? AND day = ? AND branch_id = ?", body.StartSlot, body.EndSlot, body.Day, body.Branch_id).First(&existingSlot)
	if result.Error == nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Slot Already Exists",
			"data":   nil,
		})
		return
	}

	var slots models.Time_Slot
	result = config.DB.Find(&slots)
	if result.Error == nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Slot Already Exists",
			"data":   nil,
		})
		return
	}

	var branch models.Branch_info_management
	result = config.DB.Find(&branch, "id = ?", body.Branch_id)
	if result.Error == nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Slot already exists",
			"data":   nil,
		})
		return
	}

	First_three_initials := body.Day[:3]

	usid := First_three_initials + "/" + body.StartSlot + "/" + body.EndSlot

	slot := models.Time_Slot{Start_time: body.StartSlot, End_time: body.EndSlot, Day: body.Day, Branch_id: branch.ID, Unique_slot_id: usid, Status: 1}
	result = config.DB.Create(&slot)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Slot Allready Exist",
			"data":   nil,
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

// add package by admin
func AddPackage(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Name      string
		Price     float64
		Status    int
		Branch_id int
		Slot_id   []string
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed to read b0dy %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Read Body",
			"data":   nil,
		})
		return
	}

	var packages models.Package

	result := config.DB.Find(&packages, "name=? AND branch_id=?", body.Name, body.Branch_id)
	if packages.ID != 0 {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "PAckage already exist",
			"data":   nil,
		})
		return
	}

	packageModel := &models.Package{Name: body.Name, Price: body.Price, Status: body.Status, Branch_id: body.Branch_id}
	result = config.DB.Create(&packageModel)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed t0 create package",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"success": "Package Successfully Created",
		"data":    packageModel,
	})
	Last_insert_id := packageModel.ID

	fmt.Println(Last_insert_id)

	// var slots []interface{}
	for i := 0; i < len(body.Slot_id); i++ {
		psrmodel := models.Package_slot_relationship{Package_id: Last_insert_id, Slot_id: body.Slot_id[i]}
		result = config.DB.Create(&psrmodel)
		if result.Error != nil {
			logrus.Infof("Failed to get data from DB %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Relation Failed To Create",
				"data":   nil,
			})
			// slots = append(slots, psrmodel.Slot_id)
		}
	}
}

func UpdateAdmin(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Name      string
		Email     string
		Password  string
		Contact   string
		Role      string
		Status    int
		Branch_Id int
		Branch    string
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed to read b0dy %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Read Body",
			"data":   nil,
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
		result := config.DB.First(&admin, claims["sub"])
		if result.Error != nil {
			logrus.Infof("Failed to get data from DB %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Admin Update UnSuccessfully",
				"data":   nil,
			})
			return
		}

		if admin.ID == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		}

		result = config.DB.Find(&admin).Where("id = ?", claims["sub"])
		if result.Error != nil {
			logrus.Infof("Failed to get data from DB %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Admin Update UnSuccessfully",
				"data":   nil,
			})
			return
		}

		Hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			logrus.Infof("Failed to hash %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Failed To Hashed Password",
				"data":   nil,
			})
			return
		}

		if body.Role == "Super Admin" {
			body.Status = 1
		} else {
			body.Status = 2
		}
		// if body.Branch != "" {
		// 	if admin.Role == 1 {
		// 		var branch models.Branch_info_management
		// 		result := config.DB.Find(&branch).Where("id=?", body.Branch)
		// 		if result.Error != nil {
		// 			c.JSON(http.StatusBadRequest, gin.H{
		// 				"status": 400,
		// 				"error":  "failed to fetch brach detail",
		// 				"data":   nil,
		// 			})
		// 			return
		// 		}
		// 	} else {
		// 		c.JSON(http.StatusBadRequest, gin.H{
		// 			"status": 400,
		// 			"error":  "You are not authorised for update branch",
		// 			"data":   nil,/Update/staff
		// 		})
		// 		return

		// 	}
		// }

		admins := models.Admin{Name: body.Name, Email: body.Email, Contact: body.Contact, Password: string(Hash), Role: body.Status}
		result = config.DB.Model(&admin).Where("id = ?", admin.ID).Updates(admins)
		if result.Error != nil {
			logrus.Infof("Failed to get data from DB %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Admin Update UnSuccessfully",
				"data":   nil,
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
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	Id := c.Param("id")
	var body struct {
		StartSlot string
		EndSlot   string
		Status    int
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed to read b0dy %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Read Body",
			"data":   nil,
		})
		return
	}

	admin := models.Time_Slot{Start_time: body.StartSlot, End_time: body.EndSlot, Status: body.Status}
	result := config.DB.Model(&admin).Where("id = ?", Id).Updates(admin)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Slot Update UnSuccessfully",
			"data":   nil,
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
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var slot []models.Time_Slot

	// role, _ := c.Request.Cookie("Role")
	// Role, _ := strconv.Atoi(role.Value)
	branchID, _ := c.Request.Cookie("Branch_id")
	branchid, _ := strconv.Atoi(branchID.Value)
	// if Role != 1 {

	result := config.DB.Find(&slot, "branch_id=?", branchid)

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 404,
			"error":  "failed to get all slot",
		})
		return

	}
	// } else {
	// 	result := config.DB.Find(&slot)

	// 	if result.Error != nil {
	// 		logrus.Infof("Failed to get data from DB %v\n", result.Error)
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"status": 404,
	// 			"error":  "failed to get all slot",
	// 		})
	// 		return

	// 	}
	// }

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Slot Details",
		"data":    slot,
	})

}

func Get_Slot_by_day(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Day []string
	}

	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed to read b0dy %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Read body",
			"data":   nil,
		})
		return
	}

	// Create a map to group time slots by day
	days := make(map[string][]models.Time_Slot)

	// Assuming you have a database connection configured in config.DB
	for i := 0; i < len(body.Day); i++ {

		var slot []models.Time_Slot

		//var psr_ID int

		result := config.DB.Debug().Model(&models.Time_Slot{}).
			Select("time_slots.id, time_slots.start_time, time_slots.end_time, time_slots.day, time_slots.unique_slot_id, time_slots.branch_id, package_slot_relationships.id as psr_id").
			Joins("LEFT JOIN package_slot_relationships ON time_slots.id = package_slot_relationships.slot_id").
			Where("time_slots.day = ?", body.Day[i]).
			Scan(&slot)

		if result.Error != nil {
			logrus.Infof("Failed to Get Slot By Day %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Failed To Get Slots For " + body.Day[i],
				"data":   nil,
			})
			return
		}

		fmt.Println(slot)
		days[body.Day[i]] = slot
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Slot Details By Day",
		"data":    days,
	})
}

func UpdatePackage(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	Id := c.Param("id")
	var body struct {
		Name      string ` grom:"unique"`
		Price     float64
		Status    int
		Branch_id int
		Slot_id   []string
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed to Read Body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Read Body",
			"data":   nil,
		})
		return
	}

	admin := models.Package{Name: body.Name, Price: body.Price, Status: body.Status}
	result := config.DB.Model(&admin).Where("id = ?", Id).Updates(admin)
	if result.Error != nil {
		logrus.Infof("Failed to Upadte Package Details %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{

			"status": 400,
			"error":  "Package Update Unsuccessful",
			"data":   nil,
		})
		return
	}
	ID, _ := strconv.ParseUint(Id, 10, 0)

	IDuint := uint(ID)
	//var psr models.Package_slot_relationship

	config.DB.Exec("DELETE FROM package_slot_relationships WHERE package_id = ? ", Id)

	for i := 0; i < len(body.Slot_id); i++ {
		psr := models.Package_slot_relationship{Package_id: IDuint, Slot_id: body.Slot_id[i]}
		result := config.DB.Create(&psr)
		if result.Error != nil {
			logrus.Infof("Failed to Upadte Package Details %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Package Update UnSuccessfully",
				"data":   nil,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Package Updated Successfully",
		"data":    admin,
	})

}
func AllStaff(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var staff []models.Admin
	result := config.DB.Find(&staff, "role=2 OR role=3")
	if result.Error != nil {
		logrus.Infof("Failed to Get All Admin %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Get All Admin",
			"data":   nil,
		})
		return
	}
	var response []interface{}
	for _, stf := range staff {
		var branch models.Branch_info_management
		result := config.DB.Find(&branch, stf.Turf_branch_id)
		if result.Error != nil {
			logrus.Infof("Failed to Fetch Branch Name %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 404,
				"error":  "Failed To Fetch Branch Name",
				"data":   nil,
			})
			return
		}
		staffdata := map[string]interface{}{
			"ID":          stf.ID,
			"Name":        stf.Name,
			"Contact":     stf.Contact,
			"Email":       stf.Email,
			"Branch_name": branch.Branch_name,
			"Role":        stf.Role,

			"Turf_branch_id": stf.Turf_branch_id,
			"Authorization":  stf.Authorization,
		}
		response = append(response, staffdata)

	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"Success": " All  Admin",
		"data":    response,
	})

}

func Get_Package(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")

	var Slot_id []string
	var psr []models.Package_slot_relationship

	var pack []models.Package
	config.DB.Where("id = ?", id).Find(&pack)
	fmt.Println(pack)
	config.DB.Where("package_id = ?", id).Find(&psr)

	for i := 0; i < len(psr); i++ {
		Slot_id = append(Slot_id, psr[i].Slot_id)
	}

	response := struct {
		Status  int    `json:"status"`
		Success string `json:"success"`
		Data    struct {
			Pack []models.Package `json:"pack"`
			Slot []string         `json:"slot"`
		} `json:"data"`
	}{
		Status:  200,
		Success: "Data Retrieved Successfully",
		Data: struct {
			Pack []models.Package `json:"pack"`
			Slot []string         `json:"slot"`
		}{
			Pack: pack,
			Slot: Slot_id,
		},
	}

	c.JSON(http.StatusOK, response)
}
func DeleteSlot(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")

	// Create a raw SQL query to delete the record by ID.
	sqlQuery := "DELETE FROM time_slots WHERE id = ?"
	config.DB.Exec(sqlQuery, id)

	// if err != nil {
	// 	// Handle the error if the SQL query fails.
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"status": 500,
	// 		"error":  "Failed to delete slot",
	// 		"data":   nil,
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Successfully Deleted Slot",
		"data":    nil,
	})
}

func GetAllPackage(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	// role, _ := c.Request.Cookie("Role")
	// Role, _ := strconv.Atoi(role.Value)
	branchID, _ := c.Request.Cookie("Branch_id")
	branchid, _ := strconv.Atoi(branchID.Value)
	var Pkg []models.Package
	// if Role != 1 {

	result := config.DB.Find(&Pkg, "branch_id=?", branchid)

	if result.Error != nil {
		logrus.Infof("Failed to Get All Branch Details %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Get All Branch Details ",
			"data":   nil,
		})
		return

	}

	// } else {

	// 	result := config.DB.Find(&Pkg)

	// 	if result.Error != nil {
	// 		logrus.Infof("Failed to GET Package Details %v\n", result.Error)
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"status": 400,
	// 			"error":  "Failed To Get Package Details",
	// 		})
	// 		return

	// 	}

	// }
	var response []map[string]interface{}
	for _, p := range Pkg {
		var branch models.Branch_info_management
		result := config.DB.Find(&branch, "id=?", branchid)

		if result.Error != nil {
			logrus.Infof("Failed to Get All Branch Details %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Failed To Get All Branch Details",
				"Data":   nil,
			})
			return

		}
		data := map[string]interface{}{
			"ID":        p.ID,
			"Name":      p.Name,
			"Price":     p.Price,
			"Status":    p.Status,
			"Branch_id": p.Branch_id,

			"Branch_name": branch.Branch_name, // Fill in with the actual value
		}
		response = append(response, data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Package Details",
		"data":    response, // Use the response data you built
	})
}
func GetAllPackageById(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	Id := c.Param("id")
	var pkg models.Package
	result := config.DB.Find(&pkg).Where("id=?", Id)

	if result.Error != nil {
		logrus.Infof("Failed to Get Package Detail %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Get Package",
			"data":   nil,
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Package Details",
		"data":    pkg,
	})
}
func DeletePackage(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	Id := c.Param("id")
	var packages models.Package
	result := config.DB.Model(&packages).Where("id=?", Id).Delete(&packages)
	if result.Error != nil {
		logrus.Infof("Failed to Delete Package %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Delete Package",
			"data":   nil,
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Delete Package Successfully",
		"data":    packages,
	})

}
func GetConfirmBooking(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var Pkg []models.Confirm_Booking_Table
	var slot_id []int
	var Package []interface{}
	result := config.DB.Find(&Pkg)
	if result.Error != nil {
		logrus.Infof("Failed to Get Confirm Booking Details %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 404,
			"error":  "Failed To Get All Booking Details",
			"data":   nil,
		})
		return
	}

	for i := 0; i < len(Pkg); i++ {
		result := config.DB.Model(&models.Turf_Bookings{}).Where("order_id = ?", Pkg[i].Booking_order_id).Pluck("slot_id", &slot_id)

		if result.Error != nil {
			logrus.Infof("Failed to Get slot Details %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Failed To Get Slot Detail",
				"data":   nil,
			})
			return
		}

		var pkgSlots []interface{}

		for _, s := range slot_id {
			var slt models.Time_Slot
			result := config.DB.Where("id = ? ", s).Find(&slt)

			if result.Error != nil {
				logrus.Infof("Failed To Find Slot By Start_Slot  %v\n", result.Error)
				c.JSON(http.StatusBadRequest, gin.H{
					"status": 400,
					"error":  "Failed To Find Slot By Start_Slot",
					"data":   nil,
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
		"success": "Confirmed Booking Details",
		"data":    Package,
	})
}

func UpdatecomfirmDetails(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	Id := c.Param("id")
	var body struct {
		Paid_amount             float64
		Remaining_amount_to_pay float64
		Booking_status          int
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed To Read Body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed To Read Body",
			"data":   nil,
		})
		return
	}
	// var Status int
	// if body.Booking_status == "Confirm" {
	// 	Status = 4

	// } else {
	// 	Status = 1
	// }
	confirm_booking := models.Confirm_Booking_Table{Paid_amount: body.Paid_amount, Remaining_amount_to_pay: body.Remaining_amount_to_pay, Booking_status: body.Booking_status}
	result := config.DB.Model(&models.Confirm_Booking_Table{}).Where("id = ?", Id).Updates(&confirm_booking)
	// result = config.DB.Exec(result)
	if result.Error != nil {
		logrus.Infof("Failed To Update Confirm Booking Details %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Confirm Table Update UnSuccessfully",
			"data":   nil,
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
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")

	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var users []models.User
	result := config.DB.Find(&users)
	if result.Error != nil {
		logrus.Infof("Failed To Get All User %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  "Failed to load user details",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"success": "Successfully loaded user details",
		"data":    users,
	})
}

func GetAllUsersById(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	Id := c.Param("id")
	var users models.User
	result := config.DB.Find(&users, "id=?", Id)

	if result.Error != nil {
		logrus.Infof("Failed To Get User Details By Id %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to load user details",
			"data":   nil,
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "user details",
		"data":    users,
	})
}
func DeleteUser(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	Id := c.Param("id")
	var users models.User
	result := config.DB.Model(&users).Where("id=?", Id).Delete(&users)
	if result.Error != nil {
		logrus.Infof("Failed To Deleted User %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "User Delete UnSuccessfully",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "User Delete Successfully",
		"data":    nil,
	})

}
func AddUser(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Full_Name string
		Email     string
		Password  string
		Contact   string
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed To Read Body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}

	//Hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		logrus.Infof("Failed To Hash Password %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to hash password",
			"data":   nil,
		})
		return
	}

	//create the user
	user := models.User{Full_Name: body.Full_Name, Email: body.Email, Password: string(hash), Contact: body.Contact, Account_Status: 1}

	result := config.DB.Create(&user)
	if result.Error != nil {
		logrus.Infof("Failed To Created User %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "user Unsuccessfully created",
			"data":    nil,
		})

		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  201,
		"message": "user successfully created",
		"data":    user,
	})

}

func UpdateUserDetails(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	Id := c.Param("id")
	var body struct {
		Full_Name      string
		Email          string
		Password       string
		Contact        string
		Account_Status int
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed To Read Body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}
	Hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		logrus.Infof("Failed To Hash Passowrd %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to hash password",
			"data":   nil,
		})
		return
	}

	users := models.User{Full_Name: body.Full_Name, Email: body.Email, Contact: body.Contact, Password: string(Hash), Account_Status: body.Account_Status}
	result := config.DB.Model(&users).Where("id = ?", Id).Updates(users)
	if result.Error != nil {
		logrus.Infof("Failed To User Detail Updated %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "User Update UnSuccessfully",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "User Update Successfully",
		"data":    body,
	})

}
func CountUser(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var user models.User
	var count int64
	result := config.DB.Model(&user).Count(&count)

	if result.Error != nil {
		logrus.Infof("Failed To User Count Detail %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Total User Count Unsuccessfully",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Total User Count Successfully",
		"data":    count,
	})
}

func GetMonthlyUsers(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var users []models.User

	// Get the current year
	currentYear := time.Now().Year()

	// Retrieve users for all months in the current year
	result := config.DB.
		Where("YEAR(created_at) = ?", currentYear).
		Find(&users)

	// Handle the error
	if result.Error != nil {
		logrus.Infof("Failed to retrieve monthly users: %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to retrieve monthly users",
			"data":   nil,
		})
		return
	}

	// Organize users by month
	organizedUsers := make(map[string][]map[string]interface{})

	for _, user := range users {
		createdAtMonth := user.CreatedAt.Month().String()
		userInfo := map[string]interface{}{
			"ID":             user.ID,
			"Full_Name":      user.Full_Name,
			"Email":          user.Email,
			"Contact":        user.Contact,
			"Account_Status": user.Account_Status,
			// Add more fields as needed
		}

		// Append the user to the corresponding month
		organizedUsers[createdAtMonth] = append(organizedUsers[createdAtMonth], userInfo)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Monthly Users Successfully Retrieved",
		"data":    organizedUsers,
	})
}

func Today_Total_Booking(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	now := time.Now()
	date := now.Format("02-01-2006")
	var booking models.Confirm_Booking_Table
	var count int64

	// role, _ := c.Request.Cookie("Role")
	// Role, _ := strconv.Atoi(role.Value)
	branchID, _ := c.Request.Cookie("Branch_id")
	branchid, _ := strconv.Atoi(branchID.Value)

	// if Role != 1 {
	result := config.DB.Model(&booking).Where("date=? AND booking_status = 4 AND branch_id=?", date, branchid).Count(&count)
	if result.Error != nil || count == 0 {
		logrus.Infof("Failed To Load Today Total Booking %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Today Total  Booking Unsuccessfully",
			"data":   0,
		})
		return
	}
	// } else {
	// 	result := config.DB.Model(&booking).Where("date=? AND booking_status=4", date).Count(&count)
	// 	if result.Error != nil || count == 0 {
	// 		logrus.Infof("Failed To Load Total Booking %v\n", result.Error)
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"status": 400,
	// 			"error":  "Today Total  Booking Unsuccessfully",
	// 			"data":   0,
	// 		})
	// 		return
	// 	}
	// }

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Today Total  Booking Successfully",
		"data":    count,
	})
}

// func In_live_slot(c *gin.Context){
// 	var slot models.Confirm_Booking_Table
// 	config.DB.
// }

// func GetCurrentHourSlot() (time.Time, time.Time) {
// 	currentTime := time.Now()
// 	start := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second(), currentTime.Nanosecond(), currentTime.Location())
// 	day := time.Now().Weekday()
// 	fmt.Println(day)

// 	end := start.Add(time.Hour)

// 	return start, end
// }

// func getOccupiedSlots(startTime, endTime time.Time) ([]models.Turf_Bookings, error) {
// 	var occupiedSlots []models.Turf_Bookings

// 	// Query the database to find occupied slots within the specified time range
// 	if err := config.DB.Where("start_time >= ? AND end_time <= ?", startTime, endTime).Find(&occupiedSlots).Error; err != nil {
// 		return nil, err
// 	}

//		return occupiedSlots, nil
//	}
func AdminAddScreenshot(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var body struct {
		Amount float64
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed To Read Body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Invalid Request To Payload",
			"data":   nil,
		})
		return

	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filePath := filepath.Join("./uploads/admin_uploads", file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	payment := models.Screenshot{Payment_screenshot: filePath, Booking_order_id: id}
	result := config.DB.Create(&payment)
	if result.Error != nil {
		logrus.Infof("Failed To Insert Screenshot %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to insert screeshot",
			"data":   nil,
		})
		return
	} else {
		changed_status := models.Confirm_Booking_Table{
			Booking_status: 3,
		}
		var booking models.Confirm_Booking_Table
		status := config.DB.Model(&booking).Where("booking_order_id = ?", booking.Booking_order_id).Updates(changed_status)
		if status.Error != nil {
			logrus.Infof("Failed To Update Confirm Detail %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "failed to upadte confirm boooking details",
				"data":   nil,
			})
			return
		}
		var turf_book models.Turf_Bookings

		is_booked := models.Turf_Bookings{
			Is_booked: 3,
		}
		result := config.DB.Model(&turf_book).Where("order_id = ?", booking.Booking_order_id).Updates(is_booked)
		if result.Error != nil {
			logrus.Infof("Failed To Update Boooking Detail %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "failed to upadte boooking details",
				"data":   nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "Successfully uploaded",
			"data":    payment,
		})

	}

}

// func Pending_bookings(c *gin.Context) {
// 	var pending []models.Confirm_Booking_Table
// 	config.DB.Find(&pending, "booking_status = ?", 2)
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  200,
// 		"message": "Successfully upladed",
// 		"data":    pending,
// 	})

// }

//	func Partial_payment(c *gin.Context) {
//		var partial []models.Confirm_Booking_Table
//		config.DB.Find(&partial, "remaining_amount_to_pay > 0")
//		c.JSON(http.StatusOK, gin.H{
//			"status":  200,
//			"message": "Successfully upladed",
//			"data":    partial,
//		})
//	}
func AddSlotForUser(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	Id := c.Param("id")
	ID, _ := strconv.ParseUint(Id, 10, 64)
	var body struct {
		Date string
		Slot []int
		// StartSlot string
		// EndSlot   string
	}
	var Slots []int
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed To Read Body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}

	rows := config.DB.Model(&models.Turf_Bookings{}).Where("date = ?", body.Date).Pluck("slot_id", &Slots)

	if rows.Error != nil {
		logrus.Infof("Failed To Add Slot For User %v\n", rows.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to add slot for user",
			"data":   nil,
		})
		return
	}
	availableSlots := []int{}
	for _, s := range body.Slot {
		for _, s1 := range Slots {
			if s == s1 {
				availableSlots = append(availableSlots, int(s))
			}
		}
	}
	uniqueslots := make([]int, 0)

	bMap := make(map[int]bool)
	for _, val := range availableSlots {
		bMap[val] = true

	}

	for _, val := range body.Slot {

		if !bMap[val] {
			uniqueslots = append(uniqueslots, val)

		}
	}

	fmt.Println("ava:", availableSlots)
	if len(availableSlots) == 0 {

		Booking_id, _ := uuid.NewRandom()

		B_id := Booking_id.String()

		for i := 0; i < len(body.Slot); i++ {

			var psr models.Package_slot_relationship

			config.DB.First(&psr, "slot_id=?", int(body.Slot[i]))

			//fetch the price based on package id retrieved

			var price models.Package

			config.DB.Find(&price, "id=?", psr.Package_id)

			price25 := percent.PercentFloat(25.0, price.Price)

			booking := models.Turf_Bookings{Date: body.Date, Slot_id: body.Slot[i], User_id: uint(ID), Package_slot_relation_id: int(psr.ID), Package_id: psr.Package_id, Price: price.Price, Minimum_amount_to_pay: price25, Order_id: B_id}
			result := config.DB.Create(&booking)
			if result.Error != nil {
				logrus.Infof("Failed To Add Slot For User %v\n", result.Error)
				c.JSON(http.StatusOK, gin.H{
					"status": 400,
					"error":  "Slot Allready booked",
					"data":   nil,
				})
				return
			}

		}

		var booking models.Turf_Bookings

		//confirm booking table

		config.DB.Find(&booking, "order_id = ?", B_id)

		var totalPrice float64
		var total_min_amount float64
		for p := 0; p < len(body.Slot); p++ {
			totalPrice += booking.Price
			total_min_amount += booking.Minimum_amount_to_pay
		}

		confirm_booking := models.Confirm_Booking_Table{Date: body.Date, User_id: uint(ID), Booking_order_id: B_id, Total_price: totalPrice, Total_min_amount_to_pay: total_min_amount, Booking_status: 1}

		result := config.DB.Create(&confirm_booking)
		if result.Error != nil {
			logrus.Infof("Failed To Add Slot For User %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Slot Allready Booked",
				"data":   nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"success": "Slot reserved successfully",
			"data":    booking,
		})

	} else if len(availableSlots) != 0 && len(uniqueslots) != 0 {

		Booking_id, _ := uuid.NewRandom()

		B_id := Booking_id.String()

		for i := 0; i < len(uniqueslots); i++ {

			var psr models.Package_slot_relationship

			config.DB.First(&psr, "slot_id=?", int(uniqueslots[i]))

			//fetch the price based on package id retrieved

			var price models.Package

			config.DB.Find(&price, "id=?", psr.Package_id)

			price25 := percent.PercentFloat(25.0, price.Price)

			booking := models.Turf_Bookings{Date: body.Date, Slot_id: uniqueslots[i], User_id: uint(ID), Package_slot_relation_id: int(psr.ID), Package_id: psr.Package_id, Price: price.Price, Minimum_amount_to_pay: price25, Order_id: B_id}
			result := config.DB.Create(&booking)
			if result.Error != nil {
				logrus.Infof("Failed To Add Slot For User %v\n", result.Error)
				c.JSON(http.StatusOK, gin.H{
					"status": 400,
					"error":  "Slot Allready booked",
					"data":   nil,
				})
				return
			}

		}

		var booking models.Turf_Bookings

		//confirm booking table

		config.DB.Find(&booking, "order_id = ?", B_id)

		var totalPrice float64
		var total_min_amount float64
		for p := 0; p < len(uniqueslots); p++ {
			totalPrice += booking.Price
			total_min_amount += booking.Minimum_amount_to_pay
		}

		confirm_booking := models.Confirm_Booking_Table{Date: body.Date, User_id: uint(ID), Booking_order_id: B_id, Total_price: totalPrice, Total_min_amount_to_pay: total_min_amount, Booking_status: 2}

		result := config.DB.Create(&confirm_booking)
		if result.Error != nil {
			logrus.Infof("Failed To Add Slot For User %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Slot Allready boooked",
				"data":   nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"success": "Slot reserved successfully",
			"data":    booking,
		})

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Slot is allready booked",
			"data":   nil,
		})
	}
}

func LiveUser(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	// Check for the Authorization cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Check for token expiration
	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expirationTime) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var admin models.Admin
	config.DB.Find(&admin, claims["sub"])

	if admin.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}

	// Get the current date and time
	now := time.Now()
	date := now.Format("02-01-2006")

	currentTime := now.Format("15:04")

	// Find the time slot for the current time
	var slot models.Time_Slot
	result := config.DB.Where("start_time <= ? AND end_time >= ? AND branch_id = ? ", currentTime, currentTime, admin.Turf_branch_id).First(&slot)
	if result.Error != nil {
		logrus.Infof("Failed To Get Live User %v\n", result.Error)
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  "Slot is not found",
			"data":   nil,
		})
		return
	}

	// Find the booking details for the current date and time slot
	var booking models.Turf_Bookings
	result = config.DB.Where("date = ? AND slot_id = ? AND branch_id = ? ", date, slot.ID, admin.Turf_branch_id).First(&booking)
	if result.Error != nil {
		logrus.Infof("Failed To Get Live User %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Booking details not found",
			"data":   nil,
		})
		return
	}

	// Find the user associated with the booking
	var user models.User
	if result := config.DB.Where("id = ?", booking.User_id).First(&user); result.Error != nil {
		logrus.Infof("Failed To Get Live User %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "User not found",
			"data":   nil,
		})
		return
	}

	// Create a response containing booking and user information
	live := map[string]interface{}{
		"Full_Name":  user.Full_Name,
		"Contact":    user.Contact,
		"start_time": slot.Start_time,
		"End_time":   slot.End_time,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "User fetched successfully",
		"data":    live,
	})
}

func Testimonials(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Name        string
		Designation string
		Review      string
		Image       string
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed To Read Body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "faild to read body",
			"data":   nil,
		})
		return

	}
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filePath := filepath.Join("./uploads/testimonials", file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	if filepath.Ext(filePath) != ".jpg" && filepath.Ext(filePath) != ".png" {
		logrus.Infof("Failed To Upload Image %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Upload the right file format (jpg or png)",
			"data":   nil,
		})
		return
	}

	testimonial := &models.Testi_Monial{Name: body.Name, Designation: body.Designation, Review: body.Review, Image: filePath}
	result := config.DB.Create(&testimonial)
	if result.Error != nil {
		logrus.Infof("Failed To Create Testimonials %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to create testimonials",
			"data":   nil,
		})
		return
	}
	//c.Header("Content-Type", "multipart/mixed; boundary=myboundary")
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "testimonials create successfully",
		"data":    testimonial,
	})
}
func Upadte_Testimonials(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var body struct {
		Name        string
		Designation string
		Review      string
		Image       string
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed To Read Body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "faild to read body",
			"data":   nil,
		})
		return

	}

	testimonial := &models.Testi_Monial{Name: body.Name, Designation: body.Designation, Review: body.Review}
	result := config.DB.Model(&testimonial).Where("id=?", id).Updates(&testimonial)
	if result.Error != nil {
		logrus.Infof("Failed To Update Testimonials Details %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Update testimonials Details",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Testimonials Upadted Successfully",
		"data":    testimonial,
	})

}
func UpdateImageForTestimonials(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var body struct {
		Image string
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed To Read Body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filePath := filepath.Join("./uploads/testimonials", file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	if filepath.Ext(filePath) != ".jpg" && filepath.Ext(filePath) != ".png" {
		logrus.Infof("Failed To Upload Right File Format %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Upload the right file format (jpg or png)",
			"data":   nil,
		})
		return
	}

	testimonial := &models.Testi_Monial{Image: filePath}
	result := config.DB.Model(&testimonial).Where("id = ?", id).Updates(&testimonial)
	if result.Error != nil {
		logrus.Infof("Failed To Update Testimonials Image %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to update testimonials Image",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "testimonials Image updated successfully",
		"data":    testimonial,
	})
}
func UpdateImageForTestimonials2(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Image string
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed To Read Body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filePath := filepath.Join("./uploads/testimonials", file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	if filepath.Ext(filePath) != ".jpg" && filepath.Ext(filePath) != ".png" {
		logrus.Infof("Failed To Upload Right File Format %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Upload the right file format (jpg or png)",
			"data":   nil,
		})
		return
	}
	var testimonials models.Testi_Monial
	result := config.DB.Find(&testimonials)
	if result.Error != nil {
		logrus.Infof("Failed To Update Testimonials Image %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to update testimonials",
			"data":   nil,
		})
		return
	}
	fmt.Println(testimonials.ID)
	testimonial := &models.Testi_Monial{Image: filePath}
	result = config.DB.Model(&testimonial).Where("id = ?", testimonials.ID).Updates(&testimonial)
	if result.Error != nil {
		logrus.Infof("Failed To Update Testimonials Image %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to update testimonials",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "testimonials Images updated successfully",
		"data":    testimonial,
	})
}
func AllTestimonials(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var testimonials []models.Testi_Monial
	result := config.DB.Find(&testimonials)
	if result.Error != nil {
		logrus.Infof("Failed To Get All Testimonials %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to fetch testimonials",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "testimonials details fetch successfully",
		"data":    testimonials,
	})

}
func GETTestimonialsById(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")

	var testimonials models.Testi_Monial
	result := config.DB.Find(&testimonials, "id=?", id)

	fmt.Println(testimonials.Image)

	if result.Error != nil {
		logrus.Infof("Failed To Get Testimonials By Id %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to fetch testimonial by id",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "testimonial fetch successfully by id",
		"data":    testimonials,
	})

}

func GETTestimonialsimagesById(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")

	var testimonials models.Testi_Monial
	result := config.DB.Find(&testimonials, "id=?", id)

	if result.Error != nil {
		logrus.Infof("Failed To Get Testimonials Image %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to fetch testimonial image",
			"data":   nil,
		})
		return
	}

	// Determine the file path based on the file format (you may need to store this information in your model)
	var filePath string
	if strings.HasSuffix(testimonials.Image, ".jpg") {
		filePath = testimonials.Image
		c.Header("Content-Type", "image/jpeg")
	} else if strings.HasSuffix(testimonials.Image, ".png") {
		filePath = testimonials.Image
		c.Header("Content-Type", "image/png")
	} else {
		logrus.Infof("Failed To Get  Testimonials image %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "unsupported image format",
			"data":   nil,
		})
		return
	}

	// Read the image file
	imageData, err := ioutil.ReadFile(filePath)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"error":  "internal server error",
			"data":   nil,
		})
		return
	}

	// Send the image data as the response
	c.Data(http.StatusOK, c.GetHeader("Content-Type"), imageData)
}
func DeleteTestimonials(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var testimonial models.Testi_Monial
	result := config.DB.Model(&testimonial).Where("id=?", id).Delete(&testimonial)
	if result.RowsAffected == 0 {
		logrus.Infof("Failed To Delete Testimonials %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "unsuccessfully Deleted Testimonial",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "successfully Deleted Testimonial",
		"data":    nil,
	})
}
func readJPGFile(filePath string) ([]byte, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func AdminLogout(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	// Clear the "Authorization" cookie to log out
	c.SetCookie("Authorization", "", -1, "", "", false, true)
	c.SetCookie("AID", "", -1, "", "", false, true)
	c.SetCookie("Role", "", -1, "", "", false, true)
	c.SetCookie("Branch_id", "", -1, "", "", false, true)
	c.SetCookie("logged_in", "", -1, "", "", false, true)

	// You can also clear any other session-related data if needed
	c.Set("UserID", "")

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Logged out successfully",
	})
}
func AddContent(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Heading    string
		SubHeading string
		Button     string
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed To Read Body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to read body",
			"data":   nil,
		})
		return
	}
	content := &models.Content{Heading: body.Heading, SubHeading: body.SubHeading, Button: body.Button}
	result := config.DB.Create(&content)
	if result.Error != nil {
		logrus.Infof("Failed To Create Content %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to create content",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": 201,
		"error":  "Success to create content",
		"data":   content,
	})

}
func GETContent(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var content []models.Content
	result := config.DB.Find(&content)
	if result.Error != nil {
		logrus.Infof("Failed To Get All Content %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to get content",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  "Success to get content",
		"data":   content,
	})

}
func UpdateContent(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	Id := c.Param("id")
	var body struct {
		Heading    string
		SubHeading string
		Button     string
		Status     string
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed To Read Body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to read body",
			"data":   nil,
		})
		return
	}
	content := models.Content{Heading: body.Heading, SubHeading: body.SubHeading, Button: body.Button, Status: body.Status}
	result := config.DB.Model(&content).Where("id=?", Id).Updates(&content)
	if result.Error != nil {
		logrus.Infof("Failed To Upadte Content %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to update content",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  "Success to update content",
		"data":   content,
	})

}
func GetContentById(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")

	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	Id := c.Param("id")
	var content models.Content
	result := config.DB.Find(&content, "id=?", Id)

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to get content",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success to get content",
		"data":    content,
	})
}
func ActiveContent(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var content models.Content
	result := config.DB.Find(&content, "status=1")

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to get active content",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success to get active content",
		"data":    content,
	})
}

func DeleteContent(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var content models.Content
	result := config.DB.Model(&content).Where("id=?", id).Delete(&content)
	if result.RowsAffected == 0 {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "unsuccessfully Deleted content",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "successfully Deleted content",
		"data":    nil,
	})
}
func AddImageForCarousel(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Image string
	}
	if c.Bind(&body) != nil {
		logrus.Infof("Failed to read b0dy in AddImageF0rCar0usel")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to read body",
			"data":   nil,
		})
		return
	}
	file, err := c.FormFile("image")
	if err != nil {
		logrus.Infof("err: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filePath := filepath.Join("./uploads/carousel", file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {

		logrus.Infof("Failed to save uploaded file: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	if filepath.Ext(filePath) != ".jpg" && filepath.Ext(filePath) != ".png" {
		logrus.Info("Upload the right file format (jpg or png) AddImageF0rCar0usel")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Upload the right file format (jpg or png)",
			"data":   nil,
		})
		return
	}
	carousel := &models.Carousel{Image: filePath, Status: "1"}
	result := config.DB.Create(&carousel)
	if result.Error != nil {
		logrus.Infof("Failed to add data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to add image for carousel",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  "Success to add image for carousel",
		"data":   carousel,
	})
}
func GetAllImageCarousel(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")

	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var carousel []models.Carousel
	result := config.DB.Find(&carousel)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to get all carousel",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  "Success to get all carousel",
		"data":   carousel,
	})

}
func GetActiveImageCarousel(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var carousel []models.Carousel
	result := config.DB.Find(&carousel, "status=1")

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to get active carousel",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success to get active carousel",
		"data":    carousel,
	})

}
func Upadtecarousel(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var body struct {
		Status string
	}
	if c.Bind(&body) != nil {
		logrus.Info("Invalid request payload")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to read body",
			"data":   nil,
		})
		return
	}

	carousel := models.Carousel{Status: body.Status}
	result := config.DB.Model(&carousel).Where("id=?", id).Updates(&carousel)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to update carousel",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"error":  "Success to update content",
		"data":   carousel,
	})

}
func UpadtecarouselImage(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Image string
	}
	if c.Bind(&body) != nil {
		logrus.Info("Invalid request payload")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}
	id := c.Param("id")
	file, err := c.FormFile("image")
	if err != nil {
		logrus.Infof("Failed to upload image file %s \n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filePath := filepath.Join("./uploads/carousel", file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		logrus.Infof("Failed to save uploaded file to disk: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	if filepath.Ext(filePath) != ".jpg" && filepath.Ext(filePath) != ".png" {

		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Upload the right file format (jpg or png)",
			"data":   nil,
		})
		return
	}

	carousel := &models.Carousel{Image: filePath}
	result := config.DB.Model(&carousel).Where("id = ?", id).Updates(&carousel)
	if result.Error != nil {
		logrus.WithField("err", result.Error).Warnf("Could not update carousel with id %v", id)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to update image for carousel",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "carousel updated image for successfully",
		"data":    carousel,
	})
}
func DeleteCarousel(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var carousel models.Carousel
	result := config.DB.Model(&carousel).Where("id=?", id).Delete(&carousel)
	if result.RowsAffected == 0 {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "unsuccessfully Deleted Carousel",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "successfully Deleted  Carousel",
		"data":    nil,
	})
}

//package slot relationship retrieval for slot management in packages

// func PSR_slots(c *gin.Context) {

// 	var package_name []models.Package

// 	//var psr_ID int

// 	result := config.DB.Debug().Model(&models.Package{}).
// 		Select("packages.name").
// 		Joins("LEFT JOIN package_slot_relationships ON packages.id = package_slot_relationships.package_id").
// 		Scan(&package_name)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error":   result.Error.Error(),
// 			"message": "Failed to fetch package slots",
// 		})
// 		return
// 	}
// 	var slot []models.Time_Slot
// 	result = config.DB.Debug().Model(&models.Time_Slot{}).
// 		Select("time_slots.id, time_slots.start_time, time_slots.end_time, time_slots.day time_slots.branch_id, package_slot_relationships.id as psr_id").
// 		Joins("LEFT JOIN package_slot_relationships ON time_slots.id = package_slot_relationships.slot_id").
// 		Scan(&slot)

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  200,
// 		"success": "package names",
// 		"data":    package_name,
// 	})
// }

// func PSR_slots(c *gin.Context) {
// 	var response struct {
// 		PackageName []models.Package
// 		Slot        []models.Time_Slot
// 	}

// 	result := config.DB.Debug().Model(&models.Package{}).
// 		Select("packages.id, packages.name").
// 		Joins("LEFT JOIN package_slot_relationships ON packages.id = package_slot_relationships.package_id").
// 		Scan(&response.PackageName)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error":   result.Error.Error(),
// 			"message": "Failed to fetch package slots",
// 		})
// 		return
// 	}

// 	result = config.DB.Debug().Model(&models.Time_Slot{}).
// 		Select("time_slots.id, time_slots.start_time, time_slots.end_time, time_slots.day, time_slots.branch_id, package_slot_relationships.id as psr_id").
// 		Joins("LEFT JOIN package_slot_relationships ON time_slots.id = package_slot_relationships.slot_id").
// 		Scan(&response.Slot)

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  200,
// 		"success": "package names and slots",
// 		"data":    response,
// 	})
// }

func PSR_slots(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var response struct {
		Data []interface{}
	}

	var packages []models.Package
	// var slots []models.Time_Slot

	result := config.DB.Debug().Raw(`
	SELECT p.id as ID, p.name as Name, p.price as Price, p.status as Status, p.branch_id as Branch_id, ts.start_time as Start_time, ts.end_time as End_time, ts.day as Day, ts.branch_id as Slot_Branch_id, psr.id as PSR_id, bim.branch_name as Branch_name FROM package_slot_relationships psr INNER JOIN packages p ON p.id = psr.package_id INNER JOIN time_slots ts ON psr.slot_id = ts.id INNER JOIN branch_info_managements bim ON ts.branch_id = bim.id
`).Scan(&packages)

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   result.Error.Error(),
			"message": "Failed to fetch package slots",
		})
		return
	}
	// Combine the "Package" and "Slot" data into a single array
	for _, packageData := range packages {
		response.Data = append(response.Data, packageData)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "package names and slots",
		"data":    response,
	})
}

func GETCarouselActiveImages(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")

	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var carousels []models.Carousel
	result := config.DB.Find(&carousels, "status = 1")

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to fetch active image for carousels",
			"data":   nil,
		})
		return
	}

	// Collect the image data for the first three images
	var imageData []byte
	var contentType string // Initialize the content type variable

	for i, carousel := range carousels {
		if i < 3 {
			// Determine the file path based on the file format (you may need to store this information in your model)
			var filePath string

			if strings.HasSuffix(carousel.Image, ".jpg") {
				filePath = carousel.Image
				contentType = "image/jpeg"
			} else if strings.HasSuffix(carousel.Image, ".png") {
				filePath = carousel.Image
				contentType = "image/png"
			} else {
				// Handle unsupported image formats for individual images, but continue processing others
				c.JSON(http.StatusBadRequest, gin.H{
					"status": 400,
					"error":  "unsupported image format",
					"data":   nil,
				})
				continue
			}

			// Read the image file
			imageBytes, err := ioutil.ReadFile(filePath)
			if err != nil {
				logrus.Infof("Err0r reading the image file %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": 500,
					"error":  "internal server error",
					"data":   nil,
				})
				return
			}

			// Append the image data to the imageData slice
			imageData = append(imageData, imageBytes...)
		}
	}

	// Send the combined image data with the dynamic content type
	c.Data(http.StatusOK, contentType, imageData)
}
func RemainingPaymentForUser(c *gin.Context) {
	// Set CORS headers to allow cross-origin requests
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")

	// Handle preflight OPTIONS requests
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	now := time.Now()
	date := now.Format("02-01-2006")

	// Define a slice to store booking data
	var booking []models.Confirm_Booking_Table

	// Use the WHERE clause in the Find method to filter results
	result := config.DB.Find(&booking, "date <= ? AND remaining_amount_to_pay > 0", date)

	// Check if any matching records were found
	if result.RowsAffected == 0 {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "No matching booking details found",
			"data":   nil,
		})
		return
	}

	var responseData []interface{}

	for _, bookings := range booking {
		var user models.User
		result := config.DB.Find(&user, "id=?", bookings.User_id)

		if result.Error != nil {
			logrus.Infof("Failed to get data from DB %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 404,
				"error":  "Failed to fetch user name",
			})
			return
		}

		var branch models.Branch_info_management
		result = config.DB.Find(&branch, "id=?", bookings.Branch_id)

		if result.Error != nil {
			logrus.Infof("Failed to get data from DB %v\n", result.Error)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 404,
				"error":  "Failed to fetch branch name",
			})
			return
		}

		bookingData := map[string]interface{}{
			"ID":                   bookings.ID,
			"CreatedAt":            bookings.CreatedAt,
			"UserID":               bookings.User_id,
			"UserName":             user.Full_Name,
			"Contact":              user.Contact,
			"Date":                 bookings.Date,
			"BookingOrderID":       bookings.Booking_order_id,
			"TotalPrice":           bookings.Total_price,
			"TotalMinAmountToPay":  bookings.Total_min_amount_to_pay,
			"PaidAmount":           bookings.Paid_amount,
			"RemainingAmountToPay": bookings.Remaining_amount_to_pay,

			"BranchName": branch.Branch_name,
		}

		responseData = append(responseData, bookingData)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Successfully fetched booking details",
		"data":    responseData,
	})
}

func GetCarouselimagesById(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")

	var carousel models.Carousel
	result := config.DB.Find(&carousel, "id=?", id)

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to fetch carousel image for id",
			"data":   nil,
		})
		return
	}

	// Determine the file path based on the file format (you may need to store this information in your model)
	var filePath string
	if strings.HasSuffix(carousel.Image, ".jpg") {
		filePath = carousel.Image
		c.Header("Content-Type", "image/jpeg")
	} else if strings.HasSuffix(carousel.Image, ".png") {
		filePath = carousel.Image
		c.Header("Content-Type", "image/png")
	} else {

		// Handle unsupported image formats
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "unsupported image format",
			"data":   nil,
		})
		return
	}

	// Read the image file
	imageData, err := ioutil.ReadFile(filePath)
	if err != nil {
		logrus.Infof("Err0r readin the image file %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"error":  "internal server error",
			"data":   nil,
		})
		return
	}

	// Send the image data as the response
	c.Data(http.StatusOK, c.GetHeader("Content-Type"), imageData)
}

func Cnfrm_slots(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var response struct {
		Data []interface{}
	}

	var bookings []models.Confirm_Booking_Table
	// role, _ := c.Request.Cookie("Role")
	// Role, _ := strconv.Atoi(role.Value)
	branchID, _ := c.Request.Cookie("Branch_id")
	branchid, _ := strconv.Atoi(branchID.Value)

	// if Role != 1 {
	result := config.DB.Debug().Raw(` SELECT  u.full_name as Name, u.contact as Contact, cb.date , cb.total_price , cb.total_min_amount_to_pay, cb.paid_amount, cb.remaining_amount_to_pay , cb.booking_status, cb.branch_id FROM users u INNER JOIN confirm_booking_tables cb ON u.id = cb.user_id WHERE cb.booking_status = 4 AND cb.branch_id=?`, branchid).Scan(&bookings)

	//INNER JOIN branch_info_managements bim ON ts.branch_id = bim.id

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   result.Error.Error(),
			"message": "Failed to fetch Confirm slots",
		})
		return
	}
	// } else {

	// 	// var slots []models.Time_Slot

	// 	result := config.DB.Debug().Raw(`SELECT  u.full_name as Name, u.contact as Contact, cb.date , cb.total_price , cb.total_min_amount_to_pay, cb.paid_amount, cb.remaining_amount_to_pay , cb.booking_status, cb.branch_id FROM users u INNER JOIN confirm_booking_tables cb ON u.id = cb.user_id WHERE cb.booking_status = 4`).Scan(&bookings)

	// 	//INNER JOIN branch_info_managements bim ON ts.branch_id = bim.id

	// 	if result.Error != nil {
	// 		logrus.Infof("Failed to get data from DB %v\n", result.Error)
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error":   result.Error.Error(),
	// 			"message": "Failed to fetch confirm slots",
	// 		})
	// 		return
	// 	}
	// }
	// Combine the "Package" and "Slot" data into a single array
	for _, cbData := range bookings {
		response.Data = append(response.Data, cbData)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Confirm Slot",
		"data":    response,
	})
}

func Pending_bookings(c *gin.Context) {
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var response struct {
		Data []interface{}
	}

	var bookings []models.Confirm_Booking_Table

	// role, _ := c.Request.Cookie("Role")
	// Role, _ := strconv.Atoi(role.Value)
	branchID, _ := c.Request.Cookie("Branch_id")
	branchid, _ := strconv.Atoi(branchID.Value)

	// if Role != 1 {

	// var slots []models.Time_Slot

	result := config.DB.Debug().Raw(`SELECT  u.full_name as Name, u.contact as Contact, cb.ID, cb.date , cb.total_price , cb.total_min_amount_to_pay, cb.paid_amount, cb.remaining_amount_to_pay , cb.booking_status, bim.branch_name FROM users u INNER JOIN confirm_booking_tables cb ON u.id = cb.user_id INNER JOIN branch_info_managements bim ON cb.branch_id = bim.id WHERE cb.booking_status = 3 AND cb.branch_id =?`, branchid).Scan(&bookings)

	//INNER JOIN branch_info_managements bim ON ts.branch_id = bim.id

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   result.Error.Error(),
			"message": "Failed to fetch pending booking",
		})
		return
	}
	// } else {
	// 	result := config.DB.Debug().Raw(`SELECT  u.full_name as Name, u.contact as Contact, cb.ID, cb.date , cb.total_price , cb.total_min_amount_to_pay, cb.paid_amount, cb.remaining_amount_to_pay , cb.booking_status, bim.branch_name FROM users u INNER JOIN confirm_booking_tables cb ON u.id = cb.user_id INNER JOIN branch_info_managements bim ON cb.branch_id = bim.id WHERE cb.booking_status = 3`).Scan(&bookings)

	// 	//INNER JOIN branch_info_managements bim ON ts.branch_id = bim.id

	// 	if result.Error != nil {
	// 		logrus.Infof("Failed to get data from DB %v\n", result.Error)
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error":   result.Error.Error(),
	// 			"message": "Failed to fetch pending booking",
	// 		})

	// 	}
	// }
	// Combine the "Package" and "Slot" data into a single array
	for _, cbData := range bookings {
		response.Data = append(response.Data, cbData)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "pending booking",
		"data":    response,
	})
}
func Pending_bookings_by_ID(c *gin.Context) {
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	Id := c.Param("id")
	var response struct {
		Data []interface{}
	}

	ID, _ := strconv.Atoi(Id)

	var bookings []models.Confirm_Booking_Table
	// var slots []models.Time_Slot

	result := config.DB.Debug().Raw(`
	SELECT  u.full_name as Name, u.contact as Contact, cb.ID, cb.date , cb.total_price , cb.total_min_amount_to_pay, cb.paid_amount, cb.remaining_amount_to_pay , cb.booking_status, cb.booking_order_id, bim.branch_name FROM users u INNER JOIN confirm_booking_tables cb ON u.id = cb.user_id INNER JOIN branch_info_managements bim ON cb.branch_id = bim.id WHERE cb.booking_status = 3 AND cb.ID = ?
`, ID).Scan(&bookings)

	//INNER JOIN branch_info_managements bim ON ts.branch_id = bim.id

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   result.Error.Error(),
			"message": "Failed to fetch pending booking by id",
		})
		return
	}
	// Combine the "Package" and "Slot" data into a single array
	for _, cbData := range bookings {
		response.Data = append(response.Data, cbData)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "pending booking by id",
		"data":    response,
	})
}

func GetpaymentimagesById(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	// var body struct {
	// 	Booking_order_id string
	// }
	id := c.Param("id")

	var payment models.Screenshot
	result := config.DB.Find(&payment, "booking_order_id=?", id)

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to fetch Payment Image",
			"data":   nil,
		})
		return
	}

	// Handle unsupported image formats
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Payment Image",
		"data":   payment,
	})
	return
}
func AddNavbar(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Name   string
		Link   string
		Status string
	}
	if c.Bind(&body) != nil {
		logrus.Infof("Failed to read b0dy")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}
	navbar := models.Navbar{Name: body.Name, Link: body.Link, Status: "2"}
	result := config.DB.Create(&navbar)
	if result.Error != nil {
		logrus.Infof("Failed to add navbar t0 DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Add navbar",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "navbar add successfully",
		"data":    navbar,
	})

}
func GetAllNavbar(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var navbars []models.Navbar
	result := config.DB.Find(&navbars)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Get Navbar Details",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": " Get Navabar Details successfully",
		"data":    navbars,
	})
}
func GetActiveNavbar(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var navbars []models.Navbar
	result := config.DB.Find(&navbars, "status=1")
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Get active Navbar Details",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": " Get active Navabar Details successfully",
		"data":    navbars,
	})
}
func UpadateNavbar(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var body struct {
		Name   string
		Link   string
		Status string
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed to read body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}
	navbar := models.Navbar{Name: body.Name, Link: body.Link, Status: body.Status}
	result := config.DB.Model(&navbar).Where("id=?", id).Updates(&navbar)
	if result.Error != nil {
		logrus.Infof("Failed to update data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Update navbar",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "navbar Updated successfully",
		"data":    navbar,
	})

}
func GetNavbarById(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var navbars models.Navbar
	result := config.DB.Find(&navbars, "id=?", id)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Get Navbar Details",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": " Get Navabar Details successfully",
		"data":    navbars,
	})
}
func DeleteNavbar(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var navbars models.Navbar
	result := config.DB.Model(&navbars).Where("id=?", id).Delete(&navbars)
	if result.Error != nil {
		logrus.Infof("Failed to delete data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Delete Navbar Details",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": " Delete Navabar Details successfully",
	})
}
func AddHeading(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Slider       string
		Testimonials string
		Footer       string
		Status       string
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed to read body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}
	heading := models.Heading{Slider: body.Slider, Testimonials: body.Testimonials, Footer: body.Footer, Status: "2"}
	result := config.DB.Create(&heading)
	if result.Error != nil {
		logrus.Infof("Failed to add data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Add Heading",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Heading add successfully",
		"data":    heading,
	})

}
func GetAllHeading(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var headings []models.Heading
	result := config.DB.Find(&headings)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Get Headings Details",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": " Get Headings Details successfully",
		"data":    headings,
	})
}
func GetActiveHeading(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var headings models.Heading
	result := config.DB.Find(&headings, "status=1")
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Get Avtive Heading Details",
			"data":   nil,
		})
		return
	}
	if headings.ID == 0 {

		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "No Active Heading",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": " Get Active Heading Details successfully",
		"data":    headings,
	})
}
func UpadateHeading(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var body struct {
		Slider       string
		Testimonials string
		Footer       string
		Status       string
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed to read body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}
	heading := models.Heading{Slider: body.Slider, Testimonials: body.Testimonials, Footer: body.Footer, Status: body.Status}
	result := config.DB.Model(&heading).Where("id=?", id).Updates(&heading)
	if result.Error != nil {
		logrus.Infof("Failed to update data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Update heading",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "heading Updated successfully",
		"data":    heading,
	})

}
func GetHeadingById(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var heading models.Heading
	result := config.DB.Find(&heading, "id=?", id)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Get Heading Details",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": " Get Heading Details successfully",
		"data":    heading,
	})
}
func DeleteHeading(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var heading models.Heading
	result := config.DB.Model(&heading).Where("id=?", id).Delete(&heading)
	if result.Error != nil {
		logrus.Infof("Failed to delete data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Delete Heading Details",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": " Delete Heading Details successfully",
	})
}

func Total_Revenue(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var revenue []models.Confirm_Booking_Table

	// role, _ := c.Request.Cookie("Role")
	// Role, _ := strconv.Atoi(role.Value)
	branchID, _ := c.Request.Cookie("Branch_id")
	branchid, _ := strconv.Atoi(branchID.Value)

	// fmt.Println(Role)

	// if Role != 1 {

	result := config.DB.Find(&revenue, "branch_id=?", branchid)

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 404,
			"error":  "failed to get total revenue ",
			"data":   nil,
		})
		return

	}
	// } else {
	// 	result := config.DB.Find(&revenue)
	// 	if result.Error != nil {
	// 		logrus.Infof("Failed to get data from DB %v\n", result.Error)
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"status": 404,
	// 			"error":  "failed to get total revenue ",
	// 			"data":   nil,
	// 		})
	// 		return
	// 	}
	// }

	// Calculate the total Paid_amount
	totalPaidAmount := 0.0
	for _, booking := range revenue {
		totalPaidAmount += booking.Paid_amount
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "revenue",
		"data":   totalPaidAmount,
	})
}

func Total_Remaining_amount(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var revenue []models.Confirm_Booking_Table

	// role, _ := c.Request.Cookie("Role")
	// Role, _ := strconv.Atoi(role.Value)
	branchID, _ := c.Request.Cookie("Branch_id")
	branchid, _ := strconv.Atoi(branchID.Value)

	// fmt.Println(Role)

	// if Role != 1 {

	result := config.DB.Find(&revenue, "branch_id=?", branchid)

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 404,
			"error":  "failed to get Total_Remaining_amount ",
			"data":   nil,
		})
		return

	}
	// } else {
	// 	result := config.DB.Find(&revenue)
	// 	if result.Error != nil {
	// 		logrus.Infof("Failed to get data from DB %v\n", result.Error)
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"status": 404,
	// 			"error":  "failed to get Total_Remaining_amount ",
	// 			"data":   nil,
	// 		})
	// 		return
	// 	}
	// }

	// Calculate the total Paid_amount
	totalRemainingAmount := 0.0
	for _, booking := range revenue {
		totalRemainingAmount += booking.Remaining_amount_to_pay
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Total_Remaining_amount",
		"data":   totalRemainingAmount,
	})
}

func Total_Sales(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var revenue []models.Confirm_Booking_Table

	// role, _ := c.Request.Cookie("Role")
	// Role, _ := strconv.Atoi(role.Value)
	branchID, _ := c.Request.Cookie("Branch_id")
	branchid, _ := strconv.Atoi(branchID.Value)

	// if Role != 1 {

	result := config.DB.Find(&revenue, "branch_id=?", branchid)

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 404,
			"error":  "failed to get total sales ",
			"data":   nil,
		})
		return

	}
	// } else {
	// 	result := config.DB.Find(&revenue)
	// 	if result.Error != nil {
	// 		logrus.Infof("Failed to get data from DB %v\n", result.Error)
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"status": 400,
	// 			"error":  "failed to get total sales",
	// 			"data":   nil,
	// 		})
	// 		return
	// 	}
	// }

	// Calculate the total Paid_amount
	totalSales := 0.0
	for _, booking := range revenue {
		totalSales += booking.Total_price
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "total Sales",
		"data":   totalSales,
	})
}

func Total_Monthly_revenue(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	branchID, _ := c.Request.Cookie("Branch_id")
	branchid, _ := strconv.Atoi(branchID.Value)

	sqlQuery := "SELECT * FROM `confirm_booking_tables` WHERE Year(created_at) = Year(CURRENT_DATE) AND MONTH(created_at) = MONTH(CURRENT_DATE) AND branch_id = ?"
	var monthlyRevenue []models.Confirm_Booking_Table
	result := config.DB.Raw(sqlQuery, branchid).Scan(&monthlyRevenue)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to execute query",
			"data":   nil,
		})
		return
	}

	// Calculate the sum of Paid_amount
	totalPaidAmount := 0.0
	for _, booking := range monthlyRevenue {
		totalPaidAmount += booking.Paid_amount
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "m0nthly revenue details",
		"data":   totalPaidAmount,
	})
}
func AddIcon(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var body struct {
		Name   string
		Link   string
		Status string
	}
	if c.Bind(&body) != nil {
		logrus.Infof("Failed to read body")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}
	icon := models.Icon{Name: body.Name, Link: body.Link, Status: "2"}
	result := config.DB.Create(&icon)
	if result.Error != nil {
		logrus.Infof("Failed to add data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Add Icon",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Icon Add successfully",
		"data":    icon,
	})

}
func GetAllIcon(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var icons []models.Icon
	result := config.DB.Find(&icons)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Get Icon Details",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": " Get Icon Details successfully",
		"data":    icons,
	})
}
func GetActiveIcon(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	var icons []models.Icon
	result := config.DB.Find(&icons, "status=1")
	if result.Error != nil {
		logrus.Infof("Failed to get  data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Active Icon Details",
			"data":   nil,
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": " Get Active Icon Details successfully",
		"data":    icons,
	})
}
func UpadateIcon(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var body struct {
		Name   string
		Link   string
		Status string
	}
	if c.Bind(&body) != nil {
		logrus.Infof("Failed to read body")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}
	icon := models.Icon{Name: body.Name, Link: body.Link, Status: body.Status}
	result := config.DB.Model(&icon).Where("id=?", id).Updates(&icon)
	if result.Error != nil {
		logrus.Infof("Failed to update data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Update Icon",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Icon Updated successfully",
		"data":    icon,
	})

}
func GetIconById(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var icons models.Icon
	result := config.DB.Find(&icons, "id=?", id)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Get Icon Details",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": " Get Icon Details successfully",
		"data":    icons,
	})
}
func DeleteIcon(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	id := c.Param("id")
	var icons models.Icon
	result := config.DB.Model(&icons).Where("id=?", id).Delete(&icons)
	if result.Error != nil {
		logrus.Infof("Failed to delete data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to Delete Icon Details",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": " Delete Icon Details successfully",
	})
}
func Multiple_slot_booking(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	id := c.Param("id")

	ID, _ := strconv.Atoi(id)

	fmt.Println(ID)

	Id := uint(ID)

	fmt.Println(Id)

	var body struct {
		Start_date string
		End_date   string
		Slots      []int
		Branch_id  int
	}
	if c.Bind(&body) != nil {
		logrus.Infof("Failed to read body")
		c.JSON(http.StatusBadRequest, gin.H{

			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}

	Booking_id, _ := uuid.NewRandom()

	B_id := Booking_id.String()

	// Parse start and end dates as time objects
	startDate, err := time.Parse("02-01-2006", body.Start_date)
	if err != nil {
		logrus.Infof("Failed to get start date format %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Invalid start date format",
			"data":   nil,
		})
		return
	}

	endDate, err := time.Parse("02-01-2006", body.End_date)
	if err != nil {
		logrus.Infof("Failed to get end date format", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Invalid end date format",
			"data":   nil,
		})
		return
	}

	// Loop through dates and create bookings
	for currentDate := startDate; currentDate.Before(endDate) || currentDate.Equal(endDate); currentDate = currentDate.AddDate(0, 0, 1) {
		for i := 0; i < len(body.Slots); i++ {
			var psr models.Package_slot_relationship

			config.DB.First(&psr, "slot_id=?", int(body.Slots[i]))

			// Fetch the price based on package id retrieved
			var price models.Package
			config.DB.Find(&price, "id=?", psr.Package_id)
			price25 := percent.PercentFloat(25.0, price.Price)

			booking := models.Turf_Bookings{
				User_id:                  Id,
				Date:                     currentDate.Format("02-01-2006"),
				Slot_id:                  body.Slots[i],
				Package_slot_relation_id: int(psr.ID),
				Package_id:               psr.Package_id,
				Price:                    price.Price,
				Minimum_amount_to_pay:    price25,
				Order_id:                 B_id,
				Is_booked:                4,
				Branch_id:                body.Branch_id,
			}

			result := config.DB.Create(&booking)
			if result.Error != nil {
				logrus.Infof("Failed to get data from DB %v\n", result.Error)
				c.JSON(http.StatusOK, gin.H{
					"status": 400,
					"error":  "Slot Already Exist",
					"data":   nil,
				})
				return
			}
		}
	}

	var booking models.Turf_Bookings

	// Confirm booking table
	config.DB.Find(&booking, "order_id = ?", B_id)

	var totalPrice float64
	var total_min_amount float64
	for p := 0; p < len(body.Slots); p++ {
		totalPrice += booking.Price
		total_min_amount += booking.Minimum_amount_to_pay
	}

	confirm_booking := models.Confirm_Booking_Table{
		User_id:                 Id,
		Date:                    body.Start_date,
		Booking_order_id:        B_id,
		Total_price:             totalPrice,
		Total_min_amount_to_pay: total_min_amount,
		Booking_status:          4,
		Branch_id:               body.Branch_id,
	}

	result := config.DB.Create(&confirm_booking)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Slot Already Exist",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Slots reserved successfully",
		"data":    booking,
	})
}

func Get_Available_slots_Multi_Dates(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var body struct {
		Start_date string
		End_date   string
		Branch_id  int
	}
	err := c.Bind(&body)
	if err != nil {
		logrus.Infof("Failed to read body %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   nil,
		})
		return
	}

	// Fetch all time slots
	var slots []models.Time_Slot
	result := config.DB.Find(&slots)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	// Create a slice to store the final response for all dates
	var response []gin.H

	// Parse start and end dates as time objects
	startDate, err := time.Parse("02-01-2006", body.Start_date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Invalid start date format",
			"data":   nil,
		})
		return
	}

	endDate, err := time.Parse("02-01-2006", body.End_date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "Invalid end date format",
			"data":   nil,
		})
		return
	}

	// Loop through dates within the range
	for currentDate := startDate; currentDate.Before(endDate) || currentDate.Equal(endDate); currentDate = currentDate.AddDate(0, 0, 1) {
		// Fetch booked slots for the current date
		var bookedSlots []models.Turf_Bookings
		result = config.DB.Where("date = ? AND is_booked IN (1, 2, 3, 4) AND branch_id = ?", currentDate.Format("02-01-2006"), body.Branch_id).Find(&bookedSlots)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to find booked slots for date",
			})
			return
		}

		// Create a map to store booked slots with their is_booked status
		bookedSlotMap := make(map[int]int)
		for _, bookedSlot := range bookedSlots {
			bookedSlotMap[bookedSlot.Slot_id] = bookedSlot.Is_booked
		}

		// Create a slice to store available slots for the current date
		var availableSlots []gin.H
		for _, slot := range slots {
			isBooked, exists := bookedSlotMap[int(slot.ID)]
			if !exists {
				isBooked = 1
			}

			var psr models.Package_slot_relationship
			result = config.DB.Where("slot_id = ?", slot.ID).Find(&psr)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Failed to find Package_slot_relationship for slot",
				})
				return
			}

			var price models.Package
			result = config.DB.Where("id = ?", psr.Package_id).Find(&price)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Failed to find Package for slot",
				})
				return
			}

			availableSlots = append(availableSlots, gin.H{
				"Slot":      slot,
				"Is_booked": isBooked,
				"Package":   price.Name,
				"Price":     price.Price,
			})
		}

		response = append(response, gin.H{
			"Date":            currentDate.Format("02-01-2006"),
			"Available_slots": availableSlots,
		})
	}

	// Return the available slots for each date within the range
	c.JSON(http.StatusOK, gin.H{
		"available_slots": response,
	})
}

func GetLoggedAdmin(c *gin.Context) {

	cookie, err := c.Cookie("AID")

	fmt.Println("Authorization Cookie Value:", cookie)

	if err != nil || cookie == "" {
		fmt.Println("ccookie nt found")
		return
	}

	var admin models.Admin

	result := config.DB.Find(&admin, "id=?", cookie)
	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "not f0und",
			"data":   nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "admin details",
		"data":   admin,
	})

}

func MultipleImages(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")

	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	// Get the booking_order_id from the URL
	id := c.Param("id")

	var ss []models.Screenshot
	result := config.DB.Find(&ss, "booking_order_id=?", id)

	if result.Error != nil {
		logrus.Infof("Failed to get data from DB %v\n", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to fetch screenshot image",
			"data":   nil,
		})
		return
	}

	// Loop through the images associated with the order_id
	for _, screenshot := range ss {
		if strings.HasSuffix(screenshot.Payment_screenshot, ".jpg") {
			// Set the content type for JPEG images
			c.Header("Content-Type", "image/jpeg")
		} else if strings.HasSuffix(screenshot.Payment_screenshot, ".png") {
			// Set the content type for PNG images
			c.Header("Content-Type", "image/png")
		} else {

			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "unsupported image format",
				"data":   nil,
			})
			return
		}

		// Read the image file and send it as a response
		imageData, err := ioutil.ReadFile(screenshot.Payment_screenshot)
		if err != nil {
			fmt.Println("Error reading the image file:", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"error":  "internal server error",
				"data":   nil,
			})
			return
		}

		// Log the ID and file path
		fmt.Printf("ID: %s, File Path: %s\n", id, screenshot.Payment_screenshot)

		c.Data(http.StatusOK, c.GetHeader("Content-Type"), imageData)
	}
}

func Graph_API(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Referer, Sec-Ch-Ua, Sec-Ch-Ua-Mobile, Sec-Ch-Ua-Platform, User-Agent")
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	var revenue []models.Confirm_Booking_Table
	var count int64

	// role, _ := c.Request.Cookie("Role")
	// Role, _ := strconv.Atoi(role.Value)
	branchID, _ := c.Request.Cookie("Branch_id")
	branchid, _ := strconv.Atoi(branchID.Value)

	// fmt.Println(Role)

	// if Role != 1 {

	var user models.User

	result := config.DB.Model(&user).Count(&count)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Total User Count Unsuccessfully",
			"data":   nil,
		})
		return
	}

	result = config.DB.Find(&revenue, "branch_id=?", branchid)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 404,
			"error":  "failed to get all branch details ",
		})
		return

	}

	// } else {

	// 	var user models.User

	// 	result := config.DB.Model(&user).Count(&count)

	// 	if result.Error != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"status": 400,
	// 			"error":  "Total User Count Unsuccessfully",
	// 			"data":   nil,
	// 		})
	// 		return
	// 	}

	// 	result = config.DB.Find(&revenue)
	// 	if result.Error != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"status": 400,
	// 			"error":  "failed to Delete Heading Details",
	// 			"data":   nil,
	// 		})
	// 		return
	// 	}
	// }

	//Calculate the total Paid_amount
	totalPaidAmount := 0.0
	for _, booking := range revenue {
		totalPaidAmount += booking.Paid_amount
	}

	totalSales := 0.0
	for _, booking := range revenue {
		totalSales += booking.Total_price
	}

	totalUser := float64(count)

	rati0 := (totalSales / totalUser)

	Average_Revenue_per_User := (totalPaidAmount / totalUser)

	c.JSON(http.StatusOK, gin.H{
		"status":           200,
		"msg":              "revenue",
		"data":             rati0,
		"Revenue_per_User": Average_Revenue_per_User,
	})
}
func PackageNameList(c *gin.Context) {
	var packages []models.Package
	result := config.DB.Select("DISTINCT name").Find(&packages)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Get Package Name List Unsuccessfully",
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Get Package Name List",
		"data":   packages,
	})
}

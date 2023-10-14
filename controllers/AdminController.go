package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"turf/config"
	"turf/models"

	"github.com/dariubs/percent"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func AdminSignup(c *gin.Context) {
	var body struct {
		Name        string
		Contact     string
		Password    string
		Email       string
		Role        int
		Status      string
		Branch_name string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   "null",
		})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to hash password",
			"data":   "null",
		})
		return
	}
	if body.Status == "Super Admin" {
		body.Role = 1
	} else if body.Status == "Admin" {
		body.Role = 2
	} else if body.Status == "Staff" {
		body.Role = 3
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Select a Valid Role",
			"data":   "null",
		})
		return
	}
	var branch models.Branch_info_management
	result := config.DB.Find(&branch, "branch_name=?", body.Branch_name)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Error finding branch id",
			"data":   "null",
		})
		return
	}

	bodys := models.Admin{
		Name:           body.Name,
		Contact:        body.Contact,
		Password:       string(password),
		Email:          body.Email,
		Role:           body.Role,
		Turf_branch_id: branch.ID,
	}

	result = config.DB.Create(&bodys)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Admin Allready Exist",
			"data":   "null",
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

func AdminLogin(c *gin.Context) {
	var body struct {
		Name     string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to hash password",
			"data":   "null",
		})
		return
	}

	fmt.Println(body.Name)
	fmt.Println(body.Password)

	var admin models.Admin
	config.DB.Table("admins").Select("id", "name", "password").Where("name", body.Name).Scan(&admin)
	fmt.Println(admin)
	fmt.Println(admin.ID)
	if admin.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"Error":  "Admin Does Not Exist",
			"data":   "null",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Invalid Password",
			"data":   "null",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": admin.ID,
		"exp": time.Now().Add(time.Minute * 2).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Failed to create token",
			"data":   "null",
		})
		return
	}

	// send the generated jwt token back & set it in cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 120, "", "", false, true)
	admin.LastLogin = time.Now()
	config.DB.Save(&admin)
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
// 			"data":   "null",
// 		})
// 		return

// 	}
// 	password, err := bcrypt.GenerateFromPassword([]byte(body.NewPassword), 14)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": 400,
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
// 			"status": 400,
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
// func Select_branch(c *gin.Context) {
// 	var branch models.Branch_info_management
// 	result := config.DB.Select("branch_name").Find(&branch)
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": 400,
// 			"error":  "Failed to get branch",
// 			"data":   "null",
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
	var body struct {
		Turf_name             string
		Branch_name           string
		Branch_address        string
		Branch_email          string
		Branch_contact_number string
		GST_no                string
		Status                int
		Ground_Size           string
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
	branch := models.Branch_info_management{Turf_name: body.Turf_name, Branch_name: body.Branch_name, Branch_email: body.Branch_email, Branch_contact_number: body.Branch_contact_number, Branch_address: body.Branch_address, GST_no: body.GST_no, Status: 1, Ground_Size: body.Ground_Size}
	result := config.DB.Create(&branch)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
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
func Update_Branch(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Turf_name             string
		Branch_name           string
		Branch_address        string
		Branch_email          string
		Branch_contact_number string
		GST_no                string
		Status                int
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   "null",
		})
		return

	}

	branch := models.Branch_info_management{Turf_name: body.Turf_name, Branch_name: body.Branch_name, Branch_email: body.Branch_email, Branch_contact_number: body.Branch_contact_number, Branch_address: body.Branch_address, GST_no: body.GST_no, Status: body.Status}
	result := config.DB.Model(&branch).Where("id=?", id).Updates(&branch)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Branch Update unsuccessfully",
			"data":   "null",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"success": "Branch Successfully Updated",
		"data":    branch,
	})
}

// func GET_All_Branch(c *gin.Context) {
// 	var branch []models.Branch_info_management
// 	result := config.DB.Find(&branch)
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": 400,
// 			"error":  "No Branch Found",
// 			"data":   "null",
// 		})
// 		return
// 	}

// 	//Response
// 	c.JSON(http.StatusCreated, gin.H{
// 		"status":  200,
// 		"success": "All Branch  Successfully",
// 		"data":    branch,
// 	})

// }
func AddSlot(c *gin.Context) {
	var body struct {
		StartSlot string
		EndSlot   string
		Day       string
		Branch_id int
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

	var branch models.Branch_info_management
	config.DB.First(&branch, "id = ?", body.Branch_id)

	fmt.Println(branch.ID)

	First_three_initials := body.Day[:3]

	usid := First_three_initials + "/" + body.StartSlot + "/" + body.EndSlot

	slot := models.Time_Slot{Start_time: body.StartSlot, End_time: body.EndSlot, Day: body.Day, Branch_id: branch.ID, Unique_slot_id: usid, Status: 1}
	result := config.DB.Create(&slot)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
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

// add package by admin
func AddPackage(c *gin.Context) {
	var body struct {
		Name      string
		Price     float64
		Status    int
		Branch_id int
		Slot_id   []string
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

	packageModel := &models.Package{Name: body.Name, Price: body.Price, Status: body.Status, Branch_id: body.Branch_id}
	result := config.DB.Create(&packageModel)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "package Allready Exist",
			"data":   "null",
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
		// slots = append(slots, psrmodel.Slot_id)
	}
}

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
				"status": 400,
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
				"status": 400,
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
			"status": 400,
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

func Get_Slot_by_day(c *gin.Context) {
	var body struct {
		Day []string
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
	var slot []models.Time_Slot

	var days []interface{}

	for i := 0; i < len(body.Day); i++ {

		result := config.DB.Find(&slot, "day = ?", body.Day[i])

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "failed to get all slot",
			})
			return

		}
		days = append(days, slot)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "slot details",
		"data":    days,
	})

}

func UpdatePackage(c *gin.Context) {
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
			"status": 400,
			"error":  "Package Update Unsuccessful",
			"data":   "sry",
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
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Package Update UnSuccessfully",
				"data":   "null",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Package Update Successfully",
		"data":    admin,
	})

}

func Get_Package(c *gin.Context) {
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
		Success: "Data retrieved successfully",
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
			"status": 400,
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

// func UpdateUserDetails(c *gin.Context) {
// 	Id := c.Param("id")
// 	var body struct {
// 		Full_Name      string
// 		Email          string
// 		Password       string
// 		Contact        string
// 		Account_Status string
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
// 	Hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "failed to hash password",
// 		})
// 		return
// 	}
// 	var Account_Status int
// 	switch {
// 	case body.Account_Status == "Active":
// 		Account_Status = 1
// 	case body.Account_Status == "Hold":
// 		Account_Status = 2
// 	case body.Account_Status == "Block":
// 		Account_Status = 3
// 	default:
// 		Account_Status = 0
// 	}

// 	users := models.User{Full_Name: body.Full_Name, Email: body.Email, Contact: body.Contact, Password: string(Hash), Account_Status: Account_Status}
// 	result := config.DB.Model(&users).Where("id = ?", Id).Updates(users)
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": 400,
// 			"error":  "User Update UnSuccessfully",
// 			"data":   "null",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  200,
// 		"success": "User Update Successfully",
// 		"data":    body,
// 	})

// }

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
	id := c.Param("id")
	var body struct {
		Amount float64
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Invalid Request",
			"data":   "null",
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to insert",
			"data":   "null",
		})
		return
	} else {
		changed_status := models.Confirm_Booking_Table{
			Booking_status: 3,
		}
		var booking models.Confirm_Booking_Table
		status := config.DB.Model(&booking).Where("booking_order_id = ?", booking.Booking_order_id).Updates(changed_status)
		if status.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "failed to insert",
				"data":   "null",
			})
			return
		}
		var turf_book models.Turf_Bookings

		is_booked := models.Turf_Bookings{
			Is_booked: 3,
		}
		result := config.DB.Model(&turf_book).Where("order_id = ?", booking.Booking_order_id).Updates(is_booked)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "failed to insert",
				"data":   "null",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "Successfully upladed",
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
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   "null",
		})
		return
	}

	rows := config.DB.Model(&models.Turf_Bookings{}).Where("date = ?", body.Date).Pluck("slot_id", &Slots)

	if rows.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to read body",
			"data":   "null",
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

	// availableSlots1 := []int{}
	// for _, s := range body.Slot {
	// 	for _, s1 := range availableSlots {
	// 		if s != s1 {
	// 			fmt.Println(s)
	// 			availableSlots1 = append(availableSlots1, int(s))
	// 		}
	// 	}
	// }
	// fmt.Println("ava1", availableSlots1)
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
				c.JSON(http.StatusOK, gin.H{
					"status": 400,
					"error":  "Slot Allready Exist",
					"data":   "null",
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
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Slot Allready Exist",
				"data":   "null",
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
				c.JSON(http.StatusOK, gin.H{
					"status": 400,
					"error":  "Slot Allready Exist",
					"data":   "null",
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
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"error":  "Slot Allready Exist",
				"data":   "null",
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
			"data":   "null",
		})
	}
}

func LiveUser(c *gin.Context) {
	var live []interface{}
	now := time.Now()
	date := now.Format("02:01:2006")
	time := now.Format("15:04:05")
	var slot models.Time_Slot
	result := config.DB.Where("start_time <= ? AND end_time >= ?", time, time).Find(&slot)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  "slot is not found",
			"data":   "null",
		})
		return
	}

	var booking models.Turf_Bookings
	result = config.DB.Where("date=? AND slot_id=?", date, slot.ID).Find(&booking)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  "booking detail not found",
			"data":   "null",
		})
		return
	}

	var user models.User
	result = config.DB.Where("id", booking.User_id).Find(&user)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 400,
			"error":  "user  not found",
			"data":   "null",
		})
		return
	}

	live = append(live, booking, user)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "user fetch successfully",
		"data":    live,
	})

}
func Testimonials(c *gin.Context) {
	var body struct {
		Name        string
		Designation string
		Review      string
		Image       string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "faild to read body",
			"data":   "null",
		})
		return

	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filePath := filepath.Join("./uploads/testi_monials", file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	if filepath.Ext(filePath) != ".jpg" && filepath.Ext(filePath) != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Upload the right file format (jpg or png)",
			"data":   "null",
		})
		return
	}
	testimonial := &models.Testi_Monial{Name: body.Name, Designation: body.Designation, Review: body.Review, Image: filePath}
	result := config.DB.Create(&testimonial)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to create testimonials",
			"data":   "null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "testimonials create successfully",
		"data":    testimonial,
	})
}
func Upadte_TestiMonilas(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Name        string
		Designation string
		Review      string
		Image       string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "faild to read body",
			"data":   "null",
		})
		return

	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filePath := filepath.Join("./uploads/testi_monials", file.Filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	if filepath.Ext(filePath) != ".jpg" && filepath.Ext(filePath) != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "Upload the right file format (jpg or png)",
			"data":   "null",
		})
		return
	}

	testimonial := &models.Testi_Monial{Name: body.Name, Designation: body.Designation, Review: body.Review, Image: filePath}
	result := config.DB.Model(&testimonial).Where("id=?", id).Updates(&testimonial)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to create testimonials",
			"data":   "null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "testimonials create successfully",
		"data":    testimonial,
	})

}
func AllTestimonials(c *gin.Context) {
	var testimonials []models.Testi_Monial
	result := config.DB.Find(&testimonials)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"error":  "failed to fetch testimonials",
			"data":   "null",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "testimonials fetch successfully",
		"data":    testimonials,
	})

}

func AdminLogout(c *gin.Context) {
	// Clear the "Authorization" cookie to log out
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	// You can also clear any other session-related data if needed
	c.Set("UserID", "")

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Logged out successfully",
	})
}

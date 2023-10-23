package controllers

import (
	"fmt"
	"math/rand"
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
	"gopkg.in/gomail.v2"
)

func Signup(c *gin.Context) {
	// Get the email & password from req body
	//huhdjhkejhk

	var body struct {
		Full_Name string
		Email     string
		Password  string
		Contact   string
		Is_active string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	//Hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	//create the user
	user := models.User{Full_Name: body.Full_Name, Email: body.Email, Password: string(hash), Contact: body.Contact, Account_Status: "0"}

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})

		return
	} else {
		var user models.User
		config.DB.First(&user, "email = ?", body.Email)
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid Email or Password",
			})
			return
		}
		var err error
		email := body.Email

		// otpChannel := make(chan string)

		otp := generateOTP()

		go storeOTP(email, otp)

		go func() {
			err := sendOTPEmail(email, otp)
			if err != nil {
				// Handle the error (e.g., log it)
			}
		}()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to send the email",
			})

		}

		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "user successfully created",
			"data":    user,
		})
		return
	}

}

type OTPData struct {
	Email      string
	Contact    string
	OTP        string
	Expiration time.Time
}

var otpStore = make(map[string]OTPData)

func generateOTP() string {

	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(900000) + 100000
	return strconv.Itoa(otp)
}

func storeOTP(email, otp string) {
	expiration := time.Now().Add(15 * time.Minute)
	otpStore[email] = OTPData{Email: email, OTP: otp, Expiration: expiration}
}
func sendOTPEmail(email, otp string) error {

	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	smtpUsername := "rakshawd@gmail.com"
	smtpPassword := "aoesrthiacvxnpnv"

	m := gomail.NewMessage()
	m.SetHeader("From", "rakshawd@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "OTP for Password Reset")
	m.SetBody("text/plain", "Your OTP is: "+otp)

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)

	}

	fmt.Println("Email sent successfully")

	return nil
}

// func SendEmailOTP() {

// 	var body struct {
// 		Email string
// 	}

// 	var user models.User
// 	config.DB.First(&user, "email = ?", body.Email)

// 	if user.ID == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid Email or Password",
// 		})
// 		return
// 	}
// 	var err error
// 	email := body.Email

// 	// otpChannel := make(chan string)

// 	otp := generateOTP()

// 	go storeOTP(email, otp)

// 	go func() {
// 		err := sendOTPEmail(email, otp)
// 		if err != nil {
// 			// Handle the error (e.g., log it)
// 		}
// 	}()

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "failed to send the email",
// 		})

// 	}

// }

func VerifyOTPhandler(c *gin.Context) {

	var body struct {
		Email     string
		Otp       string
		Is_active string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	otpData, ok := otpStore[body.Email]

	//var user models.User

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"error":  "OTP verification unsuccessful",
			"data":   "null",
		})
		return
	}

	if body.Otp == otpData.OTP && time.Now().Before(otpData.Expiration) {

		config.DB.Exec("UPDATE users SET is_active = 1 WHERE email = ?", body.Email)

		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"success": "OTP verfication successful",
			"data":    "null",
		})

	} else {

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid OTP please try again",
		})
		return

	}
}

func Login(c *gin.Context) {
	// Get the email & pass off req body

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	//Look up for requested user
	var user models.User
	config.DB.First(&user, "email = ? AND account_status = 1", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password ",
		})
		return
	} else if user.ID != 0 && user.Account_Status == "0" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Kindly verify your email first",
		})
		return
	}

	//Compare sent password with saved password in hash format
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	// Generate a JWT Token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 4).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// send the generated jwt token back & set it in cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*4, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Success",
		"data":    user,
	})
}

func Booking(c *gin.Context) {
	var body struct {
		Day       int
		Date      string
		Slot      []int
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

	// location, err := time.LoadLocation("Asia/Kolkata")
	// fmt.Println(location)
	// if err != nil {
	// 	// Handle the error, e.g., log it or set a default time zone
	// 	location = time.UTC // Default to UTC in case of an error
	// }

	var Slots []int

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
		var user models.User
		config.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusNotFound)
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

		fmt.Println("ava:", availableSlots)
		if len(availableSlots) == 0 {

			Booking_id, _ := uuid.NewRandom()

			B_id := Booking_id.String()

			for i := 0; i < len(body.Slot); i++ {
				var price models.Package
				var psr models.Package_slot_relationship

				config.DB.First(&psr, "slot_id=?", int(body.Slot[i]))

				//fetch the price based on package id retrieved

				config.DB.Find(&price, "id=?, status=1", psr.Package_id)

				price25 := percent.PercentFloat(25.0, price.Price)

				booking := models.Turf_Bookings{Date: body.Date, Slot_id: body.Slot[i], User_id: user.ID, Package_slot_relation_id: int(psr.ID), Package_id: psr.Package_id, Price: price.Price, Minimum_amount_to_pay: price25, Order_id: B_id, Is_booked: 2, Branch_id: body.Branch_id}
				result := config.DB.Create(&booking)
				if result.Error != nil {
					c.JSON(http.StatusOK, gin.H{
						"status": 400,
						"error":  "Slot Already Exist",
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

			confirm_booking := models.Confirm_Booking_Table{Date: body.Date, User_id: user.ID, Booking_order_id: B_id, Total_price: totalPrice, Total_min_amount_to_pay: total_min_amount, Booking_status: 2, Branch_id: body.Branch_id}

			result := config.DB.Create(&confirm_booking)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "400",
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

				booking := models.Turf_Bookings{Date: body.Date, Slot_id: uniqueslots[i], User_id: user.ID, Package_slot_relation_id: int(psr.ID), Package_id: psr.Package_id, Price: price.Price, Minimum_amount_to_pay: price25, Order_id: B_id, Is_booked: 2, Branch_id: body.Branch_id}
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

			confirm_booking := models.Confirm_Booking_Table{Date: body.Date, User_id: user.ID, Booking_order_id: B_id, Total_price: totalPrice, Total_min_amount_to_pay: total_min_amount, Booking_status: 2, Branch_id: body.Branch_id}

			result := config.DB.Create(&confirm_booking)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "400",
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
}

// func Final_booking(c *gin.Context) {
// 	var body struct {
// 		Day       string
// 		Date      string
// 		Slot      []int
// 		Branch_id int
// 	}

// 	var Slots []int
// 	err := c.Bind(&body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": 400,
// 			"error":  "failed to read body",
// 			"data":   "null",
// 		})
// 		return
// 	}

// 	tokenString, err := c.Cookie("Authorization")

// 	if err != nil {
// 		c.AbortWithStatus(http.StatusUnauthorized)
// 	}

// 	// decode & validate the same

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 		return []byte(os.Getenv("SECRET")), nil
// 	})

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		// check expiration
// 		if float64(time.Now().Unix()) > claims["exp"].(float64) {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}

// 		// find the user with token sub i.e user id
// 		var user models.User
// 		config.DB.First(&user, claims["sub"])

// 		if user.ID == 0 {
// 			c.AbortWithStatus(http.StatusNotFound)
// 		}
// 		rows := config.DB.Model(&models.Turf_Bookings{}).Where("date = ?", body.Date).Pluck("slot_id", &Slots)

// 		if rows.Error != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status": 400,
// 				"error":  "failed to read body",
// 				"data":   "null",
// 			})
// 			return
// 		}
// 		availableSlots := []int{}
// 		for _, s := range body.Slot {
// 			for _, s1 := range Slots {
// 				if s == s1 {
// 					availableSlots = append(availableSlots, int(s))
// 				}
// 			}
// 		}
// 		uniqueslots := make([]int, 0)

// 		bMap := make(map[int]bool)
// 		for _, val := range availableSlots {
// 			bMap[val] = true

// 		}

// 		for _, val := range body.Slot {

// 			if !bMap[val] {
// 				uniqueslots = append(uniqueslots, val)

// 			}
// 		}

// 		if len(availableSlots) == 0 {

// 			Booking_id, _ := uuid.NewRandom()

// 			B_id := Booking_id.String()

// 			for i := 0; i < len(body.Slot); i++ {
// 				var price models.Package
// 				var psr models.Package_slot_relationship

// 			}
// 		}
// 	}

// }

// func Get_avl_days(c *gin.Context) {
// 	var pack models.Package

// 	result := config.DB.Find(&pack, "id = 1")
// 	if result.Error != nil {
// 		fmt.Println("fdjhkdjhkd")
// 	}

// 	fmt.Println(pack.Avail_days)

// }

func Screenshot(c *gin.Context) {
	var err error
	var body struct {
		Amount float64
	}

	c.Bind(&body)
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
		var user models.User
		config.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		}

		// img := c.Request.FormValue("image")

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Define the path where the file will be saved
		filePath := filepath.Join("./uploads", file.Filename)
		// Save the file to the defined path
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		var booking models.Confirm_Booking_Table

		config.DB.Find(&booking, "user_id = ?", user.ID)

		payment := models.Screenshot{Payment_screenshot: filePath, Booking_order_id: booking.Booking_order_id}
		result := config.DB.Create(&payment)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "400",
				"error":  "failed to insert",
				"data":   "null",
			})
			return
		} else {
			changed_status := models.Confirm_Booking_Table{
				Booking_status: 3,
			}
			status := config.DB.Model(&booking).Where("booking_order_id = ?", booking.Booking_order_id).Updates(changed_status)
			if status.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "400",
					"error":  "failed to insert",
					"data":   "null",
				})
				return
			}
			var turf_book models.Turf_Bookings

			is_booked := models.Turf_Bookings{
				Is_booked: 3,
			}
			result = config.DB.Model(&turf_book).Where("order_id = ?", booking.Booking_order_id).Updates(is_booked)
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "400",
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
}

func AvailableSlot(c *gin.Context) {

	// slot go routine running
	go Slot_go_rountine()

	var body struct {
		Date time.Time
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
	var slots []models.Turf_Bookings
	var bookSlots, AllSlot []int
	var slot []models.Time_Slot
	result := config.DB.Find(&slot)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}
	currentTime := time.Now()
	date := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second(), 0, currentTime.Location())

	fmt.Println(date)

	result = config.DB.Where("date = ? AND is_booked = ?", date, 1).Find(&slots)
	if result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	for _, s := range slots {
		bookSlots = append(bookSlots, s.Slot_id)
	}

	for _, s := range slot {
		AllSlot = append(AllSlot, int(s.ID))
	}

	availableSlots := []int{}
	for _, s := range AllSlot {
		if !contains(bookSlots, s) {
			availableSlots = append(availableSlots, s)
		}
	}
	var availableSlots1 []map[string]interface{}
	// fmt.Println(availableSlots)
	Data := make(map[string]interface{})
	for _, s := range availableSlots {
		var slt models.Time_Slot
		result := config.DB.Where("id = ? ", s).Find(&slt)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to find slot by start_slot",
			})
			return

		}

		Data = map[string]interface{}{
			"id":        slt.ID,
			"starttime": slt.Start_time,
			"endtime":   slt.End_time,
		}
		availableSlots1 = append(availableSlots1, Data)

	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "Get Available Slot Successfully ",
		"data":    availableSlots1,
	})
}

func contains(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
func GetAllDetail(c *gin.Context) {
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
	var user models.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// find the user with token sub i.e user id

		config.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		}

		result := config.DB.Find(&user, claims["sub"])

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 404,
				"error":  "failed to fatch user detail",
			})
			return

		}

		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"success": "fatch user detail successfully",
			"data":    user,
		})
	}
}
func GetBookingDetail(c *gin.Context) {
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
	var user models.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// find the user with token sub i.e user id

		config.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		}
		var booking []models.Confirm_Booking_Table
		result := config.DB.Find(&booking).Where("slot_id", claims["sub"])

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 404,
				"error":  "failed to fatch booking detail",
			})
			return

		}

		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"success": "fatch booking detail successfully",
			"data":    booking,
		})
	}

}
func UpdateUser(c *gin.Context) {
	var body struct {
		Full_Name   string
		Email       string
		OldPassword string
		Password    string
		Contact     string
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
		var user models.User
		config.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusNotFound)
		}
		// OldPassword, err := bcrypt.GenerateFromPassword([]byte(body.OldPassword), 10)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"error": "failed to hash password",
		// 	})
		// 	return
		// }

		result := config.DB.Find(&user).Where("id = ?", claims["sub"])
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "400",
				"error":  "User Update UnSuccessfully",
				"data":   "null",
			})
			return
		}
		if body.OldPassword != "" {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.OldPassword))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid Email or Password",
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

			users := models.User{Full_Name: body.Full_Name, Email: body.Email, Contact: body.Contact, Password: string(Hash)}
			result = config.DB.Model(&user).Where("id = ?", claims["sub"]).Updates(users)
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

		} else {
			users := models.User{Full_Name: body.Full_Name, Contact: body.Contact}
			result = config.DB.Model(&user).Where("id = ?", claims["sub"]).Updates(users)
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

	}
}

func Slot_go_rountine() {
	for {
		config.DB.Exec("UPDATE confirm_booking_tables SET booking_status = 1, deleted_at=NOW() WHERE DATE_ADD(created_at, INTERVAL 15 MINUTE ) < NOW() AND booking_status = 2")

		config.DB.Exec("UPDATE turf_bookings SET is_booked = 1, deleted_at=NOW() WHERE DATE_ADD(created_at, INTERVAL 15 MINUTE ) < NOW() AND is_booked=2")

		fmt.Println("goroutine running")
		time.Sleep(30 * time.Second)
	}

}
func GetAllDetailbydate(c *gin.Context) {

	startdate := c.PostForm("startdate")
	enddate := c.PostForm("enddate")
	var booking []models.Turf_Bookings
	result := config.DB.Find(&booking).Where("date >= ? AND date <=?", startdate, enddate)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "400",
			"error":  "fetch booking detail UnSuccessfully",
			"data":   "null",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": "fetch booking detail Successfully",
		"data":    booking,
	})
}

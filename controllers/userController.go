package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
	"turf/config"
	"turf/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	user := models.User{Full_Name: body.Full_Name, Email: body.Email, Password: string(hash), Contact: body.Contact, Is_active: 0}

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
			"status": 200,
			"sucess": "OTP verfication successful",
			"data":   "null",
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
	config.DB.First(&user, "email = ? AND is_active=1", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password ",
		})
		return
	} else if user.ID != 0 && user.Is_active == 0 {
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
		Date      string
		Day       string
		Slot      int
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
	// AvailableSlot(body.Date)
	var slot models.Time_Slot
	result := config.DB.Where("start_time = ? AND end_time >= ?", body.StartSlot, body.EndSlot).Find(&slot)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to find slot by start_slot",
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

		booking := models.Turf_Bookings{Date: body.Date, Slot_id: int(slot.ID), User_id: user.ID}
		result = config.DB.Create(&booking)
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
			"success": "slot reserved successfully",
			"data":    booking,
		})
	}
}

// func AvailableSlot(date string) {
// 	var slots []models.Booking
// 	var bookSlots, AllSlot []int

// 	var slot []models.Slot
// 	result := config.DB.Find(&slot)
// 	if result.Error != nil {
// 		fmt.Println(result.Error)
// 		return
// 	}
// 	result = config.DB.Where("date = ?", date).Find(&slots)
// 	if result.Error != nil {
// 		fmt.Println(result.Error)
// 		return
// 	}

// 	for i := 0; i < len(slots); i++ {
// 		bookSlots = append(bookSlots, slots[i].Slot)
// 	}

// 	for i := 0; i < len(slot); i++ {
// 		AllSlot = append(AllSlot, int(slot[i].ID))
// 	}
// 	fmt.Println(bookSlots)
// 	fmt.Println(AllSlot)

// }
func AvailableSlot(c *gin.Context) {
	var body struct {
		Date string
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

	result = config.DB.Where("date = ?", body.Date).Find(&slots)
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

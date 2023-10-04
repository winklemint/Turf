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
		Email     string
		Password  string
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
	user := models.User{Email: body.Email, Password: string(hash), Is_active: 0}

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
		return
	} else {
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

func SendEmailOTP(c *gin.Context) {
	var body struct {
		Email string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

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

}

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
		// var err error
		// p, _ := uuid.NewRandom()
		// qrpid := "P" + p.String()

		// formData := url.Values{}

		// formData.Set("referralid", qrpid)

		// baseURL := "https://mgcworld88.com/registration-player"
		// fullURL := baseURL + "?" + formData.Encode()

		// imageName := qrpid + ".png"

		// imagePath := "uploads/playerQR"

		// err = qrcode.WriteFile(fullURL, qrcode.Medium, 256, imagePath, imageName)
		// if err == nil {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"message": "Failed to generate qr",
		// 	})
		// }

		// fullpath := imagePath + imageName

		// config.DB.Exec("UPDATE users SET user_qr=?, user_ref_id=?, is_active = 1 WHERE email = ?", fullpath, qrpid, body.Email)

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

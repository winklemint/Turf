package routes

import (
	"fmt"
	"net/http"
	"turf/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAdminPanelRoutes(router *gin.Engine) {

	router.LoadHTMLGlob("templates/*.html")

	formGroup := router.Group("/admin/login")
	formGroup.Use(IsAuthenticated())

	formGroup.StaticFS("/", http.Dir("templates")) // Serve all files first

}

func RegisterAdminPanelDashboard(router *gin.Engine) {
	// Load HTML templates for the admin panel
	//router.LoadHTMLGlob("templates/*.html")
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}

func RegisterAdminPanelCreateBranch(router *gin.Engine) {
	// Load HTML templates for the admin panel
	//router.LoadHTMLGlob("templates/*.html")
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/create/branch", func(c *gin.Context) {
		c.HTML(http.StatusOK, "createbranch.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}

func RegisterAdminPanelAllBranch(router *gin.Engine) {
	// Load HTML templates for the admin panel
	//router.LoadHTMLGlob("templates/*.html")
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/all/branch", func(c *gin.Context) {
		c.HTML(http.StatusOK, "allbranch.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}
func RegisterAdminPanelAllTestiMonials(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/testimonials", func(c *gin.Context) {
		c.HTML(http.StatusOK, "alltestimonials.html", nil)
	})

}
func RegisterAdminPanelUpdateTestiMonials(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/data/testimonials", func(c *gin.Context) {
		// Retrieve the "id" query parameter from the request URL
		id := c.DefaultQuery("id", "default_value_if_not_provided")

		// Now, you can use the "id" variable in your code as needed.
		// For example, you can use it to fetch data related to this ID.

		// Render your HTML template (updatetestimonials.html) with the data
		c.HTML(http.StatusOK, "updatetestimonials.html", gin.H{"id": id})
	})

}
func RegisterAdminPaneladdContent(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/content", func(c *gin.Context) {
		// Retrieve the "id" query parameter from the request URL
		id := c.DefaultQuery("id", "default_value_if_not_provided")

		// Now, you can use the "id" variable in your code as needed.
		// For example, you can use it to fetch data related to this ID.

		// Render your HTML template (updatetestimonials.html) with the data
		c.HTML(http.StatusOK, "content.html", gin.H{"id": id})
	})

}
func RegisterAdminPanelUpdatecarousel(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")
	router.GET("update/carousel", func(c *gin.Context) {
		// Retrieve the "id" query parameter from the request URL
		id := c.DefaultQuery("id", "default_value_if_not_provided")

		// Now, you can use the "id" variable in your code as needed.
		// For example, you can use it to fetch data related to this ID.

		// Render your HTML template (updatetestimonials.html) with the data
		c.HTML(http.StatusOK, "updatecarousel.html", gin.H{"id": id})
	})

}

func RegisterAdminPanelUpdateContent(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")
	router.GET("update/content", func(c *gin.Context) {
		// Retrieve the "id" query parameter from the request UR
		id := c.DefaultQuery("id", "default_value_if_not_provided")

		// Now, you can use the "id" variable in your code as needed.
		// For example, you can use it to fetch data related to this ID.

		// Render your HTML template (updatetestimonials.html) with the data
		c.HTML(http.StatusOK, "contentupdate.html", gin.H{"id": id})
	})

}
func RegisterAdminPanelUpdatebranchs(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")
	router.GET("update/branch", func(c *gin.Context) {
		// Retrieve the "id" query parameter from the request UR
		id := c.DefaultQuery("id", "default_value_if_not_provided")

		// Now, you can use the "id" variable in your code as needed.
		// For example, you can use it to fetch data related to this ID.

		// Render your HTML template (updatetestimonials.html) with the data
		c.HTML(http.StatusOK, "branchupdate.html", gin.H{"id": id})
	})

}
func RegisterAdminPanelUpdatepackage(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")
	router.GET("update/package", func(c *gin.Context) {
		// Retrieve the "id" query parameter from the request UR
		id := c.DefaultQuery("id", "default_value_if_not_provided")

		// Now, you can use the "id" variable in your code as needed.
		// For example, you can use it to fetch data related to this ID.

		// Render your HTML template (updatetestimonials.html) with the data
		c.HTML(http.StatusOK, "packageupdate.html", gin.H{"id": id})
	})

}
func RegisterAdminPanelUpdateUser(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")
	router.GET("data/user", func(c *gin.Context) {
		// Retrieve the "id" query parameter from the request UR
		id := c.DefaultQuery("id", "default_value_if_not_provided")

		// Now, you can use the "id" variable in your code as needed.
		// For example, you can use it to fetch data related to this ID.

		// Render your HTML template (updatetestimonials.html) with the data
		c.HTML(http.StatusOK, "userupdate.html", gin.H{"id": id})
	})

}

func RegisterAdminPanelAllPackages(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/packages", func(c *gin.Context) {
		c.HTML(http.StatusOK, "allpackages.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}
func RegisterAdminPanelAddCarousel(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/carousel", func(c *gin.Context) {
		c.HTML(http.StatusOK, "carousel.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}

func RegisterAdminPanelAddUser(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/add/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "useradd.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}
func RegisterAdminPanelAddTestimonials(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/add/testimonial", func(c *gin.Context) {
		c.HTML(http.StatusOK, "addtestimonial.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}
func RegisterAdminPanelAddPackages(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/add/package", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_package.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}

func RegisterAdminPanelAllSlots(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/slot", func(c *gin.Context) {
		c.HTML(http.StatusOK, "allslot.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}

func RegisterAdminPanelCreateSlots(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/create/slot", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create_slot.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}
func RegisterAdminPanelAllUser(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}

func RegisterAdminPanelPSR(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.html")

	// Define a route to serve the "dashboard.html" template
	router.GET("/psr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "package_slot.html", gin.H{
			// You can pass data to the template if needed
			//"data": "helloworld.html",
		})
	})
}


// Serve all files first

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {

		//uri := c.Param("path")
		// Check the request path
		requestedPath := c.Request.URL.Path

		// If the requested path is "/panel/index.html," allow access without login
		if requestedPath == "/panel/" {
			c.Next()
			return
		}

		// Check if the token is present in the request's cookies
		cookie, err := c.Cookie("Authorization")

		fmt.Println("Authorization Cookie Value:", cookie)

		if err != nil || cookie == "" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"Request": c,
			})
			c.Abort()
			return
		}

		middleware.RequireAdminAuth(c)

		c.Next()
	}
}

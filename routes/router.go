package routes

import (
	"github.com/gin-gonic/gin"
	"renovation/api/v1"
	"renovation/middleware"
	"renovation/utils"
	"github.com/gin-contrib/multitemplate"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	//p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.HTMLRender = createMyRender()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 32 << 20  // 8 MiB
	//静态图片资源
	r.Static("/upload", "./upload")

	r.Static("/static", "./web/front/dist")
	//r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	/*r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})*/

	//控制台页面接口
	admin := r.Group("api/v1")
	admin.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		admin.GET("admin/users", v1.GetUsers)
		admin.GET("admin/user/:id", v1.GetUserInfo)
		admin.POST("admin/add", v1.AddUser)
		//admin.PUT("admin/:id", v1.EditUser)
		admin.DELETE("admin/:id", v1.DeleteUser)
		//修改密码
		admin.PUT("admin/changepw/:id", v1.ChangeUserPassword)

		//岗位模块
		admin.POST("job/add", v1.AddJob)
		admin.PUT("job/:id", v1.EditJob)
		admin.DELETE("job/:id", v1.DeleteJob)
		//查询岗位下的所有简历
		admin.GET("job_res/:id", v1.GetJobRes)

		//简历模块
		admin.POST("resume/add", v1.AddResume)
		admin.DELETE("resume/:id", v1.DeleteRes)

		//设计团队模块
		admin.POST("team/add", v1.AddTeam)
		admin.DELETE("team/:id", v1.DeleteTeam)

		//商品图片模块
		admin.POST("commodity_pict/add", v1.AddCommodityPict)
		admin.DELETE("commodity_pict/:id", v1.DeleteCommodityPict)

		//商品模块
		admin.POST("commodity/add", v1.AddCommodity)
		admin.PUT("commodity/:id", v1.EditCommodity)
		admin.DELETE("commodity/:id", v1.DeleteCommodity)

		//logo模块
		admin.POST("logo/add", v1.AddLogo)//上线后关闭
		admin.PUT("logo/:id", v1.EditLogo)
		admin.DELETE("logo/:id", v1.DeleteLogo)//上线后关闭

		//轮播图模块
		admin.POST("carousel/add", v1.AddCarousel)
		admin.DELETE("carousel/:id", v1.DeleteCarousel)
	}

	//前端页面接口
	router := r.Group("api/v1")
	{
		/*//测试接口
		router.GET("hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"meg": "123",
			})
		})*/

		//后台登陆
		router.POST("login", v1.Login)

		//岗位模块
		router.GET("jobs", v1.GetJob)
		router.GET("job/:id", v1.GetJobInfo)

		//获取验证码
		router.GET("captcha", v1.GetCaptcha)

		//简历模块
		router.GET("resumes", v1.GetRes)
		router.GET("resume/:id", v1.GetResInfo)

		//设计团队模块
		router.GET("teams", v1.GetTeams)

		//查询商品下的所有图片
		router.GET("commodity_pict/:id", v1.GetCommPic)

		//商品模块
		router.GET("all_comm_pic", v1.GetAllCommPic)
		router.GET("commodity", v1.GetCommodity)

		//logo模块
		router.GET("logo/:id", v1.GetLogoInfo)

		//轮播图模块
		router.GET("carousels", v1.GetCarousels)
	}

	_ = r.Run(utils.HttpPort)
}

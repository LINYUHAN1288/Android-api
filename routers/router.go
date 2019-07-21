// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"Android-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/barcode2",
			beego.NSInclude(
				&controllers.Barcode2Controller{},
			),
		),

		beego.NSNamespace("/ch_dietnutri",
			beego.NSInclude(
				&controllers.ChDietnutriController{},
			),
		),

		beego.NSNamespace("/ch_dietnutri_meterial",
			beego.NSInclude(
				&controllers.ChDietnutriMeterialController{},
			),
		),

		beego.NSNamespace("/ch_dietnutricrawl",
			beego.NSInclude(
				&controllers.ChDietnutricrawlController{},
			),
		),

		beego.NSNamespace("/ch_dris",
			beego.NSInclude(
				&controllers.ChDrisController{},
			),
		),

		beego.NSNamespace("/chineseFoodGroup",
			beego.NSInclude(
				&controllers.ChineseFoodGroupController{},
			),
		),

		beego.NSNamespace("/chinesefoodgoodsnutri",
			beego.NSInclude(
				&controllers.ChinesefoodgoodsnutriController{},
			),
		),

		beego.NSNamespace("/chinesefoodnutri",
			beego.NSInclude(
				&controllers.ChinesefoodnutriController{},
			),
		),

		beego.NSNamespace("/chinesefoodnutrifromhk",
			beego.NSInclude(
				&controllers.ChinesefoodnutrifromhkController{},
			),
		),

		beego.NSNamespace("/chinesevegetable",
			beego.NSInclude(
				&controllers.ChinesevegetableController{},
			),
		),

		beego.NSNamespace("/cookbook",
			beego.NSInclude(
				&controllers.CookbookController{},
			),
		),

		beego.NSNamespace("/deliciousFoodInfo",
			beego.NSInclude(
				&controllers.DeliciousFoodInfoController{},
			),
		),

		beego.NSNamespace("/food_data",
			beego.NSInclude(
				&controllers.FoodDataController{},
			),
		),

		beego.NSNamespace("/foodfrombhcp",
			beego.NSInclude(
				&controllers.FoodfrombhcpController{},
			),
		),

		beego.NSNamespace("/foodsafenews",
			beego.NSInclude(
				&controllers.FoodsafenewsController{},
			),
		),

		beego.NSNamespace("/foodsecurityfromzccw",
			beego.NSInclude(
				&controllers.FoodsecurityfromzccwController{},
			),
		),

		beego.NSNamespace("/foodvideo",
			beego.NSInclude(
				&controllers.FoodvideoController{},
			),
		),

		beego.NSNamespace("/kitchenstrories",
			beego.NSInclude(
				&controllers.KitchenstroriesController{},
			),
		),

		beego.NSNamespace("/media2_web",
			beego.NSInclude(
				&controllers.Media2WebController{},
			),
		),

		beego.NSNamespace("/media_video",
			beego.NSInclude(
				&controllers.MediaVideoController{},
			),
		),

		beego.NSNamespace("/news",
			beego.NSInclude(
				&controllers.NewsController{},
			),
		),

		beego.NSNamespace("/sports_data",
			beego.NSInclude(
				&controllers.SportsDataController{},
			),
		),

		beego.NSNamespace("/sporttype",
			beego.NSInclude(
				&controllers.SporttypeController{},
			),
		),
		/*
		beego.NSNamespace("/user_basic",
			beego.NSInclude(
				&controllers.UserBasicController{},
			),
		),
		*/

		beego.NSNamespace("/user_goal",
			beego.NSInclude(
				&controllers.UserGoalController{},
			),
		),

		beego.NSNamespace("/user_monitor",
			beego.NSInclude(
				&controllers.UserMonitorController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.Router("/user_basic/:id", &controllers.UserBasicController{}, "get:GetOne;post:Post")
	beego.Router("/user_basic/query/:Mobile", &controllers.UserBasicController{}, "get:Get;post:Post")
	beego.Router("/user_basic/update/:Mobile/:oldpwd/:newpwd", &controllers.UserBasicController{}, "put:Update;post:Post")
}

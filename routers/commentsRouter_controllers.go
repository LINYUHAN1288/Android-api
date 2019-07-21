package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["Android-api/controllers:Barcode2Controller"] = append(beego.GlobalControllerRouter["Android-api/controllers:Barcode2Controller"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:Barcode2Controller"] = append(beego.GlobalControllerRouter["Android-api/controllers:Barcode2Controller"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:Barcode2Controller"] = append(beego.GlobalControllerRouter["Android-api/controllers:Barcode2Controller"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:Barcode2Controller"] = append(beego.GlobalControllerRouter["Android-api/controllers:Barcode2Controller"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:Barcode2Controller"] = append(beego.GlobalControllerRouter["Android-api/controllers:Barcode2Controller"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriMeterialController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriMeterialController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriMeterialController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriMeterialController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriMeterialController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriMeterialController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriMeterialController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriMeterialController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriMeterialController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutriMeterialController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutricrawlController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutricrawlController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutricrawlController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutricrawlController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutricrawlController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutricrawlController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutricrawlController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutricrawlController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDietnutricrawlController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDietnutricrawlController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDrisController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDrisController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDrisController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDrisController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDrisController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDrisController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDrisController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDrisController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChDrisController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChDrisController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChineseFoodGroupController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChineseFoodGroupController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChineseFoodGroupController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChineseFoodGroupController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChineseFoodGroupController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChineseFoodGroupController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChineseFoodGroupController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChineseFoodGroupController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChineseFoodGroupController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChineseFoodGroupController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodgoodsnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodgoodsnutriController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodgoodsnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodgoodsnutriController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodgoodsnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodgoodsnutriController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodgoodsnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodgoodsnutriController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodgoodsnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodgoodsnutriController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutriController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutriController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutriController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutriController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutriController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutriController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutrifromhkController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutrifromhkController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutrifromhkController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutrifromhkController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutrifromhkController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutrifromhkController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutrifromhkController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutrifromhkController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutrifromhkController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesefoodnutrifromhkController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesevegetableController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesevegetableController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesevegetableController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesevegetableController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesevegetableController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesevegetableController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesevegetableController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesevegetableController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:ChinesevegetableController"] = append(beego.GlobalControllerRouter["Android-api/controllers:ChinesevegetableController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:CookbookController"] = append(beego.GlobalControllerRouter["Android-api/controllers:CookbookController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:CookbookController"] = append(beego.GlobalControllerRouter["Android-api/controllers:CookbookController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:CookbookController"] = append(beego.GlobalControllerRouter["Android-api/controllers:CookbookController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:CookbookController"] = append(beego.GlobalControllerRouter["Android-api/controllers:CookbookController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:CookbookController"] = append(beego.GlobalControllerRouter["Android-api/controllers:CookbookController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:DeliciousFoodInfoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:DeliciousFoodInfoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:DeliciousFoodInfoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:DeliciousFoodInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:DeliciousFoodInfoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:DeliciousFoodInfoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:DeliciousFoodInfoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:DeliciousFoodInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:DeliciousFoodInfoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:DeliciousFoodInfoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodDataController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodDataController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodDataController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodDataController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodDataController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodDataController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodDataController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodDataController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodDataController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodDataController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodfrombhcpController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodfrombhcpController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodfrombhcpController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodfrombhcpController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodfrombhcpController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodfrombhcpController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodfrombhcpController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodfrombhcpController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodfrombhcpController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodfrombhcpController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodsafenewsController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodsafenewsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodsafenewsController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodsafenewsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodsafenewsController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodsafenewsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodsafenewsController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodsafenewsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodsafenewsController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodsafenewsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodsecurityfromzccwController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodsecurityfromzccwController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodsecurityfromzccwController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodsecurityfromzccwController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodsecurityfromzccwController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodsecurityfromzccwController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodsecurityfromzccwController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodsecurityfromzccwController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodsecurityfromzccwController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodsecurityfromzccwController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodvideoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodvideoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodvideoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodvideoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodvideoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodvideoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodvideoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodvideoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:FoodvideoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:FoodvideoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:KitchenstroriesController"] = append(beego.GlobalControllerRouter["Android-api/controllers:KitchenstroriesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:KitchenstroriesController"] = append(beego.GlobalControllerRouter["Android-api/controllers:KitchenstroriesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:KitchenstroriesController"] = append(beego.GlobalControllerRouter["Android-api/controllers:KitchenstroriesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:KitchenstroriesController"] = append(beego.GlobalControllerRouter["Android-api/controllers:KitchenstroriesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:KitchenstroriesController"] = append(beego.GlobalControllerRouter["Android-api/controllers:KitchenstroriesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:Media2WebController"] = append(beego.GlobalControllerRouter["Android-api/controllers:Media2WebController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:Media2WebController"] = append(beego.GlobalControllerRouter["Android-api/controllers:Media2WebController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:Media2WebController"] = append(beego.GlobalControllerRouter["Android-api/controllers:Media2WebController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:Media2WebController"] = append(beego.GlobalControllerRouter["Android-api/controllers:Media2WebController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:Media2WebController"] = append(beego.GlobalControllerRouter["Android-api/controllers:Media2WebController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:MediaVideoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:MediaVideoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:MediaVideoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:MediaVideoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:MediaVideoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:MediaVideoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:MediaVideoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:MediaVideoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:MediaVideoController"] = append(beego.GlobalControllerRouter["Android-api/controllers:MediaVideoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:NewsController"] = append(beego.GlobalControllerRouter["Android-api/controllers:NewsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:NewsController"] = append(beego.GlobalControllerRouter["Android-api/controllers:NewsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:NewsController"] = append(beego.GlobalControllerRouter["Android-api/controllers:NewsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:NewsController"] = append(beego.GlobalControllerRouter["Android-api/controllers:NewsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:NewsController"] = append(beego.GlobalControllerRouter["Android-api/controllers:NewsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:SportsDataController"] = append(beego.GlobalControllerRouter["Android-api/controllers:SportsDataController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:SportsDataController"] = append(beego.GlobalControllerRouter["Android-api/controllers:SportsDataController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:SportsDataController"] = append(beego.GlobalControllerRouter["Android-api/controllers:SportsDataController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:SportsDataController"] = append(beego.GlobalControllerRouter["Android-api/controllers:SportsDataController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:SportsDataController"] = append(beego.GlobalControllerRouter["Android-api/controllers:SportsDataController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:SporttypeController"] = append(beego.GlobalControllerRouter["Android-api/controllers:SporttypeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:SporttypeController"] = append(beego.GlobalControllerRouter["Android-api/controllers:SporttypeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:SporttypeController"] = append(beego.GlobalControllerRouter["Android-api/controllers:SporttypeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:SporttypeController"] = append(beego.GlobalControllerRouter["Android-api/controllers:SporttypeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:SporttypeController"] = append(beego.GlobalControllerRouter["Android-api/controllers:SporttypeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/v1/:Mobile`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserBasicController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/v2/:Mobile/:oldpwd/:newpwd`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserGoalController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserGoalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserGoalController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserGoalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserGoalController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserGoalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserGoalController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserGoalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserGoalController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserGoalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserMonitorController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserMonitorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserMonitorController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserMonitorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserMonitorController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserMonitorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserMonitorController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserMonitorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["Android-api/controllers:UserMonitorController"] = append(beego.GlobalControllerRouter["Android-api/controllers:UserMonitorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

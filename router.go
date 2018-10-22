package main

import (
	"net/http"

	. "github.com/dangyanglim/go_cnode/apis"
	"github.com/dangyanglim/go_cnode/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	//gin.SetMode(gin.ReleaseMode)
	router.LoadHTMLGlob("views/*")
	router.StaticFS("/public", http.Dir("./public"))
	router.StaticFile("/favicon.ico", "./public/images/cnode_icon_32.png")
	router.GET("/", site.Index)

	router.POST("/person", AddPersonApi)

	router.GET("/persons", GetPersonsApi)

	router.GET("/person/:id", GetPersonApi)

	router.PUT("/person/:id", ModPersonApi)

	router.DELETE("/person/:id", DelPersonApi)

	// router.GET("/property", LoginByPassword)

	//提供类似蛋贝接口
	//登陆接口
	router.POST("/loginAccess", LoginAccess)
	//查看该用户下所有设备的列表信息，管理员则查看所有列表。
	router.POST("/deviceInfo/getDeviceListAll", LoginAccess)
	//查看该用户下所有能耗设备的列表信息，管理员则查看所有能耗设备。水电表
	router.POST("/deviceInfo/getEnergyDeviceList", LoginAccess)
	//查询该客户下的某个能耗设备信息，管理员可以查看任意能耗设备信息。
	router.POST("/deviceInfo/getEnergyDeviceInfo", LoginAccess)
	//查询当天能耗设备示数。
	router.POST("/deviceInfo/getEnergyDailyReading", LoginAccess)
	//查询能耗设备区间用量。电表精确到小时，水表只取天数。
	router.POST("/deviceInfo/getEnergySectionConsumption", LoginAccess)
	//对电表进行控制，比如断电，恢复操作(此接口需要电表硬件支持)
	router.POST("/deviceInfo/energy/gateControl", LoginAccess)
	//查询电表拉合闸状态
	router.POST("/deviceInfo/energy/getGateStatus", LoginAccess)
	//查看某小区物业下所有房间关联的设备列表
	router.POST("/deviceInfo/getDeviceList", LoginAccess)
	//同步房源接口，不太理解
	router.POST("/house/sync", LoginAccess)
	//门锁接口
	router.POST("/deviceInfo/getLockList", LoginAccess)
	router.POST("/deviceInfo/getLockInfo", LoginAccess)
	router.POST("/deviceInfo/getLockPwdList", LoginAccess)
	router.POST("/deviceCtrl/lockPwd/addPwd", LoginAccess)
	router.POST("/deviceCtrl/lockPwd/editPwd", LoginAccess)
	router.POST("/deviceCtrl/lockPwd/delPwd", LoginAccess)
	router.POST("/deviceCtrl/lockPwd/updatePwd", LoginAccess)

	//下面是设备向平台汇报接口，门锁
	//密码开门请求
	router.POST("/openDoor", LoginAccess)
	router.POST("/warning", LoginAccess)
	router.POST("/access", LoginAccess)

	//自定义远程抄表接口
	//前端接口
	//用户密码
	router.GET("/app/v1/user/loginByPwd", UserLoginByPwd)
	//用户注册绑定接口
	router.GET("/app/v1/user/register", UserRegister)
	//用户重置密码
	router.GET("/app/v1/user/userResetPwd", UserResetPwd)
	// 获取验证码
	router.GET("/app/v1/user/getCode", GetUserCode)
	//用户验证码登陆
	router.GET("/app/v1/user/userLoginByCode", UserLoginByCode)
	//获取该用户信息
	router.GET("/app/v1/user", LoginAccess)
	//获取房间信息
	router.GET("/app/v1/user/getRoom", UserGetRoom)

	//获取房间
	//router.GET("/app/v1/user/getRoom", LoginAccess)
	//共享房间
	router.GET("/app/v1/admin/shareRoomToUser", LoginAccess)
	//获取水表信息
	router.GET("/app/v1/user/getWater", LoginAccess)
	//获取电表信息
	router.GET("/app/v1/user/getElectricity", LoginAccess)

	//后端接口
	//管理后台登陆接口
	router.GET("/api/v1/property/LoginByPwd", PropertyLoginByPwd)
	//获取该管理员信息
	router.GET("/api/v1/property/GetCode", GetPropertyCode)

	router.GET("/api/v1/property/Register", PropertyRegister)

	router.GET("/api/v1/property/ResetPwd", PropertyResetPwd)

	router.GET("/api/v1/property", GetProperty)

	router.GET("/api/v1/property/AddRoom", AddPropertyRoomApi)

	router.GET("/api/v1/property/DelRoom", DelRoomApi)
	//获取所有房间
	router.GET("/api/v1/property/getRooms", GetPropertyRoomApi)
	router.GET("/api/v1/property/getRoom", GetPropertyRoomByIdApi)

	router.GET("/api/v1/property/addRenter", AddPropertyRenterApi)

	router.GET("/api/v1/property/getRenters", GetPropertyRentersApi)
	router.GET("/api/v1/property/delRenter", DelPropertyRenterApi)
	//获取所有物业
	router.GET("/api/v1/property/getHouses", LoginAccess)
	//获取该物业所有楼层
	router.GET("/api/v1/property/getfloors", LoginAccess)
	router.GET("/api/v1/property/getRoomsCount", GetPropertRoomCountApi)
	router.GET("/api/v1/property/addAmmeter", AddAmmeterApi)
	router.GET("/api/v1/property/getAmmeters", GetAmmetersApi)
	router.GET("/api/v1/property/delAmmeter", DelAmmeterApi)

	router.GET("/api/v1/property/addLock", AddLockApi)
	router.GET("/api/v1/property/getLocks", GetLocksApi)
	router.GET("/api/v1/property/delLock", DelLockApi)

	router.GET("/api/v1/property/addWater", AddWaterApi)
	router.GET("/api/v1/property/getWaters", GetWatersApi)
	router.GET("/api/v1/property/delWater", DelWaterApi)

	router.GET("/api/v1/property/addGateway", AddGatewayApi)
	router.GET("/api/v1/property/getGateways", GetGatewaysApi)
	router.GET("/api/v1/property/delGateway", DelGatewayApi)

	router.GET("/api/v1/property/save", UpdatePropertyNameApi)
	//获取房间情况

	//绑定租客
	router.GET("/api/v1/property/tieUserToRoom", LoginAccess)
	//获取水表信息
	router.GET("/api/v1/property/getWater", LoginAccess)
	//获取电表信息
	router.GET("/api/v1/property/getElectricity", LoginAccess)

	return router
}

package routes

import (
	"go-skeleton-manager-rabbitmq/controllers"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var prefix = "/api/inisa"

// Build returns the application routes
func Build(e *echo.Echo) {
	RouteServices(e)
}

func RouteServices(e *echo.Echo) {
	r := e.Group(prefix)
	var Auth = middleware.JWT([]byte(os.Getenv("JWT_SECRET")))
	r.GET("/user/:id", controllers.GetUser, Auth)
	r.GET("/user", controllers.GetUser)
	r.POST("/user", controllers.CreateUser)
	r.PUT("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	//r.GET("/exportExcel", controllers.TestExport)
	// ----------------- Pembelian --------------------------------
	r.POST("/pembelian", controllers.CreatePembelian)
	r.PUT("/pembelian/:id", controllers.UpdatePembelian)
	r.PUT("/pembelian/:id/:iddetail", controllers.UpdatePembelian)
	r.DELETE("/pembelian/:id", controllers.DeletePembelian)
	r.GET("/pembelian/:id", controllers.GetListPembelianById)
	r.GET("/pembelian-detail/:id", controllers.GetDetailPembelianDetail)
	r.GET("/pembelian", controllers.GetListPembelian)
	r.GET("/pembelian-with-detail", controllers.GetPembelianWithDetail)
	r.GET("/pembelian-with-detail/:id", controllers.GetPembelianWithDetail)

	r.POST("/penjualan", controllers.CreatePenjualan)
	r.PUT("/penjualan/:id", controllers.UpdatePenjualan)
	r.DELETE("/penjualan/:id", controllers.DeletePenjualan)
	r.GET("/penjualan", controllers.GetListPenjualan)
	r.GET("/penjualan/detail", controllers.GetDetailPenjualan)

	/* ------------------ laporan ----------------------------*/
	r.GET("/produk-perkota", controllers.LapProdukPerKota)
	r.GET("/produk-pertoko", controllers.LapProdukPerToko)
	r.GET("/rekap-penjualan", controllers.LapRekapPenjualan)

	r.GET("/laporan-harian", controllers.GetLaporanHarian)
	r.GET("/laporan-harian-toko", controllers.GetLaporanHarianToko)
	r.GET("/rekap-laporan", controllers.GetDetailLaporan)

	r.GET("/form/:id", controllers.GetForm)

	r.GET("/test", controllers.UploadFile)
}

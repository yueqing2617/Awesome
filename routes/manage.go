package routes

import (
	"Awesome/app/http/controllers/manager"
	"Awesome/app/http/middleware"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Manager() {
	facades.Route.Prefix("/manage").Middleware(middleware.ErrorHandler()).Group(func(r route.Route) {
		// basic
		r.Prefix("/basic").Group(func(r route.Route) {
			// color
			r.Prefix("/color").Group(func(r route.Route) {
				colors := manager.NewBasicColor()
				r.Get("/", colors.Index)
				r.Post("/", colors.Create)
				r.Get("/:key", colors.Show)
				r.Put("/:id", colors.Edit)
				r.Delete("/", colors.Delete)
			})
			// order type
			r.Prefix("/order_type").Group(func(r route.Route) {
				orderTypes := manager.NewBasicOrderType()
				r.Get("/", orderTypes.Index)
				r.Post("/", orderTypes.Create)
				r.Get("/:key", orderTypes.Show)
				r.Put("/:id", orderTypes.Edit)
				r.Delete("/", orderTypes.Delete)
			})
			// procedure
			r.Prefix("/procedure").Group(func(r route.Route) {
				procedures := manager.NewBasicProcedure()
				r.Get("/", procedures.Index)
				r.Post("/", procedures.Create)
				r.Get("/:key", procedures.Show)
				r.Put("/:id", procedures.Edit)
				r.Delete("/", procedures.Delete)
			})
			// size
			r.Prefix("/size").Group(func(r route.Route) {
				sizes := manager.NewBasicSize()
				r.Get("/", sizes.Index)
				r.Post("/", sizes.Create)
				r.Get("/:key", sizes.Show)
				r.Put("/:id", sizes.Edit)
				r.Delete("/", sizes.Delete)
			})
			// customer
			r.Prefix("/customer").Group(func(r route.Route) {
				customers := manager.NewBasicCustomer()
				r.Get("/", customers.Index)
				r.Post("/", customers.Create)
				r.Get("/:key", customers.Show)
				r.Put("/:id", customers.Edit)
				r.Delete("/", customers.Delete)
			})
			// salesman
			r.Prefix("/salesman").Group(func(r route.Route) {
				salesmen := manager.NewBasicSalesman()
				r.Get("/", salesmen.Index)
				r.Post("/", salesmen.Create)
				r.Get("/:key", salesmen.Show)
				r.Put("/:id", salesmen.Edit)
				r.Delete("/", salesmen.Delete)
			})
		})
		// system
		r.Prefix("/system").Group(func(r route.Route) {
			// permission
			r.Prefix("/permission").Group(func(r route.Route) {
				permissions := manager.NewAdminPermission()
				r.Get("/", permissions.Index)
				r.Post("/", permissions.Create)
				r.Get("/:id", permissions.Show)
				r.Put("/:id", permissions.Edit)
				r.Delete("/", permissions.Delete)
			})
			// role
			r.Prefix("/role").Group(func(r route.Route) {
				roles := manager.NewAdminRole()
				r.Get("/", roles.Index)
				r.Post("/", roles.Create)
				r.Get("/:id", roles.Show)
				r.Put("/:id", roles.Edit)
				r.Delete("/", roles.Delete)
			})
			// admin
			r.Prefix("/admin").Group(func(r route.Route) {
				admins := manager.NewAdmin()
				r.Get("/", admins.Index)
				r.Post("/", admins.Create)
				r.Get("/:id", admins.Show)
				r.Put("/:id", admins.Edit)
				r.Delete("/", admins.Delete)
			})
		})
		// cloth
		r.Prefix("/cloth").Group(func(r route.Route) {
			// style
			r.Prefix("/style").Group(func(r route.Route) {
				styles := manager.NewClothStyle()
				r.Get("/", styles.Index)
				r.Post("/", styles.Create)
				r.Get("/:id", styles.Show)
				r.Put("/:id", styles.Edit)
				r.Delete("/", styles.Delete)
			})
			// order
			r.Prefix("/order").Group(func(r route.Route) {
				orders := manager.NewClothOrder()
				r.Get("/", orders.Index)
				r.Post("/", orders.Create)
				r.Get("/:id", orders.Show)
				r.Put("/:id", orders.Edit)
				r.Delete("/", orders.Delete)
			})
			// tailor
			r.Prefix("/tailor").Group(func(r route.Route) {
				tailors := manager.NewClothTailor()
				r.Get("/", tailors.Index)
				r.Post("/", tailors.Create)
				r.Get("/:id", tailors.Show)
				r.Put("/:id", tailors.Edit)
				r.Delete("/", tailors.Delete)
				r.Prefix("/piece").Group(func(r route.Route) {
					pieces := manager.NewClothTailorCuttingPiece()
					r.Get("/", pieces.Index)
					r.Post("/:id", pieces.Create)
					r.Get("/:id", pieces.Show)
					r.Put("/:id", pieces.Edit)
					r.Delete("/", pieces.Delete)
				})
			})
		})
		// common
		r.Prefix("/common").Group(func(r route.Route) {
			common := manager.NewCommon()
			// login
			r.Post("/login", common.Login)
			// logout
			r.Post("/logout", common.Logout)
			// info
			r.Get("/info", common.Info)
			// captcha
			r.Get("/captcha", common.Captcha)
			// refresh token
			r.Get("/refresh_token", common.Refresh)
			// upload
			r.Post("/upload", common.Upload)
			// options
			r.Middleware(middleware.AdminLoginCheck()).Get("/options", common.Options)
			// edit password
			r.Middleware(middleware.AdminLoginCheck()).Post("/edit_password", common.EditPassword)
		})
	})
}

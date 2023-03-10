// Package routes
// @file : admin.go
// @project : AGPC(Awesome Garment Production Cloud)
// @author : 周东明（Empty）
// @contact : empty@inzj.cn
// @created at: 2023/3/7 15:11
// ----------------------------------------------------------
package routes

import (
	"Awesome/app/http/controllers/admin"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func Admin() {
	facades.Route.Prefix("/admin").Group(func(r route.Route) {
		// system
		r.Prefix("/system").Group(func(r route.Route) {
			// admin
			r.Prefix("/admin").Group(func(r route.Route) {
				admins := admin.NewAdmins()
				r.Get("/", admins.Index)
				r.Post("/", admins.Create)
				r.Get("/:id", admins.Show)
				r.Put("/:id", admins.Store)
				r.Delete("/", admins.Delete)
				r.Put("/:id/password", admins.EditPassword)
				//r.Post("/restore", admins.Restore)
			})
			// permission
			r.Prefix("/permission").Group(func(r route.Route) {
				permissions := admin.NewPermissions()
				r.Get("/", permissions.Index)
				r.Post("/", permissions.Create)
				r.Get("/:id", permissions.Show)
				r.Put("/:id", permissions.Store)
				r.Delete("/", permissions.Delete)
			})
			// role
			r.Prefix("/role").Group(func(r route.Route) {
				roles := admin.NewAdminRoles()
				r.Get("/", roles.Index)
				r.Post("/", roles.Create)
				r.Get("/:id", roles.Show)
				r.Put("/:id", roles.Store)
				//r.Delete("/", roles.Delete)
			})
		})
		// common
		r.Prefix("/common").Group(func(r route.Route) {
			common := admin.NewCommon()
			r.Get("/captcha", common.Captcha)
			r.Get("/info", common.Info)
			r.Post("/login", common.Login)
			r.Post("/logout", common.Logout)
			r.Put("/refresh", common.Refresh)
		})
		// basic
		r.Prefix("/basic").Group(func(r route.Route) {
			// color
			r.Prefix("/color").Group(func(r route.Route) {
				colors := admin.NewBasicColor()
				r.Get("/", colors.Index)
				r.Post("/", colors.Create)
				r.Get("/:id", colors.Show)
				r.Put("/:id", colors.Store)
				r.Delete("/", colors.Delete)
			})
			// size
			r.Prefix("/size").Group(func(r route.Route) {
				sizes := admin.NewBasicSize()
				r.Get("/", sizes.Index)
				r.Post("/", sizes.Create)
				r.Get("/:id", sizes.Show)
				r.Put("/:id", sizes.Store)
				r.Delete("/", sizes.Delete)
			})
			// customer
			r.Prefix("/customer").Group(func(r route.Route) {
				customers := admin.NewBasicCustomer()
				r.Get("/", customers.Index)
				r.Post("/", customers.Create)
				r.Get("/:id", customers.Show)
				r.Put("/:id", customers.Store)
				r.Delete("/", customers.Delete)
			})
			// salesman
			r.Prefix("/salesman").Group(func(r route.Route) {
				salesmen := admin.NewBasicSalesman()
				r.Get("/", salesmen.Index)
				r.Post("/", salesmen.Create)
				r.Get("/:id", salesmen.Show)
				r.Put("/:id", salesmen.Store)
				r.Delete("/", salesmen.Delete)
			})
			// orderType
			r.Prefix("/order_type").Group(func(r route.Route) {
				orderTypes := admin.NewBasicOrderType()
				r.Get("/", orderTypes.Index)
				r.Post("/", orderTypes.Create)
				r.Get("/:id", orderTypes.Show)
				r.Put("/:id", orderTypes.Store)
				r.Delete("/", orderTypes.Delete)
			})
			// procedure
			r.Prefix("/procedure").Group(func(r route.Route) {
				procedures := admin.NewBasicProcedure()
				r.Get("/", procedures.Index)
				r.Post("/", procedures.Create)
				r.Get("/:id", procedures.Show)
				r.Put("/:id", procedures.Store)
				r.Delete("/", procedures.Delete)
			})
		})
	})
}

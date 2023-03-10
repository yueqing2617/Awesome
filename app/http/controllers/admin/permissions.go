package admin

import (
	"Awesome/app/http/helper"
	"Awesome/app/models"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type Permissions struct {
	//Dependent services
}

func NewPermissions() *Permissions {
	return &Permissions{
		//Inject services
	}
}

// Index is the index action of Permissions controller.
func (r *Permissions) Index(ctx http.Context) {
	mode := ctx.Request().Query("mode", "tree")

	// query
	query := facades.Orm.Query().Order("sort asc").Model(&models.Permission{})
	var list []models.Permission
	var count int64
	// 改写成 switch
	switch mode {
	case "tree":
		list, _ = (new(models.Permission)).GetPermissionList()
		count = int64(len(list))
	case "list":
		err := query.Find(&list)
		if err != nil {
			return
		}
		err = query.Count(&count)
		if err != nil {
			return
		}
	default:
		helper.RestfulError(ctx, "mode 参数错误")
		return
	}

	helper.RestfulSuccess(ctx, "success", http.Json{
		"Items": list,
		"Total": count,
	})
}

// Show is the show action of Permissions controller.
func (r *Permissions) Show(ctx http.Context) {
}

// Create is the create action of Permissions controller.
func (r *Permissions) Create(ctx http.Context) {
}

// Store is the store action of Permissions controller.
func (r *Permissions) Store(ctx http.Context) {
}

// Delete is the delete action of Permissions controller.
func (r *Permissions) Delete(ctx http.Context) {

}

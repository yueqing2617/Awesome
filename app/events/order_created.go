package events

import (
	"Awesome/app/models"
	"github.com/goravel/framework/contracts/event"
	"github.com/goravel/framework/facades"
	"github.com/pkg/errors"
)

type OrderCreated struct {
}

func (receiver *OrderCreated) Handle(args []event.Arg) ([]event.Arg, error) {
	code := args[0].Value.(string)
	styleCode := args[1].Value.(string)
	quantity := args[2].Value.(uint)
	opt := args[3].Value.(string)
	customerName := args[4].Value.(string)
	switch opt {
	case "create":
		var count int64
		if err := facades.Orm.Query().Model(&models.ClothTailor{}).Where("cloth_order_code", code).Count(&count); err != nil {
			return args, err
		}
		if count == 0 {
			if err := facades.Orm.Query().Model(&models.ClothTailor{}).Create(&models.ClothTailor{
				ClothOrderCode: code,
				ClothStyleCode: styleCode,
				Total:          quantity,
				CustomerName:   customerName,
			}); err != nil {
				return args, err
			}
		}
		return args, nil
	case "update":
		res, err := facades.Orm.Query().Model(&models.ClothTailor{}).Where("cloth_order_code", code).Updates(map[string]interface{}{
			"cloth_style_code": styleCode,
			"total":            quantity,
			"customer_name":    customerName,
		})
		if err != nil {
			return args, err
		}
		if res.RowsAffected == 0 {
			return args, errors.New("更新失败")
		}
		return args, nil
	default:
		return args, nil
	}
}

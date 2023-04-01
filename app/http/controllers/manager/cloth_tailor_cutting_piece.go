package manager

import (
	"Awesome/app/http/helper"
	"Awesome/app/http/requests"
	"Awesome/app/models"
	"errors"
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"strconv"
)

type ClothTailorCuttingPiece struct {
	//Dependent services
	Model models.ClothTailorCuttingPieces
}

func NewClothTailorCuttingPiece() *ClothTailorCuttingPiece {
	return &ClothTailorCuttingPiece{
		//Inject services
	}
}

// Index GET /dummy
func (r *ClothTailorCuttingPiece) Index(ctx http.Context) {
	page, _ := strconv.Atoi(ctx.Request().Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Request().Query("limit", "10"))
	orderCode := ctx.Request().Query("cloth_order_code", "")
	clothStyleCode := ctx.Request().Query("cloth_style_code", "")
	order := ctx.Request().Query("order", "id desc")
	// query
	query := facades.Orm.Query().Model(&r.Model)
	if orderCode != "" {
		query = query.Where("cloth_order_code", "like", "%"+orderCode+"%")
	}
	if clothStyleCode != "" {
		query = query.Where("cloth_style_code", "like", "%"+clothStyleCode+"%")
	}
	// 排序
	query = query.Order(order)
	// 分页
	var total int64
	var list []models.ClothTailorCuttingPieces
	if err := query.Paginate(page, limit, &list, &total); err != nil {
		helper.RestfulError(ctx, "查询失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "查询成功", http.Json{
		"total": total,
		"items": list,
	})
}

// Create POST /dummy
func (r *ClothTailorCuttingPiece) Create(ctx http.Context) {
	var params requests.AdminClothTailorCuttingPieceRequest
	valid, err := ctx.Request().ValidateRequest(&params)
	if err != nil {
		helper.RestfulError(ctx, "表单参数错误："+err.Error())
		return
	}
	if valid != nil {
		helper.RestfulError(ctx, "表单参数错误："+helper.GetRequestError(valid.All()))
		return
	}
	id := ctx.Request().Input("id", "")

	order, err := findOrderTailor(id)
	if err != nil {
		helper.RestfulError(ctx, err.Error())
		return
	}
	err = checkTailor(order, &params)
	if err != nil {
		helper.RestfulError(ctx, err.Error())
		return
	}
	procedures := helper.JsonToProcedures(order.ClothOrder.ProceduresStr)
	tx, err := facades.Orm.Query().Begin()
	num := params.StartNum
	for _, v := range params.Pieces {
		// 保存裁剪记录
		for i := 0; i < v.Num; i++ {
			num++
			cutting := &models.ClothTailorCuttingPieces{
				ClothTailorID:  order.ID,
				ClothStyleCode: order.ClothOrder.ClothStyleCode,
				BedNumber:      params.BedNum,
				Number:         num,
				Layer:          v.Layer,
				Color:          v.Color,
				Size:           v.Size,
				IsCompleted:    false,
			}
			if err = tx.Create(&cutting); err != nil {
				err := tx.Rollback()
				helper.RestfulError(ctx, "保存裁剪记录失败："+err.Error())
				return
			}
			process := []models.ClothTailorCuttingPieceProcess{}
			for _, procedure := range procedures {
				cuttingProcedure := &models.ClothTailorCuttingPieceProcess{
					Name:           procedure.Name,
					Sort:           procedure.Sort,
					Price:          procedure.Price,
					CuttingPieceID: cutting.ID,
					IsCompleted:    false,
					EmployeeID:     0,
					CompletedAt:    "0",
				}
				process = append(process, *cuttingProcedure)
			}
			if err = tx.Create(&process); err != nil {
				err := tx.Rollback()
				helper.RestfulError(ctx, "保存裁剪工序失败："+err.Error())
				return
			}
		}
	}
	if err = tx.Commit(); err != nil {
		helper.RestfulError(ctx, "裁剪失败："+err.Error())
		return
	}
	helper.RestfulSuccess(ctx, "裁剪成功", nil)
}

// Show GET /dummy/{id}
func (r *ClothTailorCuttingPiece) Show(ctx http.Context) {
}

// Edit GET /dummy/{id}/edit
func (r *ClothTailorCuttingPiece) Edit(ctx http.Context) {
}

// Delete DELETE /dummy
// @param id []int
func (r *ClothTailorCuttingPiece) Delete(ctx http.Context) {
	ids := ctx.Request().QueryArray("id")
	if len(ids) == 0 {
		helper.RestfulError(ctx, "id:至少选择一条记录")
		return
	}
	res, err := facades.Orm.Query().Delete(&r.Model, ids)
	if err != nil {
		helper.RestfulError(ctx, "删除失败："+err.Error())
		return
	}
	// 成功删除了 res.RowsAffected 条记录
	helper.RestfulSuccess(ctx, "成功删除了 "+strconv.Itoa(int(res.RowsAffected))+" 条记录", nil)
}

func findOrderTailor(id string) (order *models.ClothTailor, err error) {
	var tailor models.ClothTailor
	if err = facades.Orm.Query().Model(&tailor).Where("id", id).With("ClothOrder").FirstOrFail(&tailor); err != nil {
		return nil, err
	}
	if tailor.ClothOrder.ID == 0 {
		return nil, errors.New("订单不存在")
	}
	fmt.Println(tailor.ClothOrder.Status)
	if tailor.ClothOrder.Status != 1 {
		return nil, errors.New("订单已关闭或已完成")
	}
	if tailor.IsCompleted {
		return nil, errors.New("当前裁剪任务已完成")
	}
	return &tailor, nil
}

func checkTailor(order *models.ClothTailor, params *requests.AdminClothTailorCuttingPieceRequest) error {
	pieces := params.Pieces
	if len(pieces) == 0 {
		return errors.New("待裁剪的颜色尺码不能为空")
	}
	contains := helper.JSONToContains(order.ClothOrder.ContainsStr)
	// 获取当前裁剪任务已裁剪的颜色尺码中相同床号的最大扎号
	var maxPieceNum int
	if err := facades.Orm.Query().Model(&models.ClothTailorCuttingPieces{}).Where("cloth_tailor_id = ? and bed = ?", order.ID, params.BedNum).Select("COALESCE(max(number), 0) as max_piece_num").First(&maxPieceNum); err != nil {
		return err
	}
	if maxPieceNum > 0 {
		// 检查当前裁剪任务已裁剪的颜色尺码中相同床号的最大扎号是否大于当前扎号
		if maxPieceNum >= int(params.StartNum) {
			return errors.New("当前裁剪任务已裁剪的颜色尺码中相同床号的最大扎号必须大于当前扎号: " + strconv.Itoa(int(maxPieceNum)))
		}
	}

	// 检查裁剪数量
	for k, piece := range pieces {
		row := k + 1
		if piece.Color == "" {
			return errors.New("第" + strconv.Itoa(row) + "行颜色不能为空")
		}
		if piece.Size == "" {
			return errors.New("第" + strconv.Itoa(row) + "行尺码不能为空")
		}
		if piece.Num <= 0 {
			return errors.New("第" + strconv.Itoa(row) + "行裁剪数量必须大于0")
		}
		// 从订单中检查颜色尺码是否存在
		var rowContains models.ClothOrderContains
		for _, contain := range contains {
			if contain.Color == piece.Color && contain.Size == piece.Size {
				rowContains = contain
				break
			}
		}
		if rowContains.Size == "" && rowContains.Color == "" {
			return errors.New("第" + strconv.Itoa(row) + "行: 在订单中未找到颜色为：" + piece.Color + " 尺码为：" + piece.Size + "的生产需求")
		}
		// 检查裁剪数量是否超出订单需求
		if int(rowContains.Num) < piece.Num {
			return errors.New("第" + strconv.Itoa(row) + "行: 裁剪数量超出订单需求")
		}
		var count int64
		// 检查裁剪数量是否超出订单需求
		if err := facades.Orm.Query().Model(&models.ClothTailorCuttingPieces{}).Where("cloth_tailor_id", order.ID).Where("color", piece.Color).Where("size", piece.Size).Count(&count); err != nil {
			return err
		}
		if int(count)+piece.Num > int(rowContains.Num) {
			return errors.New("第" + strconv.Itoa(row) + "行: 裁剪数量超出订单需求，已裁剪数量为" + strconv.Itoa(int(count)) + "，订单需求为" + strconv.Itoa(int(rowContains.Num)))
		}
	}
	return nil
}

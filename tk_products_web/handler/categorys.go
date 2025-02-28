package handler

import (
	"context"

	"github.com/Numsina/tk_products/tk_products_srv/domain"
	"github.com/Numsina/tk_products/tk_products_srv/pb/v1"
)

func (h *Handler) CreateCategory(ctx context.Context, req *pb.CreateCategoryReq) (*pb.CreateCategoryResp, error) {
	category, err := h.svc.CreateCategory(ctx, domain.Categorys{
		Name:     req.Category.Name,
		Level:    int(req.Category.Level),
		Uid:      req.Category.Uid,
		ParentId: req.Category.ParentId,
		RootId:   req.Category.RootId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateCategoryResp{Category: &pb.Categorys{
		Id:       category.Id,
		Name:     category.Name,
		Level:    int32(category.Level),
		Uid:      category.Uid,
		ParentId: category.ParentId,
		RootId:   category.RootId,
	}}, nil
}

func (h *Handler) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryReq) (*pb.UpdateCategoryResp, error) {
	err := h.svc.UpdateCategory(ctx, domain.Categorys{
		Id:       req.Category.Id,
		Name:     req.Category.Name,
		Level:    int(req.Category.Level),
		Uid:      req.Category.Uid,
		ParentId: req.Category.ParentId,
		RootId:   req.Category.RootId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCategoryResp{}, nil
}

func (h *Handler) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryReq) (*pb.DeleteCategoryResp, error) {
	err := h.svc.DeleteCategory(ctx, req.Id, req.Uid)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCategoryResp{}, nil
}

func (h *Handler) ListCategorys(ctx context.Context, req *pb.ListCategorysReq) (*pb.ListCategorysResp, error) {
	list, err := h.svc.GetCategoryList(ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &pb.ListCategorysResp{Category: list}, nil
}

func (h *Handler) SearchCategorys(ctx context.Context, req *pb.SearChcategorytsReq) (*pb.SearchcategorysResp, error) {
	_, err := h.svc.GetCategoryByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return nil, err
	//return &pb.SearchcategorysResp{
	//	Category: &pb.Categorys{
	//		Id:             category.Id,
	//		Name:           category.Name,
	//		Level:          category.Level,
	//		Uid:            category.Uid,
	//		ParentId:       category.ParentId,
	//		ParentCategory: category.ParentCategory,
	//		RootId:         category.RootId,
	//	},
	//}, nil
}

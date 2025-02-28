package handler

import (
	"context"

	"github.com/Numsina/tk_products/tk_products_srv/domain"
	"github.com/Numsina/tk_products/tk_products_srv/pb/v1"
)

func (h *Handler) CreateBrand(ctx context.Context, req *pb.CreateBrandReq) (*pb.CreateBrandResp, error) {
	brand, err := h.svc.CreateBrand(ctx, domain.Brands{
		Name: req.Brand.Name,
		Logo: req.Brand.Logo,
		Uid:  req.Brand.Uid,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateBrandResp{Brand: &pb.Brands{
		Id:   brand.Id,
		Name: brand.Name,
		Logo: brand.Logo,
		Uid:  brand.Uid,
	}}, nil
}

func (h *Handler) UpdateBrand(ctx context.Context, req *pb.UpdateBrandReq) (*pb.UpdateBrandResp, error) {
	err := h.svc.UpdateBrand(ctx, domain.Brands{
		Id:   req.Brand.Id,
		Name: req.Brand.Name,
		Logo: req.Brand.Logo,
		Uid:  req.Brand.Uid,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateBrandResp{}, nil
}

func (h *Handler) DeleteBrand(ctx context.Context, req *pb.DeleteBrandReq) (*pb.DeleteBrandResp, error) {
	err := h.svc.DeleteBrand(ctx, req.Id, req.Uid)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBrandResp{}, nil
}

func (h *Handler) ListBrands(ctx context.Context, req *pb.ListBrandsReq) (*pb.ListBrandsResp, error) {
	list, err := h.svc.GetBrandList(ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	return &pb.ListBrandsResp{Brand: list}, nil
}

func (h *Handler) SearchBrands(ctx context.Context, req *pb.SearchBrandtsReq) (*pb.SearchBrandsResp, error) {
	_, err := h.svc.GetBrandByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	//return &pb.ListBrandsResp{Brand: list}, nil
	return &pb.SearchBrandsResp{}, nil
}

func (h *Handler) GetBrand(ctx context.Context, req *pb.GetBrandReq) (*pb.GetBrandResp, error) {
	list, err := h.svc.GetBrandByUid(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	return &pb.GetBrandResp{Brand: list}, nil
}

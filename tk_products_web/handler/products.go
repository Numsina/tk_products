package handler

import (
	"context"
	"sync"

	"github.com/Numsina/tk_products/tk_products_srv/domain"
	"github.com/Numsina/tk_products/tk_products_srv/pb/v1"
	"github.com/Numsina/tk_products/tk_products_srv/service"
)

type Handler struct {
	svc       service.Product
	rakingSvc service.RankingService
	pb.UnimplementedProductCatalogServiceServer
	pb.UnimplementedBrandServiceServer
	pb.UnimplementedCategoryServiceServer
}

func NewHandler(svc service.Product) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) GetProductRankingHots(ctx context.Context, req *pb.GetProductRankingHotsReq) (*pb.GetProductRankingHotsResp, error) {
	datas, err := h.rakingSvc.GetTopN(ctx, int(req.CategoryId))
	if err != nil {
		return nil, err
	}
	var d []*pb.Products
	for _, v := range datas {
		d = append(d, &pb.Products{
			Name:         v.Name,
			CategoryId:   v.CategoryId,
			CategoryName: v.CategoryName,
			BrandId:      v.BrandId,
			BrandName:    v.BrandName,
			Description:  v.Description,
			IsNew:        v.IsNew,
			Click:        v.Click,
			IsHot:        true,
			Favorite:     v.Favorite,
			Sale:         v.Sale,
			OnSale:       v.OnSale,
			MarkPrice:    v.MarkPrice,
			ShopPrice:    v.ShopPrice,
			Picture:      v.Picture,
			Images:       v.Images,
		})
	}
	return &pb.GetProductRankingHotsResp{Products: d}, nil
}

func (h *Handler) CreateProduct(ctx context.Context, req *pb.CreateProductReq) (*pb.CreateProductResp, error) {
	product, err := h.svc.CraeteProduct(ctx, domain.Products{
		Name:         req.Products.Name,
		CategoryId:   req.Products.CategoryId,
		CategoryName: req.Products.CategoryName,
		BrandId:      req.Products.BrandId,
		BrandName:    req.Products.BrandName,
		Description:  req.Products.Description,
		IsNew:        req.Products.IsNew,
		OnSale:       req.Products.OnSale,
		MarkPrice:    req.Products.MarkPrice,
		ShopPrice:    req.Products.ShopPrice,
		Picture:      req.Products.Picture,
		Images:       req.Products.Images,
		Uid:          req.Products.Uid,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResp{
		Id: product.Id,
		Products: &pb.Products{
			Name:         product.Name,
			CategoryId:   product.CategoryId,
			CategoryName: product.CategoryName,
			BrandId:      product.BrandId,
			BrandName:    product.BrandName,
			Description:  product.Description,
			IsNew:        product.IsNew,
			Click:        product.Click,
			Favorite:     product.Favorite,
			Sale:         product.Sale,
			OnSale:       product.OnSale,
			MarkPrice:    product.MarkPrice,
			ShopPrice:    product.ShopPrice,
			Picture:      product.Picture,
			Images:       product.Images,
		},
	}, nil
}

func (h *Handler) UpdateProduct(ctx context.Context, req *pb.UpdateProductReq) (*pb.UpdateProductResp, error) {
	product, err := h.svc.UpdateProduct(ctx, domain.Products{
		Name:         req.Products.Name,
		CategoryId:   req.Products.CategoryId,
		CategoryName: req.Products.CategoryName,
		BrandId:      req.Products.BrandId,
		BrandName:    req.Products.BrandName,
		Description:  req.Products.Description,
		IsNew:        req.Products.IsNew,
		OnSale:       req.Products.OnSale,
		MarkPrice:    req.Products.MarkPrice,
		ShopPrice:    req.Products.ShopPrice,
		Picture:      req.Products.Picture,
		Images:       req.Products.Images,
		Uid:          req.Products.Uid,
		Page:         req.Page,
		PageSize:     int32(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateProductResp{
		Id: product.Id,
	}, err
}

func (h *Handler) DeleteProduct(ctx context.Context, req *pb.DeleteProductReq) (*pb.DeleteProductResp, error) {
	err := h.svc.DeleteProduct(ctx, req.Id, req.Uid)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteProductResp{}, nil
}

func (h *Handler) ListProducts(ctx context.Context, req *pb.ListProductsReq) (*pb.ListProductsResp, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) SearchProducts(ctx context.Context, req *pb.SearchProductsReq) (*pb.SearchProductsResp, error) {
	list, err := h.svc.GetProductList(ctx, domain.Products{
		Name:         req.Products.Name,
		CategoryId:   req.Products.CategoryId,
		CategoryName: req.Products.CategoryName,
		BrandId:      req.Products.BrandId,
		BrandName:    req.Products.BrandName,
		Description:  req.Products.Description,
		IsNew:        req.Products.IsNew,
		OnSale:       req.Products.OnSale,
		MarkPrice:    req.Products.MarkPrice,
		ShopPrice:    req.Products.ShopPrice,
		Picture:      req.Products.Picture,
		Images:       req.Products.Images,
		Uid:          req.Products.Uid,
	}, 0, 0)
	if err != nil {
		return nil, err
	}
	return &pb.SearchProductsResp{Products: list}, err
}

func (h *Handler) GetProduct(ctx context.Context, req *pb.GetProductReq) (*pb.GetProductResp, error) {
	product, err := h.svc.GetProductInfoById(ctx, int32(req.Id))
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		//uid 从meta中获取
		err = h.svc.IncreateClick(ctx, int32(req.Id), 1)
		if err != nil {
			return
		}
	}()
	wg.Wait()
	return &pb.GetProductResp{Product: &pb.Products{
		Name:         product.Name,
		CategoryId:   product.CategoryId,
		CategoryName: product.CategoryName,
		BrandId:      product.BrandId,
		BrandName:    product.BrandName,
		Description:  product.Description,
		IsNew:        product.IsNew,
		Click:        product.Click,
		Favorite:     product.Favorite,
		Sale:         product.Sale,
		OnSale:       product.OnSale,
		MarkPrice:    product.MarkPrice,
		ShopPrice:    product.ShopPrice,
		Picture:      product.Picture,
		Images:       product.Images,
	}}, err
}

func (h *Handler) IncrementClick(ctx context.Context, req *pb.IncrementClickReq) (*pb.IncrementClickResp, error) {
	err := h.svc.IncreateClick(ctx, req.Id, req.Uid)
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (h *Handler) IncreateFavorite(ctx context.Context, req *pb.IncreateFavoriteReq) (*pb.IncreateFavoriteResp, error) {
	err := h.svc.IncreateFavorite(ctx, req.Id, req.Uid)
	if err != nil {
		return nil, err
	}
	return nil, err
}

package handler

import (
	"../common"
	"../domain/model"
	"../domain/service"
	category "../proto/category"
	"context"
	"github.com/rs/zerolog/log"
)

type Category struct {
	CategoryDataService service.ICategoryDataService
}

func (c *Category) CreateCategory(ctx context.Context, request *category.CategoryRequest, response *category.CreateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}
	categoryId, err := c.CategoryDataService.AddCategory(category)
	if err != nil {
		return err
	}
	response.Message = "category added."
	response.CategoryId = categoryId
	return  nil
}

func (c *Category) UpdateCategory(ctx context.Context, request *category.CategoryRequest, response *category.UpdateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(request, category)
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(category)
	if err != nil {
		return err
	}
	response.Message = "category updated."
	return nil
}

func (c *Category) DeleteCategory(ctx context.Context, request *category.DeleteCategoryRequest, response *category.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(request.CategoryId)
	if err != nil {
		return nil
	}
	response.Message = "category deleted."
	return nil
}

func (c *Category) FindCategoryByName(ctx context.Context, request *category.FindByNameRequest, response *category.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return nil
	}
	return common.SwapTo(category, response)
}

func (c *Category) FindCategoryByID(ctx context.Context, request *category.FindByIDRequest, response *category.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByID(request.CategoryId)
	if err != nil {
		return nil
	}
	return common.SwapTo(category, response)
}

func categoryToResponse(categorySlice []model.Category, response *category.FindAllResponse) {
	for _, categ := range categorySlice {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(categ, cr)
		if err != nil {
			log.Err(err)
			break
		}
		response.Category = append(response.Category, cr)
	}
}

func (c *Category) FindCategoryByLevel(ctx context.Context, request *category.FindByLevelRequest, response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByLevel(request.Level)
	if err != nil {
		return nil
	}
	categoryToResponse(categorySlice, response)
	return nil
}

func (c *Category) FindCategoryByParent(ctx context.Context, request *category.FindByParentRequest, response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByParent(request.ParentId)
	if err != nil {
		return nil
	}
	categoryToResponse(categorySlice, response)
	return nil
}

func (c *Category) FindAllCategory(ctx context.Context, request *category.FindAllRequest, response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return nil
	}
	categoryToResponse(categorySlice, response)
	return nil
}
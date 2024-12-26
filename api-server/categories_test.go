package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/stellora/airline/api-server/api"
)

func TestGetCategory(t *testing.T) {
	ctx, handler := handlerTest(t)
	categories = []api.Category{
		{Id: "A", Title: "Category A"},
		{Id: "B", Title: "Category B"},
	}

	t.Run("exists", func(t *testing.T) {
		resp, err := handler.GetCategory(ctx, api.GetCategoryRequestObject{
			Id: "A",
		})
		if err != nil {
			t.Fatal(err)
		}

		want := api.GetCategory200JSONResponse{
			Id:    "A",
			Title: "Category A",
		}
		if !reflect.DeepEqual(want, resp) {
			t.Errorf("got %v, want %v", resp, want)
		}
	})

	t.Run("does not exist", func(t *testing.T) {
		resp, err := handler.GetCategory(ctx, api.GetCategoryRequestObject{
			Id: "999",
		})
		if err != nil {
			t.Fatal(err)
		}
		if want := (&api.GetCategory404Response{}); !reflect.DeepEqual(resp, want) {
			t.Errorf("got %v, want %v", resp, want)
		}
	})
}

func TestListCategories(t *testing.T) {
	ctx, handler := handlerTest(t)
	categories = []api.Category{
		{Title: "Category 1"},
		{Title: "Category 2"},
	}

	resp, err := handler.ListCategories(ctx, api.ListCategoriesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}

	want := api.ListCategories200JSONResponse{
		api.Category{Title: "Category 1"},
		api.Category{Title: "Category 2"},
	}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}
}

func TestCreateCategory(t *testing.T) {
	ctx, handler := handlerTest(t)
	categories = []api.Category{}

	resp, err := handler.CreateCategory(ctx, api.CreateCategoryRequestObject{
		Body: &api.CreateCategoryJSONRequestBody{
			Title: "New Category",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.CreateCategory201Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkCategoryTitles(t, handler, []string{"New Category"})
}

func TestDeleteCategory(t *testing.T) {
	ctx, handler := handlerTest(t)
	categories = []api.Category{
		{Id: "1", Title: "Category 1"},
		{Id: "2", Title: "Category 2"},
	}

	resp, err := handler.DeleteCategory(ctx, api.DeleteCategoryRequestObject{
		Id: "1",
	})
	if err != nil {
		t.Fatal(err)
	}

	want := api.DeleteCategory204Response{}
	if !reflect.DeepEqual(want, resp) {
		t.Errorf("got %v, want %v", resp, want)
	}

	checkCategoryTitles(t, handler, []string{"Category 2"})
}

func checkCategoryTitles(t *testing.T, handler *Handler, want []string) {
	resp, err := handler.ListCategories(context.Background(), api.ListCategoriesRequestObject{})
	if err != nil {
		t.Fatal(err)
	}
	categories := resp.(api.ListCategories200JSONResponse)
	if len(categories) != len(want) {
		t.Errorf("got %d categories, want %d", len(categories), len(want))
	}
	for i, category := range categories {
		if category.Title != want[i] {
			t.Errorf("got title %q, want %q", category.Title, want[i])
		}
	}
}

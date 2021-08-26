package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RedHatInsights/sources-api-go/middleware"
	m "github.com/RedHatInsights/sources-api-go/model"
	"github.com/RedHatInsights/sources-api-go/util"
)

var testSourceData = []m.Source{
	{ID: 1, Name: "Source1", SourceTypeID: 1, TenantID: 1},
	{ID: 2, Name: "Source2", SourceTypeID: 1, TenantID: 1},
}

func TestApplicationSourceSubcollectionList(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/sources/v3.1/application_types/1/sources", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("limit", 99)
	c.Set("offset", -1)
	c.Set("filters", []middleware.Filter{})
	c.Set("tenantID", int64(1))
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := SourceList(c)
	if err != nil {
		t.Error(err)
	}

	if rec.Code != 200 {
		t.Error("Did not return 200")
	}

	var out util.Collection
	err = json.Unmarshal(rec.Body.Bytes(), &out)
	if err != nil {
		t.Error("Failed unmarshaling output")
	}

	if out.Meta.Limit != 99 {
		t.Error("limit not set correctly")
	}

	if out.Meta.Offset != -1 {
		t.Error("offset not set correctly")
	}

	if len(out.Data) != 1 {
		t.Error("not enough objects passed back from DB")
	}

	for _, src := range out.Data {
		s, ok := src.(map[string]interface{})
		if !ok {
			t.Error("model did not deserialize as a source")
		}

		if s["name"] != "Source1" {
			t.Error("ghosts infected the return")
		}
	}
}

func TestSourceList(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/sources/v3.1/sources", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("limit", 99)
	c.Set("offset", -1)
	c.Set("filters", []middleware.Filter{})
	c.Set("tenantID", int64(1))

	err := SourceList(c)
	if err != nil {
		t.Error(err)
	}

	if rec.Code != 200 {
		t.Error("Did not return 200")
	}

	var out util.Collection
	err = json.Unmarshal(rec.Body.Bytes(), &out)
	if err != nil {
		t.Error("Failed unmarshaling output")
	}

	if out.Meta.Limit != 99 {
		t.Error("limit not set correctly")
	}

	if out.Meta.Offset != -1 {
		t.Error("offset not set correctly")
	}

	if len(out.Data) != 2 {
		t.Error("not enough objects passed back from DB")
	}

	for _, src := range out.Data {
		s, ok := src.(map[string]interface{})
		if !ok {
			t.Error("model did not deserialize as a source")
		}

		if s["name"] != "Source1" && s["name"] != "Source2" {
			t.Error("ghosts infected the return")
		}
	}
}

func TestSourceGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/sources/v3.1/sources/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")
	c.Set("tenantID", int64(1))

	err := SourceGet(c)
	if err != nil {
		t.Error(err)
	}

	if rec.Code != 200 {
		t.Error("Did not return 200")
	}

	var outSrc m.SourceResponse
	err = json.Unmarshal(rec.Body.Bytes(), &outSrc)
	if err != nil {
		t.Error("Failed unmarshaling output")
	}

	if *outSrc.Name != "Source1" {
		t.Error("ghosts infected the return")
	}
}

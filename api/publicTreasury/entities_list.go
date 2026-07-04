package publicTreasury

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/JulianToledano/goingecko/v3/api/internal"
	"github.com/JulianToledano/goingecko/v3/api/publicTreasury/types"
)

// entitiesListOption is an interface that extends internal.Option to provide options specific to the entities list endpoint.
type entitiesListOption interface {
	internal.Option

	isEntitiesListOption()
}

// WithEntitiesEntityType returns an entitiesListOption that filters entities by type.
// Supported values are "company" and "government".
func WithEntitiesEntityType(entityType string) entitiesListOption {
	return entitiesEntityTypeOption{entityType}
}

// WithEntitiesPerPage returns an entitiesListOption that sets the number of entities returned per page.
func WithEntitiesPerPage(perPage int) entitiesListOption {
	return entitiesPerPageOption{perPage}
}

// WithEntitiesPage returns an entitiesListOption that sets the page of entities to return.
func WithEntitiesPage(page int) entitiesListOption {
	return entitiesPageOption{page}
}

// EntitiesList allows you to query all supported public treasury entities with entity ID, name, symbol, and country.
func (c *PublicTreasuryClient) EntitiesList(ctx context.Context, options ...entitiesListOption) ([]types.Entity, error) {
	params := url.Values{}
	for _, opt := range options {
		opt.Apply(&params)
	}

	rUrl := fmt.Sprintf("%s/entities/list?%s", c.URL, params.Encode())
	resp, err := c.MakeReq(ctx, rUrl)
	if err != nil {
		return nil, err
	}

	var data []types.Entity
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type entitiesEntityTypeOption struct{ entityType string }

func (o entitiesEntityTypeOption) Apply(v *url.Values) {
	v.Set("entity_type", o.entityType)
}

func (o entitiesEntityTypeOption) isEntitiesListOption() {}

type entitiesPerPageOption struct{ perPage int }

func (o entitiesPerPageOption) Apply(v *url.Values) {
	v.Set("per_page", strconv.Itoa(o.perPage))
}

func (o entitiesPerPageOption) isEntitiesListOption() {}

type entitiesPageOption struct{ page int }

func (o entitiesPageOption) Apply(v *url.Values) {
	v.Set("page", strconv.Itoa(o.page))
}

func (o entitiesPageOption) isEntitiesListOption() {}

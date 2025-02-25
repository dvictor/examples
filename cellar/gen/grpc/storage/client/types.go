// Code generated by goa v3.0.2, DO NOT EDIT.
//
// storage gRPC client types
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o
// $(GOPATH)/src/goa.design/examples/cellar

package client

import (
	"unicode/utf8"

	storagepb "goa.design/examples/cellar/gen/grpc/storage/pb"
	storage "goa.design/examples/cellar/gen/storage"
	storageviews "goa.design/examples/cellar/gen/storage/views"
	goa "goa.design/goa/v3/pkg"
)

// NewListRequest builds the gRPC request type from the payload of the "list"
// endpoint of the "storage" service.
func NewListRequest() *storagepb.ListRequest {
	message := &storagepb.ListRequest{}
	return message
}

// NewListResult builds the result type of the "list" endpoint of the "storage"
// service from the gRPC response type.
func NewListResult(message *storagepb.StoredBottleCollection) storageviews.StoredBottleCollectionView {
	result := make([]*storageviews.StoredBottleView, len(message.Field))
	for i, val := range message.Field {
		result[i] = &storageviews.StoredBottleView{
			ID:      &val.Id,
			Name:    &val.Name,
			Vintage: &val.Vintage,
		}
		if val.Description != "" {
			result[i].Description = &val.Description
		}
		if val.Rating != 0 {
			result[i].Rating = &val.Rating
		}
		if val.Winery != nil {
			result[i].Winery = protobufStoragepbWineryToStorageviewsWineryView(val.Winery)
		}
		if val.Composition != nil {
			result[i].Composition = make([]*storageviews.ComponentView, len(val.Composition))
			for j, val := range val.Composition {
				result[i].Composition[j] = &storageviews.ComponentView{
					Varietal: &val.Varietal,
				}
				if val.Percentage != 0 {
					result[i].Composition[j].Percentage = &val.Percentage
				}
			}
		}
	}
	return result
}

// NewShowRequest builds the gRPC request type from the payload of the "show"
// endpoint of the "storage" service.
func NewShowRequest(payload *storage.ShowPayload) *storagepb.ShowRequest {
	message := &storagepb.ShowRequest{
		Id: payload.ID,
	}
	return message
}

// NewShowResult builds the result type of the "show" endpoint of the "storage"
// service from the gRPC response type.
func NewShowResult(message *storagepb.ShowResponse) *storageviews.StoredBottleView {
	result := &storageviews.StoredBottleView{
		ID:      &message.Id,
		Name:    &message.Name,
		Vintage: &message.Vintage,
	}
	if message.Description != "" {
		result.Description = &message.Description
	}
	if message.Rating != 0 {
		result.Rating = &message.Rating
	}
	if message.Winery != nil {
		result.Winery = protobufStoragepbWineryToStorageviewsWineryView(message.Winery)
	}
	if message.Composition != nil {
		result.Composition = make([]*storageviews.ComponentView, len(message.Composition))
		for i, val := range message.Composition {
			result.Composition[i] = &storageviews.ComponentView{
				Varietal: &val.Varietal,
			}
			if val.Percentage != 0 {
				result.Composition[i].Percentage = &val.Percentage
			}
		}
	}
	return result
}

// NewShowNotFoundError builds the error type of the "show" endpoint of the
// "storage" service from the gRPC error response type.
func NewShowNotFoundError(message *storagepb.ShowNotFoundError) *storage.NotFound {
	er := &storage.NotFound{
		Message: message.Message_,
		ID:      message.Id,
	}
	return er
}

// NewAddRequest builds the gRPC request type from the payload of the "add"
// endpoint of the "storage" service.
func NewAddRequest(payload *storage.Bottle) *storagepb.AddRequest {
	message := &storagepb.AddRequest{
		Name:    payload.Name,
		Vintage: payload.Vintage,
	}
	if payload.Description != nil {
		message.Description = *payload.Description
	}
	if payload.Rating != nil {
		message.Rating = *payload.Rating
	}
	if payload.Winery != nil {
		message.Winery = svcStorageWineryToStoragepbWinery(payload.Winery)
	}
	if payload.Composition != nil {
		message.Composition = make([]*storagepb.Component, len(payload.Composition))
		for i, val := range payload.Composition {
			message.Composition[i] = &storagepb.Component{
				Varietal: val.Varietal,
			}
			if val.Percentage != nil {
				message.Composition[i].Percentage = *val.Percentage
			}
		}
	}
	return message
}

// NewAddResult builds the result type of the "add" endpoint of the "storage"
// service from the gRPC response type.
func NewAddResult(message *storagepb.AddResponse) string {
	result := message.Field
	return result
}

// NewRemoveRequest builds the gRPC request type from the payload of the
// "remove" endpoint of the "storage" service.
func NewRemoveRequest(payload *storage.RemovePayload) *storagepb.RemoveRequest {
	message := &storagepb.RemoveRequest{
		Id: payload.ID,
	}
	return message
}

// NewRateRequest builds the gRPC request type from the payload of the "rate"
// endpoint of the "storage" service.
func NewRateRequest(payload map[uint32][]string) *storagepb.RateRequest {
	message := &storagepb.RateRequest{}
	message.Field = make(map[uint32]*storagepb.ArrayOfString, len(payload))
	for key, val := range payload {
		tk := key
		tv := &storagepb.ArrayOfString{}
		tv.Field = make([]string, len(val))
		for i, val := range val {
			tv.Field[i] = val
		}
		message.Field[tk] = tv
	}
	return message
}

// NewMultiAddRequest builds the gRPC request type from the payload of the
// "multi_add" endpoint of the "storage" service.
func NewMultiAddRequest(payload []*storage.Bottle) *storagepb.MultiAddRequest {
	message := &storagepb.MultiAddRequest{}
	message.Field = make([]*storagepb.Bottle, len(payload))
	for i, val := range payload {
		message.Field[i] = &storagepb.Bottle{
			Name:    val.Name,
			Vintage: val.Vintage,
		}
		if val.Description != nil {
			message.Field[i].Description = *val.Description
		}
		if val.Rating != nil {
			message.Field[i].Rating = *val.Rating
		}
		if val.Winery != nil {
			message.Field[i].Winery = svcStorageWineryToStoragepbWinery(val.Winery)
		}
		if val.Composition != nil {
			message.Field[i].Composition = make([]*storagepb.Component, len(val.Composition))
			for j, val := range val.Composition {
				message.Field[i].Composition[j] = &storagepb.Component{
					Varietal: val.Varietal,
				}
				if val.Percentage != nil {
					message.Field[i].Composition[j].Percentage = *val.Percentage
				}
			}
		}
	}
	return message
}

// NewMultiAddResult builds the result type of the "multi_add" endpoint of the
// "storage" service from the gRPC response type.
func NewMultiAddResult(message *storagepb.MultiAddResponse) []string {
	result := make([]string, len(message.Field))
	for i, val := range message.Field {
		result[i] = val
	}
	return result
}

// NewMultiUpdateRequest builds the gRPC request type from the payload of the
// "multi_update" endpoint of the "storage" service.
func NewMultiUpdateRequest(payload *storage.MultiUpdatePayload) *storagepb.MultiUpdateRequest {
	message := &storagepb.MultiUpdateRequest{}
	if payload.Ids != nil {
		message.Ids = make([]string, len(payload.Ids))
		for i, val := range payload.Ids {
			message.Ids[i] = val
		}
	}
	if payload.Bottles != nil {
		message.Bottles = make([]*storagepb.Bottle, len(payload.Bottles))
		for i, val := range payload.Bottles {
			message.Bottles[i] = &storagepb.Bottle{
				Name:    val.Name,
				Vintage: val.Vintage,
			}
			if val.Description != nil {
				message.Bottles[i].Description = *val.Description
			}
			if val.Rating != nil {
				message.Bottles[i].Rating = *val.Rating
			}
			if val.Winery != nil {
				message.Bottles[i].Winery = svcStorageWineryToStoragepbWinery(val.Winery)
			}
			if val.Composition != nil {
				message.Bottles[i].Composition = make([]*storagepb.Component, len(val.Composition))
				for j, val := range val.Composition {
					message.Bottles[i].Composition[j] = &storagepb.Component{
						Varietal: val.Varietal,
					}
					if val.Percentage != nil {
						message.Bottles[i].Composition[j].Percentage = *val.Percentage
					}
				}
			}
		}
	}
	return message
}

// ValidateStoredBottleCollection runs the validations defined on
// StoredBottleCollection.
func ValidateStoredBottleCollection(message *storagepb.StoredBottleCollection) (err error) {
	for _, e := range message.Field {
		if e != nil {
			if err2 := ValidateStoredBottle(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStoredBottle runs the validations defined on StoredBottle.
func ValidateStoredBottle(message *storagepb.StoredBottle) (err error) {
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	if message.Winery != nil {
		if err2 := ValidateWinery(message.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if message.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 1900, true))
	}
	if message.Vintage > 2020 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 2020, false))
	}
	for _, e := range message.Composition {
		if e != nil {
			if err2 := ValidateComponent(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if utf8.RuneCountInString(message.Description) > 2000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", message.Description, utf8.RuneCountInString(message.Description), 2000, false))
	}
	if message.Rating < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", message.Rating, 1, true))
	}
	if message.Rating > 5 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", message.Rating, 5, false))
	}
	return
}

// ValidateWinery runs the validations defined on Winery.
func ValidateWinery(message *storagepb.Winery) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("message.region", message.Region, "(?i)[a-z '\\.]+"))
	err = goa.MergeErrors(err, goa.ValidatePattern("message.country", message.Country, "(?i)[a-z '\\.]+"))
	err = goa.MergeErrors(err, goa.ValidatePattern("message.url", message.Url, "(?i)^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	return
}

// ValidateComponent runs the validations defined on Component.
func ValidateComponent(message *storagepb.Component) (err error) {
	err = goa.MergeErrors(err, goa.ValidatePattern("message.varietal", message.Varietal, "[A-Za-z' ]+"))
	if utf8.RuneCountInString(message.Varietal) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.varietal", message.Varietal, utf8.RuneCountInString(message.Varietal), 100, false))
	}
	if message.Percentage < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.percentage", message.Percentage, 1, true))
	}
	if message.Percentage > 100 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.percentage", message.Percentage, 100, false))
	}
	return
}

// ValidateShowResponse runs the validations defined on ShowResponse.
func ValidateShowResponse(message *storagepb.ShowResponse) (err error) {
	if message.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "message"))
	}
	if utf8.RuneCountInString(message.Name) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("message.name", message.Name, utf8.RuneCountInString(message.Name), 100, false))
	}
	if message.Winery != nil {
		if err2 := ValidateWinery(message.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if message.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 1900, true))
	}
	if message.Vintage > 2020 {
		err = goa.MergeErrors(err, goa.InvalidRangeError("message.vintage", message.Vintage, 2020, false))
	}
	for _, e := range message.Composition {
		if e != nil {
			if err2 := ValidateComponent(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if message.Description != "" {
		if utf8.RuneCountInString(message.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("message.description", message.Description, utf8.RuneCountInString(message.Description), 2000, false))
		}
	}
	if message.Rating != 0 {
		if message.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", message.Rating, 1, true))
		}
	}
	if message.Rating != 0 {
		if message.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("message.rating", message.Rating, 5, false))
		}
	}
	return
}

// ValidateMultiAddResponse runs the validations defined on MultiAddResponse.
func ValidateMultiAddResponse(message *storagepb.MultiAddResponse) (err error) {
	if message.Field == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("field", "message"))
	}
	return
}

// svcStorageviewsWineryViewToStoragepbWinery builds a value of type
// *storagepb.Winery from a value of type *storageviews.WineryView.
func svcStorageviewsWineryViewToStoragepbWinery(v *storageviews.WineryView) *storagepb.Winery {
	res := &storagepb.Winery{}
	if v.Name != nil {
		res.Name = *v.Name
	}
	if v.Region != nil {
		res.Region = *v.Region
	}
	if v.Country != nil {
		res.Country = *v.Country
	}
	if v.URL != nil {
		res.Url = *v.URL
	}

	return res
}

// protobufStoragepbWineryToStorageviewsWineryView builds a value of type
// *storageviews.WineryView from a value of type *storagepb.Winery.
func protobufStoragepbWineryToStorageviewsWineryView(v *storagepb.Winery) *storageviews.WineryView {
	res := &storageviews.WineryView{
		Name:    &v.Name,
		Region:  &v.Region,
		Country: &v.Country,
	}
	if v.Url != "" {
		res.URL = &v.Url
	}

	return res
}

// protobufStoragepbWineryToStorageWinery builds a value of type
// *storage.Winery from a value of type *storagepb.Winery.
func protobufStoragepbWineryToStorageWinery(v *storagepb.Winery) *storage.Winery {
	res := &storage.Winery{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
	}
	if v.Url != "" {
		res.URL = &v.Url
	}

	return res
}

// svcStorageWineryToStoragepbWinery builds a value of type *storagepb.Winery
// from a value of type *storage.Winery.
func svcStorageWineryToStoragepbWinery(v *storage.Winery) *storagepb.Winery {
	res := &storagepb.Winery{
		Name:    v.Name,
		Region:  v.Region,
		Country: v.Country,
	}
	if v.URL != nil {
		res.Url = *v.URL
	}

	return res
}

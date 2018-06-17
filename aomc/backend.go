package aomc

import (
	"github.com/philippgille/apiomat-go/aomc/dto"
)

// GetRawBackends returns the raw backends for the given customer.
// "Raw" means the struct is mapped 1:1 to an "Application" JSON.
//
// The returned backends don't necessarily belong to the customer (he doesn't have to be the "app admin").
// They're just the ones where the customer has READ permissions.
func (client Client) GetRawBackends(customer string) ([]dto.Backend, error) {
	jsonString, err := client.Get("customers/"+customer+"/apps", nil)
	if err != nil {
		return nil, err
	}
	result, err := ConvertRawBackendsFromJSON(jsonString)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetRawBackendByName returns the raw backend for the given customer and backend name.
// "Raw" means the struct is mapped 1:1 to an "Application" JSON.
//
// The returned backend doesn't necessarily belong to the customer.
// It's just one where the customer has READ permissions.
func (client Client) GetRawBackendByName(name string, customer string) (dto.Backend, error) {
	jsonString, err := client.Get("customers/"+customer+"/apps/"+name, nil)
	if err != nil {
		return dto.Backend{}, err
	}
	result, err := ConvertRawBackendFromJSON(jsonString)
	if err != nil {
		return dto.Backend{}, err
	}
	return result, nil
}

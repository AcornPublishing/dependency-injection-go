// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func initializeDeps() (*Fetcher, error) {
	cache, err := ProvideCache()
	if err != nil {
		return nil, err
	}
	fetcher := ProvideFetcher(cache)
	return fetcher, nil
}

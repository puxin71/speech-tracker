package pkg

type MockConfigProvider struct{}

func (MockConfigProvider) GetResourcePath() string {
	return "./../../resources"
}

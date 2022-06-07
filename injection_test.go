package headwater

import "testing"

type DataAccess struct {
	Id int
}

type ApiClient struct {
	Id  int
	Url string
}

type Container struct {
	Url        Injector[string]
	DataAccess Injector[*DataAccess]
	ApiClient  Injector[*ApiClient]
}

func CreateContainer() Container {
	var id = 0
	getId := func() int {
		id += 1
		return id
	}
	var injection Container
	injection = Container{
		Url: CreateValue("http://localhost"),
		DataAccess: CreateFactory(func() *DataAccess {
			return &DataAccess{
				Id: getId(),
			}
		}),
		ApiClient: CreateSingleton(func() *ApiClient {
			ok, url := injection.Url.Get()
			if !ok {
				url = "http://127.0.0.1"
			}

			return &ApiClient{
				Id:  getId(),
				Url: url,
			}
		}),
	}
	return injection
}

func TestValueInjector(t *testing.T) {
	container := CreateContainer()
	ok, url := container.Url.Get()

	if !ok {
		t.Error("Received not ok")
	}

	want := "http://localhost"

	if url != want {
		t.Errorf("Received: %v, Expected %v", url, want)
	}
}

func TestFactoryInjector(t *testing.T) {
	container := CreateContainer()
	ok, dataAccess := container.DataAccess.Get()

	if !ok {
		t.Error("Received not ok")
	}

	if dataAccess == nil {
		t.Errorf("Received: %v", dataAccess)
	}

	ok, nextDataAccess := container.DataAccess.Get()

	if !ok {
		t.Error("Received not ok")
	}

	if dataAccess == nextDataAccess {
		t.Errorf("Received the same instance")
	}
}

func TestNestedInjector(t *testing.T) {
	container := CreateContainer()
	ok, apiClient := container.ApiClient.Get()

	if !ok {
		t.Error("Received not ok")
	}

	if apiClient == nil {
		t.Errorf("Received: %v", apiClient)
	}

	want := "http://localhost"

	if apiClient.Url != want {
		t.Errorf("Received: %v, Expected %v", apiClient.Url, want)
	}
}

func TestSingletonInjector(t *testing.T) {
	container := CreateContainer()
	ok, apiClient := container.ApiClient.Get()

	if !ok {
		t.Error("Received not ok")
	}

	ok, nextApiClient := container.ApiClient.Get()

	if !ok {
		t.Error("Received not ok")
	}

	if apiClient != nextApiClient {
		t.Error("Received multiple instances")
	}
}

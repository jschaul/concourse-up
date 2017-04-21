package concourse_test

import (
	"bitbucket.org/engineerbetter/concourse-up/config"
	"bitbucket.org/engineerbetter/concourse-up/terraform"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestConcourse(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Concourse Suite")
}

type FakeConfigClient struct {
	FakeLoad         func() (*config.Config, error)
	FakeUpdate       func(*config.Config) error
	FakeLoadOrCreate func() (*config.Config, bool, error)
	FakeStoreAsset   func(filename string, contents []byte) error
	FakeLoadAsset    func(filename string) ([]byte, error)
	FakeHasAsset     func(filename string) (bool, error)
}

func (client *FakeConfigClient) Load() (*config.Config, error) {
	return client.FakeLoad()
}

func (client *FakeConfigClient) Update(config *config.Config) error {
	return client.FakeUpdate(config)
}

func (client *FakeConfigClient) LoadOrCreate() (*config.Config, bool, error) {
	return client.FakeLoadOrCreate()
}

func (client *FakeConfigClient) StoreAsset(filename string, contents []byte) error {
	return client.FakeStoreAsset(filename, contents)
}

func (client *FakeConfigClient) LoadAsset(filename string) ([]byte, error) {
	return client.FakeLoadAsset(filename)
}

func (client *FakeConfigClient) HasAsset(filename string) (bool, error) {
	return client.FakeHasAsset(filename)
}

type FakeTerraformClient struct {
	FakeOutput  func() (*terraform.Metadata, error)
	FakeApply   func() error
	FakeDestroy func() error
	FakeCleanup func() error
}

func (client *FakeTerraformClient) Output() (*terraform.Metadata, error) {
	return client.FakeOutput()
}

func (client *FakeTerraformClient) Apply() error {
	return client.FakeApply()
}

func (client *FakeTerraformClient) Destroy() error {
	return client.FakeDestroy()
}

func (client *FakeTerraformClient) Cleanup() error {
	return client.FakeCleanup()
}

type FakeBoshInitClient struct {
	FakeDeploy  func() ([]byte, error)
	FakeDelete  func() error
	FakeCleanup func() error
}

func (client *FakeBoshInitClient) Deploy() ([]byte, error) {
	return client.FakeDeploy()
}

func (client *FakeBoshInitClient) Delete() error {
	return client.FakeDelete()
}

func (client *FakeBoshInitClient) Cleanup() error {
	return client.FakeCleanup()
}
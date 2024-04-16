package configurator

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/yaml.v3"
)

type MockedOs struct {
	mock.Mock
}

func (mio *MockedOs) ReadFile(filePath string) ([]byte, error) {
	args := mio.Called(filePath)
	return args.Get(0).([]byte), args.Error(1)
}

func (mio *MockedOs) Stat(filePath string) (os.FileInfo, error) {
	args := mio.Called(filePath)
	a0, _ := args.Get(0).(os.FileInfo)
	return a0, args.Error(1)
}

func (mio *MockedOs) WriteFile(name string, data []byte, perm os.FileMode) error {
	return mio.Called(name, data, perm).Error(0)
}

func TestReadConfigurationFromFileEmpty(t *testing.T) {
	mos := MockedOs{}
	configFile := "someFile"
	mos.On("ReadFile", configFile).Return([]byte{}, nil)
	c := NewConfigurator(&mos)

	conf, err := c.readConfigurationFromFile(configFile)

	assert.Nil(t, err)
	assert.Equal(t, &Configuration{}, conf)
	mos.AssertExpectations(t)
}

func TestIsFilePresent(t *testing.T) {
	tests := []struct {
		inErr error
		out   bool
	}{
		{nil, true},
		{os.ErrNotExist, false},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v,%v", test.inErr, test.out)
		t.Run(name, func(t *testing.T) {
			mos := MockedOs{}
			configFile := "someFile"
			mos.On("Stat", configFile).Return(nil, test.inErr)
			c := NewConfigurator(&mos)

			isPresent := c.isFilePresent(configFile)

			assert.Equal(t, test.out, isPresent)
			mos.AssertExpectations(t)
		})
	}
}

func TestSaveConfigurationToFile(t *testing.T) {
	config := Configuration{
		StartPath: "somePath",
		RealClean: false,
		IsReady:   false,
		SizeConfig: SizeConfig{
			Threshold: 2,
			CatchZero: true,
		},
		Exts: Exts{
			WhiteList: []string{"a"},
			BlackList: []string{"b"},
		},
		Files: Files{
			WhiteList: []string{"c"},
			BlackList: []string{"d"},
		},
		Contents: []string{"someContent"},
	}
	mos := MockedOs{}
	filePath := "someFilePath"
	bytes, err := yaml.Marshal(config)
	assert.Nil(t, err)
	mos.On("WriteFile", filePath, bytes, os.FileMode(0666)).Return(nil)
	c := NewConfigurator(&mos)

	err = c.SaveConfigurationToFile(&config, filePath)

	assert.Nil(t, err)
}

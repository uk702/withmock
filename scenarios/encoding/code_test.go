package withdeps

import (
	"testing"

	"fmt" // mock

	"github.com/golang/mock/gomock"

	"encoding/json" // mock
	"encoding/base64" // mock
)

func TestShow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fmt.MOCK().SetController(ctrl)

	// We need some test data
	data := "one\ntwo\nthree"

	// The test data contains three lines that we expect to be printed.
	fmt.EXPECT().Printf("%d: %s\n", 1, "one")
	fmt.EXPECT().Printf("%d: %s\n", 2, "two")
	fmt.EXPECT().Printf("%d: %s\n", 3, "three")

	// Run the function we want to test
	err := Show(data)

	if err != nil {
		t.Errorf("Unexpected error return: %s", err)
	}
}

func TestShow2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	json.MOCK().SetController(ctrl)
	base64.MOCK().SetController(ctrl)
}

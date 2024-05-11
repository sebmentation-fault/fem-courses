package projector_test

import (
	"reflect"
	"testing"

	"github.com/sebmentation-fault/projector/pkg/projector"
)

func getOpts(args []string) *projector.Opts {
	return &projector.Opts {
		Args: args,
		Config: "",
		Pwd: "",
	}
}

func testConfig(t *testing.T, args []string, expected []string, operation projector.Operation) {
	opts := getOpts(args)
	config, err := projector.NewConfig(opts)

	if err != nil {
		t.Errorf("expected to get no error but %v", err)
		return
	}

	if !reflect.DeepEqual(expected, config.Args) {
		t.Errorf("expected %+v but got %+v", expected, config.Args)
		return
	}

	if config.Operation != operation {
		t.Errorf("expected operation %+v but got %+v", operation, config.Operation)
		return
	}
}

func TestConfigPrint(t *testing.T) {
	testConfig(t, []string{}, []string{}, projector.Print)
}

func TestConfigPrintKey(t *testing.T) {
	testConfig(t, []string{"foo"}, []string{"foo"}, projector.Print)
}

func TestConfigAddKey(t *testing.T) {
	testConfig(t, []string{"add", "foo", "bar"}, []string{"foo", "bar"}, projector.Add)
}

func TestConfigRemoveKey(t *testing.T) {
	testConfig(t, []string{"rm", "foo"}, []string{"foo"}, projector.Remove)
}

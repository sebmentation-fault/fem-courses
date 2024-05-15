package projector_test

import (
	"testing"

	"github.com/sebmentation-fault/projector/pkg/projector"
)

func getData() *projector.Data {
    return &projector.Data{
        Projector: map[string]map[string]string{
            "/": {
                "foo": "bar1",
                "fem": "am_learning",
            },
            "/foo": {
                "foo": "bar2",
            },
            "/foo/bar": {
                "foo": "bar3",
            },
        },
    }
}

func getProjector(pwd string, data *projector.Data) *projector.Projector {
	config := &projector.Config{
		Args:      []string{},
		Operation: projector.Print,
		Pwd:       pwd,
		Config:    "Hello, Sebastion",
	}

	return projector.CreateProjector(config, data)
}

func testProjector(t *testing.T, proj *projector.Projector, key, value string) {
	v, ok := proj.GetValue(key)
	if !ok {
		t.Errorf("expected to find value \"%v\"", value)
	}
	if value != v {
		t.Errorf("expected to find %v but received %v", value, v)
	}
}

func TestGetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	testProjector(t, proj, "foo", "bar3")
	testProjector(t, proj, "fem", "am_learning")
}

func TestSetValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	testProjector(t, proj, "foo", "bar3")
	proj.SetValue("foo", "bar4")
	testProjector(t, proj, "foo", "bar4")

	proj.SetValue("fem", "stopped_learning")
	testProjector(t, proj, "fem", "stopped_learning")

	proj = getProjector("/", data)
	testProjector(t, proj, "fem", "am_learning")
}

func TestRemoveValue(t *testing.T) {
	data := getData()
	proj := getProjector("/foo/bar", data)

	testProjector(t, proj, "foo", "bar3")
	proj.RemoveValue("foo")
	testProjector(t, proj, "foo", "bar2")

	proj.RemoveValue("fem")
	testProjector(t, proj, "fem", "am_learning")
}

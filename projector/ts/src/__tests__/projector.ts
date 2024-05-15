import { Operation } from "../config";
import { Projector } from "../projector";

function createData() {
    return {
        projector: {
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
        }
    };
}

function getProjector(pwd: string, data = createData()): Projector {
    return new Projector({
        args: [],
        operation: Operation.Print,
        pwd,
        config: "Hello, Seb",
    }, data);
}

test("getValueAll", function() {
    const proj = getProjector("/foo/bar");
    expect(proj.getValueAll()).toEqual({
        "fem": "am_learning",
        "foo": "bar3",
    });

});

test("getValue", function() {
    let proj = getProjector("/foo/bar");
    expect(proj.getValue("foo")).toEqual("bar3");

    proj = getProjector("/foo");
    expect(proj.getValue("foo")).toEqual("bar2");
    expect(proj.getValue("fem")).toEqual("am_learning");

    proj = getProjector("/");
    expect(proj.getValue("foo")).toEqual("bar1");
});

test("setValue", function() {
    let data = createData();
    let proj = getProjector("/foo/bar", data);
    proj.setValue("foo", "baz");
    expect(proj.getValue("foo")).toEqual("baz");

    proj.setValue("fem", "stopped_learning");
    expect(proj.getValue("fem")).toEqual("stopped_learning");

    proj = getProjector("/", data);
    expect(proj.getValue("fem")).toEqual("am_learning");
});

test("removeValue", function() {
    const proj = getProjector("/foo/bar");

    proj.removeValue("fem");
    expect(proj.getValue("fem")).toEqual("am_learning");

    proj.removeValue("foo");
    expect(proj.getValue("foo")).toEqual("bar2");
});

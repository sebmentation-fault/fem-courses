enum RSEnum {
    Foo(fn() -> i32),
    Foo2(Option<i32>),
    Bar(String),
    Baz(Vec<String>),
}

fn bar -> i32 {
    return 5;
}

fn main() {
    let foo = RSEnum::Foo(bar);

    match foo {
        RSEnum::Foo(fn_value) => {},
        RSEnum::Foo2(Some(value)) => {},
        RSEnum::Foo2(None) => {},
        _ => {},
    }
}

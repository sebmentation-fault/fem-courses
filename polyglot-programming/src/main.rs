
fn error_me(throw: bool) -> Result<(), usize> {
    if throw {
        return Err(7);
    }

    return Ok(());
}

fn main() -> Result<(), usize> {

    let value = error_me(false)?;

    if error_me(true).is_ok() {

    }

    return Ok(());
}

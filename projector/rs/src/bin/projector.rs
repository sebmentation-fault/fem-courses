use clap::Parser;
use anyhow::Result;

use projectorrs::{config::Config, opts::Opts};

fn main() -> Result<()> {

    let config: Config = Opts::parse().try_into()?;

    print!("{:?}", config);

    return Ok(());
}

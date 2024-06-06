use clap::Parser;
use anyhow::Result;

use projectorrs::{config::{Config, Operation}, opts::Opts, projector::Projector};

fn main() -> Result<()> {

    let config: Config = Opts::parse().try_into()?;
    let mut proj: Projector = Projector::new_projector(config.config, config.pwd);

    // if config.pwd == PathBuf::from("") {
    //
    // }

    match config.operation {
        Operation::Print(None) => {
            let value = proj.get_value_all();
            let value = serde_json::to_string(&value)?;
            println!("{}", value);
        },
        Operation::Print(Some(key)) => {
            proj.get_value(&key).map(|x| {
                println!("{}", x);
            });
        },
        Operation::Add(key, value) => {
            proj.set_value(&key, &value);
            proj.save()?;
        },
        Operation::Remove(key) => {
            proj.remove_value(&key);
            proj.save()?;
        },
    }

    return Ok(());
}

use std::path::PathBuf;
use clap::Parser;

#[derive(Debug, Parser)]
#[clap()]
pub struct Opts {
    pub args: Vec<String>,

    #[clap(short = 'c', long = "config")]
    pub config: Option<PathBuf>,

    #[clap(short = 'p', long = "pwd")]
    pub pwd: Option<PathBuf>,

}

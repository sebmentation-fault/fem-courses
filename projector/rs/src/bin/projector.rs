use clap::Parser;
use projectorrs::opts::Opts;

fn main() {

    let opts = Opts::parse();

    print!("{:?}", opts);
}

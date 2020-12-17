use clap::{App, SubCommand};
use std::result::Result;

mod commands;

fn main() -> Result<(), anyhow::Error> {
    let matches = App::new("gpm")
        .version("0.1")
        .author("TODO <TODO@users.noreply.github.com>")
        .about("Games Package Manager utility")
        .subcommand(
            SubCommand::with_name("init")
                .version("0.1")
                .about("creates an mod project in the current directory"),
        )
        .get_matches();

    match matches.subcommand() {
        ("init", _init_arg) => commands::init::init()?,
        _ => println!("sub command unknown or unspecified"),
    };

    Ok(())
}

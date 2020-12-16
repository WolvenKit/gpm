extern crate games_package_manager;
extern crate clap;
extern crate termcolor;

use clap::{App, SubCommand};
use std::result::Result;

mod commands;
mod prompt;

fn main() -> Result<(), &'static str> {
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

  if let Some(matches) = matches.subcommand_matches("init") {
    commands::init::init()?;
  }

  Ok(())
}
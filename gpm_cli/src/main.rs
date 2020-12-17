use clap::{App, Arg, SubCommand};
use std::path::PathBuf;
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
        .subcommand(
            SubCommand::with_name("package")
                .about("create a redistributable archive of a mod")
                .arg(
                    Arg::with_name("input_dir")
                        .short("i")
                        .takes_value(true)
                        .help("the directory containing the mod to package"),
                )
                .arg(
                    Arg::with_name("output_file")
                        .short("o")
                        .takes_value(true)
                        .required(true)
                        .help("the output file to create"),
                ),
        )
        .get_matches();

    match matches.subcommand() {
        ("init", _) => commands::init::init()?,
        ("package", Some(archive_arg)) => {
            commands::package::package(commands::package::PackageParameter {
                input_dir: PathBuf::from(archive_arg.value_of("input_dir").unwrap_or(".")),
                output_file: PathBuf::from(archive_arg.value_of("output_file").unwrap()), //unwrap: output_file is required
            })?;
        }
        _ => println!("sub command unknown or unspecified"),
    };

    Ok(())
}

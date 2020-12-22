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
        .subcommand(
            SubCommand::with_name("remote")
                .about("control the list of known remotes")
                .subcommands(vec![
                    SubCommand::with_name("list").about("shows the list of known remotes"),
                    SubCommand::with_name("add")
                        .about("add the supplied remote to the list of known remotes")
                        .arg(
                            Arg::with_name("address")
                                .index(1)
                                .short("a")
                                .takes_value(true)
                                .required(true)
                                .help("the remote address that will be added"),
                        ),
                    SubCommand::with_name("remove")
                        .about("remote the supplied remote from the list of known remotes")
                        .arg(
                            Arg::with_name("address")
                                .index(1)
                                .short("a")
                                .required(true)
                                .takes_value(true)
                                .help("the remote address that will be removed"),
                        ),
                ]),
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

        ("remote", arg) => {
            match arg {
                None => {
                    println!("{}", matches.usage())
                }
                Some(c) => match c.subcommand() {
                    ("list", _) => commands::remote::remote_list()?,
                    ("add", Some(args)) => {
                        commands::remote::remote_add(args.value_of("address").unwrap())?
                    }
                    ("remove", Some(args)) => {
                        commands::remote::remote_remove(args.value_of("address").unwrap())?
                    }

                    _ => {
                        println!("{}", matches.usage())
                    }
                },
            };
        }

        _ => println!("{}", matches.usage()),
    };

    Ok(())
}

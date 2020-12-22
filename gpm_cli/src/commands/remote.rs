use anyhow::Context;
use gpm_core::config::Config;
use gpm_core::remote::Remote;

pub fn remote_list() -> anyhow::Result<()> {
    Config::write_default_config_if_not_exist()
        .context(format!("Writing default config if it does not exist"))?;

    let remotes = Remote::get_all_remotes()?;

    let remotes_string = remotes
        .iter()
        .map(|remote| format!(" - {}", remote.address()))
        .fold("remotes:".to_owned(), |acc, current_remote| {
            format!("{}\n{}", acc, current_remote)
        });

    println!("{}", remotes_string);

    Ok(())
}

pub fn remote_add(address: &str) -> anyhow::Result<()> {
    let new_remote = Remote::new(address);

    Config::write_default_config_if_not_exist()
        .context(format!("Writing default config if it does not exist"))?;

    new_remote.add()?;

    println!("added");

    Ok(())
}

pub fn remote_remove(address: &str) -> anyhow::Result<()> {
    let remote_to_remove = Remote::new(address);

    Config::write_default_config_if_not_exist()
        .context(format!("Writing default config if it does not exist"))?;

    remote_to_remove.remove()?;

    println!("removed");

    Ok(())
}

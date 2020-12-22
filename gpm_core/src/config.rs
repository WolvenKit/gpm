use crate::remote::Remote;
use anyhow::Context;
use serde::{Deserialize, Serialize};
use std::env;
use std::fs;
use std::path::{Path, PathBuf};

/// Represents the configuration used by the CLI.
#[derive(Deserialize, Serialize)]
pub struct Config {
    pub general: ConfigGeneral,
}

#[derive(Deserialize, Serialize)]
pub struct ConfigGeneral {
    pub remotes: Vec<Remote>,
}

/// Where the config file is stored depending on the OS:
/// - `%userprofile%/appdata/local/games-package-manager/config.toml` on Windows
/// - `${home}/.config/games-package-manager/config.toml` on Linux
///
/// _I don't know if this code should have been in `crate::constants` or not. I felt
/// it fit here as it's used only by the Config code. Feel free to move it if you
/// don't like the current location_
static CONFIG_RELATIVE_PATH: &str = "games-package-manager/config.toml";

fn get_config_path() -> Result<PathBuf, std::io::Error> {
    let error_if_not_found = Err(std::io::Error::new(
        std::io::ErrorKind::NotFound,
        "Could not find home directory",
    ));
    let absolute_path = dirs::config_dir().map(|p| p.join(CONFIG_RELATIVE_PATH));

    match absolute_path {
        None => error_if_not_found,
        Some(path) => Ok(path),
    }
}

impl Config {
    pub fn new(remotes: Vec<Remote>) -> Config {
        Config {
            general: ConfigGeneral { remotes: remotes },
        }
    }

    /// Loads the config from the disk and returns a `Config` with the data it read.
    /// it expect to find in a
    pub fn config_from_disk() -> anyhow::Result<Config> {
        let config_path = get_config_path().context(format!("Could not find home directory"))?;

        let config_content = fs::read_to_string(config_path).context(format!(
            "Could not read the config from the disk at [{:?}], perhaps it doesn't exist?",
            get_config_path()
        ))?;

        let config: Config = toml::from_str(&config_content).context(format!(
            "Could not parse the config at [{:?}], perhaps it doesn't follow the standard?",
            get_config_path()
        ))?;

        Ok(config)
    }

    /// Write the serialized version of `Self` to the disk.
    pub fn write_to_disk(&self) -> anyhow::Result<()> {
        let config_path = get_config_path().context(format!("Could not find home directory"))?;

        let new_config_content = toml::to_string_pretty(self)
      .context(format!("Could not save the current config to the disk as it doesn't follow the standard config"))?;

        let config_directory = config_path.parent().context(format!(
            "Could not find the parent in which the config should be written"
        ))?;

        fs::create_dir_all(config_directory).context(format!(
            "Could not create the directories for the config at [{:?}]",
            get_config_path()
        ))?;

        fs::write(config_path, new_config_content).context(format!(
            "Could not write the new config to the disk in [{:?}]",
            get_config_path()
        ))?;

        Ok(())
    }

    /// This function checks if the config can be found in the user's disk and at
    /// the pre-configured path (depends on the OS, read the Config top comment).
    /// And if it doesn't exist, write the default config to the disk and return `true`
    /// to notifiy it wrote to the disk. Otherwise return `false` (when nothing was done)
    pub fn write_default_config_if_not_exist() -> anyhow::Result<bool> {
        let config_path = get_config_path().context(format!("Could not find home directory"))?;

        if config_path.exists() {
            return Ok(false);
        }

        let default_config = Self::default_config();

        default_config
            .write_to_disk()
            .context(format!("Could not write the default config to the disk"))?;

        Ok(true)
    }

    /// Return the default config that would be written on the user's disk if it doesn't
    /// exist already.
    ///
    /// **For anyone who wants to update the default config, this is here.**
    fn default_config() -> Config {
        Config {
            general: ConfigGeneral {
                remotes: vec![Remote::new("https://todo.todo:todo/api/todo")],
            },
        }
    }
}

use crate::config::Config;
use anyhow::Context;
use serde::{Deserialize, Serialize};

/// Represents any remote server with a compatible Games Package Manager API.
/// The GPM core library comes with a default, hardcoded remote. But the user
/// can freely edit the list of known remotes to add or remove remotes.
///
/// In the config, the remotes are placed under the `[remotes]` section
#[derive(Serialize, Deserialize)]
pub struct Remote(pub String);

impl Remote {
    pub fn new(address: &str) -> Remote {
        Remote(address.to_owned())
    }

    pub fn address(&self) -> &str {
        &self.0
    }

    /// add the `&self` to the list of known remotes.
    ///
    /// If the remote is already in the known remotes list, the function do nothing.
    /// If it doesn't exist, it will add the new remote in the file.
    pub fn add(self) -> anyhow::Result<()> {
        let mut config: Config = Config::config_from_disk().context(format!(
            "Could not retrieve the current config from the disk"
        ))?;

        // as mentioned above, we do not add the remote again if it already exists
        if !config.general.remotes.contains(&self) {
            config.general.remotes.push(self);
        }

        config
            .write_to_disk()
            .context(format!("Could not add the new remote to the config"))?;

        Ok(())
    }

    /// remove the `&self` from the list of known remotes.
    /// It returns `true` if the remote was in the list of known remotes, and
    /// `false` if it wasn't.
    pub fn remove(&self) -> anyhow::Result<bool> {
        let mut config: Config = Config::config_from_disk().context(format!(
            "Could not retrieve the current config from the disk"
        ))?;

        let some_remote_index = config.general.remotes.iter().position(|r| r == self);

        let was_found = some_remote_index.is_some();

        if let Some(index) = some_remote_index {
            config.general.remotes.swap_remove(index);
        }

        config
            .write_to_disk()
            .context(format!("Could not add the new remote to the config"))?;

        Ok(was_found)
    }

    /// get the list of known remotes from the config file on the disk
    pub fn get_all_remotes() -> anyhow::Result<Vec<Remote>> {
        let config: Config = Config::config_from_disk().context(format!(
            "Could not retrieve the current config from the disk"
        ))?;

        Ok(config.general.remotes)
    }
}

impl PartialEq for Remote {
    fn eq(&self, other: &Self) -> bool {
        self.0 == other.0
    }
}

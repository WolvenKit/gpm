use std::fs::File;
use std::io;
use std::io::Read;
use std::path::{Path, PathBuf};

use crate::package::{Package, PackageInformation, PackageInformationExtraData};
use crate::TOML_CONFIG_PATH;

use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct StoredPackageInformation {
    creator: String,
    identifier: String,
    version: String,
    display_name: String,
    description: String,
    license: String,

    #[serde(default)]
    website_url: Option<String>,
    #[serde(default)]
    dependencies: Vec<String>,
    #[serde(default)]
    tags: Vec<String>,
    #[serde(default)]
    install_strategies: Vec<String>,
    #[serde(default)]
    extra_data: Vec<(String, String)>,
}

impl Into<PackageInformation> for StoredPackageInformation {
    fn into(mut self) -> PackageInformation {
        let extra_data = self
            .extra_data
            .drain(..)
            .map(|(key, value)| PackageInformationExtraData { key, value })
            .collect();

        PackageInformation {
            creator: self.creator,
            identifier: self.identifier,
            version: self.version,
            display_name: self.display_name,
            description: self.description,
            license: self.license,
            website_url: self.website_url,
            dependencies: self.dependencies,
            tags: self.tags,
            install_strategies: self.install_strategies,
            extra_data,
        }
    }
}

impl From<&PackageInformation> for StoredPackageInformation {
    fn from(package: &PackageInformation) -> Self {
        let extra_data = package
            .extra_data
            .iter()
            .map(|data| (data.key.clone(), data.value.clone()))
            .collect();
        Self {
            creator: package.creator.clone(),
            identifier: package.identifier.clone(),
            version: package.version.clone(),
            display_name: package.display_name.clone(),
            description: package.description.clone(),
            license: package.license.clone(),
            website_url: package.website_url.clone(),
            dependencies: package.dependencies.clone(),
            tags: package.tags.clone(),
            install_strategies: package.install_strategies.clone(),
            extra_data,
        }
    }
}

#[derive(thiserror::Error, Debug)]
pub enum LoadPackageFromProjectError {
    #[error("io error with the file {0}")]
    FileIOError(PathBuf, io::Error),
    #[error("error while parsing the toml file {0}")]
    TomlDecodeError(PathBuf, #[source] toml::de::Error),
}

pub fn load_package_from_project(
    project_path: &Path,
) -> Result<Package, LoadPackageFromProjectError> {
    let config_path = project_path.join(TOML_CONFIG_PATH);
    let mut config_file = File::open(&config_path)
        .map_err(|err| LoadPackageFromProjectError::FileIOError(config_path.to_path_buf(), err))?;
    let mut config_content = Vec::new();
    config_file
        .read_to_end(&mut config_content)
        .map_err(|err| LoadPackageFromProjectError::FileIOError(config_path.to_path_buf(), err))?;
    let stored_package_information = toml::from_slice::<StoredPackageInformation>(&config_content)
        .map_err(|err| {
            LoadPackageFromProjectError::TomlDecodeError(config_path.to_path_buf(), err)
        })?;
    Ok(Package {
        information: stored_package_information.into(),
    })
}

pub fn get_project_config_json(
    package_information: &PackageInformation,
) -> Result<Vec<u8>, serde_json::Error> {
    let stored_package_information = StoredPackageInformation::from(package_information);
    serde_json::to_vec_pretty(&stored_package_information)
}

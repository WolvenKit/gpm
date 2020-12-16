// Code below is used to represent a package that can be downloaded, installed,
// or created by the user and published to the store.

use std::result::Result;

pub struct PackageInformationExtraData {
  key: String,
  value: String
}

impl PackageInformationExtraData {
  pub fn new(key: &str, value: &str) -> PackageInformationExtraData {
    PackageInformationExtraData {
      key: key.to_owned(),
      value: key.to_owned()
    }
  }
}

pub struct PackageInformation {
  // required values
  creator: String,
  identifier: String,
  version: String,
  display_name: String,
  description: String,
  license: String,

  // optional values
  website_url: Option<String>,
  dependencies: Vec<u8>,
  tags: Vec<String>,
  install_strategies: Vec<String>,
  extra_data: Vec<PackageInformationExtraData>
}

impl PackageInformation {
  pub fn new (creator: &str, identifier: &str, version: &str, display_name: &str, description: &str, license: &str) -> PackageInformation {
    PackageInformation {
      creator: creator.to_owned(),
      identifier: identifier.to_owned(),
      version: version.to_owned(),
      display_name: display_name.to_owned(),
      description: description.to_owned(),
      license: license.to_owned(),

      website_url: None,
      dependencies: Vec::new(),
      tags: Vec::new(),
      install_strategies: Vec::new(),
      extra_data: Vec::new()
    }
  }
}

pub struct Package {
  information: PackageInformation
}

impl Package {
  pub fn new(package_information: PackageInformation) -> Package {
    Package {
      information: package_information
    }
  }

  pub fn init_project_directory() -> Result<(), &'static str> {
    Ok(())
  }

  pub fn publish() -> Result<(), &'static str> {
    Ok(())
  }
}
/// Code below is used to represent a package that can be downloaded, installed,
/// or created by the user and published to the store.

pub struct PackageInformationExtraData {
    pub key: String,
    pub value: String,
}

impl PackageInformationExtraData {
    pub fn new(key: &str, value: &str) -> PackageInformationExtraData {
        PackageInformationExtraData {
            key: key.to_owned(),
            value: value.to_owned(),
        }
    }
}

pub struct RequiredPublishInformation {
    pub creator: String,
    pub identifier: String,
    pub version: String,
    pub display_name: String,
    pub description: String,
    pub license: String,
}

pub struct PackageInformation {
    // required values for publied package
    //
    // remember to change get_required_publish_information when adding/renaming/deleting one of
    // those
    pub creator: Option<String>,
    pub identifier: Option<String>,
    pub version: Option<String>,
    pub display_name: Option<String>,
    pub description: Option<String>,
    pub license: Option<String>,

    // optional values
    pub website_url: Option<String>,
    pub dependencies: Vec<String>,
    pub tags: Vec<String>,
    pub install_strategies: Vec<String>,
    pub extra_data: Vec<PackageInformationExtraData>,
}

impl PackageInformation {
    pub fn new(
        creator: &str,
        identifier: &str,
        version: &str,
        display_name: &str,
        description: &str,
        license: &str,
    ) -> PackageInformation {
        PackageInformation {
            creator: Some(creator.to_owned()),
            identifier: Some(identifier.to_owned()),
            version: Some(version.to_owned()),
            display_name: Some(display_name.to_owned()),
            description: Some(description.to_owned()),
            license: Some(license.to_owned()),

            website_url: None,
            dependencies: Vec::new(),
            tags: Vec::new(),
            install_strategies: Vec::new(),
            extra_data: Vec::new(),
        }
    }

    /// return all the information required for published package
    /// ([`RequiredPublishInformation`]), otherwise, return an error containing a
    /// list of the missing field
    pub fn required_publish_information(&self) -> Option<RequiredPublishInformation> {
        Some(RequiredPublishInformation {
            creator: self.creator.clone()?,
            identifier: self.identifier.clone()?,
            version: self.version.clone()?,
            display_name: self.display_name.clone()?,
            description: self.description.clone()?,
            license: self.license.clone()?,
        })
    }

    pub fn missing_publish_field(&self) -> Vec<&'static str> {
        let mut missing_publish_field = Vec::new();
        if self.creator.is_none() {
            missing_publish_field.push("creator");
        };
        if self.identifier.is_none() {
            missing_publish_field.push("identifier");
        };
        if self.version.is_none() {
            missing_publish_field.push("version");
        };
        if self.display_name.is_none() {
            missing_publish_field.push("display_name");
        };
        if self.description.is_none() {
            missing_publish_field.push("description");
        };
        if self.license.is_none() {
            missing_publish_field.push("license");
        };
        missing_publish_field
    }
}

pub struct Package {
    pub information: PackageInformation,
}

impl Package {
    pub fn new(package_information: PackageInformation) -> Package {
        Package {
            information: package_information,
        }
    }

    pub fn init_project_directory() -> Result<(), anyhow::Error> {
        Ok(())
    }

    pub fn publish() -> Result<(), anyhow::Error> {
        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use crate::package::PackageInformation;

    #[test]
    fn test_publish_field() {
        let mut package_information = PackageInformation::new(
            "a creator",
            "an_identifier",
            "1.2.3",
            "A Display Name",
            "the description of this is a description",
            "some version of AGPL",
        );

        assert!(package_information.missing_publish_field().is_empty());
        let required_publish_information =
            package_information.required_publish_information().unwrap();
        assert_eq!(&required_publish_information.version, "1.2.3");

        package_information.version = None;
        package_information.display_name = None;
        let missing_publish_field = package_information.missing_publish_field();
        assert!(missing_publish_field.contains(&"version"));
        assert!(missing_publish_field.contains(&"display_name"));
        assert_eq!(missing_publish_field.len(), 2);
        assert!(package_information.required_publish_information().is_none());
    }
}

use anyhow::Context;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::fs::File;
use std::io::{Read, Write};
use std::path::{Path, PathBuf};

/// Allow to look for a specific version of a dependancies
#[derive(Serialize, Deserialize, Clone, PartialEq, Debug)]
#[serde(tag = "type")]
pub enum LockSource {
    /// Use the mod with the given ID and the matching version.
    ///
    /// Even if the parent contain the id, we still store it there, as it may be usefull that id
    /// differ from it's parent.
    IdVersion { identifier: String, version: String },
    /// Use a specific path on the local filesystem. If relative, it'll be based around the
    /// profile folder.
    Path { path: PathBuf },
}

/// contain a fixed set of mod dependency, with each dependency having a specific version. Mod are
/// identified by their id. They are unique.
#[derive(Serialize, Deserialize, Default, Debug)]
pub struct LockFile {
    pub dependencies: HashMap<String, LockSource>,
}

impl LockFile {
    /// create a new [`LockFile`] with no dependancies.
    pub fn new() -> Self {
        Self::default()
    }

    /// return the [`LockSource`] corresponding to a given package identifier if it is present in
    /// this [`LockFile`], None otherwise.
    pub fn dependency_source(&self, identifier: &str) -> Option<LockSource> {
        self.dependencies.get(identifier).cloned()
    }

    /// define the [`LockSource`] for a given package identifier, overwriting the current one.
    ///
    /// return the previous [`LockSource`] if overwriting it
    pub fn set_dependency_source(
        &mut self,
        identifier: String,
        source: LockSource,
    ) -> Option<LockSource> {
        self.dependencies.insert(identifier, source)
    }

    /// remove the mod with the given id. Return the previous entry if it exist, None otherwise.
    pub fn remove_dependency_source(&mut self, identifier: &str) -> Option<LockSource> {
        self.dependencies.remove(identifier)
    }

    /// iterate over all the locked dependancies
    pub fn iter_dependency_source(
        &self,
    ) -> std::collections::hash_map::Iter<'_, String, LockSource> {
        self.dependencies.iter()
    }

    /// load the lock file from input TOML stream
    pub fn load_reader<T: Read>(input: &mut T) -> Result<Self, anyhow::Error> {
        let mut buffer = Vec::new();
        input
            .read_to_end(&mut buffer)
            .context("can't load the input file in memory")?;
        toml::from_slice(&buffer).context("can't parse the TOML lock file")
    }

    /// write this [`LockFile`] to the output stream (TOML)
    pub fn write_writer<T: Write>(&self, output: &mut T) -> Result<(), anyhow::Error> {
        let buffer = toml::to_vec(&self).context("can't encode the TOML lock file")?;
        output
            .write_all(&buffer)
            .context("can't push the encoded file")
    }

    /// load a [`LockFile`] from the given file
    pub fn load_file(path: &Path) -> anyhow::Result<Self> {
        let mut file =
            File::open(path).with_context(|| format!("can't open the lock file at {:?}", &path))?;
        Ok(Self::load_reader(&mut file)
            .with_context(|| format!("can't load the TOML lock file at {:?}", &path))?)
    }

    /// write this [`LockFile`] to the given file
    pub fn write_file(&self, path: &Path) -> anyhow::Result<()> {
        let mut file = File::create(&path)
            .with_context(|| format!("can't create the lock file at {:?}", &path))?;
        Ok(Self::write_writer(&self, &mut file)
            .with_context(|| format!("can't write the TOML lock file at {:?}", &path))?)
    }
}

#[cfg(test)]
mod tests {
    use crate::lockfile::{LockFile, LockSource};

    #[test]
    fn test_lock_file() {
        let package1_source = LockSource::IdVersion {
            identifier: "package1_bis".into(),
            version: "1.0.0".into(),
        };
        let mut lock_file = LockFile::new();
        assert!(lock_file.dependency_source("package1").is_none());
        assert!(lock_file
            .set_dependency_source("package1".into(), package1_source.clone())
            .is_none());
        assert!(lock_file
            .iter_dependency_source()
            .map(|(k, _)| k)
            .collect::<Vec<_>>()
            .contains(&&"package1".to_string()));
        assert_eq!(
            lock_file.dependency_source("package1"),
            Some(package1_source.clone())
        );
        assert_eq!(
            lock_file.remove_dependency_source("package1"),
            Some(package1_source)
        );
        assert!(lock_file.remove_dependency_source("package1").is_none());
    }
}

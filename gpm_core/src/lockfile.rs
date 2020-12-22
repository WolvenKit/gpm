use anyhow::Context;
use serde::{Deserialize, Serialize};
use std::collections::HashSet;
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
/// identified by their id. They are unique. They are also ordered.
#[derive(Serialize, Deserialize, Default, Debug)]
pub struct LockFile {
    dependencies: Vec<(String, LockSource)>,
}

impl LockFile {
    /// create a new [`LockFile`] with no dependancies.
    pub fn new() -> Self {
        Self::default()
    }

    /// remove duplicated dependency.
    fn dedup(&mut self) {
        let mut already_present: HashSet<String> = HashSet::new();
        let mut to_remove: Vec<usize> = Vec::new();
        for (count, (source_id, _)) in self.dependencies.iter().enumerate() {
            if already_present.contains(source_id) {
                to_remove.push(count);
            } else {
                already_present.insert(source_id.clone());
            }
        }

        for remove_id in to_remove.iter().rev() {
            self.dependencies.remove(*remove_id);
        }
    }

    /// return the [`LockSource`] corresponding to a given package identifier if it is present in
    /// this [`LockFile`], None otherwise.
    pub fn dependency_source(&self, identifier: &str) -> Option<LockSource> {
        Some(
            self.dependencies[self.dependency_source_position(identifier)?]
                .1
                .clone(),
        )
    }

    /// return the numerical position of the given [`LockSource`], or None if not present.
    pub fn dependency_source_position(&self, identifier: &str) -> Option<usize> {
        for (count, (actual_id, _)) in self.dependencies.iter().enumerate() {
            if actual_id == identifier {
                return Some(count);
            }
        }
        None
    }

    /// define the [`LockSource`] for a given package identifier, overwriting the current one.
    ///
    /// If the mod identifier is already present, only change the source. If it isn't present,
    /// add it at the end of the dependancies. This allow to check, that, if dependencies are added
    /// before the dependent, it dependancies are loaded after the dependent whatever happen.
    /// A potential issue is:
    /// 1. for circular dependency (won't detect them, but shouldn't happen in a first time)
    /// 2. A dependency is added to one package. We would need to re-add every package for this,
    /// respecting the order. One possibility would be to create a new [`LockFile`] respecting the
    /// add order, but keeping the [`LockSource`] of the original lock file if existing.
    pub fn add_dependency(&mut self, identifier: String, source: LockSource) {
        if let Some(position) = self.dependency_source_position(&identifier) {
            self.dependencies[position] = (identifier, source);
        } else {
            self.dependencies.push((identifier, source));
        }
    }

    /// remove the mod with the given id. Return the previous entry if it exist, None otherwise.
    pub fn remove_dependency_source(&mut self, identifier: &str) -> Option<LockSource> {
        let mut to_remove = None;
        for (count, (actual_id, _)) in self.dependencies.iter().enumerate() {
            if actual_id == identifier {
                to_remove = Some(count);
            };
        }
        if let Some(remove_id) = to_remove {
            Some(self.dependencies.remove(remove_id).1)
        } else {
            None
        }
    }

    /// iterate over all the locked dependancies, in order
    pub fn iter_dependencies_source(
        &self,
    ) -> std::slice::Iter<'_, (std::string::String, LockSource)> {
        self.dependencies.iter()
    }

    /// load the lock file from input TOML stream (entrys are deduplicated if necessary)
    pub fn load_reader<T: Read>(input: &mut T) -> Result<Self, anyhow::Error> {
        let mut buffer = Vec::new();
        input
            .read_to_end(&mut buffer)
            .context("can't load the input file in memory")?;
        let mut result =
            toml::from_slice::<Self>(&buffer).context("can't parse the TOML lock file")?;
        result.dedup();
        Ok(result)
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
        lock_file.add_dependency("package1".into(), package1_source.clone());
        assert!(lock_file
            .iter_dependencies_source()
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

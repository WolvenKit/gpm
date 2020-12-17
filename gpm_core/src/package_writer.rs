use std::fs::File;
use std::io;
use std::io::{ErrorKind, Read, Seek, Write};
use std::path::{Path, PathBuf};

use crate::store_project::{
    get_project_config_json, load_package_from_project, LoadPackageFromProjectError,
};
use crate::{IGNORE_PATH, JSON_CONFIG_PATH};

use walkdir::WalkDir;
use zip::{
    write::{FileOptions, ZipWriter},
    CompressionMethod,
};

use ignore;
use ignore::gitignore::GitignoreBuilder;

#[derive(thiserror::Error, Debug)]
pub enum CreatePackageError {
    #[error("io error while reading {0}")]
    FileIOError(PathBuf, #[source] io::Error),
    #[error("error while loading the package in {0}")]
    LoadPackageError(PathBuf, #[source] LoadPackageFromProjectError),
    #[error("error while running walkdir")]
    WalkDirError(#[from] walkdir::Error),
    #[error("error stripping a the path {0} with {1}")]
    StripPrefixError(PathBuf, PathBuf, #[source] std::path::StripPrefixError),
    #[error("error while handling the zip file")]
    ZipError(#[from] zip::result::ZipError),
    #[error("error while writing to the zip file")]
    ZipWriteError(#[source] io::Error),
    #[error("can't generate the output json configuration file, but can parse the toml one. Probably internal error")]
    EncodeJsonError(#[source] serde_json::error::Error),
    #[error("error while handling the ignore file")] // this one should include path
    IgnoreFileError(#[from] ignore::Error),
}

pub fn create_package<D: Write + Seek>(
    input_dir: &Path,
    destination: &mut D,
) -> Result<(), CreatePackageError> {
    // load the package
    let package = load_package_from_project(&input_dir)
        .map_err(|err| CreatePackageError::LoadPackageError(input_dir.to_path_buf(), err))?;

    //load the ignore file
    let ignore_path = input_dir.join(IGNORE_PATH);

    let mut builder = GitignoreBuilder::new(&input_dir);
    let ignore = match builder.add(&ignore_path) {
        None => Some(builder.build()?),
        Some(err) => match err.io_error() {
            Some(io_err) => match io_err.kind() {
                ErrorKind::NotFound => {
                    println!("{:?} not found, ignoring it.", ignore_path);
                    None
                }
                _ => return Err(CreatePackageError::from(err)),
            },
            None => return Err(CreatePackageError::from(err)),
        },
    };

    // write the zip file
    let mut zip = ZipWriter::new(destination);

    let walkdir = WalkDir::new(&input_dir).follow_links(true);

    let zip_options = FileOptions::default().compression_method(CompressionMethod::Deflated);

    let mut embedded_content = Vec::new();
    for entry in walkdir {
        let entry = entry?;

        let content_abs_path = entry.path();
        let content_rel_path = content_abs_path.strip_prefix(&input_dir).map_err(|err| {
            CreatePackageError::StripPrefixError(
                content_abs_path.to_path_buf(),
                input_dir.to_path_buf(),
                err,
            )
        })?;

        if content_rel_path == PathBuf::from(JSON_CONFIG_PATH) {
            continue;
        };

        let is_file = entry.file_type().is_file();

        if let Some(ignore) = &ignore {
            if ignore
                .matched_path_or_any_parents(&content_rel_path, !is_file)
                .is_ignore()
            {
                println!("ignored {:?}", content_rel_path);
                continue;
            }
        };

        if is_file {
            println!("adding the file {:?} to the archive", content_rel_path);
            zip.start_file(content_rel_path.to_string_lossy(), zip_options)?;
            let mut embedded_file = File::open(content_abs_path).map_err(|err| {
                CreatePackageError::FileIOError(content_abs_path.to_path_buf(), err)
            })?;
            embedded_file
                .read_to_end(&mut embedded_content)
                .map_err(|err| {
                    CreatePackageError::FileIOError(content_abs_path.to_path_buf(), err)
                })?;
            zip.write_all(&embedded_content)
                .map_err(CreatePackageError::ZipWriteError)?;
            embedded_content.clear();
        } else {
            println!("adding the directory {:?} to the archive", content_rel_path);
            zip.add_directory(content_rel_path.to_string_lossy(), zip_options)?;
        }
    }
    let config_json: Vec<u8> = get_project_config_json(&package.information)
        .map_err(CreatePackageError::EncodeJsonError)?;
    zip.start_file(JSON_CONFIG_PATH, zip_options)?;
    zip.write_all(&config_json)
        .map_err(CreatePackageError::ZipWriteError)?;
    Ok(())
}

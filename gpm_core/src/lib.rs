pub mod config;
pub mod display;
pub mod package;
pub mod package_writer;
pub mod remote;
pub mod store_project;

pub mod constants {
    pub const TOML_CONFIG_PATH: &str = "config.toml";
    pub const JSON_CONFIG_PATH: &str = "config.json";
    pub const IGNORE_PATH: &str = ".modignore";
}

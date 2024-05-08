//! This module is responsible for compiling the protobuf files before running the application.
//!
//! It defines a `main` function which is the entry point of the build script. The `main` function calls the `compile_protos` function from the `tonic_build` crate to compile the `notification.proto` file.
//!
//! The `compile_protos` function takes the path to the `notification.proto` file as an argument and compiles it into Rust code. The generated Rust code is written to the `OUT_DIR` directory.
//!
//! If the protobuf files are compiled successfully, the `main` function returns `Ok(())`. If an error occurs, the `main` function returns `Err(e)`, where `e` is the error.
//!
//! # Example
//!
//! To compile the protobuf files, you can run the following command in the terminal:
//!
//! ```shell
//! cargo build
//! ```
fn main() -> Result<(), Box<dyn std::error::Error>> {
    tonic_build::compile_protos("./proto/notification.proto")?;
    Ok(())
}

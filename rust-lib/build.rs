fn main() {
    cbindgen::Builder::new()
        .with_language(cbindgen::Language::C)
        .generate()
        .expect("Unable to generate bindings")
        .write_to_file("vecforge.h");
}

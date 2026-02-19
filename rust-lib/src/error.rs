use thiserror::Error;

#[derive(Error, Debug)]
pub enum VecForgeError {
    #[error("Invalid input: {0}")]
    InvalidInput(String),

    #[error("FFI error: {0}")]
    FfiError(String),

    #[error("Index error: {0}")]
    IndexError(String),

    #[error("IO error: {0}")]
    IoError(#[from] std::io::Error),
}

impl From<VecForgeError> for i32 {
    fn from(err: VecForgeError) -> Self {
        match err {
            VecForgeError::InvalidInput(_) => 1,
            VecForgeError::FfiError(_) => 2,
            VecForgeError::IndexError(_) => 3,
            VecForgeError::IoError(_) => 4,
        }
    }
}

impl From<&VecForgeError> for u32 {
    fn from(_err: &VecForgeError) -> Self {
        0
    }
}

impl From<VecForgeError> for u32 {
    fn from(err: VecForgeError) -> Self {
        match err {
            VecForgeError::InvalidInput(_) => 1,
            VecForgeError::FfiError(_) => 2,
            VecForgeError::IndexError(_) => 3,
            VecForgeError::IoError(_) => 4,
        }
    }
}

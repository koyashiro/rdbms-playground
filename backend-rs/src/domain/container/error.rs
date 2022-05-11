use std::{error::Error, fmt::Display};

#[derive(Clone, Debug, Eq, PartialEq)]
pub struct MySQLVersionParseError;

impl Display for MySQLVersionParseError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "invalid MySQL version")
    }
}

impl Error for MySQLVersionParseError {}

#[derive(Clone, Debug, Eq, PartialEq)]
pub struct MariaDBVersionParseError;

impl Display for MariaDBVersionParseError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "invalid MariaDB version")
    }
}

impl Error for MariaDBVersionParseError {}

#[derive(Clone, Debug, Eq, PartialEq)]
pub struct PostgreSQLVersionParseError;

impl Display for PostgreSQLVersionParseError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "invalid PostgreSQL version")
    }
}

impl Error for PostgreSQLVersionParseError {}

#[derive(Clone, Debug, Eq, PartialEq)]
pub enum ContainerImageParseErrorKind {
    Name,
    MySQLVersion(MySQLVersionParseError),
    MariaDBVersion(MariaDBVersionParseError),
    PostgreSQLVersion(PostgreSQLVersionParseError),
}

#[derive(Clone, Debug, Eq, PartialEq)]
pub struct ContainerImageParseError {
    kind: ContainerImageParseErrorKind,
}

impl ContainerImageParseError {
    pub fn new(kind: ContainerImageParseErrorKind) -> Self {
        Self { kind }
    }
}

impl Display for ContainerImageParseError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match &self.kind {
            ContainerImageParseErrorKind::Name => write!(f, "invalid image name"),
            ContainerImageParseErrorKind::MySQLVersion(e) => e.fmt(f),
            ContainerImageParseErrorKind::MariaDBVersion(e) => e.fmt(f),
            ContainerImageParseErrorKind::PostgreSQLVersion(e) => e.fmt(f),
        }
    }
}

impl Error for ContainerImageParseError {}

impl From<MySQLVersionParseError> for ContainerImageParseError {
    fn from(v: MySQLVersionParseError) -> Self {
        Self {
            kind: ContainerImageParseErrorKind::MySQLVersion(v),
        }
    }
}

impl From<MariaDBVersionParseError> for ContainerImageParseError {
    fn from(v: MariaDBVersionParseError) -> Self {
        Self {
            kind: ContainerImageParseErrorKind::MariaDBVersion(v),
        }
    }
}

impl From<PostgreSQLVersionParseError> for ContainerImageParseError {
    fn from(v: PostgreSQLVersionParseError) -> Self {
        Self {
            kind: ContainerImageParseErrorKind::PostgreSQLVersion(v),
        }
    }
}

use std::{fmt::Display, str::FromStr};

use crate::domain::container::error::{
    ContainerImageParseError, ContainerImageParseErrorKind, MariaDBVersionParseError,
    MySQLVersionParseError, PostgreSQLVersionParseError,
};

#[derive(Clone, Debug, Eq, PartialEq)]
pub enum ContainerImage {
    MySQL(MySQLVersion),
    MariaDB(MariaDBVersion),
    PostgreSQL(PostgreSQLVersion),
}

impl Display for ContainerImage {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::MySQL(v) => write!(f, "mysql:{v}"),
            Self::MariaDB(v) => write!(f, "mariadb:{v}"),
            Self::PostgreSQL(v) => write!(f, "postgres:{v}"),
        }
    }
}

impl FromStr for ContainerImage {
    type Err = ContainerImageParseError;

    fn from_str(v: &str) -> Result<Self, Self::Err> {
        let s = v.split(':').collect::<Vec<&str>>();
        if s.len() != 2 {
            return Err(ContainerImageParseError::new(
                ContainerImageParseErrorKind::Name,
            ));
        }
        let (image, tag) = (s[0], s[1]);
        let image = match image {
            "mysql" => ContainerImage::MySQL(tag.parse()?),
            "mariadb" => ContainerImage::MariaDB(tag.parse()?),
            "postgres" => ContainerImage::PostgreSQL(tag.parse()?),
            _ => {
                return Err(ContainerImageParseError::new(
                    ContainerImageParseErrorKind::Name,
                ))
            }
        };
        Ok(image)
    }
}

#[derive(Clone, Debug, Eq, PartialEq)]
pub enum MySQLVersion {
    V5,
    V8,
}

impl Display for MySQLVersion {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            MySQLVersion::V5 => write!(f, "5"),
            MySQLVersion::V8 => write!(f, "8"),
        }
    }
}

impl FromStr for MySQLVersion {
    type Err = MySQLVersionParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let version = match s {
            "5" => Self::V5,
            "8" => Self::V8,
            _ => return Err(MySQLVersionParseError),
        };
        Ok(version)
    }
}

#[derive(Clone, Debug, Eq, PartialEq)]
pub enum MariaDBVersion {
    V10_2,
    V10_3,
    V10_4,
    V10_5,
    V10_6,
    V10_7,
    V10_8,
}

impl Display for MariaDBVersion {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            MariaDBVersion::V10_2 => write!(f, "10.2"),
            MariaDBVersion::V10_3 => write!(f, "10.3"),
            MariaDBVersion::V10_4 => write!(f, "10.4"),
            MariaDBVersion::V10_5 => write!(f, "10.5"),
            MariaDBVersion::V10_6 => write!(f, "10.6"),
            MariaDBVersion::V10_7 => write!(f, "10.7"),
            MariaDBVersion::V10_8 => write!(f, "10.8"),
        }
    }
}

impl FromStr for MariaDBVersion {
    type Err = MariaDBVersionParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let version = match s {
            "10.2" => Self::V10_2,
            "10.3" => Self::V10_3,
            "10.4" => Self::V10_4,
            "10.5" => Self::V10_5,
            "10.6" => Self::V10_6,
            "10.7" => Self::V10_7,
            "10.8" => Self::V10_8,
            _ => return Err(MariaDBVersionParseError),
        };
        Ok(version)
    }
}

#[derive(Clone, Debug, Eq, PartialEq)]
pub enum PostgreSQLVersion {
    V10,
    V11,
    V12,
    V13,
    V14,
}

impl Display for PostgreSQLVersion {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            PostgreSQLVersion::V10 => write!(f, "10"),
            PostgreSQLVersion::V11 => write!(f, "11"),
            PostgreSQLVersion::V12 => write!(f, "12"),
            PostgreSQLVersion::V13 => write!(f, "13"),
            PostgreSQLVersion::V14 => write!(f, "14"),
        }
    }
}

impl FromStr for PostgreSQLVersion {
    type Err = PostgreSQLVersionParseError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let version = match s {
            "10" => Self::V10,
            "11" => Self::V11,
            "12" => Self::V12,
            "13" => Self::V13,
            "14" => Self::V14,
            _ => return Err(PostgreSQLVersionParseError),
        };
        Ok(version)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn container_image_display() {
        assert_eq!(
            format!("{}", ContainerImage::MySQL(MySQLVersion::V8)),
            "mysql:8".to_owned()
        );
        assert_eq!(
            format!("{}", ContainerImage::MariaDB(MariaDBVersion::V10_8)),
            "mariadb:10.8".to_owned()
        );
        assert_eq!(
            format!("{}", ContainerImage::PostgreSQL(PostgreSQLVersion::V14)),
            "postgres:14".to_owned()
        );
    }

    #[test]
    fn container_image_from_str() {
        assert_eq!(
            ContainerImage::from_str("mysql:8").unwrap(),
            ContainerImage::MySQL(MySQLVersion::V8)
        );
        assert_eq!(
            ContainerImage::from_str("mariadb:10.8").unwrap(),
            ContainerImage::MariaDB(MariaDBVersion::V10_8)
        );
        assert_eq!(
            ContainerImage::from_str("postgres:14").unwrap(),
            ContainerImage::PostgreSQL(PostgreSQLVersion::V14)
        );
    }

    #[test]
    fn mysql_version_display() {
        assert_eq!(format!("{}", MySQLVersion::V5), "5".to_owned());
        assert_eq!(format!("{}", MySQLVersion::V8), "8".to_owned());
    }

    #[test]
    fn mysql_version_from_str() {
        assert_eq!(MySQLVersion::from_str("5").unwrap(), MySQLVersion::V5);
        assert_eq!(MySQLVersion::from_str("8").unwrap(), MySQLVersion::V8);
    }

    #[test]
    fn mariadb_version_display() {
        assert_eq!(format!("{}", MariaDBVersion::V10_2), "10.2".to_owned());
        assert_eq!(format!("{}", MariaDBVersion::V10_3), "10.3".to_owned());
        assert_eq!(format!("{}", MariaDBVersion::V10_4), "10.4".to_owned());
        assert_eq!(format!("{}", MariaDBVersion::V10_5), "10.5".to_owned());
        assert_eq!(format!("{}", MariaDBVersion::V10_6), "10.6".to_owned());
        assert_eq!(format!("{}", MariaDBVersion::V10_7), "10.7".to_owned());
        assert_eq!(format!("{}", MariaDBVersion::V10_8), "10.8".to_owned());
    }

    #[test]
    fn mariadb_version_from_str() {
        assert_eq!(
            MariaDBVersion::from_str("10.2").unwrap(),
            MariaDBVersion::V10_2
        );
        assert_eq!(
            MariaDBVersion::from_str("10.3").unwrap(),
            MariaDBVersion::V10_3
        );
        assert_eq!(
            MariaDBVersion::from_str("10.4").unwrap(),
            MariaDBVersion::V10_4
        );
        assert_eq!(
            MariaDBVersion::from_str("10.5").unwrap(),
            MariaDBVersion::V10_5
        );
        assert_eq!(
            MariaDBVersion::from_str("10.6").unwrap(),
            MariaDBVersion::V10_6
        );
        assert_eq!(
            MariaDBVersion::from_str("10.7").unwrap(),
            MariaDBVersion::V10_7
        );
        assert_eq!(
            MariaDBVersion::from_str("10.8").unwrap(),
            MariaDBVersion::V10_8
        );
    }

    #[test]
    fn postgresql_version_display() {
        assert_eq!(format!("{}", PostgreSQLVersion::V10), "10".to_owned());
        assert_eq!(format!("{}", PostgreSQLVersion::V11), "11".to_owned());
        assert_eq!(format!("{}", PostgreSQLVersion::V12), "12".to_owned());
        assert_eq!(format!("{}", PostgreSQLVersion::V13), "13".to_owned());
        assert_eq!(format!("{}", PostgreSQLVersion::V14), "14".to_owned());
    }

    #[test]
    fn postgresql_version_from_str() {
        assert_eq!(
            PostgreSQLVersion::from_str("10").unwrap(),
            PostgreSQLVersion::V10
        );
        assert_eq!(
            PostgreSQLVersion::from_str("11").unwrap(),
            PostgreSQLVersion::V11
        );
        assert_eq!(
            PostgreSQLVersion::from_str("12").unwrap(),
            PostgreSQLVersion::V12
        );
        assert_eq!(
            PostgreSQLVersion::from_str("13").unwrap(),
            PostgreSQLVersion::V13
        );
        assert_eq!(
            PostgreSQLVersion::from_str("14").unwrap(),
            PostgreSQLVersion::V14
        );
    }
}

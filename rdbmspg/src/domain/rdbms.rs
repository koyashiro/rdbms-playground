use std::fmt::Display;

#[derive(Debug, Eq, PartialEq)]
pub enum Rdbms {
    MySQL(MySQLVersion),
    MariaDB(MariaDBVersion),
    PostgreSQL(PostgreSQLVersion),
}

impl Display for Rdbms {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::MySQL(v) => {
                write!(f, "mysql:")?;
                v.fmt(f)
            }
            Self::MariaDB(v) => {
                write!(f, "mariadb:")?;
                v.fmt(f)
            }
            Self::PostgreSQL(v) => {
                write!(f, "postgres:")?;
                v.fmt(f)
            }
        }
    }
}

#[derive(Debug, Eq, PartialEq)]
pub enum MySQLVersion {
    V5_7,
    V8_0,
}

impl Display for MySQLVersion {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::V5_7 => write!(f, "5.7"),
            Self::V8_0 => write!(f, "8.0"),
        }
    }
}

#[derive(Debug, Eq, PartialEq)]
pub enum MariaDBVersion {
    V10_2,
    V10_3,
    V10_4,
    V10_5,
    V10_6,
    V10_7,
    V10_8,
    V10_9,
}

impl Display for MariaDBVersion {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::V10_2 => write!(f, "10.2"),
            Self::V10_3 => write!(f, "10.3"),
            Self::V10_4 => write!(f, "10.4"),
            Self::V10_5 => write!(f, "10.5"),
            Self::V10_6 => write!(f, "10.6"),
            Self::V10_7 => write!(f, "10.7"),
            Self::V10_8 => write!(f, "10.8"),
            Self::V10_9 => write!(f, "10.9"),
        }
    }
}

#[derive(Debug, Eq, PartialEq)]
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
            Self::V10 => write!(f, "10"),
            Self::V11 => write!(f, "11"),
            Self::V12 => write!(f, "12"),
            Self::V13 => write!(f, "13"),
            Self::V14 => write!(f, "14"),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn rdbms_version_test() {
        assert_eq!(
            Rdbms::MySQL(MySQLVersion::V5_7).to_string(),
            "mysql:5.7".to_owned()
        );
        assert_eq!(
            Rdbms::MySQL(MySQLVersion::V8_0).to_string(),
            "mysql:8.0".to_owned()
        );

        assert_eq!(
            Rdbms::MariaDB(MariaDBVersion::V10_2).to_string(),
            "mariadb:10.2".to_owned()
        );
        assert_eq!(
            Rdbms::MariaDB(MariaDBVersion::V10_3).to_string(),
            "mariadb:10.3".to_owned()
        );
        assert_eq!(
            Rdbms::MariaDB(MariaDBVersion::V10_4).to_string(),
            "mariadb:10.4".to_owned()
        );
        assert_eq!(
            Rdbms::MariaDB(MariaDBVersion::V10_5).to_string(),
            "mariadb:10.5".to_owned()
        );
        assert_eq!(
            Rdbms::MariaDB(MariaDBVersion::V10_6).to_string(),
            "mariadb:10.6".to_owned()
        );
        assert_eq!(
            Rdbms::MariaDB(MariaDBVersion::V10_7).to_string(),
            "mariadb:10.7".to_owned()
        );
        assert_eq!(
            Rdbms::MariaDB(MariaDBVersion::V10_8).to_string(),
            "mariadb:10.8".to_owned()
        );
        assert_eq!(
            Rdbms::MariaDB(MariaDBVersion::V10_9).to_string(),
            "mariadb:10.9".to_owned()
        );

        assert_eq!(
            Rdbms::PostgreSQL(PostgreSQLVersion::V10).to_string(),
            "postgres:10".to_owned()
        );
        assert_eq!(
            Rdbms::PostgreSQL(PostgreSQLVersion::V11).to_string(),
            "postgres:11".to_owned()
        );
        assert_eq!(
            Rdbms::PostgreSQL(PostgreSQLVersion::V12).to_string(),
            "postgres:12".to_owned()
        );
        assert_eq!(
            Rdbms::PostgreSQL(PostgreSQLVersion::V13).to_string(),
            "postgres:13".to_owned()
        );
        assert_eq!(
            Rdbms::PostgreSQL(PostgreSQLVersion::V14).to_string(),
            "postgres:14".to_owned()
        );
    }

    #[test]
    fn mysql_version_test() {
        assert_eq!(MySQLVersion::V5_7.to_string(), "5.7".to_owned());
        assert_eq!(MySQLVersion::V8_0.to_string(), "8.0".to_owned());
    }

    #[test]
    fn mariadb_version_test() {
        assert_eq!(MariaDBVersion::V10_2.to_string(), "10.2".to_owned());
        assert_eq!(MariaDBVersion::V10_3.to_string(), "10.3".to_owned());
        assert_eq!(MariaDBVersion::V10_4.to_string(), "10.4".to_owned());
        assert_eq!(MariaDBVersion::V10_5.to_string(), "10.5".to_owned());
        assert_eq!(MariaDBVersion::V10_6.to_string(), "10.6".to_owned());
        assert_eq!(MariaDBVersion::V10_7.to_string(), "10.7".to_owned());
        assert_eq!(MariaDBVersion::V10_8.to_string(), "10.8".to_owned());
        assert_eq!(MariaDBVersion::V10_9.to_string(), "10.9".to_owned());
    }

    #[test]
    fn postgres_version_test() {
        assert_eq!(PostgreSQLVersion::V10.to_string(), "10".to_owned());
        assert_eq!(PostgreSQLVersion::V11.to_string(), "11".to_owned());
        assert_eq!(PostgreSQLVersion::V12.to_string(), "12".to_owned());
        assert_eq!(PostgreSQLVersion::V13.to_string(), "13".to_owned());
        assert_eq!(PostgreSQLVersion::V14.to_string(), "14".to_owned());
    }
}

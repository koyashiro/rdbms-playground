use rdbms_gateway::{
    database::{Database, Value},
    execute::{Execute, ExecuteResult},
};
use sqlx::{Column, Row, SqlitePool, TypeInfo};

#[derive(Debug)]
pub struct Sqlite;

impl Database for Sqlite {
    type Error = sqlx::Error;
}

#[derive(Debug)]
pub struct SqliteExecute {
    pool: SqlitePool,
}

impl SqliteExecute {
    pub fn new(pool: SqlitePool) -> Self {
        Self { pool }
    }
}

#[async_trait::async_trait]
impl Execute<Sqlite> for SqliteExecute {
    async fn execute(&self, query: &str) -> Result<ExecuteResult, sqlx::Error> {
        let rows = sqlx::query(query).fetch_all(&self.pool).await?;

        let rows = rows
            .iter()
            .map(|row| {
                let values = row
                    .columns()
                    .iter()
                    .enumerate()
                    .map(|(i, column)| {
                        let value = match column.type_info().name() {
                            "INTEGER" => Value::I64(row.get_unchecked(i)),
                            "REAL" => Value::F64(row.get_unchecked(i)),
                            "TEXT" => Value::String(row.get_unchecked(i)),
                            "BLOB" => Value::Bytes(row.get_unchecked(i)),
                            _ => unreachable!(),
                        };
                        Some(value)
                    })
                    .collect();
                values
            })
            .collect();

        let execute_result = ExecuteResult { rows };

        Ok(execute_result)
    }
}

#[cfg(test)]
mod tests {
    use sqlx::sqlite::SqlitePoolOptions;

    use super::*;

    #[tokio::test]
    async fn execute_test() {
        let pool = SqlitePoolOptions::new()
            .connect("sqlite::memory:")
            .await
            .unwrap();

        sqlx::query(
            "
            CREATE TABLE tbl (
                i INTEGER,
                r REAL,
                t TEXT,
                b BLOB
            );
            INSERT INTO tbl (i, r, t, b) VALUES (12345, 123.45, 'text', x'0001020304050607');
            INSERT INTO tbl (i, r, t, b) VALUES (NULL, NULL, NULL, NULL);
            ",
        )
        .execute(&pool)
        .await
        .unwrap();

        let e = SqliteExecute::new(pool);
        let rows = e
            .execute(
                "
                SELECT i, r, t, b
                FROM tbl;
                ",
            )
            .await
            .unwrap();

        assert_eq!(
            rows,
            ExecuteResult {
                rows: vec![
                    vec![
                        Some(Value::I64(12345)),
                        Some(Value::F64(123.45)),
                        Some(Value::String("text".to_owned())),
                        Some(Value::Bytes(vec![
                            0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07
                        ])),
                    ],
                    vec![
                        Some(Value::I64(0_i64)),
                        Some(Value::F64(0_f64)),
                        Some(Value::String("".to_owned())),
                        Some(Value::Bytes(Vec::new())),
                    ],
                ]
            }
        );
    }
}

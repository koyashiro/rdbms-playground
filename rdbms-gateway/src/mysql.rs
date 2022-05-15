use crate::{
    database::Database,
    execute::{Execute, ExecuteResult, Value},
};
use sqlx::{MySqlPool, Row, TypeInfo, ValueRef};

#[derive(Debug)]
pub struct MySql;

impl Database for MySql {
    type Error = sqlx::Error;
}

#[derive(Debug)]
pub struct MySqlExecute {
    pool: MySqlPool,
}

impl MySqlExecute {
    pub fn new(pool: MySqlPool) -> Self {
        Self { pool }
    }
}

#[async_trait::async_trait]
impl Execute<MySql> for MySqlExecute {
    async fn execute(&self, query: &str) -> Result<ExecuteResult, sqlx::Error> {
        let rows = sqlx::query(query).fetch_all(&self.pool).await?;

        let rows = rows
            .iter()
            .map(|row| {
                (0..row.columns().len())
                    .map(|i| {
                        let value = row.try_get_raw(i).unwrap();
                        if value.is_null() {
                            None
                        } else {
                            let type_info = value.type_info();
                            let type_name = type_info.name();
                            Some(match type_name {
                                "BOOLEAN" => Value::Bool(row.get_unchecked(i)),
                                "TINYINT" => Value::I8(row.get_unchecked(i)),
                                "SMALLINT" => Value::I16(row.get_unchecked(i)),
                                "MEDIUMINT" | "INT" => Value::I32(row.get_unchecked(i)),
                                "BIGINT" => Value::I64(row.get_unchecked(i)),
                                "TINYINT UNSIGNED" => Value::U8(row.get_unchecked(i)),
                                "SMALLINT UNSIGNED" => Value::U16(row.get_unchecked(i)),
                                "MEDIUMINT UNSIGNED" | "INT UNSIGNED" => {
                                    Value::U32(row.get_unchecked(i))
                                }
                                "BIGINT UNSIGNED" => Value::U64(row.get_unchecked(i)),
                                "FLOAT" => Value::F32(row.get_unchecked(i)),
                                "DOUBLE" => Value::F64(row.get_unchecked(i)),
                                "CHAR" | "VARCHAR" | "TEXT" => Value::String(row.get_unchecked(i)),
                                "BINARY" | "VARBINARY" | "BLOB" => {
                                    Value::Bytes(row.get_unchecked(i))
                                }
                                _ => todo!("unexpected column type: {type_name}"),
                            })
                        }
                    })
                    .collect()
            })
            .collect();

        let execute_result = ExecuteResult { rows };

        Ok(execute_result)
    }
}

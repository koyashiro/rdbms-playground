use rdbms_gateway::{
    database::{Database, Value},
    execute::{Execute, ExecuteResult},
};
use sqlx::{Column, MySqlPool, Row, TypeInfo};

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
                let values = row
                    .columns()
                    .iter()
                    .enumerate()
                    .map(|(i, column)| {
                        let column_type_name = column.type_info().name();
                        match column_type_name {
                            "BOOLEAN" => Some(Value::Bool(row.get_unchecked(i))),
                            "TINYINT" => Some(Value::I8(row.get_unchecked(i))),
                            "SMALLINT" => Some(Value::I16(row.get_unchecked(i))),
                            "MEDIUMINT" | "INT" => Some(Value::I32(row.get_unchecked(i))),
                            "BIGINT" => Some(Value::I64(row.get_unchecked(i))),
                            "TINYINT UNSIGNED" => Some(Value::U8(row.get_unchecked(i))),
                            "SMALLINT UNSIGNED" => Some(Value::U16(row.get_unchecked(i))),
                            "MEDIUMINT UNSIGNED" | "INT UNSIGNED" => {
                                Some(Value::U32(row.get_unchecked(i)))
                            }
                            "BIGINT UNSIGNED" => Some(Value::U64(row.get_unchecked(i))),
                            "FLOAT" => Some(Value::F32(row.get_unchecked(i))),
                            "DOUBLE" => Some(Value::F64(row.get_unchecked(i))),
                            "CHAR" | "VARCHAR" | "TEXT" => {
                                Some(Value::String(row.get_unchecked(i)))
                            }
                            "BINARY" | "VARBINARY" | "BLOB" => {
                                Some(Value::Bytes(row.get_unchecked(i)))
                            }
                            "NULL" => None,
                            _ => todo!("unexpected column type: {column_type_name}"),
                        }
                    })
                    .collect();
                values
            })
            .collect();

        let execute_result = ExecuteResult { rows };

        Ok(execute_result)
    }
}

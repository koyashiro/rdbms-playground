use rdbms_gateway::{
    database::{Database, Value},
    execute::{Execute, ExecuteResult},
};
use sqlx::{PgPool, Row, TypeInfo, ValueRef};

#[derive(Debug)]
pub struct Pg;

impl Database for Pg {
    type Error = sqlx::Error;
}

#[derive(Debug)]
pub struct PgExecute {
    pool: PgPool,
}

impl PgExecute {
    pub fn new(pool: PgPool) -> Self {
        Self { pool }
    }
}

#[async_trait::async_trait]
impl Execute<Pg> for PgExecute {
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
                                "BOOL" => Value::Bool(row.get_unchecked(i)),
                                "INT2" => Value::I16(row.get_unchecked(i)),
                                "INT4" => Value::I32(row.get_unchecked(i)),
                                "INT8" => Value::I64(row.get_unchecked(i)),
                                "FLOAT4" => Value::F32(row.get_unchecked(i)),
                                "FLOAT8" => Value::F64(row.get_unchecked(i)),
                                "CHAR" | "VARCHAR" | "TEXT" => Value::String(row.get_unchecked(i)),
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

use std::sync::Arc;

use tonic::{Response, Status};

use crate::{
    database::Database,
    database::Value,
    execute::{Execute, ExecuteResult},
    proto,
};

impl From<Option<Value>> for proto::Value {
    fn from(v: Option<Value>) -> Self {
        let value = v.map(|v| match v {
            Value::Bool(b) => proto::value::Value::Bool(b),
            Value::I8(i) => proto::value::Value::I8(i.into()),
            Value::I16(i) => proto::value::Value::I16(i.into()),
            Value::I32(i) => proto::value::Value::I32(i),
            Value::I64(i) => proto::value::Value::I64(i),
            Value::U8(i) => proto::value::Value::U8(i.into()),
            Value::U16(i) => proto::value::Value::U16(i.into()),
            Value::U32(i) => proto::value::Value::U32(i),
            Value::U64(i) => proto::value::Value::U64(i),
            Value::F32(f) => proto::value::Value::F32(f),
            Value::F64(f) => proto::value::Value::F64(f),
            Value::String(s) => proto::value::Value::String(s),
            Value::Bytes(b) => proto::value::Value::Bytes(b),
        });
        Self { value }
    }
}

impl From<ExecuteResult> for proto::ExecuteResponse {
    fn from(er: ExecuteResult) -> Self {
        let rows = er
            .rows
            .into_iter()
            .map(|r| {
                let values = r.into_iter().map(|v| v.into()).collect();
                proto::Row { values }
            })
            .collect();
        proto::ExecuteResponse { rows }
    }
}

#[derive(Debug)]
pub struct RdbmsGatewayService<T: Database> {
    execute: Arc<dyn Execute<T>>,
}

impl<T: Database> RdbmsGatewayService<T> {
    pub fn new(execute: Arc<dyn Execute<T>>) -> Self {
        Self { execute }
    }
}

#[async_trait::async_trait]
impl<T: Database> proto::rdbms_gateway_service_server::RdbmsGatewayService
    for RdbmsGatewayService<T>
{
    async fn execute(
        &self,
        request: tonic::Request<proto::ExecuteRequest>,
    ) -> Result<Response<proto::ExecuteResponse>, Status> {
        let query = request.get_ref().query.as_str();
        let execute_result = self.execute.execute(query).await.unwrap();
        let response = Response::new(execute_result.into());
        Ok(response)
    }
}

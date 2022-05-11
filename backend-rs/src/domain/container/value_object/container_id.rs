#[derive(Clone, Debug, Eq, PartialEq)]
pub struct ContainerId(String);

impl ContainerId {
    pub fn as_str(&self) -> &str {
        &self.0
    }
}

impl From<String> for ContainerId {
    fn from(v: String) -> Self {
        Self(v)
    }
}

impl From<ContainerId> for String {
    fn from(v: ContainerId) -> Self {
        v.0
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn container_id_from_string() {
        assert_eq!(
            ContainerId::from("9d24234b-3edf-485a-999c-5a6eeab6db9f".to_owned()),
            ContainerId("9d24234b-3edf-485a-999c-5a6eeab6db9f".to_owned())
        );
    }

    #[test]
    fn string_from_container_id() {
        assert_eq!(
            String::from(ContainerId(
                "874964a2-6199-49c0-844c-0295e956fdfe".to_owned()
            )),
            "874964a2-6199-49c0-844c-0295e956fdfe".to_owned()
        );
    }
}

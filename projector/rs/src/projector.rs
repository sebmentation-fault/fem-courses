use std::{collections::HashMap, path::PathBuf};

use serde::{Deserialize, Serialize};

use crate::config::Config;

#[derive(Debug, Serialize, Deserialize)]
struct Data {
    pub projector: HashMap<PathBuf, HashMap<String, String>>,
}

pub struct Projector {
    config: Config,
    data: Data,
}

// Better implementation would be this icl
// impl From<Config> for Projector {
    // add code here
// }

fn default_data() -> Data {
    return Data {
        projector: HashMap::new(),
    }

}

impl From<Config> for Projector {
    fn from(config: Config) -> Projector {
        if std::fs::metadata(&config.config).is_ok() {
            let contents = std::fs::read_to_string(&config.config);
            let contents = contents.unwrap_or(
                String::from("{\"projector\": {}}")
            );
            let data = serde_json::from_str(&contents);
            let data = data.unwrap_or(default_data());

            return Projector {
                config,
                data,
            }
        }


        return Projector {
            config,
            data: default_data(),
        }
    }
}

impl Projector {
    pub fn get_value_all(&self) -> HashMap<&String, &String> {
        let mut paths = vec![];
        let mut curr = Some(self.config.pwd.as_path());

        while let Some(p) = curr {
            paths.push(p);
            curr = p.parent();
        }

        let mut out = HashMap::new();

        for path in paths.into_iter().rev() {
            self.data.projector.get(path).map(|x| {
                out.extend(x.iter());
            });
        }

        return out;
    }

    pub fn get_value(&self, key: &str) -> Option<String> {

        let mut curr = Some(self.config.pwd.as_path());
        let mut out = None;

        while let Some(p) = curr {
            if let Some(dir) = self.data.projector.get(p) {
                if let Some(value) = dir.get(key) {
                    out = Some(value.to_string());
                    break;
                }
            }

            curr = p.parent();
        }

        return out;
    }

    pub fn set_value(&mut self, key: &str, value: &str) {
        self.data.projector
            .entry(self.config.pwd.clone())
            .or_default()
            .insert(key.into(), value.into());
    }
    pub fn remove_value(&mut self, key: &str) {
        self.data.projector
            .get_mut(&self.config.pwd)
            .map(|x| {
                x.remove(key);
            });
    }
}

#[cfg(test)]
mod test {
    use std::{collections::HashMap, path::PathBuf};

    use collection_macros::hashmap;

    use crate::config::{Config, Operation};

    use super::{Data, Projector};

    fn get_data() -> HashMap<PathBuf, HashMap<String, String>> {
        return hashmap! {
            PathBuf::from("/") => hashmap! {
                "foo".into() => "bar1".into(),
                "fem".into() => "am_learning".into(),
            },
            PathBuf::from("/foo") => hashmap! {
                "foo".into() => "bar2".into(),
            },
            PathBuf::from("/foo/bar") => hashmap! {
                "foo".into() => "bar3".into(),
            },
        }
    }

    fn get_projector(pwd: PathBuf) -> Projector {
        return Projector {
            config: Config {
                pwd,
                config: PathBuf::from(""),
                operation: Operation::Print(None),
            },
            data: Data {
                projector: get_data(),
            }
        }
    }

    #[test]
    fn get_value() {
        let proj = get_projector("/foo/bar".into());

        assert_eq!(proj.get_value("foo"), Some(String::from("bar3")));
        assert_eq!(proj.get_value("fem"), Some(String::from("am_learning")));
    }

    #[test]
    fn set_value() {
        let mut proj = get_projector("/foo/bar".into());

        proj.set_value("foo", "bar4");
        proj.set_value("fem", "stopped_learning");

        assert_eq!(proj.get_value("foo"), Some(String::from("bar4")));
        assert_eq!(proj.get_value("fem"), Some(String::from("stopped_learning")));
    }

    #[test]
    fn remove_value() {
        let mut proj = get_projector("/foo/bar".into());

        proj.remove_value("foo");
        proj.remove_value("fem");

        assert_eq!(proj.get_value("foo"), Some(String::from("bar2")));
        assert_eq!(proj.get_value("fem"), Some(String::from("am_learning")));
    }


}

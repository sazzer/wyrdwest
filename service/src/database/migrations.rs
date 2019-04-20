use crate::database::DatabaseWrapper;
use ::dbmigrate_lib::{read_migration_files, get_driver};
use std::path::Path;

pub trait MigratableDatabase {
    fn migrate(&self) -> Result<(), String>;
}

impl MigratableDatabase for DatabaseWrapper {
    fn migrate(&self) -> Result<(), String> {
        info!("Migrating database");

        let files = read_migration_files(Path::new("migrations"))
            .map_err(|e| format!("Failed to list files: {}", e))?;

        let driver = get_driver(self.url.as_str())
            .map_err(|e| format!("Failed to load driver: {}", e))?;

        let current = driver.get_current_number();
        let max = files.keys().max().unwrap_or(&0);

        debug!("Current migration: {} of {}", current, max);

        if current == *max {
            info!("Migrations are up-to-date");
            return Ok(());
        }

        for (number, migration) in files.iter() {
            if number > &current {
                let mig_file = migration.up.as_ref().unwrap();
                debug!("Applying migration {}: {}", number, mig_file.filename);

                match driver.migrate(mig_file.content.clone().unwrap(), mig_file.number) {
                    Err(e) => Err(format!("Failed to apply migration: {}", e)),
                    Ok(_) => {
                        debug!("Applied migration {}: {}", number, mig_file.filename); 
                        Ok(())
                    },
                }?
            }
        }

        Ok(())

    }
}

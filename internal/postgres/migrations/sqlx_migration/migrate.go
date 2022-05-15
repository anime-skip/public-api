package sqlx_migration

import (
	"anime-skip.com/public-api/internal"
	"anime-skip.com/public-api/internal/log"
)

const migrationTableName = ""

func RunAllMigrations(tx internal.Tx, name string, migrations []*Migration) error {
	return RunMigrations(tx, name, migrations, len(migrations))
}

func RunMigrations(tx internal.Tx, name string, migrations []*Migration, targetVersion int) error {
	_, err := tx.Exec(`CREATE TABLE IF NOT EXISTS migrations (id text PRIMARY KEY)`)
	if err != nil {
		return err
	}
	existing, err := getExistingMigrationIds(tx)
	if err != nil {
		return err
	}

	upgrades := []*Migration{}
	downgrades := []*Migration{}
	for version, migration := range migrations {
		hasRan := hasRanMigration(existing, migration)
		if version <= targetVersion && !hasRan {
			upgrades = append(upgrades, migration)
		} else if version > targetVersion && hasRan {
			downgrades = append(downgrades, migration)
		}
	}

	if len(upgrades) == 0 && len(downgrades) == 0 {
		log.D("No %s to run", name)
		return nil
	}
	if len(upgrades) > 0 {
		log.I("Targeting database version %d...", targetVersion)
		for _, migration := range upgrades {
			log.I(migration.ID)
			err = migration.Up(tx)
			if err != nil {
				return err
			}
			err = insertMigration(tx, migration.ID)
			if err != nil {
				return err
			}
		}
	}
	if len(downgrades) > 0 {
		log.I("Downgrading database to %d...", targetVersion)
		for _, migration := range downgrades {
			log.I(migration.ID)
			err = migration.Down(tx)
			if err != nil {
				return err
			}
			err = deleteMigration(tx, migration.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getExistingMigrationIds(tx internal.Tx) ([]ExistingMigration, error) {
	rows, err := tx.Query("SELECT id FROM migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	existing := []ExistingMigration{}
	for rows.Next() {
		var m ExistingMigration
		err = rows.Scan(&m.ID)
		if err != nil {
			return nil, err
		}
		existing = append(existing, m)
	}
	return existing, nil
}

func hasRanMigration(existing []ExistingMigration, migration *Migration) bool {
	for _, e := range existing {
		if e.ID == migration.ID {
			return true
		}
	}
	return false
}

func insertMigration(tx internal.Tx, id string) error {
	_, err := tx.Exec("INSERT INTO migrations (id) VALUES ($1)", id)
	return err
}

func deleteMigration(tx internal.Tx, id string) error {
	_, err := tx.Exec("DELETE FROM migrations WHERE id=$1", id)
	return err
}

package migrations

import (
	"log/slog"

	"github.com/jinzhu/gorm"
)

// One time migrations to arrange required schema
func RunMigrations(db *gorm.DB, logger *slog.Logger) {
	// db.AutoMigrate(&dbmodels.User{}, &dbmodels.Note{}, &dbmodels.SharedNote{})
	// 	db.Debug().Model(&dbmodels.Note{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	// 	db.Exec(`
	//     ALTER TABLE notes
	//     ADD COLUMN search_vector tsvector GENERATED ALWAYS AS (
	//         to_tsvector('english', coalesce(title, '') || ' ' || coalesce(content, ''))
	//     ) STORED;

	//     CREATE INDEX notes_search_vector_index
	//     ON notes
	//     USING GIN(search_vector);

	//     CREATE TRIGGER notes_search_vector_update
	//     BEFORE INSERT OR UPDATE ON notes
	//     FOR EACH ROW EXECUTE FUNCTION tsvector_update_trigger(search_vector, 'pg_catalog.english', title, content);
	// `)

}

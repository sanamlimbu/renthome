BEGIN;
DROP TRIGGER updatePropertyKeywords;
DROP FUNCTION updatePropretyKeywords ();
DROP TABLE properties;
DROP TABLE password_hashes;
DROP TRIGGER updateUserKeywords;
DROP FUNCTION updateUserKeywords ();
DROP TABLE users;
DROP TABLE blobs;
COMMIT;
BEGIN;
DROP TRIGGER updatePropertyKeywords;
DROP FUNCTION updatePropretyKeywords ();
DROP TABLE properties;
DROP TABLE password_hashes;
DROP TRIGGER updateUserKeywords;
DROP FUNCTION updateUserKeywords ();
DROP TABLE property_blobs;
DROP TABLE user_notifications;
DROP TABLE user_privacies;
DROP TABLE notifications;
DROP TABLE privacies;
DROP TABLE images;
DROP TABLE users;
DROP TABLE blobs;
COMMIT;
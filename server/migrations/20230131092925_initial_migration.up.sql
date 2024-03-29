BEGIN;

-- blobs
CREATE TABLE blobs 
(
    id              UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    file_name       TEXT             NOT NULL,
    mime_type       TEXT             NOT NULL,
    file_size_bytes BIGINT           NOT NULL,
    extension       TEXT             NOT NULL,
    file            BYTEA            NOT NULL,
    views           INTEGER          NOT NULL DEFAULT 0,
    hash            TEXT             NOT NULL,
    public          BOOLEAN          NOT NULL DEFAULT FALSE,

    created_at      TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
); 

/*************
 *  Agencies  *
 *************/

 CREATE TABLE agencies
 (
    id          UUID        NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    name        TEXT        NOT NULL,
    color       TEXT        NOT NULL,
    logo_id     UUID REFERENCES blobs (id),  
    image_id    UUID REFERENCES blobs (id),      
    created_at  TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    deleted_at  TIMESTAMPTZ
 );

/******************
 *  Users  *
 ******************/
 CREATE TABLE users 
 (
    id                      UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    first_name              TEXT             NOT NULL DEFAULT '',
    last_name               TEXT             NOT NULL DEFAULT '',
    email                   TEXT UNIQUE,
    facebook_id             TEXT UNIQUE,
    google_id               TEXT UNIQUE,
    apple_id                TEXT UNIQUE,
    title                   TEXT                      DEFAULT '',
    description             TEXT                      DEFAULT '',
    role                    TEXT             NOT NULL CHECK (role IN ('MEMBER', 'MANAGER', 'ADMIN', 'AGENCY')), 
    mobile                  TEXT                      DEFAULT '',
    is_verified             BOOLEAN          NOT NULL DEFAULT FALSE,
    old_password_required   BOOLEAN          NOT NULL DEFAULT TRUE, -- set to false on password reset request, set back to true on password change
    avatar_id               UUID REFERENCES blobs (id),
    agency_id               UUID REFERENCES agencies (id),
    is_agency               BOOLEAN         NOT NULL DEFAULT FALSE,
    keywords                TSVECTOR,
    created_at              TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    deleted_at              TIMESTAMPTZ
 );

 -- for user text search 
CREATE INDEX idx_fts_user_vec ON users USING gin (keywords);

CREATE OR REPLACE FUNCTION emailToLowercase() 
    RETURNS TRIGGER
AS 
$emailToLowercase$
BEGIN 
    NEW.email = LOWER(NEW.email);
    RETURN NEW;
END;
$emailToLowercase$ 
    LANGUAGE plpgsql;

CREATE TRIGGER emailToLowercase
    BEFORE INSERT OR UPDATE
    ON users
    FOR EACH ROW
    EXECUTE PROCEDURE emailToLowercase();

CREATE OR REPLACE FUNCTION updateUserKeywords()
    RETURNS TRIGGER
AS
$updateUserKeywords$
DECLARE
    temp TSVECTOR;
BEGIN
    SELECT (
            SETWEIGHT(TO_TSVECTOR('english', NEW.first_name), 'A') ||
            SETWEIGHT(TO_TSVECTOR('english', NEW.last_name), 'A') ||
            SETWEIGHT(TO_TSVECTOR('english', COALESCE(NEW.email, '')), 'A')
            )
    INTO temp;
    IF TG_OP = 'INSERT' OR temp != OLD.keywords THEN
        UPDATE
            users
        SET keywords = temp
        WHERE id = NEW.id;
    END IF;
    RETURN NULL;
END;
$updateUserKeywords$
    LANGUAGE plpgsql;

CREATE TRIGGER updateUserKeywords
    AFTER INSERT OR UPDATE
    ON users
    FOR EACH ROW
    EXECUTE PROCEDURE updateUserKeywords(); 

-- password hashes
CREATE TABLE password_hashes
(
    user_id       UUID        NOT NULL REFERENCES users (id),
    password_hash TEXT        NOT NULL,
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ,
    PRIMARY KEY (user_id)
);

-- reset passowrd codes
CREATE TABLE reset_passwords
(
    user_id       UUID        NOT NULL REFERENCES users (id),
    code          TEXT        NOT NULL,
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIMESTAMPTZ,
    PRIMARY KEY (user_id)
);

/*************
 *  Properties  *
 *************/
 CREATE TABLE properties
(
    id                  UUID        NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    slug                TEXT UNIQUE NOT NULL,
    type                TEXT        NOT NULL CHECK (type IN ('Unit', 'Apartment', 'House', 'Villa', 'Townhouse')), 
    category            TEXT        NOT NULL CHECK (category IN ('Rent', 'Buy', 'Sold')), 
    street              TEXT        NOT NULL,
    suburb              TEXT        NOT NULL,
    postcode            TEXT        NOT NULL,
    state               TEXT        NOT NULL CHECK (state IN ('NSW', 'VIC', 'QLD', 'SA', 'WA', 'TAS', 'NT', 'ACT')),     
    location            TEXT        NOT NULL,
    bed_count           INTEGER     NOT NULL,
    bath_count          INTEGER     NOT NULL,
    car_count           INTEGER     NOT NULL,
    has_aircon          BOOLEAN     NOT NULL DEFAULT TRUE,
    has_dishwasher      BOOLEAN     NOT NULL DEFAULT TRUE,
    is_furnished        BOOLEAN     NOT NULL DEFAULT FALSE,
    is_pets_considered  BOOLEAN     NOT NULL DEFAULT FALSE,
    available_at        TIMESTAMPTZ,
    is_available_now    BOOLEAN    NOT NULL DEFAULT FALSE,
    open_at             TIMESTAMPTZ,
    price               INTEGER     NOT NULL,
    agency_id           UUID        NOT NULL REFERENCES agencies (id),
    manager_id          UUID        NOT NULL REFERENCES users (id),
    keywords            TSVECTOR,
    created_at          TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);

-- for properties text search
CREATE INDEX idx_fts_property_vec ON properties USING gin (keywords);

CREATE OR REPLACE FUNCTION updatePropertyKeywords()
    RETURNS TRIGGER
AS
$updatePropertyKeywords$
DECLARE
    temp TSVECTOR;
BEGIN
    SELECT (
            SETWEIGHT(TO_TSVECTOR('english', NEW.slug), 'A') ||
            SETWEIGHT(TO_TSVECTOR('english', NEW.suburb), 'A') ||
            SETWEIGHT(TO_TSVECTOR('english', NEW.postcode), 'A') ||
            SETWEIGHT(TO_TSVECTOR('english', NEW.state), 'A')
            )
    INTO temp;
    IF TG_OP = 'INSERT' OR temp != OLD.keywords THEN
        UPDATE
            properties
        SET keywords = temp
        WHERE id = NEW.id;
    END IF;
    RETURN NULL;
END;
$updatePropertyKeywords$
    LANGUAGE plpgsql;

CREATE TRIGGER updatePropertyKeywords
    AFTER INSERT OR UPDATE
    ON properties
    FOR EACH ROW
    EXECUTE PROCEDURE updatePropertyKeywords();

-- property blobs
CREATE TABLE property_blobs
(
    property_id UUID NOT NULL REFERENCES properties (id),
    blob_id     UUID NOT NULL REFERENCES blobs (id),
    PRIMARY KEY (property_id, blob_id)
);

-- images
CREATE TABLE images
(
    id              UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    path            TEXT             NOT NULL,
    file_size_bytes BIGINT           NOT NULL,
    mime_type       TEXT             NOT NULL,           
    extension       TEXT             NOT NULL,
    property_id     UUID             NOT NULL REFERENCES properties (id),
    uploader_id     UUID             NOT NULL REFERENCES users (id),
    created_at      TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ      NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ
);

/*************
 *  Notifications  *
 *************/
 CREATE TABLE notifications
(
    id                  UUID        NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    name                TEXT        NOT NULL,
    slug                TEXT UNIQUE NOT NULL,
    description         TEXT        NOT NULL,        
    method              TEXT        NOT NULL CHECK (method IN ('Email', 'Push')), 
    category            TEXT        NOT NULL CHECK (category IN ('Property journey', 'Properties', 'Property market', 'Finance')),
    created_at          TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);

-- user_notifications
CREATE TABLE user_notifications
(
    user_id             UUID NOT NULL REFERENCES users (id),
    notification_id     UUID NOT NULL REFERENCES notifications (id),
    state               TEXT NOT NULL CHECK (state IN ('On', 'Off')) DEFAULT 'On',
    created_at          TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    PRIMARY KEY (user_id, notification_id)
);

/*************
 *  Privacy  *
 *************/
 CREATE TABLE privacies
(
    id                  UUID        NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    name                TEXT        NOT NULL,
    slug                TEXT UNIQUE NOT NULL,
    description         TEXT        NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);

-- user_privacies
CREATE TABLE user_privacies
(
    user_id        UUID NOT NULL REFERENCES users (id),
    privacy_id     UUID NOT NULL REFERENCES privacies (id),
    state          TEXT NOT NULL CHECK (state IN ('On', 'Off')) DEFAULT 'On',
    created_at     TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    updated_at     TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    PRIMARY KEY (user_id, privacy_id)
);

CREATE TABLE issue_tokens
(
    id          UUID        NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID        NOT NULL REFERENCES users (id),
    device      TEXT        NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    expires_at  TIMESTAMPTZ NOT NULL,
    blacklisted BOOLEAN     NOT NULL             DEFAULT FALSE
);

/*************
 *  Locations  *
 *************/
 CREATE TABLE locations
(
    id                  UUID        NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    suburb              TEXT        NOT NULL,
    postcode            TEXT UNIQUE NOT NULL,
    state               TEXT        NOT NULL,
    description         TEXT        NOT NULL,        
    created_at          TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);

/*************
 *  Searches  *
 *************/
 CREATE TABLE searches
(
    id                  UUID        NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id             UUID        REFERENCES users (id),
    device              TEXT        NOT NULL,
    texts               TEXT[],          
    created_at          TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL             DEFAULT NOW(),
    deleted_at          TIMESTAMPTZ
);


COMMIT;
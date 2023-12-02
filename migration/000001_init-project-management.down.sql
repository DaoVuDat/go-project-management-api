ALTER TABLE "project" DROP CONSTRAINT IF EXISTS "project_user_profile_fkey";

ALTER TABLE "user_profile" DROP CONSTRAINT IF EXISTS "user_profile_id_fkey";

DROP TABLE IF EXISTS "project";
DROP TABLE IF EXISTS "user_profile";
DROP TABLE IF EXISTS "user_account";

DROP TYPE IF EXISTS "project_status";
DROP TYPE IF EXISTS "account_status";
DROP TYPE IF EXISTS "account_type";
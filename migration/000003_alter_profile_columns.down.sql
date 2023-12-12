ALTER TABLE "project"
    ALTER COLUMN "name" SET NOT NULL,
    ALTER COLUMN "description" SET NOT NULL,
    ALTER COLUMN "start_time" SET NOT NULL,
    ALTER COLUMN "end_time" SET NOT NULL;
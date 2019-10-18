CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "users" (
"id" TEXT PRIMARY KEY,
"email" TEXT NOT NULL,
"username" TEXT NOT NULL,
"password" TEXT,
"provider" TEXT,
"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
"updated_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS "tasks" (
"id" TEXT PRIMARY KEY,
"user_id" char(36) NOT NULL,
"name" TEXT NOT NULL,
"created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
"updated_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE cascade
);

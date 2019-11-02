# Go Tasks

```sql
CREATE TABLE users (
	id 			VARCHAR(64) 	PRIMARY KEY	,
	provider_id	VARCHAR(128)				,
	username	VARCHAR(32) 	NOT NULL	,
    email 		VARCHAR(255) 	NOT NULL	,
    first_name 	VARCHAR(64) 	NOT NULL	,
	last_name 	VARCHAR(64) 	NOT NULL	,
	is_admin	BIT(1)			DEFAULT(0)	,
	created_at	DATETIME		DEFAULT CURRENT_TIMESTAMP,
	update_at	DATETIME		DATETIME ON UPDATE CURRENT_TIMESTAMP
)
GO

CREATE TABLE tasks (
	id 			VARCHAR(64) 	PRIMARY KEY	,
	user_id		VARCHAR(64) 	NOT NULL	,
	parent_id	VARCHAR(64) 	NOT NULL	,
	title		VARCHAR(32)		NOT NULL	,
	tags		TEXT						,
	note		TEXT						,
	completed	DATETIME					,
	streak		SMALLINT		DEFAULT(0)	,
    created_at	DATETIME		DEFAULT CURRENT_TIMESTAMP,
	update_at	DATETIME		DATETIME ON UPDATE CURRENT_TIMESTAMP,

	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY (parent_id) REFERENCES tasks(id) ON DELETE CASCADE
)
GO
```
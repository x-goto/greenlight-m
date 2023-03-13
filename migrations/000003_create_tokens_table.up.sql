CREATE TABLE IF NOT EXISTS tokens (
	user_id text PRIMARY KEY,
	refresh_token text NOT NULL
);
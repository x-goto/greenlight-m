CREATE TABLE IF NOT EXISTS user_auth_data (
	user_id text PRIMARY KEY,
	is_activated boolean NOT NULL DEFAULT false,
	activation_link text NOT NULL
);
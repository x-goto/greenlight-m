CREATE TABLE IF NOT EXISTS user_profiles(
	user_id text PRIMARY KEY,
	user_img_url text NOT NULL DEFAULT 'https://i.imgur.com/Eyzrkg3_d.webp?maxwidth=520&shape=thumb&fidelity=high',
	firstname text,
	lastname text,
	bio text
);
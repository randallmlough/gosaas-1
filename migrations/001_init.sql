CREATE TABLE gosaas_accounts(
	id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	email TEXT UNIQUE NOT NULL,
	stripe_id TEXT NOT NULL,
	subscription_id TEXT NOT NULL,
	plan TEXT NOT NULL,
	is_yearly BOOL NOT NULL,
	subscribed_on TIMESTAMP NOT NULL,
	seats INTEGER NOT NULL,
	is_active BOOL NOT NULL
);

CREATE TABLE gosaas_users(
	id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
	account_id INTEGER REFERENCES gosaas_accounts(id) ON DELETE CASCADE,
	email TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL,
	token TEXT UNIQUE NOT NULL,
	role INTEGER NOT NULL
);
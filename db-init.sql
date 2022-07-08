create extension "uuid-ossp";


CREATE TYPE user_role AS ENUM (
    'ADMIN',
    'SUPER_USER',
    'USER'
);

CREATE TYPE project_status AS ENUM (
	'CREATED',
	'ONGOING',
	'FINISHED',
	'CANCELED',
	'SUSPENDED'
);
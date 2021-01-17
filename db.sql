CREATE TABLE IF NOT EXISTS users (
  user_id SERIAL NOT NULL PRIMARY KEY,
  first_name VARCHAR(50) NOT NULL,
  middle_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS roles (
  role_id SERIAL NOT NULL PRIMARY KEY,
  role_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS permissions (
  perm_id SERIAL NOT NULL PRIMARY KEY,
  perm_desc VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS role_perm (
  role_id SERIAL NOT NULL,
  perm_id SERIAL NOT NULL,

  FOREIGN KEY (role_id) REFERENCES roles(role_id),
  FOREIGN KEY (perm_id) REFERENCES permissions(perm_id)
);

CREATE TABLE IF NOT EXISTS user_role (
  user_id SERIAL NOT NULL,
  role_id SERIAL NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(user_id),
  FOREIGN KEY (role_id) REFERENCES roles(role_id)
);

CREATE TABLE IF NOT EXISTS casbin_rule (
    p_type VARCHAR(100),
    v0 VARCHAR(100),
    v1 VARCHAR(100),
    v2 VARCHAR(100)
);

/* https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-create-table/ */
CREATE TABLE IF NOT EXISTS messages (
  jid VARCHAR(13),
  msg VARCHAR(255),
  created_on TIMESTAMP
);
How to start and stop PostgreSQL server?
October 30, 2018

In this post, we are going to figure out how to start, stop, and restart a PostgreSQL server on macOS, Linux, and Windows.

1. On macOS

If you installed PostgreSQL via Homebrew:

To start manually:
```bash
pg_ctl -D /usr/local/var/postgres start
```
To stop manually:
```bash
pg_ctl -D /usr/local/var/postgres stop
```
To start PostgreSQL server now and relaunch at login:
```bash
brew services start postgresql
```
And stop PostgreSQL:
```bash
brew services stop postgresql
```
If you want a hassle-free way to manage the local PostgreSQL database servers, use DBngin. It’s just one click to start, another click to turn off. No dependencies, no command line required, multiple drivers, multiple versions and multiple ports. And it’s free.

DBngin local server

2. On Windows

First, you need to find the PostgreSQL database directory, it can be something like C:\Program Files\PostgreSQL\10.4\data. Then open Command Prompt and execute this command:
```bash
pg_ctl -D "C:\Program Files\PostgreSQL\9.6\data" start
```
To stop the server
```bash
pg_ctl -D "C:\Program Files\PostgreSQL\9.6\data" stop
```
To restart the server:
```bash
pg_ctl -D "C:\Program Files\PostgreSQL\9.6\data" restart
```
Another way:

Open Run Window by `Winkey + R`
Type `services.msc`
Search Postgres service based on version installed.
Click stop, start or restart the service option.

3. On Linux

Update and install PostgreSQL 10.4
```bash
sudo apt-get update
sudo apt-get install postgresql-10.4
```
By default, the postgres user has no password and can hence only connect if ran by the postgres system user. The following command will assign it:
```bash
sudo -u postgres psql -c "ALTER USER postgres PASSWORD 'postgres';"
sudo -u postgres psql -c "CREATE DATABASE testdb;"
```
Start the PostgreSQL server
```bash
sudo service postgresql start
```
Stop the PostgreSQL server:
```bash
sudo service postgresql stop
```
Need a good GUI tool for PostgreSQL on MacOS and Windows? TablePlus is a modern, native tool with an elegant GUI that allows you to simultaneously manage multiple databases such as MySQL, PostgreSQL, SQLite, Microsoft SQL Server and more.
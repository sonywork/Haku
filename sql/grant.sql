CREATE USER txiaozhe WITH PASSWORD 'haku_db_txiaozhe';

GRANT SELECT, INSERT, UPDATE ON TABLE admin TO txiaozhe;

GRANT SELECT, INSERT, UPDATE ON TABLE blog TO txiaozhe;
GRANT SELECT, INSERT, UPDATE ON TABLE visit TO txiaozhe;

GRANT SELECT, INSERT, UPDATE ON TABLE badge TO txiaozhe;

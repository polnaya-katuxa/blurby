CREATE USER targetologist IDENTIFIED WITH sha256_password BY 'password1';

CREATE ROLE targetologist_role;
GRANT SELECT ON coursework.events TO targetologist_role;
GRANT SELECT ON coursework.ad_send_times TO targetologist_role;

GRANT targetologist_role TO targetologist;

BEGIN;DELETE FROM users;
ALTER SEQUENCE users_id_seq RESTART WITH 1;
INSERT INTO users
id,
created_at,
updated_at,
first_name,
last_name,
email,
student_number,
semester,
subject,
language,
encrypted_password,
reset_password_token,
confirm_email_token,
root) VALUES
(DEFAULT,
current_timestamp,
current_timestamp,
'Infomark',
'System',
'no-reply@info2.informatik.uni-tuebingen.de',
'42',
'42',
'computing',
'en',
'nologinpassword',
NULL,
NULL,
TRUE);
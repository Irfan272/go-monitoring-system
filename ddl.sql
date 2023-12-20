CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE division (
division_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
name_division VARCHAR (50),
description VARCHAR (255),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
updated_at TIMESTAMP, 
is_deleted BOOLEAN 
);

CREATE TABLE users (
user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
username VARCHAR(50),
email VARCHAR(50) UNIQUE NOT NULL,
password VARCHAR(50) NOT NULL,
birth_date DATE,
gender VARCHAR (10),
address VARCHAR (100),
no_hp VARCHAR (15),
division_id UUID,
role VARCHAR (50),
image_user VARCHAR (255),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
updated_at TIMESTAMP, 
is_deleted BOOLEAN,

FOREIGN KEY (division_id) REFERENCES division(division_id)
); 

CREATE TABLE equipment (
equipment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
name_equipment VARCHAR(255),
equipment_brand VARCHAR(255),
producer VARCHAR(255),
division_id UUID,
production_year DATE,
limit_year DATE,
equipment_age VARCHAR(255),
condision_status VARCHAR (10),
image_equipment VARCHAR (255),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
updated_at TIMESTAMP, 
is_deleted BOOLEAN,
FOREIGN KEY (division_id) REFERENCES division(division_id)
);


CREATE TABLE repair (
repair_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
title VARCHAR (255),
equipment_id UUID,
user_id UUID,
description VARCHAR (255),
priority VARCHAR(50),
location_equipment VARCHAR(255),
image_damage VARCHAR(255),
working_date DATE,
estimasion_date DATE,
actual_date DATE,
repair_status VARCHAR (20),
solusion VARCHAR (255),
techician_id UUID,
image_repair VARCHAR (255),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
updated_at TIMESTAMP, 
is_deleted BOOLEAN,
FOREIGN KEY (equipment_id) REFERENCES equipment(equipment_id),
FOREIGN KEY (user_id) REFERENCES users(user_id),
FOREIGN KEY (techician_id) REFERENCES users(user_id)
);

CREATE TABLE maintenance (
maintenance_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
equipment_id UUID,
priority VARCHAR (255),
working_date DATE,
estimasion_date DATE,
actual_date DATE,
techician_id UUID, 
maintenance_status VARCHAR (20),
solusion VARCHAR (255),
image_maintenance VARCHAR (255),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
updated_at TIMESTAMP, 
is_deleted BOOLEAN,
FOREIGN KEY (equipment_id) REFERENCES equipment(equipment_id),
FOREIGN KEY (techician_id) REFERENCES users(user_id)
);







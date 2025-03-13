CREATE TABLE mission (
    id UUID PRIMARY KEY,
    start_date TIMESTAMP,
    mission_duration INTEGER,
    robot_amount INTEGER,
    is_physical BOOLEAN,
    traveled_distance INTEGER,
    map_path VARCHAR(255)
);

CREATE TABLE log (
    id UUID PRIMARY KEY,
    mission_id UUID REFERENCES mission(id), 
    message VARCHAR(255),
    date TIMESTAMP
);


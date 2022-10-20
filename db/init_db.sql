CREATE IF NOT EXISTS TABLE food (
    id int primary key not null,
    name varchar(255) not null,
    type varchar(255) not null,
    purchase_date date not null,
    nutrition_fact_id int
)

CREATE IF NOT EXISTS TABLE record (
    id int primary key not null,
    food_id int not null,
    description text,
    eating_date date not null,
    satisfaction_score int not null,
    photo_url text,
    FOREIGN KEY food_id REFERENCES food id
)
DROP TABLE IF EXISTS events;

CREATE TABLE events (
-- Automatically created identifier for events
    id INT NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
-- Who is arranging the event?
    host varchar(255) NOT NULL,
-- Where the event is held?
    location varchar(255),
    start TIMESTAMP NOT NULL,
    end TIMESTAMP NOT NULL,
-- Dress code
    dc varchar(255),
-- Theme of the event can be categorized
    theme varchar(255),
-- At most 9,999,999.99 so 10 million €
    price decimal(9,2),
-- Link to signup page
    signup varchar(255),

    PRIMARY KEY (id)
);

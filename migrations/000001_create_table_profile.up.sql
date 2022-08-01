CREATE TABLE profile (
    id                  SERIAL,
    name                VARCHAR(255) NOT NULL,
    title               VARCHAR(255) NOT NULL,
    description         TEXT,
    image               VARCHAR(255) NOT NULL,
    twitter_username    VARCHAR(255) NOT NULL,
    twitter_id          BIGINT NOT NULL,
    CONSTRAINT pk_profile PRIMARY KEY(id)
);

INSERT INTO profile (name, title, image, twitter_username, twitter_id, description) 
    values ('guest guessing', 'guest guest guest', 'https://loremflickr.com/cache/resized/65535_50083909818_3ae528c790_320_240_nofilter.jpg', 'golang', 113419064, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.');
CREATE TABLE IF NOT EXISTS citix_user (
    id serial PRIMARY KEY,
    phone_number varchar(255) NOT NULL UNIQUE,
    passwoed varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS survey (
    id serial PRIMARY KEY,
    survey_id varchar(255) NOT NULL,
    survey_title varchar(255) NOT NULL,
    company_logo varchar(255) NOT NULL,
    questions json NOT NULL
);

INSERT INTO survey (survey_id, survey_title, company_logo, questions)
VALUES
    ('survey_1', 'Socializing Habits', 'https://www.akorda.kz/assets/media/gerb.jpg',
     '[{"title":"How often do you socialize with friends or acquaintances outside of work or school?","options":[{"id":1,"text":"Multiple times a week"},{"id":2,"text":"Once a week"},{"id":3,"text":"A few times a month"},{"id":4,"text":"Rarely or never"}]},{"title":"Which type of social gatherings do you enjoy the most?","options":[{"id":1,"text":"Dinner parties or potlucks"},{"id":2,"text":"Concerts and live events"},{"id":3,"text":"Outdoor picnics or barbecues"},{"id":4,"text":"Movie nights or game nights at home"}]},{"title":"How do you prefer to maintain social connections with friends and family who live far away?","options":[{"id":1,"text":"Video calls (e.g., Zoom, Skype)"},{"id":2,"text":"Phone calls and text messages"},{"id":3,"text":"Social media and messaging apps"},{"id":4,"text":"Email and occasional visits"}]},{"title":"When it comes to meeting new people, which of the following best describes you?","options":[{"id":1,"text":"Very outgoing and eager to meet new people"},{"id":2,"text":"Social but somewhat reserved in new social situations"},{"id":3,"text":"Introverted and prefer smaller, intimate gatherings"},{"id":4,"text":"Extremely introverted and avoid new social interactions"}]},{"title":"How important is volunteering or participating in community events to you for socializing and networking?","options":[{"id":1,"text":"Very important, I enjoy socializing while giving back to the community.\" "},{"id":2,"text":"Not important at all; I don''t engage in such activities."},{"id":3,"text":"Somewhat important; I participate occasionally."},{"id":4,"text":"Not very important; I prefer socializing through other means."}]},{"title":"Which social media platform do you use most frequently for connecting with friends and family?","options":[{"id":1,"text":"Facebook"},{"id":2,"text":"Instagram"},{"id":3,"text":"Snapchat"},{"id":4,"text":"Other"}]}]
');

CREATE TABLE IF NOT EXISTS survey_answer (
    id serial PRIMARY KEY,
    survey_id varchar(255) NOT NULL,
    survey_title varchar(255) NOT NULL,
    answers json NOT NULL
);

CREATE TABLE IF NOT EXISTS photo_info (
    id serial PRIMARY KEY,
    image_url varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
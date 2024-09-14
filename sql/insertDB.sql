INSERT INTO accounts (id, name, password, answer) VALUES ('a0', 'a0_name', 'a0_password', 0);
INSERT INTO accounts (id, name, password, answer) VALUES ('a1', 'a1_name', 'a1_password', 1);
INSERT INTO accounts (id, name, password, answer) VALUES ('a2', 'a2_name', 'a2_password', 2);
INSERT INTO accounts (id, name, password, answer) VALUES ('a3', 'a3_name', 'a3_password', 3);
INSERT INTO accounts (id, name, password, answer) VALUES ('a4', 'a4_name', 'a4_password', -1);
INSERT INTO accounts (id, name, password, answer) VALUES ('a5', 'a5_name', 'a5_password', -1);
INSERT INTO accounts (id, name, password, answer) VALUES ('a6', 'a6_name', 'a6_password', -1);

INSERT INTO profiles (id, name, introduction, icon_num, github_url, x_url) VALUES ('u1', 'u1_name', 'u1_intro', 0, 'u1_github_url', 'u1_x_url');
INSERT INTO profiles (id, name, introduction, icon_num, github_url, x_url) VALUES ('u2', 'u2_name', 'u2_intro', 1, 'u2_github_url', 'u2_x_url');
INSERT INTO profiles (id, name, introduction, icon_num, github_url, x_url) VALUES ('u3', 'u3_name', 'u3_intro', 2, 'u3_github_url', 'u3_x_url');
INSERT INTO profiles (id, name, introduction, icon_num, github_url, x_url) VALUES ('u4', 'u4_name', 'u4_intro', 3, 'u4_github_url', 'u4_x_url');
INSERT INTO profiles (id, name, introduction, icon_num, github_url, x_url) VALUES ('u5', 'u5_name', 'u5_intro', 4, 'u5_github_url', 'u5_x_url');

INSERT INTO products (id, title, user_name, url, description) VALUES ('p1', 'p1_title', 'u1_name', 'p1_url', 'p1_description');
INSERT INTO products (id, title, user_name, url, description) VALUES ('p2', 'p2_title', 'u2_name', 'p2_url', 'p2_description');
INSERT INTO products (id, title, user_name, url, description) VALUES ('p3', 'p3_title', 'u3_name', 'p3_url', 'p3_description');
INSERT INTO products (id, title, user_name, url, description) VALUES ('p4', 'p4_title', 'u4_name', 'p4_url', 'p4_description');

INSERT INTO comments (id, user_name, product_id, message) VALUES ('c1', 'a0_name', 'p1', 'c1_p1_message');
INSERT INTO comments (id, user_name, product_id, message) VALUES ('c2', 'a0_name', 'p2', 'c2_p2_message');
INSERT INTO comments (id, user_name, product_id, message) VALUES ('c3', 'a0_name', 'p1', 'c3_p1_message');
INSERT INTO comments (id, user_name, product_id, message) VALUES ('c4', 'a1_name', 'p3', 'c4_p3_message');
INSERT INTO comments (id, user_name, product_id, message) VALUES ('c5', 'a0_name', 'p4', 'c5_p4_message');
INSERT INTO comments (id, user_name, product_id, message) VALUES ('c6', 'a3_name', 'p1', 'c6_p1_message');
INSERT INTO comments (id, user_name, product_id, message) VALUES ('c7', 'a0_name', 'p1', 'c7_p1_message');
INSERT INTO comments (id, user_name, product_id, message) VALUES ('c8', 'a1_name', 'p1', 'c8_p1_message');
INSERT INTO comments (id, user_name, product_id, message) VALUES ('c9', 'a2_name', 'p2', 'c9_p2_message');
INSERT INTO comments (id, user_name, product_id, message) VALUES ('c10', 'a1_name', 'p1', 'c10_p1_message');

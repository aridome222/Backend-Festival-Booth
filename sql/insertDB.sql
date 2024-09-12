INSERT INTO profiles (id, name, introduction, icon_num, github_url, x_url) VALUES ('u1', 'u1_name', 'u1_intro', 0, 'u1_github_url', 'u1_x_url');
INSERT INTO profiles (id, name, introduction, icon_num, github_url, x_url) VALUES ('u2', 'u2_name', 'u2_intro', 0, 'u2_github_url', 'u2_x_url');
INSERT INTO profiles (id, name, introduction, icon_num, github_url, x_url) VALUES ('u3', 'u3_name', 'u3_intro', 0, 'u3_github_url', 'u3_x_url');
INSERT INTO profiles (id, name, introduction, icon_num, github_url, x_url) VALUES ('u4', 'u4_name', 'u4_intro', 0, 'u4_github_url', 'u4_x_url');
INSERT INTO profiles (id, name, introduction, icon_num, github_url, x_url) VALUES ('u5', 'u5_name', 'u5_intro', 0, 'u5_github_url', 'u5_x_url');

INSERT INTO products (id, title, user_name, url, description) VALUES ('p1', 'p1_title', 'u1_name', 'p1_url', 'p1_description');
INSERT INTO products (id, title, user_name, url, description) VALUES ('p2', 'p2_title', 'u1_name', 'p2_url', 'p2_description');
INSERT INTO products (id, title, user_name, url, description) VALUES ('p3', 'p3_title', 'u2_name', 'p3_url', 'p3_description');
INSERT INTO products (id, title, user_name, url, description) VALUES ('p4', 'p4_title', 'u3_name', 'p4_url', 'p4_description');

INSERT INTO comments (id, product_id, message) VALUES ('c1', 'p1_product', 'c1_p1_message');
INSERT INTO comments (id, product_id, message) VALUES ('c2', 'p2_product', 'c2_p2_message');
INSERT INTO comments (id, product_id, message) VALUES ('c3', 'p1_product', 'c3_p1_message');
INSERT INTO comments (id, product_id, message) VALUES ('c4', 'p3_product', 'c4_p3_message');
INSERT INTO comments (id, product_id, message) VALUES ('c5', 'p4_product', 'c5_p4_message');
INSERT INTO comments (id, product_id, message) VALUES ('c6', 'p1_product', 'c6_p1_message');
INSERT INTO comments (id, product_id, message) VALUES ('c7', 'p1_product', 'c7_p1_message');
INSERT INTO comments (id, product_id, message) VALUES ('c8', 'p1_product', 'c8_p1_message');
INSERT INTO comments (id, product_id, message) VALUES ('c9', 'p2_product', 'c9_p2_message');
INSERT INTO comments (id, product_id, message) VALUES ('c10', 'p1_product', 'c10_p1_message');

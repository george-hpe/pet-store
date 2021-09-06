-- Create table Products.
CREATE TABLE Products(
    prod_id text PRIMARY KEY,
    prod_name char(10) NOT NULL,
    prod_category char(10) NOT NULL,
    prod_photo_url char(20) NULL,
    status char(10) NOT NULL
);

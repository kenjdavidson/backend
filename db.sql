CREATE TABLE channels (
	channelid varchar NOT NULL,
	channelname varchar NOT NULL,
	overlayid uuid NOT NULL,
	CONSTRAINT channels_pk PRIMARY KEY (channelid),
	CONSTRAINT channels_unique UNIQUE (overlayid)
);

CREATE TABLE items (
	itemid uuid NOT NULL,
	"name" varchar NOT NULL,
	rarity varchar NOT NULL,
	image varchar NOT NULL,
	previmg varchar NOT NULL,
	CONSTRAINT items_pk PRIMARY KEY (itemid)
);

CREATE TABLE users (
	userid varchar NOT NULL,
	username varchar NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (userid)
);

CREATE TABLE channelitems (
	channelid varchar NOT NULL,
	itemid uuid NOT NULL,
	CONSTRAINT channelitems_pk PRIMARY KEY (channelid, itemid),
	CONSTRAINT channelitems_unique UNIQUE (itemid),
	CONSTRAINT channelitems_channels_fk FOREIGN KEY (channelid) REFERENCES channels(channelid),
	CONSTRAINT channelitems_items_fk FOREIGN KEY (itemid) REFERENCES items(itemid)
);

CREATE TABLE defaultchannelitems (
	channelid varchar NOT NULL,
	itemid uuid NOT NULL,
	CONSTRAINT defaultchannelitems_pk PRIMARY KEY (channelid),
	CONSTRAINT defaultchannelitems_channelitems_fk FOREIGN KEY (channelid,itemid) REFERENCES channelitems(channelid,itemid)
);

CREATE TABLE owneditems (
	userid varchar NOT NULL,
	transactionid varchar NOT NULL,
	itemid uuid NOT NULL,
	channelid varchar NOT NULL,
	CONSTRAINT owneditems_pk PRIMARY KEY (userid, itemid, channelid),
	CONSTRAINT owneditems_unique UNIQUE (transactionid),
	CONSTRAINT owneditems_channelitems_fk FOREIGN KEY (channelid,itemid) REFERENCES channelitems(channelid,itemid),
	CONSTRAINT owneditems_users_fk FOREIGN KEY (userid) REFERENCES users(userid)
);
CREATE INDEX owneditems_userid_idx ON public.owneditems USING btree (userid, channelid);

CREATE TABLE schedule (
	scheduleid uuid NOT NULL,
	dayofweek varchar NOT NULL,
	itemid uuid NOT NULL,
	channelid varchar NOT NULL,
	CONSTRAINT schedule_pk PRIMARY KEY (scheduleid),
	CONSTRAINT schedule_channelitems_fk FOREIGN KEY (channelid,itemid) REFERENCES channelitems(channelid,itemid)
);

CREATE TABLE selecteditems (
	userid varchar NOT NULL,
	channelid varchar NOT NULL,
	itemid uuid NOT NULL,
	CONSTRAINT selecteditems_pk PRIMARY KEY (userid, channelid),
	CONSTRAINT selecteditems_channelitems_fk FOREIGN KEY (channelid,itemid) REFERENCES channelitems(channelid,itemid),
	CONSTRAINT selecteditems_users_fk FOREIGN KEY (userid) REFERENCES users(userid)
);
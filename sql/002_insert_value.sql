INSERT into events(description, start_time, duration)
values('some event', '1999-Jan-08 04:05:06', '60 seconds')
returning id;
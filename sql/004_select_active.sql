SELECT id, description, start_time, duration 
FROM events
WHERE '2019-11-11 13:11:06' BETWEEN start_time::timestamp AND start_time::timestamp + duration::interval

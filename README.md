# Alarm monitoring
Start mongoDB & RabbitMQ containers:
```bash
make up
```

Start monitoring service:
```bash
make run
```

Start fake Alarms producer:
```bash
make fake N
```
Where "N" is the interval of producing data, in seconds. 10 by default.

Stop mongoDB & RabbitMQ containers:
```bash
make down
```

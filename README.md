## Задание для стажировки в Avito.

#### API-система, которая работает с Redis-хранилищем, имеет три роута.

### Routes 
```
[GET]  localhost:8089/get_key?key=<value> 
[POST] localhost:8089/set_key, { "<key>": "<value>" }
[POST] localhost:8089/del_key, { "key"  : "<value>" }
```

### Docker-Compose
#### Configure environment docker-compose.yaml and start build
```
[SH] docker-compose up --build
```

### TODO
- [ ] Create Redis-authentication
- [ ] TLS
### Detalhes
- Projeto pessoal para comparações de tempo de resposta de APIs escritas em diferentes linguages e frameworks
- Teste: Ao bater no endpoint `/ping`, retorna `pong`
- Testes realizado em ambiente local - Ubuntu 20.04(Dell i15-3501-A70P - 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz - Mem: 19739 - Disk: KBG40ZNS256G NVMe KIOXIA 256GB) 

### Golang/Gin:

**Teste 001 - 100000 operações distribuidas em 5000 rotinas paralelas.**
- `Processs took 8.919223394 s`

**Teste 002 - 10000 operações distribuidas em 100 rotinas paralelas.**
- `Processs took 21.164387086 s`

**Teste 003 - 10000 operações distribuidas em 500 rotina paralela.**
- `Processs took 5.177227185 s`

**Teste 004 - 1 operação distribuida em 1 rotina paralela.**
- `Processs took 201.101177 ms`

---

### Golang/Goji:

**Teste 001 - 100000 operações distribuidas em 5000 rotinas paralelas.**
- `Processs took 8.918164613 s`

**Teste 002 - 10000 operações distribuidas em 100 rotinas paralelas.**
- `Processs took 21.179273444 s`

**Teste 003 - 10000 operações distribuidas em 500 rotina paralela.**
- `Processs took 4.872866028 s`

**Teste 004 - 1 operação distribuida em 1 rotina paralela.**
- `Processs took 201.234991 ms`

---

### Python/Flask:

**Teste 001 - 100000 operações distribuidas em 5000 rotinas paralelas.**
- `Processs took connection reset by peer`

**Teste 002 - 10000 operações distribuidas em 100 rotinas paralelas.**
- `Processs took 25.268225399 s`

**Teste 003 - 10000 operações distribuidas em 500 rotina paralela.**
- `Processs took 46.062441486 s`

**Teste 004 - 1 operação distribuida em 1 rotina paralela.**
- `Processs took 201.540055 ms`

---

### Java/Quarkus:

**Teste 001 - 100000 operações distribuidas em 5000 rotinas paralelas.**
- `Processs took 10.780763422 s`

**Teste 002 - 10000 operações distribuidas em 100 rotinas paralelas.**
- `Processs took 21.510778521 s`

**Teste 003 - 10000 operações distribuidas em 500 rotina paralela.**
- `Processs took 5.729320888 s`

**Teste 004 - 1 operação distribuida em 1 rotina paralela.**
- `Processs took 202.417531 ms`

---

### Rust/Rocket:

**Teste 001 - 100000 operações distribuidas em 5000 rotinas paralelas.**
- `Processs took 34.024275463 s`

**Teste 002 - 10000 operações distribuidas em 100 rotinas paralelas.**
- `Processs took 24.234940881 s`

**Teste 003 - 10000 operações distribuidas em 500 rotina paralela.**
- `Processs took 7.244937692 s`

**Teste 004 - 1 operação distribuida em 1 rotina paralela.**
- `Processs took 201.799862 ms`

---
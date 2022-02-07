Projeto pessoal para comparações de tempo de resposta de APIs escritas em diferentes linguages e frameworks
Teste: Ao bater no endpoint `/ping`, retorna `pong`

Testes realizado em ambiente local (Dell i15-3501-A70P - 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz - Mem: 19739 - Disk: KBG40ZNS256G NVMe KIOXIA 256GB) 

# Golang:

## Gin:

### Teste 001 - 100000 operações distribuidas em 5000 rotinas paralelas.
Tempo de processamento: 8.919223394 segundos

### Teste 002 - 10000 operações distribuidas em 100 rotinas paralelas.
Tempo de processamento: 21.164387086 segundos

### Teste 003 - 10 operações distribuidas em 1 rotina paralela.
Tempo de processamento: 2.016956596 segundos

### Teste 004 - 1 operação distribuida em 1 rotina paralela.
Tempo de processamento: 201.101177 milissegundos

---

## Goji:

### Teste 001 - 100000 operações distribuidas em 5000 rotinas paralelas.
Tempo de processamento: 8.918164613 segundos

### Teste 002 - 10000 operações distribuidas em 100 rotinas paralelas.
Tempo de processamento: 21.179273444 segundos

### Teste 003 - 10 operações distribuidas em 1 rotina paralela.
Tempo de processamento: 2.016085266 segundos

### Teste 004 - 1 operação distribuida em 1 rotina paralela.
Tempo de processamento: 201.234991 milissegundos

---

# Python:

## Flas:

### Teste 001 - 100000 operações distribuidas em 5000 rotinas paralelas.
Tempo de processamento: connection reset by peer

### Teste 002 - 10000 operações distribuidas em 100 rotinas paralelas.
Tempo de processamento: 25.268225399 segundos

### Teste 003 - 10 operações distribuidas em 1 rotina paralela.
Tempo de processamento: 2.0268312 segundos

### Teste 004 - 1 operação distribuida em 1 rotina paralela.
Tempo de processamento: 201.540055 milissegundos

---
[[Go]]

В этой главе кратко вводят в курс дела, что это за драконы такие. Как и пишет сам автор - об этом полезно знать, так как при копировании готовых решений там могут использоваться эти пакеты и нужно знать, что они позволяют делать.

Опишу кратко для чего каждый пакет:
- `reflext` позволяет манипулировать типами на этапе выполнения. Т.е мы можем получать тип на что указывает указатель, можем пробежаться индексами по полям структуры, получить ее теги. Но не зря этот пакет - дракон. Код с использованием рефлексии работает гораздо медленнее.
- `unsafe` позволяет манипулировать памятью. В большинстве случаев он используется для интеграции с другими системами и с C-кодом (стр 381). Также с помощью пакета `unsafe` можно повысить производительность при чтении данных из сети. В общем штука низкоуровневая и я пока не очень хочу с ней разбираться.
- `cgo` позволяет выполнять C-код. Только проблема в том, что работать это все будет очень медленно, так как в си и в го разная модель памяти, а из-за этого будет больше нагружаться сборщик мусора.

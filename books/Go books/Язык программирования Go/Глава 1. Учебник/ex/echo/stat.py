measures = ["45.83s", "28.28mili", "40.89mili", "182.3micro", "182.6micro", "40.77mili", "14.79mili", "14.76mili"]
name = [
    "Join через for, длинные строки",
    "Join через for, короткие строки",
    "Join через билдер. Аллокация памяти не учитывается. Длинные строки",
    "Join через билдер. Аллокация памяти не учитывается. Короткие строки",
    "Join через билдер. Аллокация памяти учитывается. Короткие строки",
    "Join через билдер. Аллокация памяти учитывается. Длинные строки",
    "Join через strings.Join. Длинные строки",
    "Join через strings.Join. Короткие строки",
]                                 

res = []

for m, n in zip(measures, name):
    if "s" in m:
        d, q = m.split("s")
        d = float(d)
        res.append([n, d])
    elif "mili" in m:
        d, q = m.split("mili")
        d = float(d)
        res.append([n, d / 1000])
    elif "micro" in m:
        d, q = m.split("micro")
        d = float(d)
        res.append([n, d / 1_000_000])

res = sorted(res, key=lambda item: item[1])

for name, time in res:
    print(name, time)

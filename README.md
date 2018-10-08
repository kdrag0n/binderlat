# BinderLat

This program measures the time elapsed for each Binder IPC transaction on Android.

Binder is the custom form of inter-process communication used by Android. Requests happen *constantly* and are esssential for any part of the system to work. Because of this, it is obviously important that Binder is fast.

# Building
`GOOS=linux GOARCH=arm64 go build -ldflags="-s -w"` should build the program and create a `binderlat` file. If you want a smaller build, run `build.sh`.

# Example Output
```Transaction 13637:13637 took   81.2us; total 371.4ms
Transaction 13637:13637 took 1834.2us; total 373.3ms
Transaction   3921:4257 took 2032.5us; total 375.3ms
Transaction   3921:5540 took  180.3us; total 375.5ms
Transaction 13637:13637 took   76.1us; total 375.5ms
Transaction   3921:5540 took   35.1us; total 375.6ms
Transaction   3921:5540 took   56.2us; total 375.6ms
Transaction 13637:13637 took 1355.7us; total 377.0ms
Transaction   3921:5540 took 1458.8us; total 378.5ms
Transaction 13637:13637 took 1744.8us; total 380.2ms
Transaction 13637:13637 took   54.0us; total 380.3ms
Transaction   3921:5540 took   37.8us; total 380.3ms
Transaction 13637:13637 took 1351.1us; total 381.6ms
Transaction   3921:5540 took  211.0us; total 381.9ms
Transaction   3493:3493 took   40.2us; total 381.9ms
Transaction   3282:3282 took   72.7us; total 382.0ms
Transaction   3921:4257 took 3641.9us; total 385.6ms
Transaction   3921:4257 took 1543.7us; total 387.2ms
Transaction 13637:13637 took 1926.2us; total 389.1ms
Transaction   3921:4257 took   39.6us; total 389.1ms
Transaction 13637:13637 took 2391.3us; total 391.5ms
Transaction   3921:3921 took 328538.5us; total 720.0ms
Transaction   3921:4257 took 6690.4us; total 726.7ms
Transaction   3506:3538 took   78.9us; total 726.8ms
Transaction   3921:4257 took 3370.9us; total 730.2ms
Transaction   3921:5540 took  324.2us; total 730.5ms
Transaction   3493:3493 took   39.6us; total 730.6ms
Transaction   3506:3538 took 3376.5us; total 733.9ms
Transaction   3921:4257 took   49.6us; total 734.0ms
Transaction   3506:3538 took   60.8us; total 734.0ms
Transaction   3921:4257 took 3861.0us; total 737.9ms
Transaction   3506:3538 took   58.5us; total 738.0ms
Transaction   3921:4257 took 5204.3us; total 743.2ms
Transaction   3506:3538 took   45.8us; total 743.2ms
Transaction   3921:4257 took  201.5us; total 743.4ms
Transaction   3921:4091 took   92.0us; total 743.5ms
Transaction   3494:4087 took 181115.6us; total 924.6ms
Transaction   3921:5540 took  545.9us; total 925.2ms
Transaction   3493:3493 took  411.7us; total 925.6ms
Transaction 13637:13637 took  179.3us; total 925.8ms
Transaction   3506:4145 took   91.3us; total 925.8ms
Transaction   3506:3538 took 1993.6us; total 927.8ms
Transaction   3506:4145 took  159.3us; total 928.0ms
Transaction 13637:13637 took  664.7us; total 928.7ms
Transaction   3921:3976 took 2693.0us; total 931.4ms
Transaction   3506:4145 took   57.8us; total 931.4ms
Transaction   3921:5540 took   47.4us; total 931.5ms
Transaction 13637:13637 took   49.8us; total 931.5ms
Transaction   3493:3493 took   30.6us; total 931.5ms
Transaction   3921:5540 took  238.8us; total 931.8ms
Transaction   3921:5540 took   39.0us; total 931.8ms
Transaction 13637:13637 took   82.3us; total 931.9ms
Transaction   3506:4145 took 1598.7us; total 933.5ms
Transaction 13637:13637 took  298.1us; total 933.8ms
Transaction   3506:4145 took   62.7us; total 933.9ms
Transaction 13637:13637 took   47.8us; total 933.9ms
Transaction   3506:4145 took   30.0us; total 933.9ms
Transaction   3490:3490 took 1290.2us; total 935.2ms
Transaction   3506:4145 took 1823.5us; total 937.1ms
Transaction 13637:13637 took  101.1us; total 937.2ms
Transaction   3506:4145 took   45.6us; total 937.2ms
Transaction   3921:5540 took   54.7us; total 937.3ms
Transaction   3493:3493 took  446.8us; total 937.7ms
Transaction   3506:4145 took  120.7us; total 937.8ms
Transaction 13637:13637 took 12655.2us; total 950.5ms
Transaction   3921:5540 took  300.3us; total 950.8ms
Transaction   3493:3493 took   31.8us; total 950.8ms
Transaction 13637:13637 took  258.7us; total 951.1ms
Transaction   3921:4257 took  129.2us; total 951.2ms
Transaction   3921:4084 took 46749.2us; total 997.9ms
Transaction   3921:5540 took   59.5us; total 998.0ms
Transaction   3493:3493 took   39.0us; total 998.0ms
Transaction   3506:4145 took 1597.0us; total 999.6ms
Transaction   3921:3976 took 47954.4us; total 1047.6ms
Transaction   3506:3538 took   39.8us; total 1047.6ms
Transaction   3921:4084 took 7563.5us; total 1055.2ms
Transaction   3506:3538 took   52.0us; total 1055.3ms
Transaction   3921:4084 took  263.0us; total 1055.5ms
Transaction   3506:3538 took   35.3us; total 1055.6ms
Transaction   3921:4084 took 2075.7us; total 1057.6ms
Transaction   3506:3538 took   70.4us; total 1057.7ms
Transaction   3921:4084 took 9733.8us; total 1067.4ms
Transaction   3506:3538 took   62.2us; total 1067.5ms
Transaction 13637:13637 took 59832.3us; total 1127.3ms
Transaction   3921:5540 took   84.3us; total 1127.4ms
Transaction   3493:3493 took   54.0us; total 1127.5ms
Transaction   3494:3494 took 180606.1us; total 1308.1ms
Transaction   3921:4091 took   87.7us; total 1308.2ms```

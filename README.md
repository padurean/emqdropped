## emqdropped

Command line utility that displays dropped messages given an [EMQTT](http://emqtt.io/) log file obtained by running EMQTT with log level on `debug`.

### Usage

`macos/emqdropped emq.sample.log` to show just the count of dropped messages per topic

OR

`macos/emqdropped emq.sample.log true` to also show the payloads of dropped messages

Sample output:

```
8 published
4 delivered
4 dropped
         asset/commands/123 : 2 dropped
                 Some setpoint 234.67
                 Some setpoint 678.90
         asset/telemetry/123 : 2 dropped
                 <Data><TimeStamp _=\"2018-09-02 14:24:02 +0100\"/><Voltage_L1_N _=\"222.7\" flag=\"1\"/><Voltage_L2_N _=\"222.4\" flag=\"1\"/><Voltage_L3_N _=\"222.3\" flag=\"1\"/><RealEnergy_L1L2L3Con _=\"22800.0\" flag=\"1\"/><Frequency _=\"50.00\" flag=\"1\"/><RealPowerPSum_P1P2P3 _=\"1010.0\" flag=\"1\"/><RealEnergy_L1L2L3Del _=\"000.0\" flag=\"1\"/><PowerSetpointActive1 _=\"0\" flag=\"2\"/><PowerSetpointActive2 _=\"0\" flag=\"2\"/><ActivePower _=\"30.3\" flag=\"2\"/><Druck_mqtt_=\"0\" flag=\"2\"/><Frequency_Modbus _=\"10\" flag=\"2\"/><DigitalInput1 _=\"1\" flag=\"2\"/><CPULoad _=\"70\" flag=\"2\"/></Data>
                 <Data><TimeStamp _=\"2018-09-02 14:24:02 +0100\"/><Voltage_L1_N _=\"222.7\" flag=\"1\"/><Voltage_L2_N_=\"222.4\" flag=\"1\"/><Voltage_L3_N _=\"222.3\" flag=\"1\"/><RealEnergy_L1L2L3Con _=\"22800.0\" flag=\"1\"/><Frequency _=\"50.00\" flag=\"1\"/><RealPowerPSum_P1P2P3 _=\"1010.0\" flag=\"1\"/><RealEnergy_L1L2L3Del _=\"000.0\" flag=\"1\"/><PowerSetpointActive1 _=\"0\" flag=\"2\"/><PowerSetpointActive2 _=\"0\" flag=\"2\"/><ActivePower _=\"40.3\" flag=\"2\"/><Druck_mqtt _=\"0\" flag=\"2\"/><Frequency_Modbus_=\"10\" flag=\"2\"/><DigitalInput1 _=\"1\" flag=\"2\"/><CPULoad _=\"70\" flag=\"2\"/></Data>
FINISHED.
 ```

### Built and cross-compiled using [Go](https://golang.org/)

Binaries for **MacOS**, **Linux** and **Windows** are available in folders with suggestive names.

In case of errors compile it yourself:

- On MacOS: `go build -o macos/emqdropped`
- On Linux: `go build -o linux-amd64/emqdropped`
- On Windows: `go build -o windows/emqdropped.exe`

Cross-compile on MacOS:

- For Linux: `env GOOS=linux GOARCH=amd64 go build -o linux-amd64/emqdropped`
- For Windows: `env GOOS=windows GOARCH=386 go build -o windows/emqdropped.exe`

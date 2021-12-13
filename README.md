# telegraf-exec-storcli

This is a simple tool to extract [StorCLI](https://www.broadcom.com/) output and output
[Influx line protocol](https://docs.influxdata.com/influxdb/cloud/reference/syntax/line-protocol/);
it is designed to be used with a
[telegraf exec plugin](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/exec).

This parses the output of `storcli show [/cx] all J` and has been developed against Ubuntu 21.10, StorCLi 007.1904, and a 9400-16i HBA.

## Inreractive Run Example

The compiled tool can be run interactively. It needs to run as root or with sudo.

```bash
sudo ./telegraf-exec-storcli
```

## Telegraf Run Example

This is a sample telegraf exec input that assumes the binary has been installed
to `/usr/local/bin/telegraf-exec-storcli`:

```toml
[[inputs.exec]]                                                                 
  commands = ["/usr/local/bin/telegraf-exec-storcli"]
  timeout = "10s"                                                                
  data_format = "influx"      
```

Then in InfluxDB, the `storcli_controller` measurement will have these tags:

```
adapter
host
model
serial
```

And these fields:

```
memory_correctable_errors
memory_uncorrectable_errors
roc_temperature
status
```

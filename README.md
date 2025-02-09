<h1 align="center">seasonpackarr</h1>
<h1 align="center">
  <a href="https://github.com/nuxencs/seasonpackarr/blob/develop/LICENSE">
    <img src="https://img.shields.io/github/license/nuxencs/seasonpackarr?style=flat-square&color=00ACD7" alt="License">
  </a>
  <a href="https://goreportcard.com/report/github.com/nuxencs/seasonpackarr">
    <img src="https://goreportcard.com/badge/github.com/nuxencs/seasonpackarr?style=flat-square" alt="Go Report">
  </a>
  <a href="https://github.com/nuxencs/seasonpackarr/actions">
    <img src="https://img.shields.io/github/actions/workflow/status/nuxencs/seasonpackarr/release.yml?style=flat-square&logo=github" alt="Build">
  </a>
    <a href="https://github.com/nuxencs/seasonpackarr/releases">
    <img src="https://img.shields.io/github/v/release/nuxencs/seasonpackarr?style=flat-square&color=00ACD7" alt="Latest Release">
  </a>
  <a href="https://trash-guides.info/discord">
    <img src="https://img.shields.io/discord/492590071455940612?style=flat-square&logo=discord&logoColor=00ACD7&label=support&color=00ACD7" alt="Discord">
  </a>
</h1>

<p align="center">
<b>seasonpackarr</b> is a companion app for <a href="https://github.com/autobrr/autobrr">autobrr</a> that automagically <b>hardlinks</b> downloaded episodes into a season folder when a season pack is
announced, eliminating the need for re-downloading existing episodes.
</p>

> [!WARNING]
> This application is currently under active development. If you encounter any bugs, please report them in the dedicated
> #seasonpackarr channel on the TRaSH-Guides [Discord server](https://trash-guides.info/discord) or create a new issue
> on GitHub, so I can fix them.

## Installation

### Linux

To download the latest release, you can use one of the following methods:

```bash
# using curl
curl -s https://api.github.com/repos/nuxencs/seasonpackarr/releases/latest | grep download | grep linux_x86_64 | cut -d\" -f4 | xargs curl -LO

# using wget
wget -qO- https://api.github.com/repos/nuxencs/seasonpackarr/releases/latest | grep download | grep linux_x86_64 | cut -d\" -f4 | xargs wget
```

Alternatively, you can download the [source code](https://github.com/nuxencs/seasonpackarr/releases/latest) and build it yourself using `go build`.

#### Unpack

Run with `root` or `sudo`. If you do not have root, or are on a shared system, place the binary somewhere in your home
directory like `~/.bin`.

```bash
tar -C /usr/bin -xzf seasonpackarr*.tar.gz
```

This will extract `seasonpackarr` to `/usr/bin`.

Afterwards you need to make the binary executable by running the following command.

```bash
chmod +x /usr/bin/seasonpackarr
```

Note: If the commands fail, prefix them with `sudo ` and run them again.

#### Systemd (Recommended)

On Linux-based systems, it is recommended to run seasonpackarr as a sort of service with auto-restarting capabilities,
in order to account for potential downtime. The most common way is to do it via systemd.

You will need to create a service file in `/etc/systemd/system/` called `seasonpackarr@.service`.

```bash
touch /etc/systemd/system/seasonpackarr@.service
```

Then place the following content inside the file (e.g. via nano/vim/ed):

```systemd title="/etc/systemd/system/seasonpackarr@.service"
[Unit]
Description=seasonpackarr service for %i
After=syslog.target network-online.target

[Service]
Type=simple
User=%i
Group=%i
ExecStart=/usr/bin/seasonpackarr start --config=/home/%i/.config/seasonpackarr

[Install]
WantedBy=multi-user.target
```

Start the service. Enable will make it startup on reboot.

```bash
sudo systemctl enable -q --now seasonpackarr@$USER
```

Make sure it's running and **active**.

```bash
sudo systemctl status seasonpackarr@$USER
```

On first run it will create a default config, `~/.config/seasonpackarr/config.yaml` that you will need to edit.

After the config is edited you need to restart the service.

```bash
sudo systemctl restart seasonpackarr@$USER.service
```

### Docker

Docker images can be found on the right under the "Packages" section.

See `docker-compose.yml` for an example.

Make sure you use the correct path you have mapped within the container in the config file. After the first start you
will need to adjust the created config file to your needs and start the container again.

## Configuration

You can configure a decent part of the features seasonpackarr provides. I will explain the most important ones here in
more detail.

### Smart Mode

Can be enabled in the config by setting `smartMode` to `true`. Works together with `smartModeThreshold` to determine if
a season pack should get grabbed or not. Here's an example that explains it pretty well:

Let's say you have 8 episodes of a season in your client released by `RlsGrpA`. You also have 12 episodes of the same
season in your client released by `RlsGrpB` and there are a total of 12 episodes in that season. If you have smart
mode enabled with a threshold set to `0.75`, only the season pack from `RlsGrpB` will get grabbed, because `8/12 = 0.67`
which is below the threshold.

### Parse Torrent

Can be enabled in the config by setting `parseTorrentFile` to `true`. This option will make sure that the season pack
folder that gets created by seasonpackarr will always have the correct name. One example that will make the benefit
of this clearer:

- Announce name: `Show.S01.1080p.WEB-DL.DDPA5.1.H.264-RlsGrp`
- Folder name: `Show.S01.1080p.WEB-DL.DDP5.1.H.264-RlsGrp`
   
Using the announce name would create the wrong folder and would lead to all the files in the torrent being downloaded
again. The issue in the given example is the additional `A` after `DDP` which is not present in the folder name. By
using the parsed folder name the files will be hardlinked into the exact folder that is being used in the torrent.

You can take a look at the [Webhook](#webhook) section to see what you would need to add in your autobrr filter to
make use of this feature.

### Fuzzy Matching

In this section, you can toggle comparing rules. I will explain each of them in more detail here.

1. **skipRepackCompare**: When set to `true`, the comparer skips checking the repack status of the season pack release
   against the episodes in your client. The episode in the example will only be accepted as a match by seasonpackarr if
   you enable this option:
   - Announce name: `Show.S01.1080p.WEB-DL.DDPA5.1.H.264-RlsGrp`
   - Episode name: `Show.S01E01.1080p.WEB-DL.REPACK.DDPA5.1.H.264-RlsGrp`

2. **simplifyHdrCompare**: If set to `true`, this option simplifies the HDR formats `HDR10`, `HDR10+`, and `HDR+` to
   just `HDR`. This increases the likelihood of matching renamed releases that specify a more advanced HDR format in the
   announce name than in the episode title:
   - Announce name: `Show.S01.2160p.WEB-DL.DDPA5.1.DV.HDR10+.H.265-RlsGrp`
   - Episode name: `Show.S01E01.2160p.WEB-DL.DDPA5.1.DV.HDR.H.265-RlsGrp`

### Recommended options

Keep in mind, these settings are suggestions based on my own use case so feel free to adjust them according to your
specific needs.

```yaml
smartMode: true
smartModeThreshold: 0.75
parseTorrentFile: true
skipRepackCompare: true
simplifyHdrCompare: false
```

These will filter out most unwanted season packs and prevent mismatches, while still making sure that
renamed season packs and episodes can get matched.

## autobrr Filter setup

Support for multiple Sonarr and qBittorrent instances with different pre import directories was added with v0.4.0, so
you will need to run multiple instances of seasonpackarr and create multiple filters to achieve the same functionality
in lower versions. If you are running v0.4.0 or above you just need to set up your filters according to [External Filters](#external-filters).
The following is a simple example filter that only allows 1080p season packs to be matched.

### Create Filter

To import it into autobrr you need to navigate to `Filters` and click on the arrow next to `+ Create Filter` to see the
option `Import filter`. Just paste the content below into the text box that appeared and click on `Import`.

```json
{
  "name": "arr-Seasonpackarr",
  "version": "1.0",
  "data": {
    "enabled": true,
    "seasons": "1-99",
    "episodes": "0",
    "resolutions": [
      "1080p",
      "1080i"
    ]
  }
}
```

In the `General` tab you will need to adjust the value of `Priority` to be set higher than all your TV show filters. For
instance, if your Sonarr filter is set at `10` and a TV filter that sends to qBittorrent is at `15`, then you should set
the `seasonpackarr` filter to at least `16`. This ensures that it will execute before the others. It's perfectly fine to
have a `cross-seed` filter positioned above the `seasonpackarr` filter.

### External Filters

After adding the filter, you need to head to the `External` tab of the filter, click on `Add new` and select `Webhook`
in the `Type` field. The `Endpoint` field should look like this, with `host` and `port` taken from your config:

```
http://host:port/api/pack
```

`HTTP Method` needs to be set to `POST`, `Expected HTTP status` has to be set to `250` and the `Data (JSON)` field needs
to look like this:

```json
{
  "name": "{{ .TorrentName }}",
  "clientname": "default"
}
```

Replace the `clientname` value, in this case `default`, with the name you gave your desired qBittorrent client in your
config under the `clients` section. If you don't specify a `clientname` in the JSON payload, seasonpackarr will try to
use the `default` client; if you renamed or removed the `default` client the request will fail.

#### API Authentication

I strongly suggest enabling API authentication by providing an API token in the config. The following command will
generate a token for you that you can copy and paste into your config:

```bash
seasonpackarr gen-token
```

After you've set the API token in your config, you'll need to either include it in the `Endpoint` field or pass it
along in the `HTTP Request Headers` of your autobrr request; if not, the request will be rejected. I recommend using
headers to pass the API token, but I'll explain both options here.

1. **Header**: Edit the `HTTP Request Headers` field and replace `api_token` with the token you set in your config.
    ```
    X-API-Token=api_token
    ```
2. **Query Parameter**: Append `?apikey=api_token` at the end of your `Endpoint` field and replace `api_token` with the
   token you've set in your config.
    ```
    http://host:port/api/pack?apikey=api_token
    ```

The external filter you just created will be disabled by default. To avoid unwanted downloads, make sure to enable it!

### Actions

Now, you need to decide whether you want to enable torrent parsing. By activating this feature, seasonpackarr will parse
the torrent file for the season pack folder name to ensure the creation of the correct folder. You can enable this
functionality by setting `parseTorrentFile` to `true` in your config file.

If you choose to enable this feature, first follow the instructions in the [Webhook](#webhook) section, and then proceed
to the [qBittorrent](#qbittorrent) section. If you leave this feature disabled, you can skip the Webhook section and go
straight to the qBittorrent section.

> [!WARNING]
> If you enable that option you need to make sure that the Webhook action is above the qBittorrent action, otherwise the
> feature won't work correctly.

#### Webhook

Navigate to the `Actions` tab, click on `Add new` and change the `Action type` of the newly added action to `Webhook`.
The `Endpoint` field should look like this, with `host`, `port` and `api_token` taken from your config:

```
http://host:port/api/parse?apikey=api_token
```

Append the API query parameter `?apikey=api_token` only if you have enabled API authentication by providing an API token
in your config.

Finally, complete the `Payload (JSON)` field as shown below. Ensure that the value of `clientname` is the same as in the `External Filter`:

```json
{
  "name":"{{ .TorrentName }}", 
  "torrent":"{{ .TorrentDataRawBytes | js }}",
  "clientname": "default"
}
```

#### qBittorrent

Navigate to the `Actions` tab, click on `Add new` and change the `Action type` of the newly added action to `qBittorrent`.
Depending on whether you intend to only send to qBittorrent or also integrate with Sonarr, you'll need to fill out different fields.

1. **Only qBittorrent**: Fill in the `Save Path` field with the directory where your torrent data resides, for instance
   `/data/torrents`, or the `Category` field with a qBittorrent category that saves to your desired location. 
2. **Sonarr Integration**: Fill in the `Category` field with the category that Sonarr utilizes for all its downloads,
   such as `tv-hd` or `tv-uhd`.

Last but not least, under `Rules`, make sure that `Skip Hash Check` remains disabled. This precaution prevents torrents
added by seasonpackarr from causing errors in your qBittorrent client when some episodes of a season are missing.

> [!WARNING]
> If you enable that option regardless, you will most likely have to deal with errored torrents, which would require you
> to manually trigger a recheck on them to fix the issue.

## Credits

Huge credit goes to [upgraderr](https://github.com/KyleSanderson/upgraderr) and specifically [@KyleSanderson](https://github.com/KyleSanderson), whose
project provided great functions that I could make use of. Additionally, I would also like to mention [@zze0s](https://github.com/zze0s), who was
really helpful regarding any question I had as well as providing me with a lot of the structure this project has now.
Credits also go to the [TVmaze API](https://www.tvmaze.com/api) for providing comprehensive data on the total number of episodes for
a show in a specific season. Last but not least, a big thank you to [JetBrains](http://www.jetbrains.com/) for providing me with free licenses to their great tools,
in this case [GoLand](https://www.jetbrains.com/go/).

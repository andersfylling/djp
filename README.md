# Discord JSON Parser (jdp)
Discord JSON Parser is tool to manipulate raw discord json data before unmarshalling. This is to avoid unecessary heap allocations for very large objects.

This project was first written as an internal component of [Disgord](https://github.com/andersfylling/disgord) such that different cache repositories (user, guild, channels, etc.) could run in parallel on incoming discord events, by reading a shared byte slice.

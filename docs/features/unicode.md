# Unicode rendering
The Terminal UI is capable of rendering multibyte [Unicode](https://home.unicode.org/) characters from many different languages, including non multi-character emojis. The concrete results depend upon the used terminal and the configured font. A monospaced [Nerd Font](https://www.nerdfonts.com/) is suggested, but not required.

## Security
To mitigate a known [vulnerability](https://nvd.nist.gov/vuln/detail/CVE-2021-42574) regarding log files, all bidirectional Unicode characters will not be processed and displayed as the `×` character.

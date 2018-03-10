# csvf

```bash
$ cat discogs-collection.csv
"44 073, RS 6349",Neil Young & Crazy Horse,Everybody Knows This Is Nowhere...
"9362-49384-9, 48497-1",Neil Young,On The Beach...
...
$ csvf discogs-collection.csv artist=\"%1\" album=\"%2\"
artist="Neil Young & Crazy Horse" album="Everybody Knows This Is Nowhere"
artist="Neil Young" album="On The Beach"
...
```
